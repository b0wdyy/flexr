package main

import (
	"api/handlers"
	"api/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// utils.ConnectToDatabase()

	router := gin.Default()
	apiRoutes := router.Group("/api")

	apiRoutes.GET("/ping", handlers.Ping)

	// Setup auth routes
	authRoutes := apiRoutes.Group("/auth")
	handlers.SetupAuthRoutes(authRoutes)

	// Setup user routes
	userRoutes := apiRoutes.Group("/users")
	userRoutes.Use(middleware.EnsureValidToken())
	handlers.SetupUserRoutes(userRoutes)

	router.Run(":" + os.Getenv("PORT"))
}
