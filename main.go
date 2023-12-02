package main

import (
	"html/template"
	"net/http"
)

func htmlPasre(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("teamplates/index.html")
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", htmlPasre)
	http.ListenAndServe(":8080", nil) //start server on localhost 8080 ggg
}
