//go:build windows
// +build windows

/*
Copied from camlistore.org/pkg/blobserver/localdisk/receive_posix.go

Copyright 2011 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package temp

// LinkOrCopy links src to dst if possible; fails back to copying
func LinkOrCopy(src, dst string) error {
	return copyFile(src, dst)
}
