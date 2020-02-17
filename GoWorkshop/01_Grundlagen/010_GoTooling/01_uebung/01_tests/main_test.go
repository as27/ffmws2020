package main

import "testing"

func TestGetHelloString(t *testing.T) {
	expect := "Hello Germany!"
	got := getHelloString()
	if got != expect {
		t.Errorf("\nExpect: %s \nGot: %s", expect, got)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
