package main

import (
	"fmt"
	"net/http"
	"taskmaster/handlers"
	"taskmaster/repository"
)

// main function
func main() {

	// Initialize the database
	db, err := repository.InitDB("taskmaster.db")
	if err != nil {
		panic(err)
	}

	// this defer statement will close the database connection when the main function exits, very powerful.
	defer db.Close()

	http.HandleFunc("/users/register", handlers.RegisterUserHandler(db))
	http.HandleFunc("/users/login", handlers.LoginUserHandler(db))

	fmt.Println("Starting server on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
