package headerfile

import (
	"strings"
	"testing"
)

// Test reading a file that doesn't exist
func TestMissing(t *testing.T) {

	h := New("_test/missing/file/goes/here")

	_, err := h.Headers()
	if err == nil {
		t.Errorf("expected an error reading a missing file; saw none. [1/2]")
	}

	b := New("_test/missing/file/goes/here")
	_, err2 := b.Body()
	if err2 == nil {
		t.Errorf("expected an error reading a missing file; saw none. [2/2]")
	}

}

// Test basic usage
func TestBasicBlogHeaderFirst(t *testing.T) {

	// We test two input-files with different separators.
	files := []string{"_test/blog.txt", "_test/blog2.txt"}

	for _, file := range files {
		h := New(file)

		headers, err := h.Headers()
		if err != nil {
			t.Errorf("error reading headers of blog-post: %s\n", err.Error())
		}

		body := ""
		body, err = h.Body()
		if err != nil {
			t.Errorf("error reading body of blog-post: %s\n", err.Error())
		}

		//
		// Ok we should have a body, and some headers.
		//
		for header, value := range headers {

			switch header {
			case "subject", "date", "tags":
				// nop
			default:
				t.Errorf("unknown header '%s' (value:%s) in file", header, value)
			}
		}

		//
		// Expected body
		//
		if body != "This is my blog post ..\n\n" {
			t.Errorf("body did not match expectations: '%s'\n", body)
		}
	}
}

// Test basic usage
func TestBasicBlogBodyFirst(t *testing.T) {

	// We test two input-files with different separators.
	files := []string{"_test/blog.txt", "_test/blog2.txt"}

	for _, file := range files {
		h := New(file)

		// Parse body first
		body, err := h.Body()
		if err != nil {
			t.Errorf("error reading body of blog-post: %s\n", err.Error())
		}

		// Now parse headers
		headers, err2 := h.Headers()
		if err2 != nil {
			t.Errorf("error reading headers of blog-post: %s\n", err.Error())
		}

		//
		// Ok we should have a body, and some headers.
		//
		for header, value := range headers {

			switch header {
			case "subject", "date", "tags":
				// nop
			default:
				t.Errorf("unknown header '%s' (value:%s) in file", header, value)
			}
		}

		//
		// Expected body
		//
		if body != "This is my blog post ..\n\n" {
			t.Errorf("body did not match expectations: '%s'\n", body)
		}
	}
}

// Test that a malformed header-entry is found
func TestMalformedHeader(t *testing.T) {

	// Test in the header-fetching
	h := New("_test/malformed.header.txt")

	// Fetch the headers
	_, err := h.Headers()
	if err == nil {
		t.Errorf("expected an error, saw none")
	}

	// Test that the error-message matches what we expect
	if !strings.Contains(err.Error(), "malformed header") {
		t.Errorf("error message didn't match what we expected")
	}

	// Test in the body-fetching
	b := New("_test/malformed.header.txt")

	// Fetch the body
	_, err2 := b.Body()
	if err2 == nil {
		t.Errorf("expected an error, saw none")
	}

	// Test that the error-message matches what we expect
	if !strings.Contains(err2.Error(), "malformed header") &&
		!strings.Contains(err2.Error(), "#") {
		t.Errorf("error message didn't match what we expected")
	}
}
