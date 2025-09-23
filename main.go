package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	file_server := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", file_server))

	http.HandleFunc("/", index_page)
    http.HandleFunc("/articles/", articles_page)
    http.HandleFunc("/post/", read_article_page)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(templateName string, pageData PageVars, writer http.ResponseWriter, request *http.Request) {
	templatePath := fmt.Sprintf("templates/%s.html", templateName)
	file, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("Failed to parse template %s", err)
	}

	err = file.Execute(writer, pageData)
	if err != nil {
		log.Printf("Failed to fill template %s", err.Error())
	}
}
