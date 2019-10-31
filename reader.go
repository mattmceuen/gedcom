/*
Copyright © 2019 Matt McEuen

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
package gedcom

import (
	"io"
)

// TODO
type Reader struct {
	r io.Reader
}

// NewReader returns a new Reader that reads from r
func NewReader(r io.Reader) *Reader {
	var greader Reader
	greader.r = r
	return &greader
}

// Read the contents of the GEDCOM file into a Genealogy.
// As not all content is currently readable, any skipped data
// will be returned in the `warn` slice.
func (r *Reader) Read() (g *Genealogy, warn []string, err error) {
	return nil, []string{}, nil
}
