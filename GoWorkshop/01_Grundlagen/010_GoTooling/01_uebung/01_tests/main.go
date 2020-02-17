/*
In der ersten Übung ist ein Test vorbereitet, welcher einen Fehler
ausgibt.
Als erstes solltest Du einen Testlauf mit

go test

durchführen und sehen wie der Output ist. Danach kannst Du die Funktion
so anpassen, dass der Test erfolgreich durchläuft.

Ziel der Übung ist nicht den Code zu verstehen. Wir wollen einfach mal
mit einem Test arbeiten.
*/
package main

import "fmt"

func main() {
	var s string
	s = getHelloString()
	fmt.Println(s)
}

func getHelloString() string {
	return "Hello World!"
}
