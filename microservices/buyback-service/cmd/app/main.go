package main

import (
	"buyback-service/internal/app"
	"buyback-service/internal/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	c := config.NewConfig()
	router := c.Router

	router.HandleFunc("/api/buyback", app.Buyback).Methods(http.MethodPost)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), router)
}
