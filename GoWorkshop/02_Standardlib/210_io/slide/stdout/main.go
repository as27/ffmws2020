package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// START OMIT
func main() {
	w := os.Stdout
	writePage(w)
}

func writePage(w io.Writer) {
	body := loadSpiegel()
	w.Write(body)
}

func loadSpiegel() []byte {
	resp, _ := http.Get("http://www.spiegel.de/")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

// END OMIT
