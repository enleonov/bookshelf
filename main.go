package main

import (
		"fmt"
		"net/http"
		"log"
		"github.com/gorilla/mux"
		"database/sql"
		_ "github.com/lib/pq"
		"github.com/BurntSushi/toml"
		"os"
		)

type Book struct {
	Id int
	Name string
	Author string
	LibraryId int
}

type Page struct {
	Title	string
	Books []Book
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}


var db *sql.DB

type ConfigDB struct {
    DB_USER string
    DB_PASSWORD string
    DB_NAME string
}


// Reads db info from config file
func ReadConfigDB() ConfigDB {
	var configfile = ".bookshelfrc"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config ConfigDB
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}


func main() {

    var configDB = ReadConfigDB()
    //fmt.Print(configDB.DB_NAME)

    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
        configDB.DB_USER, configDB.DB_PASSWORD, configDB.DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    _, err = db.Exec("CREATE TABLE IF NOT EXISTS " +
      `library("id" SERIAL PRIMARY KEY,` +
      `"name" varchar(100))`)
    checkErr(err)

    _, err = db.Exec("CREATE TABLE IF NOT EXISTS " +
      `book("id" SERIAL PRIMARY KEY,` +
      `"name" varchar(100), "author" varchar(100), "library_id" integer NOT NULL REFERENCES public.library(id))`)
    checkErr(err)

/*
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS " +
      `library_book("id" SERIAL PRIMARY KEY,` +
      `"library_id" integer NOT NULL REFERENCES library(id), "book_id" integer NOT NULL REFERENCES public.book(id))`)
    checkErr(err)
*/

    router := mux.NewRouter().StrictSlash(true)
    //router.Methods("GET", "POST").HandleFunc("/products/{key}", ProductHandler)
    router.HandleFunc("/", Index).Methods("GET")
    router.HandleFunc("/books", CreateBook).Methods("POST")
    //router.GET("/books/:id", getBook)
    //router.PUT("/books/:id", updateBook)
    //router.DELETE("/books/:id", deleteBook)
    //router.GET("/api/v1/records", getRecords)

    router.HandleFunc("/todos", TodoIndex)
    router.HandleFunc("/todos/{todoId}", TodoShow)


    log.Fatal(http.ListenAndServe(":9013", router))

}



