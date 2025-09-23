package main

import (
	"log"
	"net/http"
	"time"
)

type PageVars interface { }

type IndexPageData struct {
	ServerTime string
	Messages []string
	Socials []Social
}

type ArticlePostPageData struct {
	Content string
}

type ArticlesPageData struct {
	Blogs []Blog
}

func index_page(writer http.ResponseWriter, request *http.Request) {
	config := get_config()

	page_data := IndexPageData {}
	page_data.ServerTime = time.Now().Format("03:04:05 PM MST")
	page_data.Messages = config.Messages
	page_data.Socials = config.Socials

	handler("index", page_data, writer, request)
}

func read_article_page(writer http.ResponseWriter, request *http.Request) {
	page_data := ArticlePostPageData {}
	contents, err := get_article_by_url(request.URL.Path)
	if err != nil {
		log.Println(err)
	}
	page_data.Content = contents
	handler("article_view", page_data, writer, request)
}

func articles_page(writer http.ResponseWriter, request *http.Request) {
	articles, err := get_articles()
	if err != nil {
		log.Println(err)
	}

	page_data := ArticlesPageData {}
	page_data.Blogs = articles

	handler("articles", page_data, writer, request)
}
