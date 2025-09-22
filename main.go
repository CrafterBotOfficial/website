package main 

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	file_server := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", file_server))

    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	file, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Failed to parse template")
	}

	err = file.Execute(writer, get_page_data())
	if err != nil {
		log.Printf("Failed to fill template %s", err.Error())
	}
}
