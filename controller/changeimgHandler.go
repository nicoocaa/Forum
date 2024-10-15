package controllers

import (
	"database/sql"
	"fmt"
	models "forum/model"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func ChangeimgHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userID := models.ReceiveCookie(r)
		file, handler, err := r.FormFile("avatar")
		if err != nil {
			fmt.Println("Erreur en récupérant le fichier")
			fmt.Println(err)
			return
		}
		defer file.Close()
	
		ext := filepath.Ext(handler.Filename)

		filePath := "../src/pp/" + strconv.Itoa(userID) + ext
	
		dst, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Erreur en créant le fichier")
			fmt.Println(err)
			return
		}
		defer dst.Close()
	
		_, err = io.Copy(dst, file)
		if err != nil {
			fmt.Println("Erreur lors de la copie du fichier")
			fmt.Println(err)
			return
		}

		db, err := sql.Open("sqlite3", "forum.db")
		if err != nil {
			fmt.Println("Error : opening SQLite database:", err)
			return 
		}

		insertPost := "UPDATE login SET url = ? WHERE id = ?"
        _, err = db.Exec(insertPost, filePath, userID)
        if err != nil {
            fmt.Println("Error inserting into login table:", err)
            return
        }

		defer db.Close()
		http.Redirect(w, r, "/account", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}