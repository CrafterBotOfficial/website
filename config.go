package main

import (
	"encoding/json"
	"log"
	"os"
)

type Social struct {
	Icon string
	Url  string
}

type Config struct {
	DBName   string
	Port 	 int
	Messages []string
	Socials  []Social
}

var cachedConfig *Config

func GetConfig() Config {
	if cachedConfig != nil {
		return *cachedConfig
	}

	json_file, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("No config.json file")
	}

	var config Config
	err = json.Unmarshal(json_file, &config)
	if err != nil {
		log.Fatal("Failed to parse config")
	}

	log.Printf("Loaded config")
	cachedConfig = &config
	return config 
}
