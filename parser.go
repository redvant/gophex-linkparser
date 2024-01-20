package linkparser

import (
	"io"
)

func LinkParse(r io.Reader) string {
	b, _ := io.ReadAll(r)
	return string(b)
}
