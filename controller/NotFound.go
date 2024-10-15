package controllers

import (
	"net/http"
	"text/template"
)

func NotFound(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)

    tmpl, err := template.ParseFiles("../html/404.html")
    if err != nil {
        http.Error(w, "Erreur lors du chargement de la page 404", http.StatusInternalServerError)
        return
    }
    
    err = tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, "Erreur lors de l'affichage de la page 404", http.StatusInternalServerError)
    }
}