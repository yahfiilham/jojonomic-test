package main

import (
	"fmt"
	"log"
	"net/http"

	"topup-service/internal/app"
	"topup-service/internal/config"
)

func main() {
	c := config.NewConfig()
	router := c.Router

	router.HandleFunc("/api/topup", app.Topup).Methods(http.MethodPost)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", c.Port), router)
}
