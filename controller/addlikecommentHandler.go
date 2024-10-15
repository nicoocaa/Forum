package controllers

import (
	models "forum/model"
	"net/http"
	"strconv"
)

func AddlikecommentHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		action := r.FormValue("action")
		commentID, err := strconv.Atoi(action)
		if err != nil {
			http.Error(w, "ID post non valide", http.StatusBadRequest)
			return
		}
		userID := models.ReceiveCookie(r)
		postID := models.Addorremovecommentlike(commentID, userID)
		http.Redirect(w, r, "/post/"+strconv.Itoa(postID)+"#"+strconv.Itoa(commentID), http.StatusSeeOther)
	}
}