package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"website/services"
)

type videoData struct {
	Videos 		[]services.VideoData
	PrevPageUrl string
	NextPageUrl string
}

func RequestVideos(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	page := queryValues.Get("page")
	var pageIndex int = 0
	if len(page) != 0 {
		var err error
		pageIndex, err = strconv.Atoi(page)
		if err != nil {
			log.Println("Bad request")
			return
		}
	}

	videos, err := services.GetChunkedVideos(6, pageIndex)
	if err != nil {
		http.Error(w, "Failed to find videos in database", http.StatusInternalServerError)
		return
	}

	d := videoData {
		Videos: videos[pageIndex],
		PrevPageUrl: getPreviousPage(pageIndex),
		NextPageUrl: getNextPage(pageIndex, len(videos)),
	}

	tmpl := template.Must(template.ParseFiles("templates/video_container.html", "templates/shared/page_buttons.html"))
	tmpl.ExecuteTemplate(w, "video_container.html", d)
}

func getPreviousPage(index int) string {
	if index == 0 {
		return ""
	}
	return fmt.Sprintf("/api/get-videos?page=%d", index - 1)
}

func getNextPage(index int, maxPages int) string {
	if maxPages <= index + 1 {
		return ""
	}
	return fmt.Sprintf("/api/get-videos?page=%d", index + 1)
}
