package main

import "fmt"

type Celsius int
type Fahrenheit int

func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(float32(c)*1.8 + 32)
}

func main() {
	var c Celsius = 26
	fmt.Println(CelsiusToFahrenheit(c))
}
