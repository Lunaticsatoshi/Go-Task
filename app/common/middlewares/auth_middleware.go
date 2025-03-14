package middlewares

import (
	"net/http"
	"slices"
	"strings"

	"github.com/Lunaticsatoshi/go-task/app/common/constants"
	"github.com/Lunaticsatoshi/go-task/app/common/interfaces"
	"github.com/Lunaticsatoshi/go-task/app/services"
	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) string {
	var token string
	cookie, err := c.Request.Cookie("user_auth_token")
	if err != nil {
		token = c.GetHeader("Authorization")
		if token == "" {
			token = c.Query("user_auth_token")
		}
	} else {
		token = cookie.Value
	}
	return token
}

func Authenticate(jwtService services.JWTService, roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := GetToken(c)
		if token == "" {
			response := interfaces.CreateFailResponse("No token found", "", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		token = strings.ReplaceAll(token, "Bearer ", "")

		jwtToken, err := jwtService.ValidateToken(token)
		if err != nil {
			response := interfaces.CreateFailResponse("Invalid token", "", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !jwtToken.Valid {
			response := interfaces.CreateFailResponse("Invalid token", "", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		idRes, roleRes, err := jwtService.GetAttrByToken(token)
		if err != nil {
			response := interfaces.CreateFailResponse("Failed to process request", "", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		} else if roleRes != constants.EnumRoleAdmin && !slices.Contains(roles, roleRes) {
			response := interfaces.CreateFailResponse("Action unauthorized", "", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}
		c.Set("UserID", idRes)
		c.Next()
	}
}
