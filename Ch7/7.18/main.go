// This file sorts XML file into parent/child elements and attributes
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// Node is either a CharData or *Element
type Node interface{}

// CharData represents XML character data (raw text), in which XML
// escape sequences have been replaced by the characters they represent.
type CharData string

// Element represents an XML node with attributes and child nodes.
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {

	dec := xml.NewDecoder(os.Stdin)

	// add a root node we can return an Element containing all the attr/child info
	root := &Element{Type: xml.Name{Local: "root"}}
	stack := []*Element{root}

	for {

		tok, err := dec.Token()

		if err == io.EOF {
			print(root)
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "error constructing: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			child := &Element{
				Type: tok.Name,
				Attr: tok.Attr,
			}
			parent := stack[len(stack)-1]
			parent.Children = append(parent.Children, child)
			stack = append(stack, child)
		case xml.CharData:
			parent := stack[len(stack)-1]
			parent.Children = append(parent.Children, CharData(tok))

		case xml.EndElement:
			stack = stack[:len(stack)-1]
		}
	}
}

// print starts from the specified node and prints nodes and traverses to child nodes
func print(node Node) {
	switch n := node.(type) {
	case *Element:
		fmt.Printf("element: %s has %d child(ren):\n", n.Type.Local, len(n.Children))
		for _, child := range n.Children {
			print(child)
		}
	case CharData:
		data := strings.Trim(strings.Trim(string([]byte(n)), " "), "\n")
		if len(data) > 1 {
			fmt.Printf("\tdata: %s\n", data)
		}
		return
	}
}
