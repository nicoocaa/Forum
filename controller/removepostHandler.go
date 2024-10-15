package controllers

import (
	"database/sql"
	"fmt"
	models "forum/model"
	"log"
	"net/http"
	"os"
	"strconv"
)

func RemovepostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
        // Receive values
		action := r.FormValue("action")
		postID, err := strconv.Atoi(action)
		if err != nil {
			fmt.Println("Error converting postID to integer:", err)
			return
		}
		userID := models.ReceiveCookie(r)
		if (models.Getloginidpost(postID) == userID){

			db, err := sql.Open("sqlite3", "forum.db")
			if err != nil {
				// Debug SQL database
				fmt.Println("Error : opening SQLite database:", err)
				return
			}

			var filePath string

			pathfile := "SELECT url FROM post WHERE id = ?"
			errNumber := db.QueryRow(pathfile, postID).Scan(&filePath)
			if errNumber != nil {
				fmt.Println("Erreur lors de la vérification de l'ID:", errNumber)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if filePath != "" {
				// Suppression de l'image
				errpath := os.Remove(filePath)
				if errpath != nil {
					log.Fatal(errpath)
				}
			}

			// SQL Command Delete Post
			deletePost := "DELETE FROM post WHERE id = ?"
			_, err = db.Exec(deletePost, postID)
			if err != nil {
				fmt.Println("Error delete into post table:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			defer db.Close()

			// SQL Command Delete Like Commentaire
			deleteLikeComment := "DELETE FROM like WHERE id_commentaire IN ( SELECT id FROM commentaire WHERE id_post = ? )"
			_, err = db.Exec(deleteLikeComment, postID)
			if err != nil {
				fmt.Println("Error delete into post table:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			defer db.Close()

			// SQL Command Delete Commentaire
			deleteComment := "DELETE FROM commentaire WHERE id_post = ?"
			_, err = db.Exec(deleteComment, postID)
			if err != nil {
				fmt.Println("Error delete into post table:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			defer db.Close()

			// SQL Command Delete Like Post
			deleteLike := "DELETE FROM like WHERE id_post = ?"
			_, err = db.Exec(deleteLike, postID)
			if err != nil {
				fmt.Println("Error delete into post table:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			defer db.Close()
			
		}
		// Redirect to the login page
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}