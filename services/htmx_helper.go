package services

import (
	"html/template"
	"net/http"
)

func RespondError(w http.ResponseWriter, e string) {
	tmpl := template.Must(template.ParseFiles("templates/shared/error.html"))
	tmpl.ExecuteTemplate(w, "error.html", e)
}
