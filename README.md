[![GoDoc](https://godoc.org/github.com/skx/headerfile?status.svg)](http://godoc.org/github.com/skx/headerfile)
[![Go Report Card](https://goreportcard.com/badge/github.com/skx/headerfile)](https://goreportcard.com/report/github.com/skx/headerfile)
[![license](https://img.shields.io/github/license/skx/headerfile.svg)](https://github.com/skx/headerfile/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/skx/headerfile.svg)](https://github.com/skx/headerfile/releases/latest)


# headerfile

This is a simple package which allows you to process files which consist of a key-value header, and then content.


## Use Case

The initial-use case was for a simple blog-compiler, which consumes a series of files containing posts.  As you might expect each post has some associated meta-data, such as a title, a set of tags, etc.

This library allows you to read the two parts of this file seperately, and cleanly:

```
Subject: This is my post
Date: 10th March 1980
Tags: foo, bar, baz

This is my blog post ..
```

Once parsed you can receive :

* The body of the post.
* A map containing the (string) keys and values present in the header.
  * The header-values may be separated by either `:` or `=`.


## Github Setup

This repository is configured to run tests upon every commit, and when
pull-requests are created/updated.  The testing is carried out via
[.github/run-tests.sh](.github/run-tests.sh) which is used by the
[github-action-tester](https://github.com/skx/github-action-tester) action.

If test-coverage drops beneath 100% this is a bug.  The package is simple
enough that this should not be an undue burdon.

Steve
--
