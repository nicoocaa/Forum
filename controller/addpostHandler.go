package controllers

import (
	"database/sql"
	"fmt"
	models "forum/model"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func AddpostHandler(w http.ResponseWriter, r *http.Request) {
	// Receive userID
	userID := models.ReceiveCookie(r)
	
	if r.Method == "POST" {
        // Receive values
        contenu := r.FormValue("content")
        categorie := r.FormValue("categorie")

		db, err := sql.Open("sqlite3", "forum.db")
		if err != nil {
			// Debug SQL database
			fmt.Println("Error : opening SQLite database:", err)
			return 
		}

		// SQL Command ID Post
		numberPost := "SELECT id FROM post ORDER BY id DESC LIMIT 1"
		var Idpost int
		errPost := db.QueryRow(numberPost).Scan(&Idpost)
		if errPost != nil {
			fmt.Println("Error querying ID:", errPost)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		defer db.Close()

		var filePath string

        file, handler, err := r.FormFile("image")
        if err != nil {
            if err != http.ErrMissingFile {
                fmt.Println("Erreur en récupérant le fichier:", err)
                return
            }
			filePath = ""
        } else {
            defer file.Close()

            ext := filepath.Ext(handler.Filename)
			Idpostnow := Idpost + 1
            filePath = "../src/uploads/" + strconv.Itoa(Idpostnow) + ext

            dst, err := os.Create(filePath)
            if err != nil {
                fmt.Println("Erreur en créant le fichier:", err)
                return
            }
            defer dst.Close()

            _, err = io.Copy(dst, file)
            if err != nil {
                fmt.Println("Erreur lors de la copie du fichier:", err)
                return
            }
        }

		rand.Seed(time.Now().UnixNano())

		var randomNumber int
		var idnumber int
		var errNumber error

		for {
			randomNumber = rand.Intn(9000) + 1000
	
			number := "SELECT COUNT(*) FROM post WHERE id = ?"
			errNumber = db.QueryRow(number, randomNumber).Scan(&idnumber)
			if errNumber != nil {
				fmt.Println("Erreur lors de la vérification de l'ID:", errNumber)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
	
			if idnumber == 0 {
				break
			}
		}

		now := time.Now()
		unixTime := now.Unix()

		// SQL Command addpost
		insertPost := "INSERT INTO post (id, id_login, contenu, id_categorie, url, unix) VALUES (?, ?, ?, ?, ?, ?)"
		_, err = db.Exec(insertPost, randomNumber, userID, contenu, categorie, filePath, unixTime)
		if err != nil {
			fmt.Println("Error inserting into login table:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	
		defer db.Close()

		// Redirect to the login page
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}