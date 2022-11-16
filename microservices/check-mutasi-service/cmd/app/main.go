package main

import (
	"check-mutasi-service/internal/app"
	"check-mutasi-service/internal/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	c := config.NewConfig()
	router := c.Router

	router.HandleFunc("/api/mutasi", app.CheckMutasi).Methods(http.MethodGet)

	log.Printf("api running in port %d", c.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), router)
}
