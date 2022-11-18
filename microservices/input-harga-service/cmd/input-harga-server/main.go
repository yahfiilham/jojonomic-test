package main

import (
	"fmt"
	"log"
	"net/http"

	"input-harga-service/configs"
	"input-harga-service/internal/app"
)

func main() {
	c := configs.NewConfig()
	router := c.Router

	router.HandleFunc("/api/input-harga", app.InputHarga).Methods(http.MethodPost)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), router)
}
