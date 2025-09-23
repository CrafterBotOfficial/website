package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", index_page)
    http.HandleFunc("/articles/", articles_page)
    http.HandleFunc("/post/", read_article_page)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
