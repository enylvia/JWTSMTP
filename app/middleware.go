package app

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"jwtsmtp/helper"
	"jwtsmtp/service"
	"net/http"
	"strings"
)


func IsAuthorized(authService service.JWTService, userService service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer"){
			response := helper.ApiResponse(http.StatusUnauthorized,"error","Token Not Found",nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
			return
		}
		tokenString := ""

		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateJWT(tokenString)
		if err != nil {
			response := helper.ApiResponse(http.StatusUnauthorized,"error","Failed to process request",nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
			return
		}

		_ , ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.ApiResponse(http.StatusUnauthorized,"error","Invalid Token",nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
			return
		}

	}
}
