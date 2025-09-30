package routers

import (
	"net/http"
	"website/routers/api"
	"website/routers/web"
)

func Init() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/api/send-idea", api.SendIdea)

	http.HandleFunc("/", web.Index)
	http.HandleFunc("/contact", web.Contact)
	http.HandleFunc("/projects", web.Project)
    http.HandleFunc("/post/", web.ReadArticle)
    http.HandleFunc("/articles/", web.ListArticle)
}
