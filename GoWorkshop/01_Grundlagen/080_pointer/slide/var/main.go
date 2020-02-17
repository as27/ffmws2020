package main

import "fmt"

func main() {
	// START OMIT
	a := 1
	b := &a
	fmt.Println(a, b, *b)
	c := b
	*c = 5
	fmt.Println(a, b, *b)
	// END OMIT
}
