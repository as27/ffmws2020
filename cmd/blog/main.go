package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	flagServerPort     = flag.String("port", ":8344", "set the port for the server")
	flagDataPort       = flag.String("dataPort", ":8345", "set the port for the server")
	flagTemplateFolder = flag.String("templateFolder", "templates", "the folders, where the templates are")
)

func main() {
	flag.Parse()
	srv := &server{
		dataLoader: httpLoader{
			addr: fmt.Sprintf("http://localhost%s", *flagDataPort),
		},
		addr:      *flagServerPort,
		router:    http.NewServeMux(),
		templates: *flagTemplateFolder,
	}
	srv.run()
}
