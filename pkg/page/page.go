package page

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

type PageHeader struct {
	ID        string    `json:"id,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

// Page definiert eine Seite
// Die kann später auch um weitere Eigenschaften
// erweitert werden.
type Page struct {
	PageHeader `json:"page_header,omitempty"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
}

// Load loads a page from a given filepath
func Load(fpath string) (Page, error) {
	var p Page
	fd, err := os.Open(fpath)
	if err != nil {
		return p, fmt.Errorf("loadPage Open(fpath): %w", err)
	}
	defer fd.Close()
	finfo, err := fd.Stat()
	if err != nil {
		return p, fmt.Errorf("loadPage fd.Stat(): %w", err)
	}
	ph := PageHeader{
		ID:        fd.Name(),
		Timestamp: finfo.ModTime(),
	}
	return readPage(fd, ph)
}

// readPage reads the conten from an io.Reader. Additional
// properties are also added.
func readPage(r io.Reader, ph PageHeader) (Page, error) {
	buf := &bytes.Buffer{}
	var p Page
	p.PageHeader = ph
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		l := scanner.Text()
		if p.Title == "" {
			p.Title = l
		}
		fmt.Fprintln(buf, l)
	}
	p.Content = buf.String()
	return p, nil
}
