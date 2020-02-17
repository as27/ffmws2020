/*
Package simpletext konvertiert text zu HTML.

Das simpletext Format ist dabei eine einfache Form von
Markdown. Dabei werden jedoch nur ein kleiner Teil der
Befehle unterstützt.

  # Überschrift 1
  ## Überschrift 2

  * Liste 1
  * Liste 2

  [Linktext](http://linkurl.com)


*/
package simpletext

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
)

var (
	reLink = regexp.MustCompile(`\[(.+?)\]\((.+?)\)`)
)

// Write schreibt das konvertierte HTML in einen io.Writer.
// Die Datenquelle für die Konvertierung ist ein io.Reader.
// Die Funktion kann somit sowohl mit einer Datei als auch
// mit einem http Request oder Response umgehen.
func Write(w io.Writer, r io.Reader) {
	var lastTag string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := scanner.Text()
		w.Write([]byte("\n"))
		switch {
		case strings.HasPrefix(t, "##"):
			wrapTag(w,
				parseLine(t[2:]),
				"h2",
			)
			lastTag = "h2"
		case strings.HasPrefix(t, "#"):
			wrapTag(w,
				parseLine(t[1:]),
				"h1",
			)
			lastTag = "h1"
		case strings.HasPrefix(t, "*"):
			if lastTag == "" {
				fmt.Fprintln(w, "<ul>")
			}
			wrapTag(w,
				parseLine(t[1:]),
				"li")
			lastTag = "li"
		case lastTag == "li" && t == "":
			// End of list
			fmt.Fprintln(w, "</ul>")
			lastTag = "ul"
		case t != "":
			// simple paragraph
			wrapTag(w, parseLine(t), "p")
		default:
			lastTag = ""
		}
	}
}

// Run funktioniert analog wie Write, nur das als
// In- und Output ein []byte verwendet wird.
func Run(input []byte) []byte {
	out := &bytes.Buffer{}
	in := bytes.NewBuffer(input)
	Write(in, out)
	return out.Bytes()
}

func parseLine(s string) string {
	s = strings.Trim(s, " ")
	s = reLink.ReplaceAllString(s, "<a href=\"$2\">$1</a>")
	return s
}

func wrapTag(w io.Writer, s, tag string) {
	fmt.Fprintf(w, "<%[1]s>%[2]s</%[1]s>", tag, s)
}
