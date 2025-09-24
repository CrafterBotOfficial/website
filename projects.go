package main

import (
	"log"
	"time"
)

type Project struct {
	Id int
	Title string
	Icon string
	DownloadUrl string
	Summary string
	Date time.Time
}

func GetProjects() ([]Project, error) {
	db := GetDatabase()
	var r []Project

	rows, err := db.Query("SELECT * FROM mods")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var project Project
		if err := rows.Scan(&project.Id, &project.Title, &project.Icon, &project.DownloadUrl, &project.Summary, &project.Date); err != nil {
			return nil, err	
		}
		r = append(r, project)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {
		return r, nil
	}
	log.Printf("No mods in db")
	return nil, nil
}
