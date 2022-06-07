package middlewares

import (
	"finalproject/helpers"
	"finalproject/params"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, params.Response{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
				Error:   "Token Not Found",
			})
			return
		}

		fmt.Println(token)

		bearer := strings.HasPrefix(token, "Bearer")
		if !bearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, params.Response{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
				Error:   "No Bearer Prefix",
			})
			return
		}

		tokenStr := strings.Split(token, "Bearer ")[1]
		if tokenStr == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, params.Response{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
				Error:   "Token Not Found",
			})
			return
		}

		claims, err := helpers.VerifyToken(tokenStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, params.Response{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
				Error:   err.Error(),
			})
			return
		}

		var data = claims.(jwt.MapClaims)

		ctx.Set("email", data["email"])
		ctx.Set("id", data["id"])
		ctx.Next()
	}
}
