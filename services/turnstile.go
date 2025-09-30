package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type turnstileResponse struct {
	Success bool `json:"success"`
}

type turnstileRequest struct {
	Secret   string `json:"secret"`
	UserToken string `json:"response"`
	Ip string `json:"remoteip"`
}

func ValidateTurnstile(token string, ip string) bool {
	secret := os.Getenv("CLOUDFLARE_TUNNEL_SECRET")
	if len(secret) == 0 {
		log.Println("Turnstile token not set")
		return false
	}

	jsonBytes, _ := json.Marshal(turnstileRequest {
		Secret: secret,
		UserToken: token,
		Ip: ip,
	})

	req, err := http.NewRequest("POST", "https://challenges.cloudflare.com/turnstile/v0/siteverify", bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Println(err)
		return false
	}

	defer req.Body.Close()
	var bytes []byte
	req.Body.Read(bytes)

	var resp turnstileResponse
	json.Unmarshal(bytes, &resp)
	return resp.Success
}
