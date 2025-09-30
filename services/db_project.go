package services

import (
	"log"
	"time"
)

type Project struct {
	Id int64 
	Title string
	Icon string
	DownloadUrl string
	Summary string
	Date time.Time
}

func GetProjects() ([]Project, error) {
	db := GetDatabase()

	rows, err := db.Query("SELECT * FROM mods")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer rows.Close()

	var r []Project
	for rows.Next() {
		var p Project
		rows.Scan(&p.Id, &p.Title, &p.Icon, &p.DownloadUrl, &p.Summary, p.Date)
		r = append(r, p)
	}

	return r, nil
}
