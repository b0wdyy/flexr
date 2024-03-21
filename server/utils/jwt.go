package utils

import "github.com/gin-gonic/gin"

func ExtractJWT(c *gin.Context) string {
	jwt := c.GetHeader("Authorization")

	return jwt
}
