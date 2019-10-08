# headerfile

The header file reader is a simple package which allows you to process files containing key-value headers.


## Use Case

Consider an input-file for a blog-post, it might look like this:

```
Subject: This is my post
Date: 10th March 1980
Tags: foo, bar, baz

This is my blog post ..
```

To process this file you wish to have the header-values available, as
well as the content.  This is what this package allows.


## Usage

Usage is as simple as you would expect:


     file := headerfile.New( "/path/to/file" )

     headers, err := file.Headers()

     body, err := file.Body()


Steve
--
