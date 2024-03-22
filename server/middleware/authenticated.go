package middleware

import (
	"api/utils"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

func TestValidToken() *jwtmiddleware.JWTMiddleware {
	issuerURL, err := url.Parse(os.Getenv("AUTH0_ISSUER"))
	audience := os.Getenv("AUTH0_AUDIENCE")

	fmt.Println(issuerURL)
	fmt.Println(audience)

	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, time.Duration(5*time.Minute))

	jwtValidator, _ := validator.New(provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{audience},
	)

	jwtMiddleware := jwtmiddleware.New(jwtValidator.ValidateToken)

	return jwtMiddleware
}

// EnsureValidToken is a middleware that will check the validity of our JWT.
func EnsureValidToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")

		if err != nil {
			log.Fatalf("Failed to parse the issuer url: %v", err)
		}

		provider := jwks.NewCachingProvider(issuerURL, 86400)

		jwtValidator, err := validator.New(
			provider.KeyFunc,
			validator.RS256,
			issuerURL.String(),
			[]string{os.Getenv("AUTH0_AUDIENCE")},
			validator.WithAllowedClockSkew(time.Minute),
		)
		if err != nil {
			log.Fatalf("Failed to set up the jwt validator")
		}

		jwt, err := utils.ExtractJWT(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No JWT provided"})
			c.Abort()
			return
		}

		fmt.Print(jwt)
		// Validate the JWT
		token, err := jwtValidator.ValidateToken(c, jwt)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		fmt.Println(token)

		c.Next()
	}

}
