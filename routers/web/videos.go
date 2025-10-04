package web

import (
	"html/template"
	"net/http"
)

type VideosPage struct {
	Title 		 string
}

func Video(w http.ResponseWriter, r *http.Request) {
	p := VideosPage {}
	p.Title = "Trail Videos"

	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/videos.html",
	))
	tmpl.ExecuteTemplate(w, "videos.html", p)
}
