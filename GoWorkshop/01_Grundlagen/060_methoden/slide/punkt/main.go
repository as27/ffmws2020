package main

import (
	"fmt"
	"math"
)

type Punkt struct{ X, Y float64 }

func (p Punkt) Entfernung(q Punkt) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func main() {
	p1 := Punkt{1, 2}
	p2 := Punkt{4, 6}
	fmt.Println("Entfernung p1 p2:", p1.Entfernung(p2))
}
