package main

import (
	"database/sql"
	"log"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func GetInfoRoot() string {
	return os.Getenv("WEBSITE_INFO_DIRECTORY")
}

func GetDatabasePath() string {
	c := GetConfig()
	return path.Join(GetInfoRoot(), c.DBName)
}

func GetDatabase() *sql.DB {
	if database == nil {
		ConnectDatabase()
	}
	return database 
}

func ConnectDatabase() {
	p := GetDatabasePath();
	log.Printf("Opening database at %s", p)

	var err error
	if _, err = os.Stat(p); err != nil {
		log.Fatalf("Database doesn't exist in directory %s", GetInfoRoot())
	}

	db, err := sql.Open("sqlite3", p)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	database = db
}
