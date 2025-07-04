package main

import (
	"log"
	"main/internal/app"
	"main/internal/config"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db_url := os.Getenv("DB_URL")
	cfg := config.MustLoad()
	a := app.New(cfg, db_url)
	a.Run()
}
