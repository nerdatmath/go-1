// Code generated by qtc from "content.xml.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line content.xml.qtpl:1
package ods

//line content.xml.qtpl:1
import (
	"encoding/xml"
	"strings"

//line content.xml.qtpl:2

//line content.xml.qtpl:5

	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line content.xml.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line content.xml.qtpl:5
func StreamXML(qw422016 *qt422016.Writer, s string) {
//line content.xml.qtpl:7
	var buf strings.Builder
	_ = xml.EscapeText(&buf, []byte(s))

//line content.xml.qtpl:10
	qw422016.N().S(buf.String())
//line content.xml.qtpl:11
}

//line content.xml.qtpl:11
func WriteXML(qq422016 qtio422016.Writer, s string) {
//line content.xml.qtpl:11
	qw422016 := qt422016.AcquireWriter(qq422016)
//line content.xml.qtpl:11
	StreamXML(qw422016, s)
//line content.xml.qtpl:11
	qt422016.ReleaseWriter(qw422016)
//line content.xml.qtpl:11
}

//line content.xml.qtpl:11
func XML(s string) string {
//line content.xml.qtpl:11
	qb422016 := qt422016.AcquireByteBuffer()
//line content.xml.qtpl:11
	WriteXML(qb422016, s)
//line content.xml.qtpl:11
	qs422016 := string(qb422016.B)
//line content.xml.qtpl:11
	qt422016.ReleaseByteBuffer(qb422016)
//line content.xml.qtpl:11
	return qs422016
//line content.xml.qtpl:11
}

//line content.xml.qtpl:14
func StreamBeginSheets(qw422016 *qt422016.Writer) {
//line content.xml.qtpl:14
	qw422016.N().S(`<?xml version="1.0" encoding="UTF-8"?>

<office:document-content xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0" xmlns:style="urn:oasis:names:tc:opendocument:xmlns:style:1.0" xmlns:text="urn:oasis:names:tc:opendocument:xmlns:text:1.0" xmlns:table="urn:oasis:names:tc:opendocument:xmlns:table:1.0" xmlns:draw="urn:oasis:names:tc:opendocument:xmlns:drawing:1.0" xmlns:fo="urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:meta="urn:oasis:names:tc:opendocument:xmlns:meta:1.0" xmlns:number="urn:oasis:names:tc:opendocument:xmlns:datastyle:1.0" xmlns:svg="urn:oasis:names:tc:opendocument:xmlns:svg-compatible:1.0" xmlns:chart="urn:oasis:names:tc:opendocument:xmlns:chart:1.0" xmlns:dr3d="urn:oasis:names:tc:opendocument:xmlns:dr3d:1.0" xmlns:config="urn:oasis:names:tc:opendocument:xmlns:config:1.0" xmlns:math="http://www.w3.org/1998/Math/MathML" xmlns:form="urn:oasis:names:tc:opendocument:xmlns:form:1.0" xmlns:script="urn:oasis:names:tc:opendocument:xmlns:script:1.0" xmlns:ooo="http://openoffice.org/2004/office" xmlns:ooow="http://openoffice.org/2004/writer" xmlns:oooc="http://openoffice.org/2004/calc" xmlns:tableooo="http://openoffice.org/2009/table" xmlns:of="urn:oasis:names:tc:opendocument:xmlns:of:1.2" xmlns:dom="http://www.w3.org/2001/xml-events" xmlns:xforms="http://www.w3.org/2002/xforms" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:gnm="http://www.gnumeric.org/odf-extension/1.0" xmlns:css3t="http://www.w3.org/TR/css3-text/" xmlns:loext="urn:org:documentfoundation:names:experimental:office:xmlns:loext:1.0" xmlns:calcext="urn:org:documentfoundation:names:experimental:calc:xmlns:calcext:1.0" office:version="1.2">
  <office:scripts/>
  <office:font-face-decls/>
  <office:automatic-styles>
    <style:style style:name="ta-0" style:family="table" style:master-page-name="ta-mp-0">
      <style:table-properties table:display="true" style:writing-mode="lr-tb"/>
    </style:style>
    <style:style style:name="AC-weight100" style:family="text">
      <style:text-properties fo:font-weight="100"/>
    </style:style>
    <style:style style:name="AC-weight200" style:family="text">
      <style:text-properties fo:font-weight="200"/>
    </style:style>
    <style:style style:name="AC-weight300" style:family="text">
      <style:text-properties fo:font-weight="300"/>
    </style:style>
    <style:style style:name="AC-weight400" style:family="text">
      <style:text-properties fo:font-weight="normal"/>
    </style:style>
    <style:style style:name="AC-weight500" style:family="text">
      <style:text-properties fo:font-weight="500"/>
    </style:style>
    <style:style style:name="AC-weight600" style:family="text">
      <style:text-properties fo:font-weight="600"/>
    </style:style>
    <style:style style:name="AC-weight700" style:family="text">
      <style:text-properties fo:font-weight="bold"/>
    </style:style>
    <style:style style:name="AC-weight800" style:family="text">
      <style:text-properties fo:font-weight="800"/>
    </style:style>
    <style:style style:name="AC-weight900" style:family="text">
      <style:text-properties fo:font-weight="900"/>
    </style:style>
    <style:style style:name="AC-weight1000" style:family="text">
      <style:text-properties fo:font-weight="900"/>
    </style:style>
    <style:style style:name="AC-italic" style:family="text">
      <style:text-properties fo:font-style="italic"/>
    </style:style>
    <style:style style:name="AC-roman" style:family="text">
      <style:text-properties fo:font-style="normal"/>
    </style:style>
    <style:style style:name="AC-subscript" style:family="text">
      <style:text-properties style:text-position="sub 83%"/>
    </style:style>
    <style:style style:name="AC-superscript" style:family="text">
      <style:text-properties style:text-position="super 83%"/>
    </style:style>
    <style:style style:name="AC-script" style:family="text">
      <style:text-properties style:text-position="0% 100%"/>
    </style:style>
    <style:style style:name="AC-strikethrough-solid" style:family="text">
      <style:text-properties style:text-line-through-type="single" style:text-line-through-style="solid"/>
    </style:style>
    <style:style style:name="AC-strikethrough-none" style:family="text">
      <style:text-properties style:text-line-through-type="none" style:text-line-through-style="none"/>
    </style:style>
    <style:style style:name="AC-underline-none" style:family="text">
      <style:text-properties style:text-underline-type="none" style:text-underline-style="none" style:text-underline-width="auto"/>
    </style:style>
    <style:style style:name="AC-underline-single" style:family="text">
      <style:text-properties style:text-underline-type="single" style:text-underline-style="solid" style:text-underline-width="auto"/>
    </style:style>
    <style:style style:name="AC-underline-double" style:family="text">
      <style:text-properties style:text-underline-type="double" style:text-underline-style="solid" style:text-underline-width="auto"/>
    </style:style>
    <style:style style:name="AC-underline-low" style:family="text">
      <style:text-properties style:text-underline-type="single" style:text-underline-style="solid" style:text-underline-width="bold"/>
    </style:style>
    <style:style style:name="AC-underline-error" style:family="text">
      <style:text-properties style:text-underline-type="single" style:text-underline-style="wave" style:text-underline-width="auto"/>
    </style:style>
    <style:style style:name="AROW-0" style:family="table-row">
      <style:table-row-properties style:row-height="12.75pt" style:use-optimal-row-height="true"/>
    </style:style>
    <style:style style:name="ACOL-0" style:family="table-column">
      <style:table-column-properties style:column-width="48pt" style:use-optimal-column-width="true"/>
    </style:style>
    <style:style style:name="ACE-0" style:family="table-cell" style:data-style-name="General">
      <style:table-cell-properties fo:background-color="transparent" fo:border-top="0.000cm none #c7c7c7" fo:border-bottom="0.000cm none #c7c7c7" fo:border-left="0.000cm none #c7c7c7" fo:border-right="0.000cm none #c7c7c7" style:diagonal-bl-tr="0.000cm none #c7c7c7" style:diagonal-tl-br="0.000cm none #c7c7c7" style:vertical-align="bottom" fo:wrap-option="no-wrap" style:shrink-to-fit="false" style:writing-mode="page" style:glyph-orientation-vertical="auto" style:cell-protect="protected" style:rotation-align="none" style:rotation-angle="0" style:print-content="true" style:decimal-places="13" style:text-align-source="fix" style:repeat-content="false"/>
      <style:paragraph-properties style:writing-mode-automatic="true" fo:text-align="right" fo:margin-left="0pt"/>
      <style:text-properties text:display="true" fo:font-weight="bold" fo:font-style="normal" style:text-line-through-type="none" style:text-line-through-style="none" style:text-underline-type="none" style:text-underline-style="none" style:text-underline-width="auto" style:text-underline-color="font-color" style:text-underline-mode="continuous" style:text-position="0% 100%" fo:font-size="10pt" fo:color="#000000" fo:font-family="Verdana"/>
    </style:style>
    <style:style style:name="ACE-1" style:family="table-cell" style:data-style-name="General">
      <style:table-cell-properties fo:background-color="transparent" fo:border-top="0.000cm none #c7c7c7" fo:border-bottom="0.000cm none #c7c7c7" fo:border-left="0.000cm none #c7c7c7" fo:border-right="0.000cm none #c7c7c7" style:diagonal-bl-tr="0.000cm none #c7c7c7" style:diagonal-tl-br="0.000cm none #c7c7c7" style:vertical-align="bottom" fo:wrap-option="no-wrap" style:shrink-to-fit="false" style:writing-mode="page" style:glyph-orientation-vertical="auto" style:cell-protect="protected" style:rotation-align="none" style:rotation-angle="0" style:print-content="true" style:decimal-places="13" style:text-align-source="fix" style:repeat-content="false"/>
      <style:paragraph-properties style:writing-mode-automatic="true" fo:text-align="right" fo:margin-left="0pt"/>
      <style:text-properties text:display="true" fo:font-weight="normal" fo:font-style="normal" style:text-line-through-type="none" style:text-line-through-style="none" style:text-underline-type="none" style:text-underline-style="none" style:text-underline-width="auto" style:text-underline-color="font-color" style:text-underline-mode="continuous" style:text-position="0% 100%" fo:font-size="10pt" fo:color="#000000" fo:font-family="Verdana"/>
    </style:style>
    <style:style style:name="ACE-2" style:family="table-cell" style:data-style-name="General">
      <style:table-cell-properties fo:background-color="transparent" fo:border-top="0.000cm none #c7c7c7" fo:border-bottom="0.000cm none #c7c7c7" fo:border-left="0.000cm none #c7c7c7" fo:border-right="0.000cm none #c7c7c7" style:diagonal-bl-tr="0.000cm none #c7c7c7" style:diagonal-tl-br="0.000cm none #c7c7c7" style:vertical-align="bottom" fo:wrap-option="no-wrap" style:shrink-to-fit="false" style:writing-mode="page" style:glyph-orientation-vertical="auto" style:cell-protect="protected" style:rotation-align="none" style:rotation-angle="0" style:print-content="true" style:decimal-places="13" style:text-align-source="value-type" style:repeat-content="false"/>
      <style:paragraph-properties style:writing-mode-automatic="true" fo:margin-left="0pt"/>
      <style:text-properties text:display="true" fo:font-weight="normal" fo:font-style="normal" style:text-line-through-type="none" style:text-line-through-style="none" style:text-underline-type="none" style:text-underline-style="none" style:text-underline-width="auto" style:text-underline-color="font-color" style:text-underline-mode="continuous" style:text-position="0% 100%" fo:font-size="10pt" fo:color="#000000" fo:font-family="Sans"/>
    </style:style>
    <style:style style:name="ACE-3" style:family="table-cell" style:data-style-name="ND-0">
      <style:table-cell-properties fo:background-color="transparent" fo:border-top="0.000cm none #c7c7c7" fo:border-bottom="0.000cm none #c7c7c7" fo:border-left="0.000cm none #c7c7c7" fo:border-right="0.000cm none #c7c7c7" style:diagonal-bl-tr="0.000cm none #c7c7c7" style:diagonal-tl-br="0.000cm none #c7c7c7" style:vertical-align="bottom" fo:wrap-option="no-wrap" style:shrink-to-fit="false" style:writing-mode="page" style:glyph-orientation-vertical="auto" style:cell-protect="protected" style:rotation-align="none" style:rotation-angle="0" style:print-content="true" style:decimal-places="13" style:text-align-source="fix" style:repeat-content="false"/>
      <style:paragraph-properties style:writing-mode-automatic="true" fo:text-align="right" fo:margin-left="0pt"/>
      <style:text-properties text:display="true" fo:font-weight="normal" fo:font-style="normal" style:text-line-through-type="none" style:text-line-through-style="none" style:text-underline-type="none" style:text-underline-style="none" style:text-underline-width="auto" style:text-underline-color="font-color" style:text-underline-mode="continuous" style:text-position="0% 100%" fo:font-size="10pt" fo:color="#000000" fo:font-family="Verdana"/>
    </style:style>
    <style:style style:name="ACE-4" style:family="table-cell" style:data-style-name="ND-1">
      <style:table-cell-properties fo:background-color="transparent" fo:border-top="0.000cm none #c7c7c7" fo:border-bottom="0.000cm none #c7c7c7" fo:border-left="0.000cm none #c7c7c7" fo:border-right="0.000cm none #c7c7c7" style:diagonal-bl-tr="0.000cm none #c7c7c7" style:diagonal-tl-br="0.000cm none #c7c7c7" style:vertical-align="bottom" fo:wrap-option="no-wrap" style:shrink-to-fit="false" style:writing-mode="page" style:glyph-orientation-vertical="auto" style:cell-protect="protected" style:rotation-align="none" style:rotation-angle="0" style:print-content="true" style:decimal-places="13" style:text-align-source="fix" style:repeat-content="false"/>
      <style:paragraph-properties style:writing-mode-automatic="true" fo:text-align="right" fo:margin-left="0pt"/>
      <style:text-properties text:display="true" fo:font-weight="bold" fo:font-style="normal" style:text-line-through-type="none" style:text-line-through-style="none" style:text-underline-type="none" style:text-underline-style="none" style:text-underline-width="auto" style:text-underline-color="font-color" style:text-underline-mode="continuous" style:text-position="0% 100%" fo:font-size="10pt" fo:color="#000000" fo:font-family="Verdana"/>
    </style:style>
    <style:style style:name="ACE-5" style:family="table-cell" style:data-style-name="ND-1">
      <style:table-cell-properties fo:background-color="transparent" fo:border-top="0.000cm none #c7c7c7" fo:border-bottom="0.000cm none #c7c7c7" fo:border-left="0.000cm none #c7c7c7" fo:border-right="0.000cm none #c7c7c7" style:diagonal-bl-tr="0.000cm none #c7c7c7" style:diagonal-tl-br="0.000cm none #c7c7c7" style:vertical-align="bottom" fo:wrap-option="no-wrap" style:shrink-to-fit="false" style:writing-mode="page" style:glyph-orientation-vertical="auto" style:cell-protect="protected" style:rotation-align="none" style:rotation-angle="0" style:print-content="true" style:decimal-places="13" style:text-align-source="fix" style:repeat-content="false"/>
      <style:paragraph-properties style:writing-mode-automatic="true" fo:text-align="right" fo:margin-left="0pt"/>
      <style:text-properties text:display="true" fo:font-weight="normal" fo:font-style="normal" style:text-line-through-type="none" style:text-line-through-style="none" style:text-underline-type="none" style:text-underline-style="none" style:text-underline-width="auto" style:text-underline-color="font-color" style:text-underline-mode="continuous" style:text-position="0% 100%" fo:font-size="10pt" fo:color="#000000" fo:font-family="Verdana"/>
    </style:style>
    <style:style style:name="ACE-6" style:family="table-cell" style:data-style-name="ND-0">
      <style:table-cell-properties fo:background-color="transparent" fo:border-top="0.000cm none #c7c7c7" fo:border-bottom="0.000cm none #c7c7c7" fo:border-left="0.000cm none #c7c7c7" fo:border-right="0.000cm none #c7c7c7" style:diagonal-bl-tr="0.000cm none #c7c7c7" style:diagonal-tl-br="0.000cm none #c7c7c7" style:vertical-align="bottom" fo:wrap-option="no-wrap" style:shrink-to-fit="false" style:writing-mode="page" style:glyph-orientation-vertical="auto" style:cell-protect="protected" style:rotation-align="none" style:rotation-angle="0" style:print-content="true" style:decimal-places="13" style:text-align-source="fix" style:repeat-content="false"/>
      <style:paragraph-properties style:writing-mode-automatic="true" fo:text-align="right" fo:margin-left="0pt"/>
      <style:text-properties text:display="true" fo:font-weight="bold" fo:font-style="normal" style:text-line-through-type="none" style:text-line-through-style="none" style:text-underline-type="none" style:text-underline-style="none" style:text-underline-width="auto" style:text-underline-color="font-color" style:text-underline-mode="continuous" style:text-position="0% 100%" fo:font-size="10pt" fo:color="#000000" fo:font-family="Verdana"/>
    </style:style>
    <style:style style:name="ACOL-1" style:family="table-column"/>
    <style:style style:name="AROW-1" style:family="table-row">
      <style:table-row-properties style:row-height="13.5pt" style:use-optimal-row-height="true"/>
    </style:style>
    <style:style style:name="AROW-2" style:family="table-row"/>
  </office:automatic-styles>
  <office:body>
    <office:spreadsheet>
      <table:calculation-settings table:null-year="1930" table:automatic-find-labels="false" table:case-sensitive="false" table:precision-as-shown="false" table:search-criteria-must-apply-to-whole-cell="true" table:use-regular-expressions="false" table:use-wildcards="false">
        <table:null-date table:date-value="1899-12-30" table:value-type="date"/>
        <table:iteration table:maximum-difference="0.001" table:status="enable" table:steps="100"/>
      </table:calculation-settings>
`)
//line content.xml.qtpl:142
}

//line content.xml.qtpl:142
func WriteBeginSheets(qq422016 qtio422016.Writer) {
//line content.xml.qtpl:142
	qw422016 := qt422016.AcquireWriter(qq422016)
//line content.xml.qtpl:142
	StreamBeginSheets(qw422016)
//line content.xml.qtpl:142
	qt422016.ReleaseWriter(qw422016)
//line content.xml.qtpl:142
}

//line content.xml.qtpl:142
func BeginSheets() string {
//line content.xml.qtpl:142
	qb422016 := qt422016.AcquireByteBuffer()
//line content.xml.qtpl:142
	WriteBeginSheets(qb422016)
//line content.xml.qtpl:142
	qs422016 := string(qb422016.B)
//line content.xml.qtpl:142
	qt422016.ReleaseByteBuffer(qb422016)
//line content.xml.qtpl:142
	return qs422016
//line content.xml.qtpl:142
}

//line content.xml.qtpl:144
func (t Table) StreamBegin(qw422016 *qt422016.Writer) {
//line content.xml.qtpl:144
	qw422016.N().S(`<table:table table:name="`)
//line content.xml.qtpl:144
	StreamXML(qw422016, t.Name)
//line content.xml.qtpl:144
	qw422016.N().S(`" table:style-name="ta-0" table:print="true">
		`)
//line content.xml.qtpl:145
	if t.Style != "" {
//line content.xml.qtpl:145
		qw422016.N().S(`<table:table-column table:style-name="`)
//line content.xml.qtpl:145
		StreamXML(qw422016, t.Style)
//line content.xml.qtpl:145
		qw422016.N().S(`" table:number-columns-repeated="`)
//line content.xml.qtpl:145
		qw422016.N().D(t.ColCount)
//line content.xml.qtpl:145
		qw422016.N().S(`"/>`)
//line content.xml.qtpl:145
	}
//line content.xml.qtpl:145
	qw422016.N().S(`
		`)
//line content.xml.qtpl:146
	t.Heading.StreamXML(qw422016)
//line content.xml.qtpl:146
	qw422016.N().S(`
`)
//line content.xml.qtpl:147
}

//line content.xml.qtpl:147
func (t Table) WriteBegin(qq422016 qtio422016.Writer) {
//line content.xml.qtpl:147
	qw422016 := qt422016.AcquireWriter(qq422016)
//line content.xml.qtpl:147
	t.StreamBegin(qw422016)
//line content.xml.qtpl:147
	qt422016.ReleaseWriter(qw422016)
//line content.xml.qtpl:147
}

//line content.xml.qtpl:147
func (t Table) Begin() string {
//line content.xml.qtpl:147
	qb422016 := qt422016.AcquireByteBuffer()
//line content.xml.qtpl:147
	t.WriteBegin(qb422016)
//line content.xml.qtpl:147
	qs422016 := string(qb422016.B)
//line content.xml.qtpl:147
	qt422016.ReleaseByteBuffer(qb422016)
//line content.xml.qtpl:147
	return qs422016
//line content.xml.qtpl:147
}

//line content.xml.qtpl:149
func (row Row) StreamXML(qw422016 *qt422016.Writer) {
//line content.xml.qtpl:150
	if len(row.Cells) != 0 {
//line content.xml.qtpl:150
		qw422016.N().S(`<table:table-row table:style-name="`)
//line content.xml.qtpl:150
		StreamXML(qw422016, row.Style)
//line content.xml.qtpl:150
		qw422016.N().S(`">`)
//line content.xml.qtpl:151
		for _, cell := range row.Cells {
//line content.xml.qtpl:151
			cell.StreamXML(qw422016)
//line content.xml.qtpl:152
		}
//line content.xml.qtpl:152
		qw422016.N().S(`</table:table-row>`)
//line content.xml.qtpl:153
	}
//line content.xml.qtpl:153
	qw422016.N().S(`
`)
//line content.xml.qtpl:154
}

//line content.xml.qtpl:154
func (row Row) WriteXML(qq422016 qtio422016.Writer) {
//line content.xml.qtpl:154
	qw422016 := qt422016.AcquireWriter(qq422016)
//line content.xml.qtpl:154
	row.StreamXML(qw422016)
//line content.xml.qtpl:154
	qt422016.ReleaseWriter(qw422016)
//line content.xml.qtpl:154
}

//line content.xml.qtpl:154
func (row Row) XML() string {
//line content.xml.qtpl:154
	qb422016 := qt422016.AcquireByteBuffer()
//line content.xml.qtpl:154
	row.WriteXML(qb422016)
//line content.xml.qtpl:154
	qs422016 := string(qb422016.B)
//line content.xml.qtpl:154
	qt422016.ReleaseByteBuffer(qb422016)
//line content.xml.qtpl:154
	return qs422016
//line content.xml.qtpl:154
}

//line content.xml.qtpl:156
func (cell Cell) StreamXML(qw422016 *qt422016.Writer) {
//line content.xml.qtpl:156
	qw422016.N().S(`<table:table-cell table:style-name="`)
//line content.xml.qtpl:156
	StreamXML(qw422016, cell.Style)
//line content.xml.qtpl:156
	qw422016.N().S(`" office:value-type="`)
//line content.xml.qtpl:156
	qw422016.N().S(cell.Type.String())
//line content.xml.qtpl:156
	qw422016.N().S(`"`)
//line content.xml.qtpl:157
	if cell.Type == FloatType {
//line content.xml.qtpl:157
		qw422016.N().S(` office:value="`)
//line content.xml.qtpl:157
		StreamXML(qw422016, cell.Value)
//line content.xml.qtpl:157
		qw422016.N().S(`"`)
//line content.xml.qtpl:158
	} else if cell.Type == DateType {
//line content.xml.qtpl:158
		qw422016.N().S(` office:date-value="`)
//line content.xml.qtpl:158
		StreamXML(qw422016, cell.Value)
//line content.xml.qtpl:158
		qw422016.N().S(`"`)
//line content.xml.qtpl:159
	}
//line content.xml.qtpl:159
	qw422016.N().S(`><text:p>`)
//line content.xml.qtpl:159
	StreamXML(qw422016, cell.Value)
//line content.xml.qtpl:159
	qw422016.N().S(`</text:p></table:table-cell>`)
//line content.xml.qtpl:159
}

//line content.xml.qtpl:159
func (cell Cell) WriteXML(qq422016 qtio422016.Writer) {
//line content.xml.qtpl:159
	qw422016 := qt422016.AcquireWriter(qq422016)
//line content.xml.qtpl:159
	cell.StreamXML(qw422016)
//line content.xml.qtpl:159
	qt422016.ReleaseWriter(qw422016)
//line content.xml.qtpl:159
}

//line content.xml.qtpl:159
func (cell Cell) XML() string {
//line content.xml.qtpl:159
	qb422016 := qt422016.AcquireByteBuffer()
//line content.xml.qtpl:159
	cell.WriteXML(qb422016)
//line content.xml.qtpl:159
	qs422016 := string(qb422016.B)
//line content.xml.qtpl:159
	qt422016.ReleaseByteBuffer(qb422016)
//line content.xml.qtpl:159
	return qs422016
//line content.xml.qtpl:159
}

//line content.xml.qtpl:161
func StreamEndTable(qw422016 *qt422016.Writer) {
//line content.xml.qtpl:161
	qw422016.N().S(`
      </table:table>
`)
//line content.xml.qtpl:163
}

//line content.xml.qtpl:163
func WriteEndTable(qq422016 qtio422016.Writer) {
//line content.xml.qtpl:163
	qw422016 := qt422016.AcquireWriter(qq422016)
//line content.xml.qtpl:163
	StreamEndTable(qw422016)
//line content.xml.qtpl:163
	qt422016.ReleaseWriter(qw422016)
//line content.xml.qtpl:163
}

//line content.xml.qtpl:163
func EndTable() string {
//line content.xml.qtpl:163
	qb422016 := qt422016.AcquireByteBuffer()
//line content.xml.qtpl:163
	WriteEndTable(qb422016)
//line content.xml.qtpl:163
	qs422016 := string(qb422016.B)
//line content.xml.qtpl:163
	qt422016.ReleaseByteBuffer(qb422016)
//line content.xml.qtpl:163
	return qs422016
//line content.xml.qtpl:163
}

//line content.xml.qtpl:165
func StreamEndSheets(qw422016 *qt422016.Writer) {
//line content.xml.qtpl:165
	qw422016.N().S(`
    </office:spreadsheet>
  </office:body>
</office:document-content>
`)
//line content.xml.qtpl:169
}

//line content.xml.qtpl:169
func WriteEndSheets(qq422016 qtio422016.Writer) {
//line content.xml.qtpl:169
	qw422016 := qt422016.AcquireWriter(qq422016)
//line content.xml.qtpl:169
	StreamEndSheets(qw422016)
//line content.xml.qtpl:169
	qt422016.ReleaseWriter(qw422016)
//line content.xml.qtpl:169
}

//line content.xml.qtpl:169
func EndSheets() string {
//line content.xml.qtpl:169
	qb422016 := qt422016.AcquireByteBuffer()
//line content.xml.qtpl:169
	WriteEndSheets(qb422016)
//line content.xml.qtpl:169
	qs422016 := string(qb422016.B)
//line content.xml.qtpl:169
	qt422016.ReleaseByteBuffer(qb422016)
//line content.xml.qtpl:169
	return qs422016
//line content.xml.qtpl:169
}
