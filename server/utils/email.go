package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"server/global"
	"strings"

	"github.com/jordan-wright/email"
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

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error

	hostAddr := fmt.Sprintf("%s:%d", host, port)

	// 根据配置的是否使用 SSL 来选择邮件发送方法
	if isSSL {
		// 使用带 TLS 的邮件发送
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		// 使用普通的邮件发送
		err = e.Send(hostAddr, auth)
	}
	// e.Send()
	return err
}
