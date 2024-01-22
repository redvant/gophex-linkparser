package linkparser

import (
	"golang.org/x/net/html"
	"io"
)

type Link struct {
	Href string
	Text string
}

func HtmlLinkParse(r io.Reader) []Link {
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	return linkParse(doc)
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
		}
	}
	return ""
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
