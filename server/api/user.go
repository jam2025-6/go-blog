package api

import (
	"errors"
	"log"
	"server/global"
	"server/model/database"
	"server/model/request"
	"server/model/response"
	"server/utils"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type UserApi struct {
}

// Register 注册
func (userApi *UserApi) Register(ctx *gin.Context) {
	var req request.Register
	// 就是把 HTTP 请求体里的 JSON 数据，自动解析并填充到你定义好的 Go 结构体 req 中，如果解析失败就返回错误。
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	// 从当前请求的上下文（ctx）中获取或创建一个默认的 Session 对象，并将其赋值给变量 session，以便后续进行读取、写入或删除会话数据的操作。
	session := sessions.Default(ctx)
	// 两次邮箱一致性判断
	savedEmail := session.Get("email")
	log.Println(savedEmail)
	if savedEmail == nil || savedEmail.(string) != req.Email {
		response.FailWithMessage("This email doesn't match the email to be verified", ctx)
		return
	}
	// 获取会话中存储的邮箱验证码
	savedCode := session.Get("verification_code")
	log.Println(savedCode)
	if savedCode == nil || savedCode.(string) != req.VerificationCode {
		response.FailWithMessage("Invalid verification code", ctx)
		return
	}

	// 判断邮箱验证码是否过期
	savedTime := session.Get("expire_time")
	if savedTime.(int64) < time.Now().Unix() {
		response.FailWithMessage("The verification code has expired, please resend it", ctx)
		return
	}

	u := database.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	user, err := userService.Register(u)
	if err != nil {
		global.Log.Error("Failed to register user:", zap.Error(err))
		response.FailWithMessage("Failed to register user", ctx)
		return
	}

	// 注册成功后，生成token并返回
	userApi.TokenNext(ctx, user)
}

func (userApi *UserApi) TokenNext(ctx *gin.Context, user database.User) {
	if user.Freeze {
		response.FailWithMessage("账户已冻结，请联系管理员解冻", ctx)
		return
	}
	baseClaims := request.BaseClaims{
		UserID: user.ID,
		UUID:   user.UUID,
		RoleID: user.RoleID,
	}
	j := utils.NewJWT()
	// 创建访问令牌
	accessClaims := j.CreateAccessClaims(baseClaims)
	accessToken, err := j.CreateAccessToken(accessClaims)
	if err != nil {
		global.Log.Error("Failed to create access token:", zap.Error(err))
		response.FailWithMessage("AccessToken 生成失败", ctx)
		return
	}
	// 创建刷新令牌
	refreshClaims := j.CreateRefreshClaims(baseClaims)
	refreshToken, err := j.CreateRefreshToken(refreshClaims)
	if err != nil {
		global.Log.Error("Failed to create refresh token:", zap.Error(err))
		response.FailWithMessage("RefreshToken 生成失败", ctx)
		return
	}

	// 是否开启了多地点登录拦截 未开启
	if !global.Config.System.UseMultipoint {
		utils.SetRefreshToken(ctx, refreshToken, int(refreshClaims.ExpiresAt.Unix()-time.Now().Unix()))
		ctx.Set("user_id", user.ID)
		response.OkWithDetailed(response.Login{
			User:                 user,
			AccessToken:          accessToken,
			AccessTokenExpiresAt: int64(accessClaims.ExpiresAt.Unix()) * 100,
		}, "登录成功", ctx)
		return
	}
	// 开启了多地点登录拦截 开启的话，需要判断是否有多个JWT,如果有，需要删除旧的JWT
	// 检查redis中是否已存在该用户的jwt
	if jwtStr, err := jwtService.GetRedisJWT(user.UUID); errors.Is(err, redis.Nil) {
		// 不存在就设置新的
		if err = jwtService.SetRedisJWT(refreshToken, user.UUID); err != nil {
			global.Log.Error("Failed to set login status:", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", ctx)
			return
		}
		utils.SetRefreshToken(ctx, refreshToken, int(refreshClaims.ExpiresAt.Unix()-time.Now().Unix()))
		ctx.Set("user_id", user.ID)
		response.OkWithDetailed(response.Login{
			User:                 user,
			AccessToken:          accessToken,
			AccessTokenExpiresAt: int64(accessClaims.ExpiresAt.Unix()) * 100,
		}, "登录成功", ctx)
	} else if err != nil {
		global.Log.Error("Failed to get login status:", zap.Error(err))
		response.FailWithMessage("获取登录状态失败", ctx)
		return
	} else {
		var blacklist database.JwtBlacklist
		blacklist.Jwt = jwtStr
		if err := jwtService.JoinInBlacklist(blacklist); err != nil {
			global.Log.Error("Failed to join in blacklist:", zap.Error(err))
			response.FailWithMessage("加入黑名单失败", ctx)
			return
		}
		if err := jwtService.SetRedisJWT(refreshToken, user.UUID); err != nil {
			global.Log.Error("Failed to set login status:", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", ctx)
			return
		}

		utils.SetRefreshToken(ctx, refreshToken, int(refreshClaims.ExpiresAt.Unix()-time.Now().Unix()))
		ctx.Set("user_id", user.ID)
		response.OkWithDetailed(response.Login{
			User:                 user,
			AccessToken:          accessToken,
			AccessTokenExpiresAt: int64(accessClaims.ExpiresAt.Unix()) * 100,
		}, "登录成功", ctx)
	}

}
