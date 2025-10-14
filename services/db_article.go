package services

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"slices"
	"strings"
	"time"
)

type Article struct {
	Id 				int64 
	Title 			string
	Author 			string
	Summary 		string
	Date 			time.Time
	IsPublic 		bool

	// not in db
	SanitizedName 	string
	Contents		string
}

func EditArticle(i int64, c string) bool {
	a, err := GetArticleById(i)
	if err != nil {
		return false
	}
	p := getPathFromArticle(a)

	f, err := os.OpenFile(p, os.O_RDWR, os.FileMode(0644))
	if err != nil {
		log.Printf("Failed to open article for write at %s", p)
		return false
	}
	defer f.Close()

	_, err = f.WriteString(c)
	if err != nil {
		log.Println("Failed to write to file")
		log.Println(err.Error())
		return false
	}

	// git 

	dir := os.Getenv("WEBSITE_INFO_DIRECTORY")
	if !runGitCommand(dir, "add", ".") {
		return false 
	}
	if !runGitCommand(dir, "commit", "-m", time.Now().Format("2006-01-02 15:04:05")) {
		return false 
	}
	if !runGitCommand(dir, "push") {
		return false 
	}

	return true
}

func GetArticleById(id int64) (*Article, error) {
	articles, err := GetArticles()
	if err != nil {
		return nil, err
	}

	idx := slices.IndexFunc(articles, func (a Article) bool { return a.Id == id }) 
	if idx == -1 {
		return nil, errors.New("Article not found")
	}

	a := articles[idx]
	// if !a.IsPublic {
	// 	return nil, errors.New("Unauthorized")
	// }

	p := getPathFromArticle(&a)
	c, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	a.Contents = string(c)
	return &a, nil
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

func getPathFromArticle(a *Article) string {
	return path.Join(GetInfoRoot(), "articles", fmt.Sprintf("%s.md", a.SanitizedName))
}

func runGitCommand(d string, c ...string) bool {
	cmd := exec.Command("git", c...)
	cmd.Dir = d
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err == nil
}
