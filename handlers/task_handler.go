package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"taskmaster/models"
	"taskmaster/services"
)

// registerUserHandler function
func RegisterUserHandler(db *sql.DB) http.HandlerFunc {

	// Create a new user struct
	var user models.User

	return func(w http.ResponseWriter, r *http.Request) {

		// Decode the request body into the user struct
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if r.Method == "POST" {

			// Check if the user already exists in the database
			var count int
			err = db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", user.Email).Scan(&count)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// If the user already exists, return an error
			if count > 0 {
				http.Error(w, "User already exists", http.StatusBadRequest)
				fmt.Println("User already exists", user)
				return
			} else {

				// hash the password
				user.Password, err = services.HashPassword(user.Password)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// Insert the user into the database
				_, err = db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", user.Email, user.Password)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// Print the user struct
				fmt.Fprint(w, "User: ", user.ID, " registered successfully!")
				fmt.Println("User created:", user)
			}
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}

}
