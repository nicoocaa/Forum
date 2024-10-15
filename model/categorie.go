package models

import (
	"database/sql"
	"fmt"
	"log"
)

func Getcategorie()([]Categorie){
	var categories []Categorie

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return categories
	}

	rows, err := db.Query("SELECT id, contenu, url FROM categorie")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var categorie Categorie
		err := rows.Scan(&categorie.Idcategorie, &categorie.Name, &categorie.Urlcategorie)
		if err != nil {
			log.Fatal(err)
		}
		categories = append(categories, categorie)
	}
	return categories
}

func Getinfocategorie(id_categorie int)(string, string){
	var categorie, urlcategorie string

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return categorie, urlcategorie
	}

	query := "SELECT contenu, url FROM categorie WHERE id = ?"
	rows, err := db.Query(query, id_categorie)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&categorie, &urlcategorie)
		if err != nil {
			log.Fatal(err)
		}
	}
	return categorie, urlcategorie
}