package controllers

import (
	models "forum/model"
	"log"
	"net/http"
	"text/template"
)

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	userID := models.ReceiveCookie(r)
	post := models.Getpostlogin(userID, userID)
	categorie := models.Getcategorie()
	login := models.Getinfologin(userID)
	tendance := models.Gettendance()

	if r.Method == "POST" {
		recherche := r.FormValue("recherche")
		post=models.Getsearchpostlogin(recherche, userID, userID)
	}

	accountInfo := models.Data{
		Categories: categorie,
		Posts: post,
		Logins: login,
		Tendances: tendance,
	}

	if !(userID == 0) {
		tmpl, err := template.ParseFiles("../html/profil.html")
		if err != nil {
			log.Println("Erreur lors de la lecture du fichier HTML:", err)
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, accountInfo)
		if err != nil {
			log.Println("Erreur lors de l'exécution du modèle HTML:", err)
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	
}