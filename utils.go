package main

import (
	"database/sql"
	"fmt"
)

//Modified createConnection to return an error, allowing better error handling and removed the print statement for the SQL open error, as it's already being returned.
func createConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		return nil, err
	}
	return db, nil
}
