package routers

import (
	"log"
	"net/http"
	"os"
	"strings"
	"website/routers/api"
	"website/routers/web"
)

func Init() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	initFileServerVideo()

	http.HandleFunc("/api/send-idea", api.SendIdea)
	http.HandleFunc("/api/get-videos", api.RequestVideos)

	http.HandleFunc("/", web.Index)
	http.HandleFunc("/contact", web.Contact)
	http.HandleFunc("/projects", web.Project)
	http.HandleFunc("/videos", web.Video)
	http.HandleFunc("/post/", web.ReadArticle)
	http.HandleFunc("/articles/", web.ListArticle)
}

func initFileServerVideo() {
	dir := os.Getenv("VIDEO_DIR")
	if len(dir) == 0 {
		log.Fatal("VIDEO_DIR isn't setup")
	}

	if !strings.Contains(dir, "trailcam") {
		log.Fatal("Bad VIDEO_DIR, make sure its correct!")
	}

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/trailcam/", http.StripPrefix("/trailcam/", fs))
}
