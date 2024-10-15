package controllers

import (
	models "forum/model"
	"log"
	"net/http"
	"text/template"
)

func ExplorerHandler(w http.ResponseWriter, r *http.Request) {
	userID := models.ReceiveCookie(r)
	categorie := models.Getcategorie()
	login := models.Getinfologin(userID)
	tendance := models.Gettendance()
	explorer := models.Getuser()

	if r.Method == "POST" {
		recherche := r.FormValue("recherche")
		explorer = models.Getsearchuser(recherche)
	}

	accountInfo := models.Data{
		Categories: categorie,
		Logins: login,
		Explorers : explorer,
		Tendances: tendance,
	}

	tmpl, err := template.ParseFiles("../html/user.html")
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
}