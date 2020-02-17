package main

import "fmt"

func main() {
	// START OMIT
	a := map[string]int{
		"eins": 1,
		"zwei": 2,
		"drei": 3,
		"fünf": 5,
		"zehn": 10,
	}
	for k, v := range a {
		fmt.Println("key:", k, "value:", v)
	}
	// END OMIT
}
