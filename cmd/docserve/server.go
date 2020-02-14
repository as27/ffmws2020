package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type server struct {
	addr   string
	router *http.ServeMux
	pages  []Page
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

	s.router.HandleFunc("/index", s.handleIndex())
	s.router.HandleFunc("/page/", s.handlePage())
}

func (s *server) loadPages(folder string) error {
	pages, err := loadPages(folder)
	if err != nil {
		return fmt.Errorf("error server.loadPages: %w", err)
	}
	s.pages = pages
	return nil
}

func (s *server) getPage(id string) (Page, bool) {
	for _, p := range s.pages {
		if id == p.ID {
			return p, true
		}
	}
	return Page{}, false
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		enc.SetIndent("  ", "")
		enc.Encode(s.pages)
	}
}

func (s *server) handlePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pageID := r.RequestURI[len("/page/"):]
		page, ok := s.getPage(pageID)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Page not found!"))
			return
		}
		enc := json.NewEncoder(w)
		enc.SetIndent("  ", "")
		enc.Encode(page)
	}
}
