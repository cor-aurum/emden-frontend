package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
)

var templateStart *template.Template
var extended bool = false

const APIUrl string = "http://localhost:8080/EMDEN/"

//const APIUrl string = "http://localhost:8080/Suche-Backend-0.0.1-SNAPSHOT/"

type Page struct {
	ChangeSearchType string
	Results          []interface{}
	Search           string
	Logo             string
	Submit           string
	Title            string
	Datafields       []string
}

func getJson(url string, target interface{}) error {
	var c = &http.Client{}
	r, err := c.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getString(url string) string {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return string(bodyBytes)
	}
	return ""
}

func getSearch(r *http.Request) (bool, string) {
	suc := false
	searchstring := r.URL.RawQuery //.Get("search")
	if len(searchstring) != 0 {
		suc = true
	}
	return suc, searchstring
}

func start(w http.ResponseWriter, r *http.Request) {
	extended = false
	handlePage(w, r, "Erweiterte Suche")
}

func extendedSearch(w http.ResponseWriter, r *http.Request) {
	extended = true
	handlePage(w, r, "Einfache Suche")
}

func handlePage(w http.ResponseWriter, r *http.Request, searchType string) {
	page := Page{ChangeSearchType: searchType, Search: "Suchbegriff"}
	page.Logo = APIUrl + "company/logo"
	page.Submit = "LOS!"
	page.Title = "EMDEN | " + getString(APIUrl+"company/name")
	if extended {
		getJson(APIUrl+"company/datafields", &page)
	}
	search, searchstring := getSearch(r)
	if search {
		getJson(APIUrl+"search?"+url.PathEscape(searchstring), &page)
	}

	templateStart.Execute(w, page)
}

func main() {
	templateStart = template.Must(template.ParseFiles("template/start.html"))
	http.Handle("/res/", http.StripPrefix("/res", http.FileServer(http.Dir("res/"))))
	http.HandleFunc("/", start)
	http.HandleFunc("/extended", extendedSearch)
	fmt.Println("Start")
	http.ListenAndServe(":8000", nil)
}
