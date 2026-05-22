package config

type JWT struct {
	AccessTokenSecret      string `json:"access_token_secret" yaml:"access_token_secret"`             // 访问令牌密钥
	RefreshTokenSecret     string `json:"refresh_token_secret" yaml:"refresh_token_secret"`           // 刷新令牌密钥
	AccessTokenExpiryTime  string `json:"access_token_expiry_time" yaml:"access_token_expiry_time"`   // 访问令牌过期时间，例如 "2h"
	RefreshTokenExpiryTime string `json:"refresh_token_expiry_time" yaml:"refresh_token_expiry_time"` // 刷新令牌过期时间，例如 "7d"
	Issuer                 string `json:"issuer" yaml:"issuer"`                                       // JWT 发布者（Audience）名称
}
