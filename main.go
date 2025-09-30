package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	config := GetConfig()

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", IndexPage)
    http.HandleFunc("/projects", ProjectPage)
    http.HandleFunc("/articles/", ListArticlePage)
    http.HandleFunc("/post/", ReadArticlePage)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}
