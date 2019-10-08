package headerfile

import (
	"strings"
	"testing"
)

// Test basic usage
func TestBasicBlog(t *testing.T) {

	h := New("_test/blog.txt")

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
