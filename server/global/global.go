package global

import (
	"server/config"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config     *config.Config // 指向 Config 结构体的指针
	Log        *zap.Logger    // 指向 zap 日志器的指针
	DB         *gorm.DB
	BlackCache local_cache.Cache // 用于存储黑名单数据的本地缓存实例
	Redis      redis.Client
	ESClient   *elasticsearch.TypedClient
)
