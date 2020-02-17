package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// START OMIT
func main() {
	r := strings.NewReader("abc,def,hij")
	fmt.Println(split(r))
}

func split(r io.Reader) []string {
	b, _ := ioutil.ReadAll(r)
	s := string(b)
	return strings.Split(s, ",")
}

// END OMIT

func fromFile(fname string) []string {
	f, _ := os.Open(fname)
	defer f.Close()
	return split(f)
}

// END2 OMIT
