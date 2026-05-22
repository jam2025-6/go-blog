package utils

import (
	"server/global"
	"strings"
)

func Email(To, subject string, body string) error {
	to := strings.Split(To, ",")
	return send(to, subject, body)
}

func send(to []string, subject string, body string) error {
	emailCfg := global.Config.Email // 获取全局邮箱配置

	from := emailCfg.From
	nickname := emailCfg.Nickname
	secret := emailCfg.Secret
	host := emailCfg.Host
	port := emailCfg.Port
	isSSL := emailCfg.IsSSL
	return nil
}
