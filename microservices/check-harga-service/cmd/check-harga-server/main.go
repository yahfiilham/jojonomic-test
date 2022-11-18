package main

import (
	"fmt"
	"log"
	"net/http"

	"check-harga-service/configs"
	"check-harga-service/internal/app"
)

func main() {
	c := configs.NewConfig()
	router := c.Router

	router.HandleFunc("/api/check-harga", app.CheckHarga).Methods(http.MethodGet)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), router)
}
