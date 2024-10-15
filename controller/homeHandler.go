package controllers

import (
	models "forum/model"
	"log"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
        NotFound(w, r, http.StatusNotFound)
        return
    }

	userID := models.ReceiveCookie(r)
	post := models.Getpost(userID)
	categorie := models.Getcategorie()
	login := models.Getinfologin(userID)
	tendance := models.Gettendance()

	if r.Method == "POST" {
		recherche := r.FormValue("recherche")
		post=models.Getsearch(recherche,userID)
	}

	postInfo := models.Data{
		Categories: categorie,
		Posts: post,
		Logins: login,
		Tendances: tendance,
	}

	tmpl, err := template.ParseFiles("../html/index.html")
	if err != nil {
		log.Println("Erreur lors de la lecture du fichier HTML:", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, postInfo)
	if err != nil {
		log.Println("Erreur lors de l'exécution du modèle HTML:", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
}