package controllers

import (
	"database/sql"
	"fmt"
	models "forum/model"
	"net/http"
	"strconv"
)

func EditpostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Receive values
		content := r.FormValue("content")
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

			// SQL Command Delete Post
			editPost := "UPDATE post SET contenu = ? WHERE id = ?"
			_, err = db.Exec(editPost, content, postID)
			if err != nil {
				fmt.Println("Error UPDATE into post table:", err)
				return
			}

			defer db.Close()

		}
		// Redirect to the login page
		http.Redirect(w, r, "/post/"+strconv.Itoa(postID), http.StatusSeeOther)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}