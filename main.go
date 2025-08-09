package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/patil-prathamesh/Blogify/database"
	"github.com/patil-prathamesh/Blogify/routes"
)

func main() {
	err := godotenv.Load(".env")
	port := os.Getenv("PORT")
	if err != nil {
		log.Panic("Error while loading .env")
	}

	database.ConnectDatabase()

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.PostRoutes(router)

	router.Run(":" + port)

}
