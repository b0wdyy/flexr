package main

import (
	"api/internal/handlers"
	"api/internal/middleware"
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

	// Middleware
	jwtMiddleware := middleware.GinJWTMiddleware()

	// Setup user routes
	userRoutes := apiRoutes.Group("/users")
	userRoutes.Use(jwtMiddleware)
	handlers.SetupUserRoutes(userRoutes)

	// Setup post routes
	postRoutes := apiRoutes.Group("/posts")
	handlers.SetupPostRoutes(postRoutes)

	router.Run(":" + os.Getenv("PORT"))
}
