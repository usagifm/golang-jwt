package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/usagifm/golang-jwt/helper"
	"github.com/usagifm/golang-jwt/service"
)

// AuthorizeJWT validatess the toke user given, return 401 is not valid
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process the request", "No Token Found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		splitToken := strings.Split(authHeader, "Bearer ")
		authHeader = splitToken[1]

		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id] :", claims["user_id"])
			log.Println("Claims[issuer] :", claims["issuer"])

		} else {

			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

	}

}
