package main

import (
	"html/template"
	"net/http"
)

var templateStart *template.Template

func start(w http.ResponseWriter, r *http.Request) {
	templateStart.Execute(w, nil)
}

func main() {
	templateStart = template.Must(template.ParseFiles("template/start.html"))
	http.HandleFunc("/", start)
	http.ListenAndServe(":8000", nil)
}
