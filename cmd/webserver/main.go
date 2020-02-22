package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var flagPageFolder = flag.String("p", "seiten", "Der Ordern unserer md Dateien")

func main() {
	flag.Parse()
	http.HandleFunc("/page/", pageHandlerFunc)
	http.HandleFunc("/api/", apiHandler)
	http.HandleFunc("/", indexHandlerFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	dateien, err := ioutil.ReadDir(*flagPageFolder)
	if err != nil {
		log.Println("Kann Verzeichnis nicht lesen!")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Seite nicht gefunden"))
	}
	for _, fi := range dateien {
		if fi.IsDir() {
			continue
		}
		fmt.Fprintf(w, "<p><a href=\"/page/%s\">%s</a></p>", fi.Name(),
			fi.Name())
	}
}

func pageHandlerFunc(w http.ResponseWriter, r *http.Request) {
	pageID := r.URL.Path[len("/page/"):]

	f, err := os.Open(filepath.Join(*flagPageFolder, pageID))
	defer f.Close()
	if err != nil {
		log.Println("pageHandlerFunc: os.Open():", pageID, err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Seite nicht gefunden"))
		return
	}
	io.Copy(w, f)
}

type Page struct {
	ID      string `json:"id"`
	Content string `json:"seite"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	var Pages []Page
	dateien, err := ioutil.ReadDir(*flagPageFolder)
	if err != nil {
		log.Println("Kann Verzeichnis nicht lesen!")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Seite nicht gefunden"))
	}
	for _, fi := range dateien {
		b, err := ioutil.ReadFile(
			filepath.Join(*flagPageFolder,
				fi.Name()),
		)
		if err != nil {
			log.Println("api handler:", err)
			continue
		}
		p := Page{
			ID:      fi.Name(),
			Content: string(b),
		}
		Pages = append(Pages, p)
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	err = enc.Encode(Pages)
	if err != nil {
		log.Println("Fehler beim Encodieren!")
	}
}
