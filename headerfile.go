// Package headerfile allows you to read a file which contains a
// small header of "key:value" lines.
//
// These kinds of files are very common when processing simple blogs,
// and similar files.
//
package headerfile

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// HeaderFile is the structure that encapsulates our object.
type HeaderFile struct {

	// File is the filename the user wished us to read.
	File string

	// Headers is a map of the key/values we found in the header,
	// if any were present.
	headers map[string]string

	// Body is the body of the text, after the headers were read.
	body string

	// Parsed records whether we've parsed the file already
	parsed bool
}

// New creates a new object.
func New(filename string) *HeaderFile {
	return &HeaderFile{File: filename,
		body:    "",
		parsed:  false,
		headers: make(map[string]string)}
}

// parse parses the input file we were constructed with.
//
// This is called the first time a header/body is requested
// from the package.  We don't re-read the file on-demand.
func (h *HeaderFile) parse() error {

	// Read the file
	bytes, err := ioutil.ReadFile(h.File)
	if err != nil {
		return err
	}

	// Compile our regular expression
	headerRegex := regexp.MustCompile("^([^:=]+)[:=](.*)$")

	// We're in the header by default
	header := true

	// Split by newlines - because we want to process the header
	// specially, splitting it into fields.
	lines := strings.Split(string(bytes), "\n")

	// Process each line
	for _, line := range lines {

		// If we're in the header ..
		if header {

			// Empty line?  Then header-time is over now.
			if len(line) < 1 {
				header = false
				continue
			}

			// Find the key + value which we expect to see
			// in the header.
			header := headerRegex.FindStringSubmatch(line)

			// If we did then we're good.
			if len(header) == 3 {
				// Save the key + value
				key := header[1]
				val := header[2]

				// Normalize keys & values.
				key = strings.ToLower(key)
				key = strings.TrimSpace(key)
				val = strings.TrimSpace(val)

				h.headers[key] = val
			} else {
				return fmt.Errorf("malformed header '%s' in %s", line, h.File)
			}
		} else {
			h.body += line + "\n"
		}
	}

	return nil
}

// Headers returns the headers associated with the file.
//
// The map will have all the header-names downcased, and the values
// will be stripped of any leading/trailing whitespace.
func (h *HeaderFile) Headers() (map[string]string, error) {

	if !h.parsed {
		err := h.parse()
		if err != nil {
			return nil, err
		}

		h.parsed = true
	}

	return h.headers, nil
}

// Body returns the body of the file, which is the region after
// the header.
//
// The body will contain an extra trailing newline.
func (h *HeaderFile) Body() (string, error) {
	if !h.parsed {
		err := h.parse()
		if err != nil {
			return "", err
		}

		h.parsed = true
	}

	return h.body, nil
}
