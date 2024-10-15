package controllers

import (
	models "forum/model"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func AccountsHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/accounts/"):]
	accountID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID post non valide", http.StatusBadRequest)
		return
	}
	userID := models.ReceiveCookie(r)
	if accountID == userID {
		http.Redirect(w, r, "/account", http.StatusSeeOther)
	}
	post := models.Getpostlogin(accountID, userID)
	categorie := models.Getcategorie()
	login := models.Getinfologin(userID)
	user := models.Getinfouser(accountID)
	tendance := models.Gettendance()

	if r.Method == "POST" {
		recherche := r.FormValue("recherche")
		post=models.Getsearchpostlogin(recherche, accountID, userID)
	}

	accountInfo := models.Data{
		Categories: categorie,
		Posts: post,
		Users: user,
		Logins: login,
		Tendances: tendance,
	}

	tmpl, err := template.ParseFiles("../html/profils.html")
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