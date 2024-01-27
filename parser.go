package link

import (
	"golang.org/x/net/html"
	"io"
	"strings"
)

// Link represents a link (<a href="...">) in an HTML
// document.
type Link struct {
	Href string
	Text string
}
// test
// test 2
// Parse receives an HTML document and returns a slice
// of Links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return linkParse(doc), nil
}

func linkParse(n *html.Node) []Link {
	var links []Link
	if n.Type == html.ElementNode && n.Data == "a" {
		l := Link{}
		l.Href = getHref(n)
		l.Text = getText(n)
		links = append(links, l)
		return links
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, linkParse(c)...)
	}
	return links
}

func getHref(n *html.Node) string {
	for _, at := range n.Attr {
		if at.Key == "href" {
			return at.Val
			break
		}
	}
	return ""
}

func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += getText(c)
	}
	return strings.Join(strings.Fields(text), " ")
}
