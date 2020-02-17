package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	list := []string{"Null", "Eins", "Zwei", "Drei", "Vier"}
	for i, val := range list {
		wg.Add(1)
		go func(nr int, v string) {
			fmt.Println(nr, v)
			wg.Done()
		}(i, val)
	}
	wg.Wait()
}
