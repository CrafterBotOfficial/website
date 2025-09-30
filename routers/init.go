package routers

import (
	"net/http"
	"website/routers/web"
)

func Init() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", web.Index)
	http.HandleFunc("/projects", web.Project)
    http.HandleFunc("/post/", web.ReadArticle)
    http.HandleFunc("/articles/", web.ListArticle)
}
