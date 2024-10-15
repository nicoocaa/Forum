package controllers

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
        NotFound(w, r, http.StatusNotFound)
        return
    }

    if r.Method == "GET" {
		renderTemplate(w, "../html/register.html", nil)
    } else if r.Method == "POST" {
        // Receive values
		name := r.FormValue("name")
		surname := r.FormValue("surname")
        username := r.FormValue("username")
        email := r.FormValue("email")
        password := r.FormValue("password")

		db, err := sql.Open("sqlite3", "forum.db")
		if err != nil {
			// Debug SQL database
			fmt.Println("Error : opening SQLite database:", err)
			return 
		}

		var verifnumber int
		var randomNumber int
		var idnumber int
		var errNumber error

		exist := "SELECT COUNT(*) FROM login WHERE username = ?"
		errNumber = db.QueryRow(exist, username).Scan(&verifnumber)
		if errNumber != nil {
			fmt.Println("Erreur lors de la vérification de l'ID:", errNumber)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	
		if verifnumber == 0 {

			// Hashed password
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				fmt.Println("Error generating hashed password:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			for {
				randomNumber = rand.Intn(9000) + 1000
		
				number := "SELECT COUNT(*) FROM login WHERE id = ?"
				errNumber = db.QueryRow(number, randomNumber).Scan(&idnumber)
				if errNumber != nil {
					fmt.Println("Erreur lors de la vérification de l'ID:", errNumber)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
		
				if idnumber == 0 {
					break
				}
			}

			// SQL Command register
			_, err = db.Exec("INSERT INTO login (id, url, name, surname, username, adresse_mail, password) VALUES (?, ?, ?, ?, ?, ?, ?)", randomNumber, "../src/pp/0.webp", name, surname, username, email, hashedPassword)
			if err != nil {
				fmt.Println("Error inserting into login table:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			defer db.Close()

			// Log hashed password
			fmt.Println("Hashed Password:", hashedPassword)

			// SQL Command verification
			verif := "SELECT COUNT(*) FROM login WHERE username = ?"
			var verifusername int
			errVerif := db.QueryRow(verif, username).Scan(&verifusername)
			if errVerif != nil {
				fmt.Println("Error: Verification impossible", errVerif)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			defer db.Close()

			// Verif Register
			if verifusername > 0 {
				fmt.Println("Le compte a bien été créé")
			} else {
				fmt.Println("Error: Le compte n'a pas été créé")
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// Redirect to the login page
			http.Redirect(w, r, "/login", http.StatusSeeOther)

		} else {
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			fmt.Println("l'username existe déja")
		}
    }
}