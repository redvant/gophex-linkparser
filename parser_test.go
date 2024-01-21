package linkparser

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

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
