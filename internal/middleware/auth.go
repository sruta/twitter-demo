package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"twitter-uala/internal/helpers"
	"twitter-uala/pkg"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		bearer := strings.Split(header, " ")
		if len(bearer) < 2 {
			msg := "request does not contain an access token"
			apiError := pkg.NewUnauthorizedApiError(pkg.NewUnauthorizedError(msg, nil))
			ctx.JSON(apiError.GetStatus(), apiError.GetResponse())
			ctx.Abort()
			return
		}

		claims, err := helpers.VerifyJwtToken(bearer[1])
		if err != nil {
			msg := "could not verify access token"
			apiError := pkg.NewUnauthorizedApiError(pkg.NewUnauthorizedError(msg, err))
			ctx.JSON(apiError.GetStatus(), apiError.GetResponse())
			ctx.Abort()
			return
		}

		ctx.Set("userID", claims.ID)
		ctx.Next()
	}
}
