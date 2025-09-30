package api

import (
	"log"
	"net/http"
	"os"
	"website/services"

	"github.com/microcosm-cc/bluemonday"
)

func SendIdea(w http.ResponseWriter, r *http.Request) {
	ip := services.GetAddressFromHeader(r)
	if len(ip) == 0 {
		htmxError("Server not behind tunnel", w)
		return
	}

	p := bluemonday.UGCPolicy()

	token := p.Sanitize(r.FormValue("cf-turnstile-response"))
	if len(token) == 0 {
		htmxError("Bad token", w)
		return
	}

	limited := services.ValidateTurnstile(token, ip)
	if limited {
		htmxError("Failed anti-robot check", w)
	}

	idea := r.FormValue("idea")
	msg := p.Sanitize(idea)

	if len(idea) == 0 {
		htmxError("Please write something to send.", w)
		return
	}

	url := os.Getenv("DISCORD_WEBHOOK")
	if len(url) == 0 {
		htmxError("Server not configured.", w)
		return
	}

	err := services.SendWebhook(url,  services.WebhookPayload { Content: msg })
	if err != nil {
		htmxError(err.Error(), w)
	}

	response := "<label class=\"\">Successfully sent message</label>"
	w.Write([]byte(response))
}

func htmxError(err string, w http.ResponseWriter) {
	log.Println(err)
	err = "Error occured:<br>" + err
	w.Write([]byte(err))
}
