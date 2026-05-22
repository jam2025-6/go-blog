package initialize

import (
	"os"
	"server/global"
	"server/task"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

func (z *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	z.logger.Info(msg, zap.Any("keyAndValues", keysAndValues))
}

func (z *ZapLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	z.logger.Error(msg, zap.Error(err), zap.Any("keyAndValues", keysAndValues))
}

func NewZapLogger() *ZapLogger {
	return &ZapLogger{
		logger: global.Log,
	}
}

// InitCron 初始化定时任务
func InitCron() {
	// 创建一个默认配置的调度器
	c := cron.New(cron.WithLogger(NewZapLogger()))
	err := task.RegisterScheduledTasks(c)
	if err != nil {
		global.Log.Error("Error scheduling cron job:", zap.Error(err))
		os.Exit(1)
	}
	c.Start()
}
