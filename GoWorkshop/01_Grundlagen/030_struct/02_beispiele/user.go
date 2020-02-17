package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	var andreas = User{
		Name:  "Andreas",
		Email: "andreas@mail.com",
		Age:   40,
	}
	fmt.Println(andreas)
}
