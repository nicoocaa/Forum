package main

import (
	"database/sql"
	"fmt"
	controllers "forum/controller"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Debug SQL database
		fmt.Println("Error : opening SQLite database:", err)
		return
	}
	defer db.Close()
	
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../css/"))))
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("../html/"))))
	http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("../src/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("../js/"))))

	http.HandleFunc("/", controllers.HomeHandler)
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/register", controllers.RegisterHandler)
	http.HandleFunc("/logout", controllers.LogoutHandler)
	http.HandleFunc("/account", controllers.AccountHandler)
	http.HandleFunc("/explorer", controllers.ExplorerHandler)
	http.HandleFunc("/accounts/", controllers.AccountsHandler)
	http.HandleFunc("/changeimg", controllers.ChangeimgHandler)
	http.HandleFunc("/post/", controllers.PostHandler)
	http.HandleFunc("/categorie/", controllers.CategorieHandler)
	http.HandleFunc("/addlike", controllers.AddlikeHandler)
	http.HandleFunc("/addlikepost", controllers.AddlikepostHandler)
	http.HandleFunc("/addlikecomment", controllers.AddlikecommentHandler)
	http.HandleFunc("/addlikeaccount", controllers.AddlikeaccountHandler)
	http.HandleFunc("/addlikeaccounts", controllers.AddlikeaccountsHandler)
	http.HandleFunc("/addcomment", controllers.AddCommentHandler)
	http.HandleFunc("/editcomment", controllers.EditcommentHandler)
	http.HandleFunc("/removecomment", controllers.RemovecommentHandler)
	http.HandleFunc("/addpost", controllers.AddpostHandler)
	http.HandleFunc("/editpost", controllers.EditpostHandler)
	http.HandleFunc("/removepost", controllers.RemovepostHandler)

	port := 80
	fmt.Printf("Server listening on :%d...\n", port)

	serverErr := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if serverErr != nil {
		fmt.Println("Error starting the server:", serverErr)
	}


}