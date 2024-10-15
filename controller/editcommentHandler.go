package controllers

import (
	"database/sql"
	"fmt"
	models "forum/model"
	"net/http"
	"strconv"
)

func EditcommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Receive values
		content := r.FormValue("content")
		action := r.FormValue("action")
		commentID, err := strconv.Atoi(action)
		if err != nil {
			fmt.Println("Error converting postID to integer:", err)
			return
		}
		userID := models.ReceiveCookie(r)

		db, err := sql.Open("sqlite3", "forum.db")
		if err != nil {
			// Debug SQL database
			fmt.Println("Error : opening SQLite database:", err)
			return
		}
		
		selectpostID := "SELECT id_post FROM commentaire WHERE id = ?"
		var postID int
		query := db.QueryRow(selectpostID, commentID).Scan(&postID)
		if query != nil {
			fmt.Println("Error querying postID:", err)
			return
		}
		defer db.Close()
		
		if (models.Getloginidcomment(commentID) == userID){

			// SQL Command Delete Post
			editComment := "UPDATE commentaire SET contenu = ? WHERE id = ?"
			_, err = db.Exec(editComment, content, commentID)
			if err != nil {
				fmt.Println("Error delete into post table:", err)
				return
			}

			defer db.Close()

		}
		// Redirect to the login page
		http.Redirect(w, r, "/post/"+strconv.Itoa(postID)+"#"+strconv.Itoa(commentID), http.StatusSeeOther)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}