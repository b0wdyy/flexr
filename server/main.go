package main

import (
	"api/handlers"
	"api/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
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
	jwtMiddleware := middleware.TestValidToken()

	// Setup user routes
	userRoutes := apiRoutes.Group("/users")
	userRoutes.Use(adapter.Wrap(jwtMiddleware.CheckJWT))
	handlers.SetupUserRoutes(userRoutes)

	// Setup post routes
	postRoutes := apiRoutes.Group("/posts")
	handlers.SetupPostRoutes(postRoutes)

	router.Run(":" + os.Getenv("PORT"))
}
