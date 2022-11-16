package main

import (
	"fmt"
	"log"
	"net/http"

	"input-harga-service/internal/app"
	"input-harga-service/internal/config"
)

func main() {
	c := config.NewConfig()
	router := c.Router

	router.HandleFunc("/api/input-harga", app.InputHarga).Methods(http.MethodPost)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", c.Port), router)
}
