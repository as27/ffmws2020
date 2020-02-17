package main

import "fmt"

// START OMIT
type DB struct {
	Names []string
}

func (db *DB) Add(names ...string) {
	db.Names = append(db.Names, names...)
}

func (db *DB) String() string {
	var s string
	for i, name := range db.Names {
		s = fmt.Sprintf("%s%03d: %s\n", s, i, name)
	}
	return s
}
func main() {
	myDB := &DB{}
	myDB.Add("Erster Name", "Zweiter Name", "Dritter Name")
	fmt.Print(myDB)
}

// END OMIT
