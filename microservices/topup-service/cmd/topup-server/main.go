package main

import (
	"fmt"
	"log"
	"net/http"

	"topup-service/configs"
	"topup-service/internal/app"
)

func main() {
	c := configs.NewConfig()
	router := c.Router

	router.HandleFunc("/api/topup", app.Topup).Methods(http.MethodPost)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), router)
}
