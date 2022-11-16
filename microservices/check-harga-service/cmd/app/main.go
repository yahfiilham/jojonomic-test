package main

import (
	"fmt"
	"log"
	"net/http"

	"check-harga-service/internal/app"
	"check-harga-service/internal/config"
)

func main() {
	c := config.NewConfig()
	router := c.Router

	router.HandleFunc("/api/check-harga", app.CheckHarga).Methods(http.MethodGet)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", c.Port), router)
}
