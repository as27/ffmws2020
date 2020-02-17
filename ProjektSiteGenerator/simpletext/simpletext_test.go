package simpletext_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/as27/dpunkt_go_buch/code/projekte/projekt1_site_generator/simpletext"
)

const example = `
# Überschrift 1

## Überschrift 2

Eine Zeile kommt hier rein.
Das ist ein neuer Paragraph

* Liste 1
* Liste 2

Mit [einem link](/a/seite2.htm) und [noch einem](/a/seite1.htm).

`

func TestWrite(t *testing.T) {
	r := strings.NewReader(example)
	buf := &bytes.Buffer{}
	simpletext.Write(buf, r)
	body := buf.String()
	wantLines := []string{
		"<h1>Überschrift 1</h1>",
		"<h2>Überschrift 2</h2>",
		"<p>Eine Zeile kommt hier rein.</p>",
		"<p>Das ist ein neuer Paragraph</p>",
		"<ul>\n<li>Liste 1</li>",
		"<li>Liste 2</li>\n</ul>",
		"Mit <a href=\"/a/seite2.htm\">einem link</a>",
		"<a href=\"/a/seite1.htm\">noch einem</a>",
	}

	for _, want := range wantLines {
		if !strings.Contains(body, want) {
			t.Errorf("Body %s does not contain: %s", body, want)
		}
	}
}
