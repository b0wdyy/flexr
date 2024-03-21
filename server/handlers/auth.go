package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", Login)
	router.POST("/logout", Logout)
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "logging in",
	})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "logging out",
	})
}
