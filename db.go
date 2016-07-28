package main

import (
	"database/sql"
//	"log"
)


func insert(name, author string, library int) (int, error) {
        var err error
        if db == nil {
                return -1, err
        }
        var row *sql.Rows
        row, err = db.Query("INSERT INTO book (id, name, author, library_id) 
		VALUES ( default, $1, $2, $3) RETURNING id", name, author, library)
        if err != nil {
                return -1, err
        }
        defer row.Close()

        var id int
	id = -1
	row.Next()
	if err = row.Scan(&id); err != nil {
		return -1, err
	}
	return id, err
}


func deleteBookById(id int) (int, error) {
	res, _ := db.Exec("DELETE FROM book WHERE id=$1", id)
	cnt, err := res.RowsAffected()
	return int(cnt), err
}

 
func update(id int, name, author string, libraryId int) (sql.Result, error) {
	return db.Exec("UPDATE book SET name = $1, author = $2 WHERE id=$3",
		name, author, id)
}

func readOneBook(id int) (Book, error) {
	var rec Book
	row := db.QueryRow("SELECT * FROM book WHERE id=$1 ORDER BY id", id)
	return rec, row.Scan(&rec.Id, &rec.Name, &rec.Author)
}



func readBooks(str string) ([]Book, error) {
	var err error
	if db == nil {
		return nil, err
	}
	var rows *sql.Rows
	if str != "" {
		rows, err = db.Query("SELECT book.id, book.name, book.author, library.id, 
			library.name FROM book, library WHERE name LIKE $1 
			and book.library_id = library.id ORDER BY book.id",
			"%"+str+"%")
	} else {
		rows, err = db.Query("SELECT book.id, book.name, book.author, library.id, 
			library.name FROM book, library WHERE book.library_id = library.id 
			ORDER BY book.id")
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
 
	var rs = make([]Book, 0)
	var rec Book
	for rows.Next() {
		if err = rows.Scan(&rec.Id, &rec.Name, &rec.Author, &rec.LibraryId, &rec.Library); err != nil {
			return nil, err
		}
		rs = append(rs, rec)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rs, nil
}


func readLibraries() ([]Library, error) {
        var err error
        if db == nil {
                return nil, err
        }
        var rows *sql.Rows
        rows, err = db.Query("SELECT * FROM library ORDER BY id")
        if err != nil {
                return nil, err
        }
        defer rows.Close()
 
        var rs = make([]Library, 0)
        var rec Library
        for rows.Next() {
                if err = rows.Scan(&rec.Id, &rec.Name); err != nil {
                        return nil, err
                }
                rs = append(rs, rec)
        }
        if err = rows.Err(); err != nil {
                return nil, err
        }
        return rs, nil
}

func getBookById(id int) (Book, error) {
	var book Book
	var row *sql.Rows
        var err error
	row, err = db.Query("SELECT book.id, book.name, book.author, library.id, 
		library.name FROM book, library WHERE book.library_id = library.id 
		and book.id = $1", id)
	if err != nil {
		return book, err
	}
	defer row.Close()
	row.Next()
	if err = row.Scan(&book.Id, &book.Name, &book.Author, &book.LibraryId, &book.Library); err != nil {
		return book, err
        }
	return book, nil
}

