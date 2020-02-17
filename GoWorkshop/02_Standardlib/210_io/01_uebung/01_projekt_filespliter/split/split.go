// Package split wird verwendet um Bytes aus einem io.Reader in
// kleinere Teile sogenannte chunks aufzuteilen.
// Einrückungen hier erzeugen einen Beispielcodeblock in der
// Dokumentation
//   c, _ := Split(myFile, 1024)
package split

import "io"

// Chunk ist ein Teil der ursprünglichen Datei
type Chunk []byte

// Chunks ist eine Sammlung meherer einzelnen Chunk. In der Regel
// stellt Chunks eine geteilte Datei dar.
type Chunks []Chunk

// Split bekommt als Input einen io.Reader und erzeugt daraus chunks
// in der größe von size
func Split(r io.Reader, size int) (Chunks, error) {
	return Chunks{}, nil
}

// Merge ist die Umkehrfunktion von Split. Als Input bekommt diese
// neben Chunks auch einen io.Writer. Dieser wird verwendet, um die
// in die einzelnen Bytes zu schreiben.
func Merge(c Chunks, w io.Writer) error {
	// Bevor du den Code für Merge schreibst, solltest Du einen
	// Test erstellen.
	// Um die Daten, welche in den io.Writer geschrieben werden
	// auszuwerten bietet sich der Typ bytes.Buffer an, da dieser
	// auch einen io.Writer implementiert.
	return nil
}
