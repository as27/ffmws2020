package main

import "fmt"

// START1 OMIT
type Geo interface {
	Name() string
	Area() float64
}

func PrintGeo(g Geo) {
	fmt.Println(g.Name())
	fmt.Println("Fläche:", g.Area())
}

// END1 OMIT
// START2 OMIT
type Rect struct {
	width  float64
	height float64
}

func (Rect) Name() string {
	return "Rechteck"
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

// END2 OMIT
// START3 OMIT
type Triangle struct {
	width  float64
	height float64
}

func (Triangle) Name() string {
	return "Dreieck"
}

func (t Triangle) Area() float64 {
	return t.width * t.height * 0.5
}

func main() {
	r := Rect{10.0, 5.0}
	PrintGeo(r)
	d := Triangle{10.0, 5.0}
	PrintGeo(d)
}

// END3 OMIT
