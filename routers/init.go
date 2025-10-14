package routers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"website/routers/api"
	"website/routers/web"
)

func Init() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	initFileServerVideo()
	initFileServerAssets()

	http.HandleFunc("/api/send-idea", api.SendIdea)
	http.HandleFunc("/api/get-videos", api.RequestVideos)
	http.HandleFunc("/api/edit-article", api.RequestEditArticleMenu)
	http.HandleFunc("/api/submit-edit-article", api.EditArticle)

	http.HandleFunc("/", web.Index)
	http.HandleFunc("/contact", web.Contact)
	http.HandleFunc("/projects", web.Project)
	http.HandleFunc("/videos", web.Video)
	http.HandleFunc("/post/", web.ReadArticle)
	http.HandleFunc("/articles/", web.ListArticle)
}

func initFileServerVideo() {
	infoDir := os.Getenv("WEBSITE_INFO_DIRECTORY")
	if len(infoDir) == 0 {
		log.Fatal("WEBSITE_INFO_DIRECTORY dir isn't setup")
	}
	dir := filepath.Join(infoDir, "trailcam")

	if !strings.Contains(dir, "trailcam") {
		log.Fatal("Bad WEBSITE_INFO_DIRECTORY, make sure its correct!")
	}

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/trailcam/", http.StripPrefix("/trailcam/", fs))
}

func initFileServerAssets() {
	infoDir := os.Getenv("WEBSITE_INFO_DIRECTORY")
	dir := filepath.Join(infoDir, "assets")
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/info/", http.StripPrefix("/info/", fs))
}
