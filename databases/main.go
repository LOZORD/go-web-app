package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	db := NewDB()
	loc := ":8080"
	log.Println("Listening on " + loc)
	http.ListenAndServe(loc, ShowBooks(db))
}

func ShowBooks(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var title string
		var author string

		err := db.QueryRow("select title, author from books").Scan(&title, &author)

		if err != nil {
			panic(err)
		}

		fmt.Fprintf(rw, "The first book is '%s' by '%s'", title, author)
	})
}

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "example.sqlite")

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("create table if not exists books(title text, author text)")

	if err != nil {
		panic(err)
	}

	return db
}

/* TODO
func addBook(book *Book, db *sql.DB) {
	dupQ := "select (count(*) > 0) from books where title=? and author=?"
	err := db.QueryRow(dupQ, book.title, book.author)
}
*/
