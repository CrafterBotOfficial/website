package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", IndexPage)
    http.HandleFunc("/projects", ProjectPage)
    http.HandleFunc("/articles/", ListArticlePage)
    http.HandleFunc("/post/", ReadArticlePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
