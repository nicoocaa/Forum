package controllers

import (
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		redirect := r.FormValue("action")
		cookie := http.Cookie{
			Name:   "userID",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, redirect, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}