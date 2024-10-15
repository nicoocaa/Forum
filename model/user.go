package models

import (
	"database/sql"
	"fmt"
	"log"
)

func Getuser()([]Explorer){
	var explorers []Explorer

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return explorers
	}

	rows, err := db.Query("SELECT id, name, surname, username, url FROM login")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var explorer Explorer
		err := rows.Scan(&explorer.Id ,&explorer.Name, &explorer.Surname, &explorer.Pseudo, &explorer.Urlpseudo)
		if err != nil {
			log.Fatal(err)
		}

		explorers = append(explorers, explorer)
	}
	return explorers
}

func Getinfouser(id int)(User){
	var users User

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return users
	}

	query := "SELECT name, surname, username, url, adresse_mail FROM login WHERE id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Name, &user.Surname, &user.Pseudo, &user.Urlpseudo, &user.Adressemail)
		if err != nil {
			log.Fatal(err)
		}
		user.Id = id
		users = user
	}

	return users
}