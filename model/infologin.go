package models

import (
	"database/sql"
	"fmt"
	"log"
)

func Getinfologin(userID int)(Login){
	var logins Login

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return logins
	}

	query := "SELECT name, surname, username, url, adresse_mail FROM login WHERE id = ?"
	rows, err := db.Query(query, userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var login Login
		err := rows.Scan(&login.Name, &login.Surname, &login.Pseudo, &login.Urlpseudo, &login.Adressemail)
		if err != nil {
			log.Fatal(err)
		}
		login.Id = userID
		logins = login
	}

	return logins
}