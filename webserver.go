package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"./datalocal"
)

var templ = template.Must(template.ParseFiles("./htmltemplate/htmlTemplate.html"))
var temp = template.Must(template.ParseFiles("./htmltemplate/index.html"))
var story datalocal.StoryF

func handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	// param := strings.Split(url, "/")
	story = datalocal.UseDecode("./datalocal/gopher.json")
	chapter := strings.Split(url, "/")[1]
	for key := range story {
		if key == chapter {
			err := templ.Execute(w, story[chapter])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

	}
	err := temp.Execute(w, story)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", mux))
}
