package main

import (
	"fmt"
	"input-harga-service/internal/app"
	"input-harga-service/internal/config"
	"log"
	"net/http"
)

func main() {
	c := config.NewConfig()
	router := c.Router

	router.HandleFunc("/api/input-harga", app.InputHarga).Methods(http.MethodPost)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), router)
}
