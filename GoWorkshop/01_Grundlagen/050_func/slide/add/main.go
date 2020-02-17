package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func main() {
	fmt.Println(add(2, 3))
	fmt.Println(sub(200, 123))
}
