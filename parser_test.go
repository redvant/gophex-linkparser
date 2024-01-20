package linkparser

import (
	"strings"
	"testing"
)

func TestLinkParse(t *testing.T) {
	r := strings.NewReader("test")

	got := LinkParse(r)
	want := "test"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
