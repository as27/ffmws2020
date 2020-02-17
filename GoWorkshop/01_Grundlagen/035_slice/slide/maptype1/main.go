package main

import "fmt"

// START OMIT
type MyMap map[string]int

func main() {
	var a MyMap
	a["eins"] = 1
	fmt.Println(a)

}
// END OMIT
