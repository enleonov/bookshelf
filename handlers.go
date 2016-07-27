package main

import (
		"fmt"
		"net/http"
		"github.com/gorilla/mux"
		"html/template"
		"encoding/json"
		)




func Index(w http.ResponseWriter, r *http.Request) {

	recs, err := readBooks("")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")


//	if err = json.NewEncoder(w).Encode(recs); err != nil {
//		w.WriteHeader(500)
//	}


	title := "FFFFF"// r.URL.Path[lenPath:]
	p := &Page{Title: title, Books: recs}
	renderTemplate(w, "index", p)
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl+".tmpl")
	t.Execute(w, p)
}


func CreateBook(w http.ResponseWriter, r *http.Request) {
	var rec Book
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil || rec.Name == "" || rec.Author == "" {
		w.WriteHeader(400)
		return
	}
	if _, err := insert(rec.Name, rec.Author, rec.LibraryId); err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(201)
}



func TodoIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}





