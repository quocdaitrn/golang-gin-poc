package middleware

import (
	"log"
	"net/http"
	
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github/quocdaitrn/golang-gin-poc/helpers"
)

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := helpers.NewJWTService().ValidateToken(tokenString)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[Admin]: ", claims["admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])
		}
	}
}
