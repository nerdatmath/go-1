// Copyright 2013, 2022 Tamás Gulácsi. All rights reserved.
// Use of this source code is governed by an Apache 2.0
// license that can be found in the LICENSE file.

package i18nmail

import (
	"bufio"
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/textproto"
	"net/url"
	"strings"
	"sync/atomic"

	"github.com/emersion/go-message"
	tp "github.com/emersion/go-message/textproto"

	// charsets
	_ "github.com/emersion/go-message/charset"

	"github.com/go-logr/logr"
	"github.com/tgulacsi/go/iohlp"
)

// MaxWalkDepth is the maximum depth Walk will descend.
const (
	MaxWalkDepth  = 32
	bodyThreshold = 1 << 20
)

var (
	logger = logr.Discard()

	// CheckEncoding is true if we should check Base64 encodings
	CheckEncoding = true

	// SaveBadInput is true if we should save bad input
	SaveBadInput = false

	// ErrStopWalk shall be returned by the TodoFunc to stop the walk silently.
	ErrStopWalk = errors.New("stop the walk")
)

// SetLogger sets the global logger
func SetLogger(lgr logr.Logger) { logger = lgr }

// TodoFunc is the type of the function called by Walk and WalkMultipart.
type TodoFunc func(mp MailPart) error

// sequence is a global sequence for numbering mail parts.
var sequence uint64

func nextSeq() uint64 {
	return atomic.AddUint64(&sequence, 1)
}
func nextSeqInt() int {
	return int(nextSeq() % uint64(1<<31))
}

// HashKeyName is the header key name for the hash
const HashKeyName = "X-HashOfFullMessage"

// MailPart is part of a mail or multipart message.
type MailPart struct {
	entity *message.Entity
	body   *io.SectionReader
	Header textproto.MIMEHeader
	// MediaType is the parsed media type.
	MediaType map[string]string
	// Parent of this part.
	Parent *MailPart
	// ContenType for the part.
	ContentType string
	// Level is the depth level.
	Level int
	// Seq is a sequence number
	Seq int
}

// NewMailPart parses the io.Reader as a full email message and returns it as a MailPart.
func NewMailPart(r io.Reader) (MailPart, error) {
	var mp MailPart
	msg, err := message.Read(r)
	if err != nil {
		return mp, err
	}
	return mp.WithEntity(msg)
}

// Entity returns the part as a *message.Entity
func (mp MailPart) Entity() *message.Entity {
	return mp.entity
}

// String returns some string representation of the part.
func (mp MailPart) String() string {
	pseq := -1
	if mp.Parent != nil {
		pseq = mp.Parent.Seq
	}
	return fmt.Sprintf("%d:::{%s %s %s}", pseq, mp.ContentType, mp.MediaType, mp.Header)
}

// Spawn returns a descendant of the MailPart (Level+1, Parent=*mp, next sequence).
func (mp MailPart) Spawn() MailPart {
	return MailPart{Parent: &mp, Level: mp.Level + 1, Seq: nextSeqInt()}
}

// GetBody returns an intact *io.SectionReader of the body.
func (mp MailPart) GetBody() *io.SectionReader {
	return io.NewSectionReader(mp.body, 0, mp.body.Size())
}

// NewEntity returns a new *message.Entity from the header map and body readers.
func NewEntity(header map[string][]string, body io.Reader) (*message.Entity, error) {
	hdr := message.HeaderFromMap(header)
	return message.New(fixHeader(hdr), body)
}

// NewEntityFromReaders returns a new *message.Entity from the header and body readers.
func NewEntityFromReaders(header, body io.Reader) (*message.Entity, error) {
	hdr, err := tp.ReadHeader(bufio.NewReader(io.LimitReader(header, 1<<20)))
	if err != nil {
		return nil, err
	}
	return message.New(fixHeader(message.Header{Header: hdr}), body)
}

func fixHeader(hdr message.Header) message.Header {
	if cte := hdr.Get("Content-Transfer-Encoding"); cte != "" && strings.ToLower(cte) != cte {
		hdr.Set("Content-Transfer-Encoding", cte)
	}
	return hdr
}

// WithReader returns a MailPart parsing the io.Reader as a full email.
func (mp MailPart) WithReader(r io.Reader) (MailPart, error) {
	entity, err := message.Read(r)
	if err != nil {
		return mp, err
	}
	return mp.WithEntity(entity)
}

// WithBody replaces only the body part.
func (mp MailPart) WithBody(r io.Reader) (MailPart, error) {
	entity := *mp.entity
	entity.Body = r
	return mp.WithEntity(&entity)
}

// WithEntity populates MailPart with the parsed *message.Entity.
func (mp MailPart) WithEntity(entity *message.Entity) (MailPart, error) {
	mp.entity = entity
	hdr := mp.entity.Header
	ct, params, err := hdr.ContentType()
	if err != nil {
		return mp, err
	}
	if ct == "" {
		ct = "message/rfc822"
	}
	mp.ContentType, mp.MediaType = ct, params
	hsh := sha512.New512_224()
	body, err := MakeSectionReader(io.TeeReader(mp.entity.Body, hsh), bodyThreshold)
	if err != nil {
		return mp, fmt.Errorf("part %s: %w", mp, err)
	}
	var a [sha512.Size224]byte

	mp.body, mp.entity.Body = body, io.NewSectionReader(body, 0, body.Size())
	if !mp.entity.Header.Has(HashKeyName) {
		mp.entity.Header.Set(HashKeyName, base64.URLEncoding.EncodeToString(hsh.Sum(a[:0])))
	}

	mp.Header = headersAsMap(mp.entity.Header)
	return mp, err
}
func headersAsMap(headers message.Header) textproto.MIMEHeader {
	fields := headers.Fields()
	m := make(map[string][]string, fields.Len())
	for fields.Next() {
		k := textproto.CanonicalMIMEHeaderKey(fields.Key())
		s, err := fields.Text()
		if err != nil {
			b, _ := fields.Raw()
			s = string(b)
		}
		m[k] = append(m[k], s)
	}
	return textproto.MIMEHeader(m)
}

// FileName returns the file name from the Content-Disposition header.
func (mp MailPart) FileName() string {
	_, params, _ := mp.entity.Header.ContentDisposition()
	return params["filename"]
}

// MakeSectionReader reads the reader and returns the byte slice.
//
// If the read length is below the threshold, then the bytes are read into memory;
// otherwise, a temp file is created, and mmap-ed.
func MakeSectionReader(r io.Reader, threshold int) (*io.SectionReader, error) {
	return iohlp.MakeSectionReader(r, threshold)
}

// WalkMessage walks over the parts of the email, calling todo on every part.
// The part.Body given to todo is reused, so read if you want to use it!
//
// By default this is recursive, except dontDescend is true.
func WalkMessage(msg *message.Entity, todo TodoFunc, dontDescend bool, parent *MailPart) error {
	var child MailPart
	if parent != nil {
		child = parent.Spawn()
	} else {
		child.Level = 1
		child.Seq = nextSeqInt()
	}
	var err error
	if child, err = child.WithEntity(msg); err != nil {
		return err
	}
	return Walk(child, todo, dontDescend)
}

// Walk over the parts of the email, calling todo on every part.
//
// By default this is recursive, except dontDescend is true.
func Walk(part MailPart, todo TodoFunc, dontDescend bool) error {
	return part.entity.Walk(func(path []int, entity *message.Entity, err error) error {
		if err != nil {
			logger.Error(err, "go-message.Walk", "path", path)
			return err
		}
		ct, params, _ := entity.Header.ContentType()
		if logger.V(1).Enabled() {
			logger.V(1).Info("entity", "path", path, "headers", headersAsMap(entity.Header), "contentType", ct, "params", params)
		}
		if path != nil && strings.HasPrefix(ct, "multipart/") && params["boundary"] != "" {
			logger.Info("SKIP", "contentType", ct, "params", params, "path", path)
			return nil
		}
		child, err := part.Spawn().WithEntity(entity)
		if err != nil {
			return err
		}
		fn := child.FileName()
		if fn == "" {
			ext, _ := mime.ExtensionsByType(child.ContentType)
			fn = fmt.Sprintf("%d.%d%s", child.Level, child.Seq, append(ext, ".dat")[0])
		}
		child.Header.Add("X-FileName", safeFn(fn, true))
		if err = todo(child); err != nil {
			return fmt.Errorf("todo(%q): %w", fn, err)
		}
		return todo(child)
	})
}

// WalkMultipart walks a multipart/ MIME parts, calls todo on every part
// mp.Body is reused, so read if you want to use it!
//
// By default this is recursive, except dontDescend is true.
func WalkMultipart(mp MailPart, todo TodoFunc, dontDescend bool) error {
	return Walk(mp, todo, dontDescend)
}

// HashBytes returns a hash (sha512_224 atm) for the given bytes
func HashBytes(data []byte) string {
	h := sha512.New512_224()
	_, _ = h.Write(data)
	var a [sha512.Size224]byte
	return base64.URLEncoding.EncodeToString(h.Sum(a[:0]))
}

func safeFn(fn string, maskPercent bool) string {
	fn = url.QueryEscape(
		strings.Replace(strings.Replace(fn, "/", "-", -1),
			`\`, "-", -1))
	if maskPercent {
		fn = strings.Replace(fn, "%", "!P!", -1)
	}
	return fn
}

// DecodeHeaders decodes the headers.
func DecodeHeaders(hdr map[string][]string) map[string][]string {
	for k, vv := range hdr {
		for i, v := range vv {
			vv[i] = HeadDecode(v)
		}
		hdr[k] = vv
	}
	return hdr
}

func hasBoundaries(body *io.SectionReader, prefixS string) bool {
	// Strip boundary from the beginning and the end.
	// "Boundary delimiters must not appear within the encapsulated material, and must be no longer than 70 characters, not counting the two leading hyphens."
	// -- https://www.rfc-editor.org/rfc/rfc2046#section-5.1.1
	// The preceeding -- is mandatory for every boundary used in the message and the trailing -- is mandatory for the closing boundary (close-delimiter).
	var pa, sa [128]byte
	n, _ := body.ReadAt(pa[:], 0)
	prefix := pa[:n]
	if !bytes.HasPrefix(prefix, []byte("--")) {
		if logger.V(1).Enabled() {
			logger.V(1).Info("line does not start with dash", "line", string(prefix))
		}
		return false
	}
	if prefixS != "" && !bytes.HasPrefix(prefix, []byte(prefixS+"\r\n")) {
		if logger.V(1).Enabled() {
			logger.V(1).Info("line does not start with the given", "prefix", prefixS, "line", string(prefix))
		}
		return false
	}
	i := bytes.Index(prefix, []byte("\r\n"))
	if i < 0 {
		if logger.V(1).Enabled() {
			logger.V(1).Info("line does not have EOL", "line", string(prefix))
		}
		return false
	}
	if prefix = prefix[:i]; body.Size() < 2*int64(len(prefix))+4 {
		if logger.V(1).Enabled() {
			logger.V(1).Info("body is not long enough", "body", body.Size())
		}
		return false
	}
	n, _ = body.ReadAt(sa[:], body.Size()-int64(len(prefix))-4)
	suffix := append(prefix, '-', '-')
	j := bytes.Index(sa[:n], suffix)
	if j < 0 {
		if logger.V(1).Enabled() {
			logger.V(1).Info("no end", "suffix", suffix, "line", string(sa[:n]))
		}
		return false
	}
	return true
}
