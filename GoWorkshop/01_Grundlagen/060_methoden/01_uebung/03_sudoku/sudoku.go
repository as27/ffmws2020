package sudoku

import "fmt"

// Field stellt das Sudoku Spielfeld dar. Das Feld hat 9 Reiehn
// und 9 Spalten. Intern ist hat die erste Spalte bzw. Reihe den
// Index 0
type Field struct {
	Values [][]int // [row][col]Value
}

// String implementiert das Stringer interface, damit die
// Ausgabe einem Sudoku Feld entspricht
func (f *Field) String() string {
	s := "\n"
	for _, r := range f.Values {
		s = s + fmt.Sprintln(r)
	}
	return s
}

// GetRow ermittelt die Werte einer Reihe
func (f *Field) GetRow(r int) []int {
	return f.Values[r]
}

// GetCol ermittelt die Werte einer Spalte und liefert diese
// als []int zurück
func (f *Field) GetCol(c int) []int {

	return []int{}
}

// GetQuad ermittelt ein Quadrat einer Zelle und liefert die
// Werte des Quadrates in eiem []int zurück. Dabei werden die
// einzelnen Zeilen einfach hintereinander gehängt.
// aus:
//   123
//   456
//   789
// Wird: 123456789
func (f *Field) GetQuad(r, c int) []int {

	return []int{}
}

// MissingValues ermittelt zu einer Zelle die erlaubten Values
func (f *Field) MissingValues(r, c int) []int {

	return []int{}
}

// Solve löst das Sudoku nach einem einfachen Algorithmus
//
// Es werden alle Felder durchgegangen und pro Feld mit Wert
// 0 die fehlenden Werte ermittelt. Sollte an einem Feld nur
// ein Wert fehlen, so wird dieser Wert für die Zelle gesetzt.
// Nachdem einmal das Feld durchgeganen wurde, wird die
// Methode noch einmal aufgerufen.
// Sollte bei einem Durchgang kein neue Wert gefunden werden,
// so wird die Funktion abgebrochen und ein entsprechender
// error Wert zurück geliefert.
func (f *Field) Solve() error {
	return nil
}

// missingValues ist eine Hilfsfunktion, welche die fehlenden
// Werte ermittelt. Im Input dürfen auch doppelte Werte vorhanden
// sein. Die Tests gehen von einer Sortierten Ausgabe aus.
// Bei 9 8 7 9 6 5 4 3 0 0 sollte die Funktion 1 2 zurück liefern.
func missingValues(values []int) []int {
	missing := []int{}

	return missing
}
