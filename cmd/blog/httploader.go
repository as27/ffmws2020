package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type httpLoader struct {
	addr string
}

func (l httpLoader) getPage(id string) (Page, bool) {
	var p Page
	url := fmt.Sprintf("%s/page/%s", l.addr, id)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error loading %s: %s", url, err)
		return p, false
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Status %d for %s", resp.StatusCode, url)
		return p, false
	}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&p)
	if err != nil {
		log.Printf("Error decoding %s: %s", url, err)
		return p, false
	}
	return p, true
}

func (l httpLoader) getIndex() []Page {
	var p []Page
	url := fmt.Sprintf("%s/index", l.addr)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error loading %s: %s", url, err)
		return p
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Status %d for %s", resp.StatusCode, url)
		return p
	}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&p)
	if err != nil {
		log.Printf("Error decoding %s: %s", url, err)
		return p
	}
	return p
}
