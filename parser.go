package linkparser

import (
	"golang.org/x/net/html"
	"io"
)

func LinkParse(r io.Reader) string {
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	return getText(doc)
}

func getText(n *html.Node) string {
	var text string
	if n.Type == html.TextNode {
		text = n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += getText(c)
	}
	return text
}
