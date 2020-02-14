package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/as27/ffmws2020/pkg/page"
)

var flagPagesFolder = flag.String("p", "pages", "set the pages folder")

func main() {
	d, err := ioutil.ReadDir(*flagPagesFolder)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, f := range d {
		if f.IsDir() {
			continue
		}
		fmt.Println(f.Name())
		p, err := page.Load(filepath.Join(*flagPagesFolder, f.Name()))
		if err != nil {
			fmt.Println("Fehler page.Load:", err)
		}
		fmt.Printf("%#v", p)
	}
}
