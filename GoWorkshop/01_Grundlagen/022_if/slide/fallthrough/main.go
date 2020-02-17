package main

import "fmt"

// START OMIT
func main() {
	foo(2)
	foo(10)
}

func foo(nr int) {
	switch {
	case nr == 2:
		fmt.Println("Ah")
		fallthrough
	case nr < 5:
		fmt.Println("jetzt ja")
	default:
		fmt.Println("eine Insel")
	}

}

// END OMIT
