package config

type ES struct {
	URL            string `json:"url" yaml:"url"`                           // Elasticsearch 服务器 URL
	Username       string `json:"username" yaml:"username"`                 // Elasticsearch 用户名
	Password       string `json:"password" yaml:"password"`                 // Elasticsearch 密码
	IsConsolePrint bool   `json:"is_console_print" yaml:"is_console_print"` // 是否在控制台打印日志
}
