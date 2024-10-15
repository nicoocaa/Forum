package controllers

import (
	"database/sql"
	"fmt"
	models "forum/model"
	"net/http"
	"strconv"
)

func RemovecommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
			// Receive values
			action := r.FormValue("action")
			commentID, err := strconv.Atoi(action)
			if err != nil {
				fmt.Println("Error converting commentID to integer:", err)
				return
			}
			action2 := r.FormValue("action2")
			postID, err := strconv.Atoi(action2)
			if err != nil {
				fmt.Println("Error converting postID to integer:", err)
				return
			}
			userID := models.ReceiveCookie(r)
			if (models.Getloginidcomment(commentID) == userID){
	
				db, err := sql.Open("sqlite3", "forum.db")
				if err != nil {
					// Debug SQL database
					fmt.Println("Error : opening SQLite database:", err)
					return
				}
	
				// SQL Command Delete Like Commentaire
				deleteLikeComment := "DELETE FROM like WHERE id_commentaire IN ( SELECT id FROM commentaire WHERE id_post = ? )"
				_, err = db.Exec(deleteLikeComment, postID)
				if err != nil {
					fmt.Println("Error delete into post table:", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
	
				defer db.Close()
	
				// SQL Command DELETE
				deleteComment := "DELETE FROM commentaire WHERE id = ?"
				_, err = db.Exec(deleteComment, commentID)
				if err != nil {
					fmt.Println("Error delete into commentaire table:", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
	
				defer db.Close()
			}
			// Redirect to the login page
			http.Redirect(w, r, "/post/"+strconv.Itoa(postID), http.StatusSeeOther)
		}
	}