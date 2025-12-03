package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestOutline(t *testing.T) {
	input := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Hello World Page</title>
	</head>
	<body>
		<h1>Hello World!</h1>
		<p>This is a simple "Hello World" example in HTML.</p>
	</body>
	</html>
	`

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Error(err)
	}

	var out io.Writer = os.Stdout
	out = new(bytes.Buffer)
	forEachNode(doc, startElement, endElement)

	_, err = html.Parse(out.(*bytes.Buffer))
	if err != nil {
		t.Error(err)
	}
}
