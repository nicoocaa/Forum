package models

import (
	"database/sql"
	"fmt"
	"math/rand"
)

func Addorremovepostlike(id_post int, userID int){
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return
	}
	if Getlikepostlogin(id_post,userID){
		id := "SELECT id FROM like WHERE id_login = ? AND id_post = ?"
		errNumber := db.QueryRow(id, userID, id_post).Scan(&idlike)
		if errNumber != nil {
			fmt.Println("Error: Nombre de like indisponible", errNumber)
			return
		}
		defer db.Close()
		// SQL Command DELETE
		_, err = db.Exec("DELETE FROM like WHERE id = ?", idlike)
		if err != nil {
			fmt.Println("Error inserting into like table:", err)
			return
		}
		defer db.Close()
	} else {
		var randomNumber int
		var idnumber int
		var errNumber error

		for {
			randomNumber = rand.Intn(9000) + 1000
	
			number := "SELECT COUNT(*) FROM like WHERE id = ?"
			errNumber = db.QueryRow(number, randomNumber).Scan(&idnumber)
			if errNumber != nil {
				fmt.Println("Erreur lors de la vérification de l'ID:", errNumber)
				return
			}
	
			if idnumber == 0 {
				break
			}
		}
		// SQL Command INSERT
		_, err = db.Exec("INSERT INTO like (id, id_login, id_post, id_commentaire) VALUES (?, ?, ?, ?)", randomNumber,userID, id_post, 0)
		if err != nil {
			fmt.Println("Error inserting into like table:", err)
			return
		}
		defer db.Close()
	}
}

func Addorremovelike(id_post int, userID int){
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return
	}
	if Getlikepostlogin(id_post,userID){
		id := "SELECT id FROM like WHERE id_login = ? AND id_post = ?"
		errNumber := db.QueryRow(id, userID, id_post).Scan(&idlike)
		if errNumber != nil {
			fmt.Println("Error: Nombre de like indisponible", errNumber)
			return
		}
		defer db.Close()
		// SQL Command DELETE
		_, err = db.Exec("DELETE FROM like WHERE id = ?", idlike)
		if err != nil {
			fmt.Println("Error inserting into like table:", err)
			return
		}
		defer db.Close()
		id_post=0
	} else {
		var randomNumber int
		var idnumber int
		var errNumber error

		for {
			randomNumber = rand.Intn(9000) + 1000
	
			number := "SELECT COUNT(*) FROM like WHERE id = ?"
			errNumber = db.QueryRow(number, randomNumber).Scan(&idnumber)
			if errNumber != nil {
				fmt.Println("Erreur lors de la vérification de l'ID:", errNumber)
				return
			}
	
			if idnumber == 0 {
				break
			}
		}
		// SQL Command INSERT
		_, err = db.Exec("INSERT INTO like (id, id_login, id_post, id_commentaire) VALUES (?, ?, ?, ?)", randomNumber, userID, id_post, 0)
		if err != nil {
			fmt.Println("Error inserting into like table:", err)
			return
		}
		defer db.Close()
		id_post=0
	}
}

func Addorremoveaccountlike(id_post int, userID int){
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return
	}
	if Getlikepostlogin(id_post,userID){
		id := "SELECT id FROM like WHERE id_login = ? AND id_post = ?"
		errNumber := db.QueryRow(id, userID, id_post).Scan(&idlike)
		if errNumber != nil {
			fmt.Println("Error: Nombre de like indisponible", errNumber)
			return
		}
		defer db.Close()
		// SQL Command DELETE
		_, err = db.Exec("DELETE FROM like WHERE id = ?", idlike)
		if err != nil {
			fmt.Println("Error inserting into like table:", err)
			return
		}
		defer db.Close()
	} else {
		var randomNumber int
		var idnumber int
		var errNumber error

		for {
			randomNumber = rand.Intn(9000) + 1000
	
			number := "SELECT COUNT(*) FROM like WHERE id = ?"
			errNumber = db.QueryRow(number, randomNumber).Scan(&idnumber)
			if errNumber != nil {
				fmt.Println("Erreur lors de la vérification de l'ID:", errNumber)
				return
			}
	
			if idnumber == 0 {
				break
			}
		}
		// SQL Command INSERT
		_, err = db.Exec("INSERT INTO like (id, id_login, id_post, id_commentaire) VALUES (?, ?, ?, ?)", randomNumber, userID, id_post, 0)
		if err != nil {
			fmt.Println("Error inserting into like table:", err)
			return
		}
		defer db.Close()
	}
}

func Addorremoveaccountslike(id_post int, userID int){
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return
	}
	if Getlikepostlogin(id_post,userID){
		id := "SELECT id FROM like WHERE id_login = ? AND id_post = ?"
		errNumber := db.QueryRow(id, userID, id_post).Scan(&idlike)
		if errNumber != nil {
			fmt.Println("Error: Nombre de like indisponible", errNumber)
			return
		}
		defer db.Close()
		// SQL Command DELETE
		_, err = db.Exec("DELETE FROM like WHERE id = ?", idlike)
		if err != nil {
			fmt.Println("Error inserting into like table:", err)
			return
		}
		defer db.Close()
	} else {
		var randomNumber int
		var idnumber int
		var errNumber error

		for {
			randomNumber = rand.Intn(9000) + 1000
	
			number := "SELECT COUNT(*) FROM like WHERE id = ?"
			errNumber = db.QueryRow(number, randomNumber).Scan(&idnumber)
			if errNumber != nil {
				fmt.Println("Erreur lors de la vérification de l'ID:", errNumber)
				return
			}
	
			if idnumber == 0 {
				break
			}
		}
		// SQL Command INSERT
		_, err = db.Exec("INSERT INTO like (id, id_login, id_post, id_commentaire) VALUES (?, ?, ?, ?)", randomNumber, userID, id_post, 0)
		if err != nil {
			fmt.Println("Error inserting into like table:", err)
			return
		}
		defer db.Close()
	}
}

func Addorremovecommentlike(id_comment int, userID int)int{
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
	}
	if Getlikecommentlogin(id_comment, userID){
		id := "SELECT id FROM like WHERE id_login = ? AND id_commentaire = ?"
		errNumber := db.QueryRow(id, userID, id_comment).Scan(&idlike)
		if errNumber != nil {
			fmt.Println("Error: Nombre de like indisponible", errNumber)
		}
		defer db.Close()
		// SQL Command DELETE
		_, err = db.Exec("DELETE FROM like WHERE id = ?", idlike)
		if err != nil {
			fmt.Println("Error inserting into like table:", err)
		}
		defer db.Close()
	} else {
		var randomNumber int
		var idnumber int
		var errNumber error

		for {
			randomNumber = rand.Intn(9000) + 1000
	
			number := "SELECT COUNT(*) FROM like WHERE id = ?"
			errNumber = db.QueryRow(number, randomNumber).Scan(&idnumber)
			if errNumber != nil {
				fmt.Println("Erreur lors de la vérification de l'ID:", errNumber)
				return 0
			}
	
			if idnumber == 0 {
				break
			}
		}
		// SQL Command INSERT
		_, err = db.Exec("INSERT INTO like (id, id_login, id_post, id_commentaire) VALUES (?, ?, ?, ?)", randomNumber, userID, 0, id_comment)
		if err != nil {
			fmt.Println("Error inserting into like table:", err)
		}
		defer db.Close()
	}
	var postID int
	// SQL Command SELECT
	id := "SELECT id_post FROM commentaire WHERE id = ?"
		errNumber := db.QueryRow(id, id_comment).Scan(&postID)
		if errNumber != nil {
			fmt.Println("Error: Nombre de like indisponible", errNumber)
		}
		defer db.Close()
	return postID
}

func Getlikenumberpost(id_post int)(int){
	var likenumber int

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return likenumber
	}

	number := "SELECT COUNT(*) FROM like WHERE id_post = ?"
	errNumber := db.QueryRow(number, id_post).Scan(&likenumber)
	if errNumber != nil {
		fmt.Println("Error: Nombre de like indisponible", errNumber)
		return 0
	}
	defer db.Close()
	return likenumber
}

func Getlikepostlogin(id_post int, userID int)(bool){
	var like bool
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return like
	}

	number := "SELECT COUNT(*) FROM like WHERE id_post = ? AND id_login = ?"
	var likenumber int
	errNumber := db.QueryRow(number, id_post, userID).Scan(&likenumber)
	if errNumber != nil {
		fmt.Println("Error: Nombre de like indisponible", errNumber)
		return false
	}
	defer db.Close()
	if likenumber > 0 {
		like = true
	} else {
		like = false
	}
	return like
}

func Getlikenumbercomment(id_commentaire int)(int){
	var likenumber int

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return likenumber
	}

	number := "SELECT COUNT(*) FROM like WHERE id_commentaire = ?"
	errNumber := db.QueryRow(number, id_commentaire).Scan(&likenumber)
	if errNumber != nil {
		fmt.Println("Error: Nombre de like indisponible", errNumber)
		return 0
	}
	defer db.Close()

	return likenumber
}

func Getlikecommentlogin(id_commentaire int, userID int)(bool){
	var like bool

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return like
	}

	number := "SELECT COUNT(*) FROM like WHERE id_commentaire = ? AND id_login = ?"
	var likenumber int
	errNumber := db.QueryRow(number, id_commentaire, userID).Scan(&likenumber)
	if errNumber != nil {
		fmt.Println("Error: Nombre de like indisponible", errNumber)
		return false
	}
	defer db.Close()
	if likenumber > 0 {
		like = true
	} else {
		like = false
	}
	return like
}