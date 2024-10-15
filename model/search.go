package models

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
)

func Getsearch(keyword string, userID int)([]Post) {
	var posts []Post

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return posts
	}

	rows, err := db.Query("SELECT id, id_login, id_categorie, contenu, url FROM post WHERE contenu LIKE ?", "%"+keyword+"%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	svglike := "<svg class='icone-like' xmlns='http://www.w3.org/2000/svg' height='24' viewBox='0 -960 960 960' width='24'><path d='m480-120-58-52q-101-91-167-157T150-447.5Q111-500 95.5-544T80-634q0-94 63-157t157-63q52 0 99 22t81 62q34-40 81-62t99-22q94 0 157 63t63 157q0 46-15.5 90T810-447.5Q771-395 705-329T538-172l-58 52Zm0-108q96-86 158-147.5t98-107q36-45.5 50-81t14-70.5q0-60-40-100t-100-40q-47 0-87 26.5T518-680h-76q-15-41-55-67.5T300-774q-60 0-100 40t-40 100q0 35 14 70.5t50 81q36 45.5 98 107T480-228Zm0-273Z'/></svg>"

	svglikeactive := "<svg class='icone-like' xmlns='http://www.w3.org/2000/svg' height='24' viewBox='0 -960 960 960' width='24' > <path d='m480-120-58-52q-101-91-167-157T150-447.5Q111-500 95.5-544T80-634q0-94 63-157t157-63q52 0 99 22t81 62q34-40 81-62t99-22q94 0 157 63t63 157q0 46-15.5 90T810-447.5Q771-395 705-329T538-172l-58 52' stroke='red' fill='red' /></svg>"

	for rows.Next() {
		var post Post
		err := rows.Scan(&idpost, &idloginpost, &idcategorie, &post.Content, &post.Url)
		if err != nil {
			log.Fatal(err)
		}

		name, surname, pseudo, urlpseudo := Getinfologinpost(idloginpost)

		categorie, urlcategorie := Getinfocategorie(idcategorie)

		likenumber := Getlikenumberpost(idpost)

		if Getlikepostlogin(idpost,userID){
			post.Svglike = template.HTML(svglikeactive)
		} else {
			post.Svglike = template.HTML(svglike)
		}

		commentnumber := Getcommentnumber(idpost)

		post.Idpost = idpost
		post.Idlogin = idloginpost
		post.Name = name
		post.Surname = surname
		post.Pseudo = pseudo
		post.Urlpseudo = urlpseudo
		post.Categorie = categorie
		post.Urlcategorie = template.HTML(urlcategorie)
		post.Likenumber = likenumber
		post.Commentnumber = commentnumber

		posts = append(posts, post)
	}
	return posts
}

func Getsearchpostlogin(keyword string, id_login int, userID int)([]Post) {
	var posts []Post

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return posts
	}

	rows, err := db.Query("SELECT id, id_login, id_categorie, contenu, url FROM post WHERE id_login = ? AND contenu LIKE ?", id_login, "%"+keyword+"%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	svglike := "<svg class='icone-like' xmlns='http://www.w3.org/2000/svg' height='24' viewBox='0 -960 960 960' width='24'><path d='m480-120-58-52q-101-91-167-157T150-447.5Q111-500 95.5-544T80-634q0-94 63-157t157-63q52 0 99 22t81 62q34-40 81-62t99-22q94 0 157 63t63 157q0 46-15.5 90T810-447.5Q771-395 705-329T538-172l-58 52Zm0-108q96-86 158-147.5t98-107q36-45.5 50-81t14-70.5q0-60-40-100t-100-40q-47 0-87 26.5T518-680h-76q-15-41-55-67.5T300-774q-60 0-100 40t-40 100q0 35 14 70.5t50 81q36 45.5 98 107T480-228Zm0-273Z'/></svg>"

	svglikeactive := "<svg class='icone-like' xmlns='http://www.w3.org/2000/svg' height='24' viewBox='0 -960 960 960' width='24' > <path d='m480-120-58-52q-101-91-167-157T150-447.5Q111-500 95.5-544T80-634q0-94 63-157t157-63q52 0 99 22t81 62q34-40 81-62t99-22q94 0 157 63t63 157q0 46-15.5 90T810-447.5Q771-395 705-329T538-172l-58 52' stroke='red' fill='red' /></svg>"

	for rows.Next() {
		var post Post
		err := rows.Scan(&idpost, &idloginpost, &idcategorie, &post.Content, &post.Url)
		if err != nil {
			log.Fatal(err)
		}

		name, surname, pseudo, urlpseudo := Getinfologinpost(idloginpost)

		categorie, urlcategorie := Getinfocategorie(idcategorie)

		likenumber := Getlikenumberpost(idpost)

		if Getlikepostlogin(idpost,userID){
			post.Svglike = template.HTML(svglikeactive)
		} else {
			post.Svglike = template.HTML(svglike)
		}

		commentnumber := Getcommentnumber(idpost)

		post.Idpost = idpost
		post.Idlogin = idloginpost
		post.Name = name
		post.Surname = surname
		post.Pseudo = pseudo
		post.Urlpseudo = urlpseudo
		post.Categorie = categorie
		post.Urlcategorie = template.HTML(urlcategorie)
		post.Likenumber = likenumber
		post.Commentnumber = commentnumber

		posts = append(posts, post)
	}
	return posts
}

func Getsearchuser(keyword string)([]Explorer) {
	var explorers []Explorer

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return explorers
	}

	rows, err := db.Query("SELECT id, name, surname, username, url FROM login WHERE username LIKE ?", "%"+keyword+"%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var explorer Explorer
		err := rows.Scan(&explorer.Id, &explorer.Name, &explorer.Surname, &explorer.Pseudo, &explorer.Urlpseudo)
		if err != nil {
			log.Fatal(err)
		}

		explorers = append(explorers, explorer)
	}
	return explorers
}