package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templateStart *template.Template

const APIUrl string = "http://localhost:8080/Suche-Backend-0.0.1-SNAPSHOT/"

type Result struct {
	Title   string
	Url     string
	Urltext string
	Text    string
}

type Page struct {
	ChangeSearchType string
	SearchBar        string
	Results          []Result
	APIUrl           string
	Search           string
}

func start(w http.ResponseWriter, r *http.Request) {
	page := Page{APIUrl: APIUrl, ChangeSearchType: "Erweiterte Suche", Search: "Suchbegriff"}
	page.SearchBar = "search"
	page.Results = append([]Result{Result{Title: "Titel1", Text: "Long1"}}, page.Results...)
	page.Results = append([]Result{Result{Title: "Titel2", Text: "Long2", Url: "code.recondita.de", Urltext: "Git-Server"}}, page.Results...)
	templateStart.Execute(w, page)
}

func main() {
	templateStart = template.Must(template.ParseFiles("template/start.html"))
	http.Handle("/res/", http.StripPrefix("/res", http.FileServer(http.Dir("res/"))))
	http.HandleFunc("/", start)
	fmt.Println("Start")
	http.ListenAndServe(":8000", nil)
}
