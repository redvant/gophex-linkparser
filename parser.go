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

	return getString(doc)
}

func getString(n *html.Node) string {
	var text string
	if n.Type == html.TextNode {
		text = n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += getString(c)
	}
	return text
}
