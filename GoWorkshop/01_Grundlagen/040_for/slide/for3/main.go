package main

import "fmt"

func main() {
	// START OMIT
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 10 {
			break
		}
	}
	// END OMIT
}
