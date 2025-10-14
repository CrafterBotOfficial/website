package web

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"website/services"

	"github.com/russross/blackfriday/v2"
)

type ArticlePostPageData struct {
	Id			int64
	Title 		string
	Author 		string
	Content 	template.HTML 
	IsAdmin		bool
}

func ReadArticle(w http.ResponseWriter, r *http.Request) {
	u := strings.Split(r.URL.Path, "/")
	endpoint := u[len(u) - 1]
	var idx int; var err error
	if idx, err = strconv.Atoi(endpoint); err != nil {
		http.Error(w, "invalid url", http.StatusInternalServerError)
		return
	}

	pageData := ArticlePostPageData {}
	article, err := services.GetArticleById(int64(idx))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	admin := services.IsAuthorized(r)
	if !article.IsPublic && !admin {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	markdown := blackfriday.Run([]byte(article.Contents), blackfriday.WithExtensions(blackfriday.LaxHTMLBlocks))
	pageData.Id = article.Id
	pageData.Title = article.Title
	pageData.Content = template.HTML(markdown)
	pageData.IsAdmin = admin 

	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/articles/read_article.html",
	))
	tmpl.ExecuteTemplate(w, "read_article.html", pageData)
}
