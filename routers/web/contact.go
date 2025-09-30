package web

import (
	"html/template"
	"net/http"
	"website/services"
)

type ContactPage struct {
	Title 		 string
	TurnstileKey string
}

func Contact(w http.ResponseWriter, r *http.Request) {
	config := services.GetConfig()

	p := ContactPage {}
	p.Title = "Contact"
	p.TurnstileKey = config.TurnstileKey
	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/contact.html",
	))
	tmpl.ExecuteTemplate(w, "contact.html", p)
}
