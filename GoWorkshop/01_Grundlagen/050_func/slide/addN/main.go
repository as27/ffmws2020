package main

import "fmt"

func main() {
	add5 := addN(5)
	add10 := addN(10)
	fmt.Println(add5(1), add5(10))
	fmt.Println(add10(1), add10(10))
}

func addN(n int) func(int) int {
	return func(x int) int {
		return n + x
	}
}
