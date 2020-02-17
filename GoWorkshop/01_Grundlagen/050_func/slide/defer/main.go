package main

import "fmt"

func main() {
	foo()
}

func foo() {
	fmt.Println("Eins")
	defer fmt.Println("Zwei")
	fmt.Println("Drei")
	defer fmt.Println("Vier")
}
