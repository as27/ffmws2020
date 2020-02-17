package main

import "fmt"

func main() {
	// START OMIT
	a := []int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Println("value:", v)
	}
	// END OMIT
}
