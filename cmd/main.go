package main

import (
	"auth/configs"
	"auth/internal"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	configs.InitDB()
	r := internal.NewRouter()
	r.Run()
}
