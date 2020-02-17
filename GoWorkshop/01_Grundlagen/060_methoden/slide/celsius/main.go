package main

import "fmt"

type Celsius int
type Fahrenheit int

func (c Celsius) String() string {
	return fmt.Sprintf("%d °C", c)
}

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(float32(c)*1.8 + 32)
}

func main() {
	c := Celsius(24)
	fmt.Println(c.String())
	fmt.Printf("C: %s\nF: %d", c, c.ToFahrenheit())
}
