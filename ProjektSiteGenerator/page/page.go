package page

import (
	"bufio"
	"io"
	"strings"
)

type Header struct {
	Title string
}

type Page struct {
	Src     string
	Path    string
	Header  Header
	Content []byte
}

// ReadHeader liest den Header aus einem io.Reader und
// liefert den Header und den Inhalt (ohne Frontmatter)
// als io.Reader zurück.
func ReadHeader(r io.Reader) (Header, io.Reader) {
	var header Header
	br := bufio.NewReader(r)
	for {
		t, err := br.ReadString('\n')
		if err != nil {
			break
		}
		switch {
		case strings.HasPrefix(t, "Title:"):
			header.Title = strings.Trim(t[6:], " \n")
		case strings.HasPrefix(t, "---"):
			return header, br
		}
	}
	return header, br
}
