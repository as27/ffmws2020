package main

type User struct {
	Vorname  string
	Nachname string
	Age      int
}

type Users []User

var MyUsers = Users{
	User{"Max", "Mustermann", 33},
	User{"Nico", "Niemand", 44},
	User{"Olaf", "Oberst", 55},
}

func main() {
	// Erstelle einen for loop, welcher alle MyUsers ausgibt.
	// Folgendes Ergebnis wird erwartet:
	// Max Mustermann (33)
	// Nico Niemand (44)
	// Olaf Oberst (55)
}
