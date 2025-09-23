package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Blog struct {
	Id int
	Title string
	Author string
}

func get_article_by_url(url string) (string, error) {
	names := strings.Split(url, "/")
	articlePath := fmt.Sprintf("%s/%s.md", os.Getenv("ARTICLES_PATH"), names[len(names) - 1])
	if _, err := os.Stat(articlePath); err != nil {
		log.Printf("Article markdown not found")
		return "Article not found", err
	}

	contents, err := os.ReadFile(articlePath)
	if err != nil {
		return "Couldn't read article", err
	}

	return string(contents), nil
}

func get_articles() ([]Blog, error) {
	db = get_database()
	var result []Blog

	rows, err := db.Query("SELECT * FROM manifest")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var blog Blog
		if err := rows.Scan(&blog.Id, &blog.Title, &blog.Author); err != nil {
			return nil, err	
		}
		result = append(result, blog)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return result, nil
}
