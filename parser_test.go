package linkparser

import (
	"golang.org/x/net/html"
	"reflect"
	"strings"
	"testing"
)

func TestLinkParse(t *testing.T) {
	cases := []struct {
		Description   string
		Html          string
		ExpectedParse []Link
	}{
		{
			Description:   "parse link",
			Html:          "<a href=\"test.url\">test</a>",
			ExpectedParse: []Link{{Href: "test.url", Text: "test"}},
		},
		{
			Description:   "parse multiple links",
			Html:          "<a href=\"test.link\">link1</a><a>link2</a>",
			ExpectedParse: []Link{
				{Href: "test.link", Text: "link1"},
				{Text: "link2"}},
		},
		{
			Description: "return nil if no anchor tags",
			Html: "<p>test</p><h1>moreinfo</h1><span>1</span>",
			ExpectedParse: nil,
		},
		{
			Description: "only parse anchor tags",
			Html: "<p>test</p><a href=\"test.url\">test<span>1</span></a><span>1</span>",
			ExpectedParse: []Link{{Href: "test.url", Text: "test1"}},
		},
	}
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			r := strings.NewReader(test.Html)
			got := HtmlLinkParse(r)
			if !reflect.DeepEqual(got, test.ExpectedParse) {
				t.Errorf("got %v, want %v", got, test.ExpectedParse)
			}
		})
	}
}

func TestGetText(t *testing.T) {
	cases := []struct {
		Description  string
		Html         string
		ExpectedText string
	}{
		{
			Description:  "get text from a single tag",
			Html:         "<p>link text</p>",
			ExpectedText: "link text",
		},
		{
			Description:  "get text from multiple sibling tags",
			Html:         "<p>text </p><span>text2</span>",
			ExpectedText: "text text2",
		},
		{
			Description:  "get text from tested tags",
			Html:         "<p>text <span>text2 <bold>text3</bold> text4</span></p>",
			ExpectedText: "text text2 text3 text4",
		},
	}
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			r := strings.NewReader(test.Html)
			doc, err := html.Parse(r)
			if err != nil {
				t.Fatalf("didn't expected an error and got: %v", err)
			}
			got := getText(doc)
			if got != test.ExpectedText {
				t.Errorf("got %q, want %q", got, test.ExpectedText)
			}
		})
	}
}
