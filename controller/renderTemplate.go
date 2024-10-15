package controllers

import (
	"fmt"
	"net/http"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmplPath string, data interface{}) {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}