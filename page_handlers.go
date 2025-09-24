package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/russross/blackfriday/v2"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	config := GetConfig()

	pageData := IndexPageData {}
	pageData.ServerTime = time.Now().Format("03:04:05 PM MST")
	pageData.Messages = config.Messages
	pageData.Socials = config.Socials

	handler("index", pageData, w, r)
}

func ProjectPage(w http.ResponseWriter, r *http.Request) {
	ps, err := GetProjects()
	if err != nil {
		http.Error(w, "Failed to get projects", http.StatusInternalServerError)
		log.Printf("%s", err)
		return
	}
	p := ProjectsPageData {}
	p.Title = "My Projects"
	p.Projects = ps
	ineritHandler("projects", p, w, r)
}

func ReadArticlePage(w http.ResponseWriter, r *http.Request) {
	pageData := ArticlePostPageData {}
	contents, err := GetArticleById(r.URL.Path)
	if err != nil {
		log.Println(err)
	}

	markdown := blackfriday.Run([]byte(contents), blackfriday.WithExtensions(blackfriday.LaxHTMLBlocks))
	pageData.Content = template.HTML(markdown)
	ineritHandler("article_view", pageData, w, r)
}

func ListArticlePage(w http.ResponseWriter, r *http.Request) {
	articles, err := GetArticles()
	if err != nil {
		log.Println(err)
	}

	pageData := ArticlesPageData {}
	pageData.Title = "Articles"
	pageData.Blogs = articles

	ineritHandler("articles", pageData, w, r)
}

func ineritHandler(n string, pageData PageVars, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("templates/*.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	t, err = t.ParseGlob("templates/articles/*.html")
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
