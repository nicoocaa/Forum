package models

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
)

func Getpost(userID int)([]Post){
	var posts []Post

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return posts
	}

	rows, err := db.Query("SELECT id, id_login, id_categorie, contenu, url FROM post ORDER BY unix DESC")
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

func Getpostcategorie(id_categorie int, userID int)([]Post){
	var posts []Post

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return posts
	}

	query := "SELECT id, id_login, contenu, url FROM post WHERE id_categorie = ? ORDER BY unix DESC"
	rows, err := db.Query(query, id_categorie)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	svglike := "<svg class='icone-like' xmlns='http://www.w3.org/2000/svg' height='24' viewBox='0 -960 960 960' width='24'><path d='m480-120-58-52q-101-91-167-157T150-447.5Q111-500 95.5-544T80-634q0-94 63-157t157-63q52 0 99 22t81 62q34-40 81-62t99-22q94 0 157 63t63 157q0 46-15.5 90T810-447.5Q771-395 705-329T538-172l-58 52Zm0-108q96-86 158-147.5t98-107q36-45.5 50-81t14-70.5q0-60-40-100t-100-40q-47 0-87 26.5T518-680h-76q-15-41-55-67.5T300-774q-60 0-100 40t-40 100q0 35 14 70.5t50 81q36 45.5 98 107T480-228Zm0-273Z'/></svg>"

	svglikeactive := "<svg class='icone-like' xmlns='http://www.w3.org/2000/svg' height='24' viewBox='0 -960 960 960' width='24' > <path d='m480-120-58-52q-101-91-167-157T150-447.5Q111-500 95.5-544T80-634q0-94 63-157t157-63q52 0 99 22t81 62q34-40 81-62t99-22q94 0 157 63t63 157q0 46-15.5 90T810-447.5Q771-395 705-329T538-172l-58 52' stroke='red' fill='red' /></svg>"

	for rows.Next() {
		var post Post
		err := rows.Scan(&idpost, &idloginpost, &post.Content, &post.Url)
		if err != nil {
			log.Fatal(err)
		}

		name, surname, pseudo, urlpseudo := Getinfologinpost(idloginpost)

		categorie, urlcategorie := Getinfocategorie(id_categorie)

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

func Getpostlogin(id_login int, userID int)([]Post){
	var posts []Post

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return posts
	}

	query := "SELECT id, id_categorie, contenu, url FROM post WHERE id_login = ? ORDER BY unix DESC"
	rows, err := db.Query(query, id_login)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	svglike := "<svg class='icone-like' xmlns='http://www.w3.org/2000/svg' height='24' viewBox='0 -960 960 960' width='24'><path d='m480-120-58-52q-101-91-167-157T150-447.5Q111-500 95.5-544T80-634q0-94 63-157t157-63q52 0 99 22t81 62q34-40 81-62t99-22q94 0 157 63t63 157q0 46-15.5 90T810-447.5Q771-395 705-329T538-172l-58 52Zm0-108q96-86 158-147.5t98-107q36-45.5 50-81t14-70.5q0-60-40-100t-100-40q-47 0-87 26.5T518-680h-76q-15-41-55-67.5T300-774q-60 0-100 40t-40 100q0 35 14 70.5t50 81q36 45.5 98 107T480-228Zm0-273Z'/></svg>"

	svglikeactive := "<svg class='icone-like' xmlns='http://www.w3.org/2000/svg' height='24' viewBox='0 -960 960 960' width='24' > <path d='m480-120-58-52q-101-91-167-157T150-447.5Q111-500 95.5-544T80-634q0-94 63-157t157-63q52 0 99 22t81 62q34-40 81-62t99-22q94 0 157 63t63 157q0 46-15.5 90T810-447.5Q771-395 705-329T538-172l-58 52' stroke='red' fill='red' /></svg>"

	for rows.Next() {
		var post Post
		err := rows.Scan(&idpost, &idcategorie, &post.Content, &post.Url)
		if err != nil {
			log.Fatal(err)
		}

		name, surname, pseudo, urlpseudo := Getinfologinpost(id_login)

		categorie, urlcategorie := Getinfocategorie(idcategorie)

		likenumber := Getlikenumberpost(idpost)

		if Getlikepostlogin(idpost,userID){
			post.Svglike = template.HTML(svglikeactive)
		} else {
			post.Svglike = template.HTML(svglike)
		}

		commentnumber := Getcommentnumber(idpost)

		post.Idpost = idpost
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

func Getpostidpost(id_post int, userID int)([]Post){
	var posts []Post

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return posts
	}

	query := "SELECT id_login, id_categorie, contenu, url FROM post WHERE id = ?"
	rows, err := db.Query(query, id_post)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	svglike := "<svg class='icone-like' xmlns='http://www.w3.org/2000/svg' height='24' viewBox='0 -960 960 960' width='24'><path d='m480-120-58-52q-101-91-167-157T150-447.5Q111-500 95.5-544T80-634q0-94 63-157t157-63q52 0 99 22t81 62q34-40 81-62t99-22q94 0 157 63t63 157q0 46-15.5 90T810-447.5Q771-395 705-329T538-172l-58 52Zm0-108q96-86 158-147.5t98-107q36-45.5 50-81t14-70.5q0-60-40-100t-100-40q-47 0-87 26.5T518-680h-76q-15-41-55-67.5T300-774q-60 0-100 40t-40 100q0 35 14 70.5t50 81q36 45.5 98 107T480-228Zm0-273Z'/></svg>"

	svglikeactive := "<svg class='icone-like' xmlns='http://www.w3.org/2000/svg' height='24' viewBox='0 -960 960 960' width='24' > <path d='m480-120-58-52q-101-91-167-157T150-447.5Q111-500 95.5-544T80-634q0-94 63-157t157-63q52 0 99 22t81 62q34-40 81-62t99-22q94 0 157 63t63 157q0 46-15.5 90T810-447.5Q771-395 705-329T538-172l-58 52' stroke='red' fill='red' /></svg>"

	for rows.Next() {
		var post Post
		err := rows.Scan(&idloginpost, &idcategorie, &post.Content, &post.Url)
		if err != nil {
			log.Fatal(err)
		}

		name, surname, pseudo, urlpseudo := Getinfologinpost(idloginpost)

		categorie, urlcategorie := Getinfocategorie(idcategorie)

		likenumber := Getlikenumberpost(id_post)

		if Getlikepostlogin(id_post,userID){
			post.Svglike = template.HTML(svglikeactive)
		} else {
			post.Svglike = template.HTML(svglike)
		}

		commentnumber := Getcommentnumber(id_post)

		post.Idpost = id_post
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

func Getinfologinpost(id_login int)(string, string, string, string){
	var name, surname, pseudo, urlpseudo string

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return name, surname, pseudo, urlpseudo
	}

	query := "SELECT name, surname, username, url FROM login WHERE id = ?"
	rows, err := db.Query(query, id_login)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&name, &surname, &pseudo, &urlpseudo)
		if err != nil {
			log.Fatal(err)
		}
	}
	return name, surname, pseudo, urlpseudo
}

func Getloginidpost(id_post int)int{
	var id int

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return id
	}

	// SQL Command password
	selectID := "SELECT id_login FROM post WHERE id = ?"
	query := db.QueryRow(selectID, id_post).Scan(&id)
	if query != nil {
		fmt.Println("Error querying password:", err)
		return id
	}

	defer db.Close()

	return id
}