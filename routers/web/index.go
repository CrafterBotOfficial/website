package web

import (
	"html/template"
	"net/http"
	"time"
	"website/services"
)

type IndexPageData struct {
	ServerTime string
	Messages []string
	Socials []services.Social
}

func Index(w http.ResponseWriter, r *http.Request) {
	config := services.GetConfig()

	pageData := IndexPageData {}
	pageData.ServerTime = time.Now().Format("03:04:05 PM MST")
	pageData.Messages = config.Messages
	pageData.Socials = config.Socials

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, pageData)
}
