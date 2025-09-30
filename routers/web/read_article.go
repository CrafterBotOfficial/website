package web

import (
	"html/template"
	"log"
	"net/http"
	"website/services"

	"github.com/russross/blackfriday/v2"
)

type ArticlePostPageData struct {
	Title string
	Author string
	Content template.HTML 
}

func ReadArticle(w http.ResponseWriter, r *http.Request) {
	pageData := ArticlePostPageData {}
	contents, err := services.GetArticleById(r.URL.Path)
	if err != nil {
		log.Println(err)
	}

	markdown := blackfriday.Run([]byte(contents), blackfriday.WithExtensions(blackfriday.LaxHTMLBlocks))
	pageData.Title = "Article"
	pageData.Content = template.HTML(markdown)

	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/articles/article_view.html",
	))
	tmpl.ExecuteTemplate(w, "article_view.html", pageData)
}
