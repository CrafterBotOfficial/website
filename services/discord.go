package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type WebhookPayload struct {
	Content 	string		`json:"content"`
}

func SendWebhook(url string, payload WebhookPayload) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return errors.New("Failed to send parse your message.")
	}

	client := http.Client {}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return errors.New("Failed to send parse your message.")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	log.Printf("%s\n", resp.Status)
	defer resp.Body.Close()
	return nil
}

func GetAddressFromHeader(r *http.Request) string {
	var result string

	keys := make([]string, 0, len(r.Header))
	for k := range r.Header {
		keys = append(keys, k)
	}
	
	for _,str  := range keys {
		if str == "Cf-Connecting-Ip" {
			result = r.Header[str][0]
			break
		}
	}

	return result
}

