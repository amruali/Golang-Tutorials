package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

func connectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./Chinook.db")
	if err != nil {
		fmt.Printf("failed to connect to database because %v", err)
	}
	return db
}

func main() {
	db := connectDB()
	defer db.Close()
	rows, _ := db.Query("SELECT ALBUMID, TITLE FROM ALBUM")
	var id int
	var title string
	for rows.Next() {
		rows.Scan(&id, &title)
		fmt.Println(strconv.Itoa(id), title)
	}
}
