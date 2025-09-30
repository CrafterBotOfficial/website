package services

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

type Article struct {
	Id int64 
	Title string
	Author string
	Summary string
	Date time.Time
	IsPublic bool

	SanitizedName string
}

func GetArticleById(u string) (string, error) {
	url := strings.Split(u, "/")
	endpoint := url[len(url) - 1]
	path := path.Join(GetInfoRoot(), "articles")
	p := fmt.Sprintf("%s/%s.md", path, endpoint)
	if _, err := os.Stat(p); err != nil {
		log.Printf("Article markdown not found")
		return "Article not found", err
	}

	s, err := os.ReadFile(p)
	if err != nil {
		return "Couldn't read article", err
	}

	return string(s), nil
}

func GetArticles() ([]Article, error) {
	db := GetDatabase()

	rows, err := db.Query("SELECT * FROM blog")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer rows.Close()

	var r []Article
	for rows.Next() {
		var b Article
		rows.Scan(&b.Id, &b.Title, &b.Author, &b.Summary, &b.Date, &b.IsPublic)
		b.SanitizedName = strings.TrimSpace(strings.ToLower(strings.ReplaceAll(b.Title, " ", "_")))
		r = append(r, b)
	}

	return r, nil
}
