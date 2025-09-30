package main

import (
	"fmt"
	"log"
	"net/http"
	"website/routers"
	"website/services"
)

func main() {
	config := services.GetConfig()

	routers.Init()
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}
