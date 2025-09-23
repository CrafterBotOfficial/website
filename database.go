package main

import (
	"database/sql"
	"log"
	"os"
 	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func get_database() *sql.DB {
	if db == nil {
		connect_database()
	}
	return db
}

func connect_database() {
	// var global_config = get_config()

	// config := mysql.NewConfig()
	// config.User = global_config.Database.User
	// config.Passwd = os.Getenv("DBPASS")
	// config.Addr = global_config.Database.Address
	// config.DBName = global_config.Database.Name

	var err error
	db, err = sql.Open("sqlite3", os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
