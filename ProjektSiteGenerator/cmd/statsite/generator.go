package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/as27/ffmws2020/ProjektSiteGenerator/page"
	"github.com/as27/ffmws2020/ProjektSiteGenerator/simpletext"
)

func loadPages(fpaths ...string) ([]page.Page, error) {
	var pages []page.Page
	for _, fpath := range fpaths {
		f, err := os.Open(fpath)
		if err != nil {
			return nil, fmt.Errorf("fehler beim öffnen von %s: %s", fpath, err)
		}
		var newPage page.Page
		var content io.Reader
		newPage.Header, content = page.ReadHeader(f)
		newPage.Src = fpath
		base := filepath.Base(fpath)
		newPage.Path = base[:len(base)-len(filepath.Ext(base))]
		buf := &bytes.Buffer{}
		simpletext.Write(buf, content)
		newPage.Content = buf.Bytes()
		pages = append(pages, newPage)
	}
	return pages, nil
}
