package convert

import (
	"bytes"
	"fmt"
	"testing"
)

/*
In dieser Datei befinden sich nur Benchmarks. Es werden drei unterschiedliche
Möglichkeiten verglichen, wie ein string zu einem string hinzugefügt wird

Wechsle in des Verzeichnis und führe dort den folgenden Befehl aus:

go test -bench .

Was ist das Ergebnis?
*/

var a string = "Erster Teil"
var a2 string = "Zweiter Teil"

func BenchmarkSprintf(b *testing.B) {
	a1 := a
	for i := 0; i < b.N; i++ {
		a1 = fmt.Sprintf("%s%s", a1, a2)
	}
}

func BenchmarkPlus(b *testing.B) {
	a1 := a
	for i := 0; i < b.N; i++ {
		a1 = a1 + a2
	}
}

func BenchmarkBuffer(b *testing.B) {
	var buffer bytes.Buffer
	buffer.WriteString(a)
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer

		buffer.WriteString(a2)

	}
}
