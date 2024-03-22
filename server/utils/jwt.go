package utils

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractJWT(c *gin.Context) (string, error) {
	// Extract the JWT from the Authorization header
	jwt := c.GetHeader("Authorization")

	if jwt == "" {
		return "", fmt.Errorf("NO JWT PROVIDED")
	}

	token := strings.Split(jwt, " ")[1]

	return token, nil
}
