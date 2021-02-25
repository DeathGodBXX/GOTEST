package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("Template/wwwroot/tmpl.html")
		tmpl.Execute(w, "Templates")

		tmple, _ := template.ParseFiles("Template/wwwroot/tmpl.html",
			"Template/wwwroot/home.html")
		tmple.ExecuteTemplate(w, "home.html", "Home Page")
	})

	http.ListenAndServe(":8080", nil)
}
