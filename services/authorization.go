package services

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
)

const AUTH_HEADER = "ADMIN_AUTH"

func IsAuthorized(r *http.Request) bool {
	c, err := r.Cookie(AUTH_HEADER)
	if err != nil {
		return false
	}

	t, err := getAuthorizationToken()
	if err != nil {
		log.Println("Failed to find admin.env file")
		return false
	}

	if c.Value == t {
		return true
	}

	return false
}

func getAuthorizationToken() (string, error) {
	b := os.Getenv("ADMIN_TOKEN")
	s := strings.TrimSpace(string(b))
	if len(s) < 2 {
		return "", errors.New("Bad auth token in admin.env")
	}

	return s, nil
}
