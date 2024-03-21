package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup) {
	router.GET("/", GetUsers)
	router.POST("/", CreateUser)
	router.GET("/:id", GetUser)
	router.PUT("/:id", UpdateUser)
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "getting users",
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "creating user",
	})
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "getting user",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "updating user",
	})
}
