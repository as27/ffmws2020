package main

import (
	"fmt"
)

func div(a, b float32) (float32, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func main() {
	d, ok := div(10, 0)
	if !ok {
		fmt.Println("Division nicht erlaubt!", "d: ", d)
	}
	e, _ := div(10, 2)
	fmt.Println("e: ", e)
}
