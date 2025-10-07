package main

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func Play(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./pages/play.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/play", Play)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
