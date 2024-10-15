package controllers

import (
	"database/sql"
	"fmt"
	models "forum/model"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
        NotFound(w, r, http.StatusNotFound)
        return
    }

	if r.Method == "GET" {
		renderTemplate(w, "../html/login.html", nil)
    } else if r.Method == "POST" {
        username := r.FormValue("username")
        password := r.FormValue("password")

		db, err := sql.Open("sqlite3", "forum.db")
		if err != nil {
			// Debug SQL database
			fmt.Println("Error : opening SQLite database:", err)
			return
		}

		// SQL Command password
		selectPassword := "SELECT password FROM login WHERE username = ?"
		var hashedPassword string
		query := db.QueryRow(selectPassword, username).Scan(&hashedPassword)
		if query != nil {
			fmt.Println("Error querying password:", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		defer db.Close()

		// Compare password
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if err != nil {
			fmt.Println("Invalid username or password:", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// SQL Command ID
		selectID := "SELECT id FROM login WHERE username = ?"
		var userID int
		errID := db.QueryRow(selectID, username).Scan(&userID)
		if errID != nil {
			fmt.Println("Error querying ID:", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		defer db.Close()

		// Debug UserID
		fmt.Println("User ID:", userID)

		// Create session
		models.CreateCookie(w, r, userID)

		http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}