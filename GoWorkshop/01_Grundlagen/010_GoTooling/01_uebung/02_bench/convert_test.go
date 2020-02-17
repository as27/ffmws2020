package convert

import (
	"fmt"
	"strconv"
	"testing"
)

/*
In dieser Datei befinden sich nur Benchmarks. Es werden drei unterschiedliche
Möglichkeiten verglichen, wie ein int zu einem string umgewandelt werden kann.

Wechsle in des Verzeichnis und führe dort den folgenden Befehl aus:

go test -bench .

Was ist das Ergebnis?
*/

var a int = 1234

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", a)
	}
}

func BenchmarkPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprint(a)
	}
}

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Itoa(a)
	}
}
