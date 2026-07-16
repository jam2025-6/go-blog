package middleware

import (
	"errors"
	"server/global"
	"server/model/database"
	"server/model/request"
	"server/model/response"
	"server/service"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := utils.GetAccessToken(ctx)
		refreshToken := utils.GetRefreshToken(ctx)

		if service.ServiceGroupApp.JwtService.IsInBlacklist(refreshToken) {
			// 如果Refresh Token在黑名单中，返回401
			utils.ClearRefreshToken(ctx)
			response.NoAuth("Account logged in from another location or token is invalid", ctx)
			// 中止后续中间件或路由处理函数的执行
			ctx.Abort()
			return
		}

		j := utils.NewJWT()
		claims, err := j.ParseAccessToken(accessToken)
		if err != nil {
			// errors.Is 判断一个错误链中是否包含指定的特定错误（utils.TokenExpired）。
			if accessToken == "" || errors.Is(err, utils.TokenExpired) {
				// 如果请求中不包含accessToken，或者accessToken已过期，则尝试使用refreshToken刷新accessToken
				refreshClaims, err := j.ParseRefreshToken(refreshToken) // 解析refreshToken
				if err != nil {
					utils.ClearRefreshToken(ctx)
					response.NoAuth("Refresh token is invalid", ctx)
					// 中止后续中间件或路由处理函数的执行
					ctx.Abort()
					return
				}
				var user database.User
				// SELECT uuid, role_id FROM users WHERE id = refreshClaims.UserID LIMIT 1;
				if err = global.DB.Select("uuid", "role_id").Take(&user, refreshClaims.UserID).Error; err != nil {
					utils.ClearRefreshToken(ctx)
					response.NoAuth("The user does not exist", ctx)
					// 中止后续中间件或路由处理函数的执行
					ctx.Abort()
					return
				}

				newAccessClaims := j.CreateAccessClaims(request.BaseClaims{
					UserID: refreshClaims.UserID,
					UUID:   user.UUID,
					RoleID: user.RoleID,
				})
				newAccessToken, err := j.CreateAccessToken(newAccessClaims)
				if err != nil {
					utils.ClearRefreshToken(ctx)
					response.NoAuth("Failed to create access token", ctx)
					// 中止后续中间件或路由处理函数的执行
					ctx.Abort()
					return
				}
				ctx.Header("new-access-token", newAccessToken)
				// new-access-expires-at 自定义响应头字段名。不是标准 HTTP 头（如 Authorization），用于告知前端新 token 的过期时刻。
				// strconv.FormatInt(..., 10)将 int64 数字转为 十进制字符串。因为 HTTP Header 的值必须是字符串。
				ctx.Header("new-access-expires-at", strconv.FormatInt(newAccessClaims.ExpiresAt.Unix(), 10))
				ctx.Set("claims", &newAccessClaims)
				ctx.Next()
				return
			}
			utils.ClearRefreshToken(ctx)
			response.NoAuth("Invalid access token", ctx)
			// 中止后续中间件或路由处理函数的执行
			ctx.Abort()
			return

		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
