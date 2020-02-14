package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var flagPagesFolder = flag.String("p", "pages", "set the pages folder")
var flagServerPort = flag.String("port", ":8345", "set the port for the server")

func main() {
	flag.Parse()
	srv := &server{
		addr:   *flagServerPort,
		router: http.NewServeMux(),
	}
	err := srv.loadPages(*flagPagesFolder)
	if err != nil {
		fmt.Println("Fehler loadPages:", err)
		os.Exit(1)
	}
	err = srv.run()
	fmt.Println(err)
}
