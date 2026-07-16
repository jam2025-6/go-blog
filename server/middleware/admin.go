package middleware

import (
	"server/model/appTypes"
	"server/model/response"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleID := utils.GetRoleID(ctx)
		if roleID != appTypes.Admin { // 1 is admin role
			response.Forbidden("Access denied. Adimin privileges are required", ctx)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
