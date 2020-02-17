package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	if fmt.Sprintf("%T", B) != "main.Anzahl" {
		fmt.Println("Definiere Variable B als Typ Anzahl!")
	}
}
