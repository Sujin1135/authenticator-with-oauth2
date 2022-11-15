package main

import (
	"auth/configs"
	"auth/internal/api"
	"auth/internal/oauth"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	configs.InitDB()
	oauth.InitAuthProviders()
	r := api.NewRouter()
	r.Run()
}
