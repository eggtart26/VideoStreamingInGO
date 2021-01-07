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
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		p := &HomePage{Name: "Welcome to page"}
		t, e := template.ParseFiles("./templates/home.html")
		if e != nil {
			log.Printf("Parsing templates home.html error: %s", e)
			return

		}

		t.Execute(w, p)
		return
	}

	if len(cname.Value) != 0 && len(sid.Value) != 0 {
	 http.Redirect(w, r, "/userhome", http.StatusFound)
	 return
 	}
}