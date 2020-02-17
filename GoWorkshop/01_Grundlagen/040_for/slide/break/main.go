package main

import "fmt"

func main() {
	// START OMIT
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, val := range values {
		if val == 5 {
			break
		}
		fmt.Println(val)
	}
	// END OMIT
}
