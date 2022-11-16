package main

import (
	"check-harga-service/internal/app"
	"check-harga-service/internal/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	c := config.NewConfig()
	router := c.Router

	router.HandleFunc("/api/check-harga", app.CheckHarga).Methods(http.MethodGet)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), router)
}
