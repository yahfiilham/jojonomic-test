package main

import (
	"fmt"
	"log"
	"net/http"

	"buyback-service/internal/app"
	"buyback-service/internal/config"
)

func main() {
	c := config.NewConfig()
	router := c.Router

	router.HandleFunc("/api/buyback", app.Buyback).Methods(http.MethodPost)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", c.Port), router)
}
