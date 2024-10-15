package models

import (
	"database/sql"
	"fmt"
	"log"
)

func Gettendance()([]Tendance){
	var tendances []Tendance

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return tendances
	}

	rows, err := db.Query("SELECT c.id, c.contenu, c.url FROM categorie c JOIN post p ON c.id = p.id_categorie GROUP BY c.id ORDER BY COUNT(p.id) DESC LIMIT 3")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var tendance Tendance
		err := rows.Scan(&tendance.Idcategorie, &tendance.Name, &tendance.Urlcategorie)
		if err != nil {
			log.Fatal(err)
		}
		tendances = append(tendances, tendance)
	}
	return tendances
}