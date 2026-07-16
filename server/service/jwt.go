package service

import (
	"context"
	"server/global"
	"server/model/database"
	"server/utils"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

type JwtService struct {
}

// SetRedisJWT 将JWT存储到Redis中
func (jwtService *JwtService) SetRedisJWT(jwt string, uuid uuid.UUID) error {
	//
	dr, err := utils.ParseDuration(global.Config.JWT.RefreshTokenExpiryTime)
	if err != nil {
		return err
	}
	// 设置JWT在Redis中的过期时间
	return global.Redis.Set(context.Background(), uuid.String(), jwt, dr).Err()
}

// GetRedisJWT 从Redis中获取JWT
func (jwtService *JwtService) GetRedisJWT(uuid uuid.UUID) (string, error) {
	return global.Redis.Get(context.Background(), uuid.String()).Result()
}

// JsonInBlacklist 将JWT添加到黑名单
func (jwtService *JwtService) JoinInBlacklist(jwtList database.JwtBlacklist) error {
	if err := global.DB.Create(&jwtList).Error; err != nil {
		return err
	}
	// 将JWT添加到内存中的黑名单缓存
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return nil
}

// IsBlacklist 检查JWT是否在黑名单中
func (jwtService *JwtService) IsInBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

// LoadAll 从数据库加载所有的JWT黑名单并加入缓存
func LoadAll() {
	var data []string
	// 从数据库中获取所有的黑名单JWT
	// .Pluck("jwt", &data)	查询指定列（jwt 字段）的所有值，存入 data 变量
	if err := global.DB.Model(&database.JwtBlacklist{}).Pluck("jwt", &data).Error; err != nil {
		global.Log.Error("Failed to load JWT blacklist from the database", zap.Error(err))
		return
	}
	// 将所有的JWT添加到BlackCache缓存中
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	}
}
