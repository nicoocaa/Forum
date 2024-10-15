package controllers

import (
	"database/sql"
	"fmt"
	models "forum/model"
	"math/rand"
	"net/http"
	"strconv"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
        // Receive values
		userID := models.ReceiveCookie(r)
        content := r.FormValue("content")
		action := r.FormValue("action")
		postID, err := strconv.Atoi(action)
		if err != nil {
			http.Error(w, "ID post non valide", http.StatusBadRequest)
			return
		}

		db, err := sql.Open("sqlite3", "forum.db")
		if err != nil {
			// Debug SQL database
			fmt.Println("Error : opening SQLite database:", err)
			return
		}

		var randomNumber int
		var idnumber int
		var errNumber error

		for {
			randomNumber = rand.Intn(9000) + 1000
	
			number := "SELECT COUNT(*) FROM commentaire WHERE id = ?"
			errNumber = db.QueryRow(number, randomNumber).Scan(&idnumber)
			if errNumber != nil {
				fmt.Println("Erreur lors de la vérification de l'ID:", errNumber)
				return
			}
	
			if idnumber == 0 {
				break
			}
		}

        // SQL Command addcomment
        insertComment := "INSERT INTO commentaire (id, id_login, id_post, contenu) VALUES (?, ?, ?, ?)"
        _, err = db.Exec(insertComment, randomNumber, userID, postID, content)
        if err != nil {
            fmt.Println("Error inserting into login table:", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

		defer db.Close()
		// Redirect to the login page
		http.Redirect(w, r, "/post/"+strconv.Itoa(postID), http.StatusSeeOther)
	}
}