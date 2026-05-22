package config

// 邮箱配置
type Email struct {
	Host     string `json:"host" yaml:"host"`         // 邮件服务器主机名
	Port     int    `json:"port" yaml:"port"`         // 邮件服务器端口号
	From     string `json:"from" yaml:"from"`         // 发送方邮箱
	Nickname string `json:"nickname" yaml:"nickname"` // 发送方昵称
	Secret   string `json:"secret" yaml:"secret"`     // 发送方邮箱密码
	IsSSL    bool   `json:"is_ssl" yaml:"is_ssl"`     // 是否使用 SSL 加密
}
