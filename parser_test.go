package linkparser

import (
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestHtmlLinkParse(t *testing.T) {
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
			Description: "parse multiple links",
			Html:        "<a href=\"test.link\">link1</a><a>link2</a>",
			ExpectedParse: []Link{
				{Href: "test.link", Text: "link1"},
				{Text: "link2"}},
		},
		{
			Description:   "return nil if there isn't anchor tags",
			Html:          "<p>test</p><h1>moreinfo</h1><span>1</span>",
			ExpectedParse: nil,
		},
		{
			Description:   "get text from tested tags",
			Html:          "<a><p>text <span>text2 <bold>text3</bold> text4</span></p></a>",
			ExpectedParse: []Link{{Text: "text text2 text3 text4"}},
		},
		{
			Description:   "don't get text inside comments",
			Html:          "<a>text to get <!-- commented text --></a>",
			ExpectedParse: []Link{{Text: "text to get"}},
		},
	}
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			r := strings.NewReader(test.Html)
			assertParse(t, r, test.ExpectedParse)
		})
	}
}

func TestParseExampleFiles(t *testing.T) {
	cases := []struct {
		Description   string
		Path          string
		ExpectedParse []Link
	}{
		{
			Description:   "test example 1",
			Path:          "testingexamples/ex1.html",
			ExpectedParse: []Link{{Href: "/other-page", Text: "A link to another page"}},
		},
		{
			Description: "test example 2",
			Path:        "testingexamples/ex2.html",
			ExpectedParse: []Link{
				{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"},
				{Href: "https://github.com/gophercises", Text: "Gophercises is on Github!"},
			},
		},
		{
			Description: "test example 3",
			Path:        "testingexamples/ex3.html",
			ExpectedParse: []Link{
				{Href: "#", Text: "Login"},
				{Href: "/lost", Text: "Lost? Need help?"},
				{Href: "https://twitter.com/marcusolsson", Text: "@marcusolsson"},
			},
		},
		{
			Description:   "test example 4",
			Path:          "testingexamples/ex4.html",
			ExpectedParse: []Link{{Href: "/dog-cat", Text: "dog cat"}},
		},
	}
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			f, err := os.Open(test.Path)
			if err != nil {
				t.Fatalf("got an error not expected: %v", err)
			}
			assertParse(t, f, test.ExpectedParse)
		})
	}
}

func assertParse(t testing.TB, r io.Reader, want []Link) {
	got := HtmlLinkParse(r)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
