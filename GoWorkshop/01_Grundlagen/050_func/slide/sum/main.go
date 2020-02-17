package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3), sum(4), sum(), sum(5, 6, 7, 8, 9))
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(sum(a...))
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
