package main

import (
	"fmt"
	"log"
	"net/http"

	"check-saldo-service/configs"
	"check-saldo-service/internal/app"
)

func main() {
	c := configs.NewConfig()
	router := c.Router

	router.HandleFunc("/api/saldo", app.CheckSaldo).Methods(http.MethodGet)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), router)
}
