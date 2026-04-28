package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // 注册sqlite3
)

// Book is a placeholder for book
type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func dbOperations(db *sql.DB) {
	// Create
	statement, _ := db.Prepare("INSERT INTO books (name, author, isbn) VALUES (?, ?, ?)")
	_, _ = statement.Exec("A Tale of Two Cities", "Charles Dickens", 140430547)
	log.Println("Inserted the book into database!")

	// Read
	rows, _ := db.Query("SELECT id, name, author FROM books")
	defer rows.Close() // 遍历完所有行且没有提前退出, 可以省略

	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.Id, &tempBook.Name, &tempBook.Author)
		log.Printf("ID:%d, Book:%s, Author:%s\n", tempBook.Id, tempBook.Name, tempBook.Author)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	// Update
	statement, _ = db.Prepare("update books set name=? where id=?")
	statement.Exec("The Tale of Two Cities", 1)
	log.Println("Successfully updated the book in database!")

	// Delete
	statement, _ = db.Prepare("delete from books where id=?")
	statement.Exec(1)
	log.Println("Successfully deleted the book in database!")
}

func main() {
	db, err := sql.Open("sqlite3", "books.db")
	if err != nil {
		log.Println(err)
	}

	// Create table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books(id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64),name VARCHAR(64) NULL)")
	if err != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfully created table books!")
	}

	statement.Exec()

	// CRUD
	dbOperations(db)
}
