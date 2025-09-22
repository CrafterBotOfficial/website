package main

import (
	"time"
)

type PageData struct {
	ServerTime string
	Messages []string
	Socials []Social
}

func get_page_data() PageData {
	config := get_config()

	page_data := PageData {}
	page_data.ServerTime = time.Now().Format("03:00 PM MST")
	page_data.Messages = config.Messages
	page_data.Socials = config.Socials
	return page_data;
}
