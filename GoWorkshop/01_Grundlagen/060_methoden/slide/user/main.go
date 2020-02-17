package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) String() string {
	return fmt.Sprintf("Name: %s (%d)\nEmail: %s", u.Name, u.Age, u.Email)
}

func main() {
	me := User{"Andreas", "andreas@mail.com", 40}
	fmt.Print(me)
}
