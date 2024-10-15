package models

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
)

func Getcomment(id_post int, userID int)([]Comment){
	var comments []Comment

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return comments
	}

	query := "SELECT id, id_login, contenu FROM commentaire WHERE id_post = ?"
	rows, err := db.Query(query, id_post)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	svglike := "<svg class='icone-like' xmlns='http://www.w3.org/2000/svg' height='24' viewBox='0 -960 960 960' width='24'><path d='m480-120-58-52q-101-91-167-157T150-447.5Q111-500 95.5-544T80-634q0-94 63-157t157-63q52 0 99 22t81 62q34-40 81-62t99-22q94 0 157 63t63 157q0 46-15.5 90T810-447.5Q771-395 705-329T538-172l-58 52Zm0-108q96-86 158-147.5t98-107q36-45.5 50-81t14-70.5q0-60-40-100t-100-40q-47 0-87 26.5T518-680h-76q-15-41-55-67.5T300-774q-60 0-100 40t-40 100q0 35 14 70.5t50 81q36 45.5 98 107T480-228Zm0-273Z'/></svg>"

	svglikeactive := "<svg class='icone-like' xmlns='http://www.w3.org/2000/svg' height='24' viewBox='0 -960 960 960' width='24' > <path d='m480-120-58-52q-101-91-167-157T150-447.5Q111-500 95.5-544T80-634q0-94 63-157t157-63q52 0 99 22t81 62q34-40 81-62t99-22q94 0 157 63t63 157q0 46-15.5 90T810-447.5Q771-395 705-329T538-172l-58 52' stroke='red' fill='red' /></svg>"

	for rows.Next() {
		var comment Comment
		err := rows.Scan(&idcomment, &idloginpost, &comment.Content)
		if err != nil {
			log.Fatal(err)
		}

		name, surname, pseudo, urlpseudo := Getinfologinpost(idloginpost)

		likenumber := Getlikenumbercomment(idcomment)

		if Getlikecommentlogin(idcomment, userID){
			comment.Svglike = template.HTML(svglikeactive)
		} else {
			comment.Svglike = template.HTML(svglike)
		}
		comment.Name = name
		comment.Surname = surname
		comment.Idcomment = idcomment
		comment.Idlogin = idloginpost
		comment.Pseudo = pseudo
		comment.Urlpseudo = urlpseudo
		comment.Likenumber = likenumber

		comments = append(comments, comment)
	}
	return comments
}

func Getcommentnumber(id_post int)(int){
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return 0
	}

	number := "SELECT COUNT(*) FROM commentaire WHERE id_post = ?"
	var commentnumber int
	errNumber := db.QueryRow(number, id_post).Scan(&commentnumber)
	if errNumber != nil {
		fmt.Println("Error: Nombre de like indisponible", errNumber)
		return 0
	}
	defer db.Close()
	return commentnumber
}

func Getloginidcomment(id_comment int)int{
	var id int

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return id
	}

	// SQL Command password
	selectID := "SELECT id_login FROM commentaire WHERE id = ?"
	query := db.QueryRow(selectID, id_comment).Scan(&id)
	if query != nil {
		fmt.Println("Error querying password:", err)
		return id
	}

	defer db.Close()

	return id
}

func Getinfologincomment(id_login int)(string, string){
	var pseudo, urlpseudo string

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return pseudo, urlpseudo
	}

	query := "SELECT username, url FROM login WHERE id = ?"
	rows, err := db.Query(query, id_login)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&pseudo, &urlpseudo)
		if err != nil {
			log.Fatal(err)
		}
	}
	return pseudo, urlpseudo
}