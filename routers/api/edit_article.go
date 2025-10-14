package api

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"website/services"
)

func EditArticle(w http.ResponseWriter, r *http.Request) {
	if !services.IsAuthorized(r) {
		http.Error(w, "Unathorized", http.StatusUnauthorized)
		return
	}

	i := getIdFromRequest(w, r)
	if i == -1 {
		services.RespondError(w, "Failed to find article")
		return
	}

	a, err := services.GetArticleById(i)
	if err != nil {
		return
	}

	if !services.EditArticle(i, r.FormValue("text")) {
		w.Write([]byte("Failed to write to file"))
	}

	p := r.FormValue("Public") // if so must be on
	b := len(p) > 0 

	if b != a.IsPublic {
		log.Println("Changing visiblity of article")
		a.IsPublic = b
		d := services.GetDatabase()

		q := "UPDATE blog SET IsPublic = ? WHERE Id = ?"
		_, err := d.Exec(q, b, a.Id)
		if err != nil {
			services.RespondError(w, "Failed to update visiblity")
			return
		}
	}

	w.Write([]byte("<p>Submitted</p>"))
}

func RequestEditArticleMenu(w http.ResponseWriter, r *http.Request) {
	if !services.IsAuthorized(r) {
		services.RespondError(w, "Unauthorized")
		// http.Error(w, "Unathorized", http.StatusUnauthorized)
		return
	}

	i := getIdFromRequest(w, r)
	if i == -1 {
		return
	}

	a, err := services.GetArticleById(i)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/shared/edit_text.html"))
	tmpl.ExecuteTemplate(w, "edit_text.html", 
		struct {
			Id 		int
			Content string
			Public  bool
		} {
			int(a.Id),
			a.Contents,
			a.IsPublic,
		})
}

func getIdFromRequest(w http.ResponseWriter, r *http.Request) int64 {
	i, err := strconv.Atoi(r.FormValue("articleId"))
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return -1
	}
	return int64(i)
}
