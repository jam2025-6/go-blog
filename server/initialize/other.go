package initialize

import (
	"os"
	"server/global"
	"server/utils"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
)

// 执行其他配置初始化
func OtherInit() {
	// 解析刷新令牌过期时间
	refreshTokenExpiry, err := utils.ParseDuration(global.Config.JWT.RefreshTokenExpiryTime)
	if err != nil {
		global.Log.Error("Failed to parse refresh token expiry time configuration:", zap.Error(err))
		os.Exit(1)
	}

	_, err = utils.ParseDuration(global.Config.JWT.AccessTokenExpiryTime)
	if err != nil {
		global.Log.Error("Failed to parse access token expiry time configuration:", zap.Error(err))
		os.Exit(1)
	}

	// 让缓存中的黑名单条目和它所对应的 JWT/Refresh Token 在业务上同时失效
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(refreshTokenExpiry),
	)
}
