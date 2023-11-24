package middleware

import (
	"mtsaudio/internal/tokens"
	"mtsaudio/internal/utils/httputils"
	"time"

	"strings"

	"github.com/gin-gonic/gin"
)

func CheckAuthification() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		authHeaderArray := strings.Split(authHeader, " ")
		if len(authHeaderArray) != 2 {
			httputils.NewResponseError(ctx, 401, "invalid authorization header")
			return
		}
		if authHeaderArray[1] == "" {
			httputils.NewResponseError(ctx, 401, "invalid jwt token")
		}
		token := authHeaderArray[1]
		if tokens.IsInBlackList(token) {
			httputils.NewResponseError(ctx, 401, "unauthorized")
			return
		}

		tokenData, err := tokens.ParseToken(token)
		if err != nil {
			httputils.NewResponseError(ctx, 401, err.Error())
			return
		}

		if int(time.Now().Unix())-tokenData.ExpIn > tokenData.CreatedAt {
			httputils.NewResponseError(ctx, 401, "token is expired")
			return
		}

		ctx.Set("id", tokenData.Id)
		ctx.Set("username", tokenData.Username)
		ctx.Next()
	}
}
