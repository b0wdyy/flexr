package handlers

import "github.com/gin-gonic/gin"

func SetupPostRoutes(router *gin.RouterGroup) {
	router.GET("/", GetPosts)
	router.POST("/", CreatePost)
	router.GET("/:id", GetPost)
	router.PUT("/:id", UpdatePost)
	router.DELETE("/:id", DeletePost)
}

func GetPosts(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get all posts",
	})
}

func CreatePost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create a post",
	})
}

func GetPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get a post",
	})
}

func UpdatePost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update a post",
	})
}

func DeletePost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete a post",
	})
}
