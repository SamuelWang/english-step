package main

import (
	"english-step/website-api/database"
	"english-step/website-api/middlewares"
	"log"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file:", err)
	}

	db, err := database.Init()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	database.MigrateDev(db)

	router := gin.Default()

	router.Use(middlewares.DBContextMiddleware(db))
}
