package models

import "html/template"

var idloginpost int
var idcategorie int
var idpost int
var idcomment int
var idlike int

type Post struct{
	Idpost  int
	Idlogin int
	Content string
	Name string
	Surname string
	Pseudo string
	Urlpseudo string
	Svglike template.HTML
	Categorie string
	Url string
	Urlcategorie template.HTML
	Likenumber int
	Commentnumber int
}

type Comment struct{
	Idcomment int
	Idlogin int
	Name string
	Surname string
	Pseudo string
	Content string
	Urlpseudo string
	Svglike template.HTML
	Likenumber int
}

type Categorie struct{
	Idcategorie int
	Name string
	Urlcategorie template.HTML
}

type Login struct{
	Id int
	Name string
	Surname string
	Pseudo string
	Urlpseudo string
	Adressemail string
}

type User struct{
	Id int
	Name string
	Surname string
	Pseudo string
	Urlpseudo string
	Adressemail string
}

type Tendance struct{
	Idcategorie int
	Name string
	Urlcategorie template.HTML
}

type Explorer struct{
	Id int
	Name string
	Surname string
	Pseudo string
	Urlpseudo string
}

type Data struct{
	Categories []Categorie
	Posts []Post
	Logins Login
	Users User
	Explorers []Explorer
	Comments []Comment
	Tendances []Tendance
}