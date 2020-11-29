package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	tmpl := loadTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Path[1:]
		t := tmpl.Lookup(filename)
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Fatalln(err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.Handle("/css/", http.FileServer(http.Dir("wwwroot")))
	http.Handle("/img/", http.FileServer(http.Dir("wwwroot")))

	http.ListenAndServe(":8080", nil)
}

func loadTemplates() *template.Template {
	tmpl := template.New("templates")
	tmpl, _ = tmpl.ParseGlob("templates/*.html")
	return tmpl
}
