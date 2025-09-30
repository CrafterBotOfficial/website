package web

import (
	"html/template"
	"log"
	"net/http"
	"website/services"
)

type ArticlesPageData struct {
	Title string
	Articles []services.Article
}

func ListArticle(w http.ResponseWriter, r *http.Request) {
	articles, err := services.GetArticles()
	if err != nil {
		log.Println(err)
	}

	pageData := ArticlesPageData {}
	pageData.Title = "Articles"
	pageData.Articles = articles

	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/articles/articles.html",
	))
	tmpl.ExecuteTemplate(w, "articles.html", pageData)
}
