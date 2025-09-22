package main

import (
	"encoding/json"
	"log"
	"os"
)

type Social struct {
	Icon string
	Url string
}

type Config struct {
	Messages []string
	Socials []Social
}

var cached_config *Config

func get_config() Config {
	if cached_config != nil {
		return *cached_config
	}

	json_file, e := os.ReadFile("config.json")
	if e != nil {
		log.Fatal("No config.json file")
	}

	var config_json Config
	e = json.Unmarshal(json_file, &config_json)
	if e != nil {
		log.Fatal("Failed to parse config")
	}

	log.Printf("Loaded config")
	cached_config = &config_json
	return config_json
}
