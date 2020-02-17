package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// START1 OMIT
type Sensor struct {
	ID     int     `json:"sensor_id"`
	Name   string  `json:"name,omitempty"`
	Values []Value `json:"values"`
}

type Value struct {
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	Timestamp int     `json:"timestamp"`
}

// END1 OMIT
// START2 OMIT
func main() {
	s := Sensor{
		1,
		"Sensor1",
		[]Value{
			Value{"T", 24, 201706010800},
			Value{"T", 28, 201706011200},
		},
	}
	b, _ := json.Marshal(s)
	fmt.Println(string(b))
	w := &bytes.Buffer{}
	encodeSensor(w, s)
	fmt.Println(w.String())
}

func encodeSensor(w io.Writer, s Sensor) error {
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	return e.Encode(s)
}

// END2 OMIT
