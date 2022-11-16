package main

import (
	"check-saldo-service/internal/app"
	"check-saldo-service/internal/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	c := config.NewConfig()
	router := c.Router

	router.HandleFunc("/api/saldo", app.CheckSaldo).Methods(http.MethodGet)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", c.Port), router)
}
