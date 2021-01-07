package main

import (
	"net/http"
	"html/template"
	"log"
	"github.com/julienschmidt/httprouter"
)

type HomePage struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p := &HomePage{Name: "Welcome to page"}
	t, e := template.ParseFiles("./templates/home.html")
	if e != nil {
		log.Printf("Parsing templates home.html error: %s", e)
		return
	}

	t.Execute(w, p)
	return
}