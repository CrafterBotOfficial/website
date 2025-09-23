package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/russross/blackfriday/v2"
)

func index_page(w http.ResponseWriter, r *http.Request) {
	config := GetConfig()

	pageData := IndexPageData {}
	pageData.ServerTime = time.Now().Format("03:04:05 PM MST")
	pageData.Messages = config.Messages
	pageData.Socials = config.Socials

	handler("index", pageData, w, r)
}

func read_article_page(w http.ResponseWriter, r *http.Request) {
	pageData := ArticlePostPageData {}
	contents, err := GetArticleById(r.URL.Path)
	if err != nil {
		log.Println(err)
	}

	markdown := blackfriday.Run([]byte(contents), blackfriday.WithExtensions(blackfriday.LaxHTMLBlocks))
	pageData.Content = template.HTML(markdown)
	inherit_handler("article_view", pageData, w, r)
}

func articles_page(w http.ResponseWriter, r *http.Request) {
	articles, err := GetArticles()
	if err != nil {
		log.Println(err)
	}

	pageData := ArticlesPageData {}
	pageData.Blogs = articles

	inherit_handler("articles", pageData, w, r)
}

func inherit_handler(n string, pageData PageVars, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("templates/articles/*.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = t.ExecuteTemplate(w, n + ".html", pageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handler(templateName string, pageData PageVars, w http.ResponseWriter, r *http.Request) {
	templatePath := fmt.Sprintf("templates/%s.html", templateName)
	file, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("Failed to parse template %s", err)
	}

	err = file.Execute(w, pageData)
	if err != nil {
		log.Printf("Failed to fill template %s", err.Error())
	}
}
