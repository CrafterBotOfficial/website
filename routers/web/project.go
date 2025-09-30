package web

import (
	"html/template"
	"log"
	"net/http"
	"website/services"
)

type ProjectsPageData struct {
	Title string
	Projects []services.Project
}

func Project(w http.ResponseWriter, r *http.Request) {
	ps, err := services.GetProjects()
	if err != nil {
		http.Error(w, "Failed to get projects", http.StatusInternalServerError)
		log.Printf("%s", err)
		return
	}
	p := ProjectsPageData {}
	p.Title = "My Projects"
	p.Projects = ps

	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/projects.html",
	))
	tmpl.ExecuteTemplate(w, "projects.html", p)
}
