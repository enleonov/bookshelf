package main

import (
		"net/http"
		"html/template"
		"encoding/json"
		"log"
		"github.com/gorilla/mux"
		"strconv"
)



func index(w http.ResponseWriter, r *http.Request) {
	books, err := readBooks("")
	if err != nil {
		w.WriteHeader(500)
		return
	}

	libraries, err := readLibraries()

	title := "bookshelf"
	p := &Page{Title: title, Books: books, Libraries: libraries}
	renderTemplate(w, "index", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl+".tmpl")
	t.Execute(w, p)
}


func deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
        if err != nil || id <= 0{
                log.Println(err.Error())
                w.WriteHeader(404)
                return
        }

	var res int
        if res, err = deleteBookById(id); err != nil {
                log.Println(err.Error())
                w.WriteHeader(500)
                return
        }
	if res <= 0 {
		w.WriteHeader(500)
		log.Println("Failed delete book")
		return
	}
	w.WriteHeader(204)

}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil || book.Name == "" || book.Author == "" {
		log.Println(err.Error())
		w.WriteHeader(400)
		return
	}

	var id int
	if id, err = insert(book.Name, book.Author, book.LibraryId); err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		return
	}

	var newBook Book
	newBook, err = getBookById(id)
	log.Println("ID", newBook.Id)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
        if err = json.NewEncoder(w).Encode(newBook); err != nil {
              w.WriteHeader(500)
		return
        }

	w.WriteHeader(201)
}





