package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"check-harga-service/internal/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Config struct {
	Router *mux.Router
	DB     *gorm.DB
	Port   int
}

func NewConfig() *Config {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rt := mux.NewRouter()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Harga{})

	port := os.Getenv("PORT")
	portInt, _ := strconv.Atoi(port)

	app := &Config{
		Router: rt,
		Port:   portInt,
		DB:     db,
	}

	return app
}
