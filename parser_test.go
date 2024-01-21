package linkparser

import (
	"strings"
	"testing"
)

func TestLinkParse(t *testing.T) {
	//html := "<a href=\"www.google.com\">link text</a>"
	html := "<p>link text</p>"
	r := strings.NewReader(html)

	got := LinkParse(r)
	want := "link text"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
