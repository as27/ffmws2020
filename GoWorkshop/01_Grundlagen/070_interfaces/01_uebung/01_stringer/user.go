package user

type User struct {
	Name    string
	City    string
	Country string
	Age     int
}

/*
	Aufgabe implementiere das Stringer Interface. Als Output für den String soll
	folgende Logik gelten:
	Name (Age)
	City
	Country
	----------
	Also für
	User{
		"Max",
		"Musterstadt",
		"Musterland",
		66,
	}
	soll folgender String erzeugt werden:
	Max (66)
	Musterstadt
	Musterland
	----------
*/
