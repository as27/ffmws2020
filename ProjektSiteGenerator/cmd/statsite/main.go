package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/as27/ffmws2020/ProjektSiteGenerator/page"
)

var (
	flagSrcFolder  = flag.String("src", "./content/", "src folder")
	flagOutFolder  = flag.String("out", "./site/", "output folder")
	flagTmplFolder = flag.String("tmpl", "./templates/", "template folder")
)

func main() {
	flag.Parse()

	files, err := ioutil.ReadDir(*flagSrcFolder)
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "content folder %s does not exist\n", *flagSrcFolder)
		os.Exit(1)
	}
	srcpages := []string{}
	for _, finfo := range files {
		name := finfo.Name()
		if filepath.Ext(name) == ".txt" {
			srcpages = append(srcpages, filepath.Join(*flagSrcFolder, name))
		}
	}
	pages, err := loadPages(srcpages...)

	for _, page := range pages {
		outFolder := filepath.Join(*flagOutFolder, page.Path)
		os.MkdirAll(outFolder, 0777)
		f, err := os.OpenFile(filepath.Join(outFolder, "index.html"),
			os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
			0777)
		if err != nil {
			fmt.Println("error opening file", outFolder)
		}
		data := struct {
			Title   string
			Content template.HTML
		}{
			page.Header.Title,
			template.HTML(page.Content),
		}
		renderPage(f, data, "page.tmpl.html")
		f.Close()
	}
	f, _ := os.OpenFile(filepath.Join(*flagOutFolder, "index.html"),
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
		0777)
	data := struct {
		Title string
		Pages []page.Page
	}{
		"Homepage",
		pages,
	}
	renderPage(f, data, "home.tmpl.html")
	f.Close()

}

func renderPage(w io.Writer, data interface{}, content string) {
	tmpl, err := template.ParseFiles(
		filepath.Join(*flagTmplFolder, "layout", "base.tmpl.html"),
		filepath.Join(*flagTmplFolder, "layout", "header.tmpl.html"),
		filepath.Join(*flagTmplFolder, "layout", "footer.tmpl.html"),
		filepath.Join(*flagTmplFolder, content),
	)
	if err != nil {
		fmt.Println("error parsing template", err)
	}
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		fmt.Println("error execute template", err)
	}
}
