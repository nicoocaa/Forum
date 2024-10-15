package controllers

import (
	models "forum/model"
	"net/http"
	"strconv"
)

func AddlikeHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		action := r.FormValue("action")
		postID, err := strconv.Atoi(action)
		if err != nil {
			http.Error(w, "ID post non valide", http.StatusBadRequest)
			return
		}
		userID := models.ReceiveCookie(r)
		models.Addorremovelike(postID,userID)
		// Redirect to the login page
		http.Redirect(w, r, "/#"+strconv.Itoa(postID), http.StatusSeeOther)
	}
}