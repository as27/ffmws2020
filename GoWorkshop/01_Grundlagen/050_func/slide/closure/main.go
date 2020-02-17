package main

import "fmt"

func main() {
	a := counter()
	b := counter()
	fmt.Println(a(), a())
	fmt.Println(b(), b(), b())
	fmt.Println(a(), a())
}

func counter() func() int {
	c := 0
	return func() int {
		c++
		return c
	}
}
