package main

import (
	"english-step/website-api/database"
	"log"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	_ "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file:", err)
	}

	database.Init()
}
