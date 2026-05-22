package config

import (
	"strconv"
	"strings"

	"gorm.io/gorm/logger"
)

type Mysql struct {
	Host         string `json:"host" yaml:"host"`                     // 数据库机地址
	Port         int    `json:"port" yaml:"port"`                     // 数据库端口
	Config       string `json:"config" yaml:"config"`                 // 数据库配置
	DBName       string `json:"db_name" yaml:"db_name"`               // 数据库名称
	Username     string `json:"username" yaml:"username"`             // 数据库用户名
	Password     string `json:"password" yaml:"password"`             // 数据库密码
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns"` // 数据库最大空闲连接数
	MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns"` // 数据库最大打开连接数
	LogMode      string `json:"log_mode" yaml:"log_mode"`             // 数据库日志模式
}

func (m Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DBName + "?" + m.Config
}

func (m Mysql) LogLevel() logger.LogLevel {
	switch strings.ToLower(m.LogMode) {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}
