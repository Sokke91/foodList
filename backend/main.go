package main

import (
	"foodList/authentication"
	"foodList/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	startServer()
}

func loadDatabase() {
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Env konnte nicht geladen werden")
	}
}

func startServer() {
	router := gin.Default()
	router.POST("/login", controllers.Login)
	protectedApi := router.Group("/api")
	protectedApi.Use(authentication.AuthMiddleware())
	protectedApi.GET("verify", controllers.ProtectedRoute)
	router.Run(":8080")
}
