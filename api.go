package main

import (
	"database/sql"
	"fmt"
	"net/http"
)
// Renamed setupJsonApi to setupJSONAPI for consistency and adherence to Go naming conventions (camelCase).
func setupJSONAPI() {
	http.HandleFunc("/createUser", func(w http.ResponseWriter, r *http.Request) {
		// create MySQL connection
		conn, err := createConnection()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		name := r.FormValue("name")
		email := r.FormValue("email")
		// Used parameterized queries (?) in SQL statements to prevent SQL injection vulnerabilities.
		query := "INSERT INTO users (name, email) VALUES (?, ?)"
		result, err := conn.Exec(query, name, email)
		if err != nil {
			fmt.Println("Error executing query:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Println("result ", result)
		w.Write([]byte("Created user successfully!"))
	})

	http.HandleFunc("/updateUser", func(w http.ResponseWriter, r *http.Request) {
		// create MySQL connection
		conn, err := createConnection()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		name := r.FormValue("name")
		email := r.FormValue("email")
		//Used parameterized queries (?) in SQL statements to prevent SQL injection vulnerabilities.
		query := "UPDATE users SET name=?, email=? WHERE id=?"
		result, err := conn.Exec(query, name, email, r.FormValue("id"))
		if err != nil {
			fmt.Println("Error executing query:", err)
			//Added error handling for database connections and query executions. It returns an HTTP 500 status and an error message if something goes wrong.
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Println("result ", result)
		w.Write([]byte("User updated successfully!"))
	})
}
