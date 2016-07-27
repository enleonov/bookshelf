package main

import (
	"database/sql"
)


func insert(name, author string, library int) (sql.Result, error) {
        return db.Exec("INSERT INTO book VALUES (default, $1, $2, $3)",
                name, author, library)
}

/*


func insert(name, author string, library int) (sql.Result, error) {
	return db.Exec("INSERT INTO book VALUES (default, $1, $2, $3)",
		name, author, library)
}

func remove(id int) (sql.Result, error) {
	return db.Exec("DELETE FROM book WHERE id=$1", id)
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

*/

func readBooks(str string) ([]Book, error) {
	var err error
	if db == nil {
		return nil, err
	}
	var rows *sql.Rows
	if str != "" {
		rows, err = db.Query("SELECT * FROM book WHERE name LIKE $1 ORDER BY id",
			"%"+str+"%")
	} else {
		rows, err = db.Query("SELECT * FROM book ORDER BY id")
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
 
	var rs = make([]Book, 0)
	var rec Book
	for rows.Next() {
		if err = rows.Scan(&rec.Id, &rec.Name, &rec.Author, &rec.LibraryId); err != nil {
			return nil, err
		}
		rs = append(rs, rec)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rs, nil
}


