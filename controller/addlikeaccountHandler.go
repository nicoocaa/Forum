package controllers

import (
	models "forum/model"
	"net/http"
	"strconv"
)

func AddlikeaccountHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		action := r.FormValue("action")
		postID, err := strconv.Atoi(action)
		if err != nil {
			http.Error(w, "ID post non valide", http.StatusBadRequest)
			return
		}
		userID := models.ReceiveCookie(r)
		models.Addorremoveaccountlike(postID, userID)
		// Redirect to the login page
		http.Redirect(w, r, "/account#"+strconv.Itoa(postID), http.StatusSeeOther)
	}
}