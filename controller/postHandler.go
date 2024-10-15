package controllers

import (
	"fmt"
	models "forum/model"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/post/"):]
	postID, err := strconv.Atoi(id)
	if err != nil {
		NotFound(w, r, http.StatusNotFound)
		fmt.Println("id invalide")
		http.Error(w, "ID post non valide", http.StatusBadRequest)
        return
	}
	userID := models.ReceiveCookie(r)
	post := models.Getpostidpost(postID,userID)
	categorie := models.Getcategorie()
	login := models.Getinfologin(userID)
	comment := models.Getcomment(postID, userID)
	tendance := models.Gettendance()

	postInfo := models.Data{
		Categories: categorie,
		Posts: post,
		Logins: login,
		Comments: comment,
		Tendances: tendance,
	}

	tmpl, err := template.ParseFiles("../html/post.html")
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