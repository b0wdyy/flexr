package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

type CustomClaims struct {
	Scope string `json:"scope"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// GinJWTMiddleware adapts the JWT middleware to be used with Gin.
func GinJWTMiddleware() gin.HandlerFunc {
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")

	if err != nil {
		log.Fatalf("Failed to parse the issuer URL: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)
	fmt.Println(os.Getenv("AUTH0_DOMAIN"))
	fmt.Println(os.Getenv("AUTH0_AUDIENCE"))

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		validator.WithCustomClaims(func() validator.CustomClaims {
			return &CustomClaims{}
		}),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the JWT validator: %v", err)
	}

	return func(c *gin.Context) {
		// Extract token from the Authorization header
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
			return
		}

		tokenString := authHeader[len("Bearer "):] // Remove "Bearer " prefix
		_, err := jwtValidator.ValidateToken(context.Background(), tokenString)

		if err != nil {
			log.Printf("Encountered error while validating JWT: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Failed to validate JWT"})
			return
		}

		c.Next()
	}
}
