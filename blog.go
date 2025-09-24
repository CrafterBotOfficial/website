package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

type Blog struct {
	Id int
	Title string
	Author string
	Summary string
	Date time.Time
	IsPublic bool

	SanitizedName string
}

func GetArticleById(u string) (string, error) {
	a := strings.Split(u, "/")
	n := a[len(a) - 1]
	d := path.Join(GetInfoRoot(), "articles")
	p := fmt.Sprintf("%s/%s.md", d, n)
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

func GetArticles() ([]Blog, error) {
	db := GetDatabase()
	var r []Blog

	c := GetConfig()
	rows, err := db.Query("SELECT * FROM " + c.Database.ManifestTable)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var blog Blog
		if err := rows.Scan(&blog.Id, &blog.Title, &blog.Author, &blog.Summary, &blog.Date, &blog.IsPublic); err != nil {
			return nil, err	
		}
		blog.SanitizedName = strings.TrimSpace(strings.ToLower(strings.ReplaceAll(blog.Title, " ", "_")))
		r = append(r, blog)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	if len(r) > 0 {
		return r, nil
	}
	log.Printf("No articles in db")
	return nil, nil
}
