package main

import (
	"fmt"
)

func main() {
	// START OMIT
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%#v\n", a[1])
	fmt.Printf("%#v\n", a[3:7])
	fmt.Printf("%#v\n", a[:3])
	fmt.Printf("%#v\n", a[5:])
	fmt.Printf("%#v\n", a[:])
	// END OMIT
}
