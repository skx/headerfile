package headerfile

import "testing"

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
