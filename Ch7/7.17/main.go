// This file parses XML file and filters the XML for provided name and attributes
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func main() {

	dec := xml.NewDecoder(os.Stdin)

	var match bool
	var nests []xml.StartElement

	for {

		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect:  %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			if match {
				nests = append(nests, tok)
				continue
			}
			match = matchElement(tok, os.Args[1:])
		case xml.EndElement:
			if len(nests) > 0 {
				nests = nests[:len(nests)-1]
			} else if len(nests) == 0 {
				match = false
			}
		case xml.CharData:
			if match {
				fmt.Println(string(tok))
			}
		}
	}
}

// matchElements parses element for provided type and attrs
func matchElement(x xml.StartElement, y []string) bool {
	var found bool
	if x.Name.Local == y[0] {
		y = y[1:]
		for _, a := range x.Attr {
			if a.Name.Local != y[0] || a.Value != y[1] {
				found = false
				break
			} else {
				found = true
				y = y[2:]
			}
		}
	}
	return found
}

// <div class="constraint">
// 		<p class="prefix">
// 			<a name="vc-roottype" id="vc-roottype"/>
// 			<b>
// 				Validity constraint: Root Element Type
// 			</b>
// 		</p>
// 		<p>The
// 			<a href="#NT-Name">Name</a>
// 			in the document type declaration
// 			<em class="rfc2119" title="Keyword in RFC 2119 context">
// 				MUST
// 			</em>
// 			match the element type of the
// 			<a title="Root Element" href="#dt-root">
// 				root element
// 			</a>
// 			.
// 		</p>
// </div>
//
// ------>>>>>>
//
// Validity constraint: Root Element Type
// The
// Name
//
// in the document type declaration
// MUST
//  match the element type of the
// root element
// .
