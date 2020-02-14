package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// Page definiert eine Seite
// Die kann später auch um weitere Eigenschaften
// erweitert werden.
type Page struct {
	ID        string    `json:"id,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
}

func loadPages(folder string) ([]Page, error) {
	var pages []Page
	d, err := ioutil.ReadDir(folder)
	if err != nil {
		return pages, fmt.Errorf("loadPages: ReadDir: %w", err)
	}
	for _, f := range d {
		if f.IsDir() {
			continue
		}
		p, err := loadPage(filepath.Join(folder, f.Name()))
		if err != nil {
			return pages, fmt.Errorf("loadPages: loadPage %s:%w", f.Name(), err)
		}
		pages = append(pages, p)
	}
	return pages, nil
}

// loadPage loads a page from a given filepath
func loadPage(fpath string) (Page, error) {
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
	p = Page{
		ID:        filepath.Base(fd.Name()),
		Timestamp: finfo.ModTime(),
		Title:     filepath.Base(fd.Name()),
	}
	b, err := ioutil.ReadAll(fd)
	if err != nil {
		return p, fmt.Errorf("LoadPage: Read from file: %w", err)
	}
	p.Content = string(b)
	return p, nil
}
