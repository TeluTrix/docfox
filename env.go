package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(){
	err := godotenv.Load()
  	if err != nil {
    	slog.Error("Error loading .env file")
		os.Exit(1)
  	}
}

