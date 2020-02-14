package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
	"time"
)

type dataLoader interface {
	getIndex() []Page
	getPage(id string) (Page, bool)
}

type server struct {
	dataLoader
	addr      string
	router    *http.ServeMux
	templates string
}

func (s *server) run() error {
	s.routes()
	srv := &http.Server{
		Handler:      s.router,
		Addr:         s.addr,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	return srv.ListenAndServe()
}

func (s *server) routes() {
	s.router.HandleFunc("/page/", s.handlePage())
	s.router.HandleFunc("/", s.handleIndex())
}

func (s *server) handleIndex() http.HandlerFunc {
	tmpl, err := s.parseFiles("index.tmpl.html")
	if err != nil {
		fmt.Println("handleIndex: Cannot parse files", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ps := s.dataLoader.getIndex()
		err = tmpl.ExecuteTemplate(w, "base", ps)
		if err != nil {
			fmt.Println("execute index template: ", err)
		}
	}
}

func (s *server) handlePage() http.HandlerFunc {
	tmpl, err := s.parseFiles("page.tmpl.html")
	if err != nil {
		fmt.Println("handlePage: Cannot parse files", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		pageID := r.RequestURI[len("/page/"):]
		ps, ok := s.dataLoader.getPage(pageID)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Page not found!"))
			return
		}
		err = tmpl.ExecuteTemplate(w, "base", ps)
		if err != nil {
			fmt.Println("execute index template: ", err)
		}
	}
}

func (s *server) parseFiles(content string) (*template.Template, error) {
	return template.ParseFiles(
		filepath.Join(s.templates, "base.tmpl.html"),
		filepath.Join(s.templates, "header.tmpl.html"),
		filepath.Join(s.templates, "footer.tmpl.html"),
		filepath.Join(s.templates, content),
	)
}
