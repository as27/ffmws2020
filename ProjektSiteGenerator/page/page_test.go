package page_test

import (
	"strings"
	"testing"

	"github.com/as27/dpunkt_go_buch/code/projekte/projekt1_site_generator/page"
)

const testPage = `
Title:   Das ist der Titel

---
# Überschrift 1

Hier steht noch mehr Text.
`

func TestReadHeader(t *testing.T) {
	r := strings.NewReader(testPage)
	header, _ := page.ReadHeader(r)
	wantTitle := "Das ist der Titel"
	if header.Title != wantTitle {
		t.Errorf("ReadHeader() = '%v', want %v", header.Title, wantTitle)
	}
}
