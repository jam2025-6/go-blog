package config

type Config struct {
	System  System  `json:"system" yaml:"system"`
	Mysql   Mysql   `json:"mysql" yaml:"mysql"`
	Zap     Zap     `json:"zap" yaml:"zap"`
	JWT     JWT     `json:"jwt" yaml:"jwt"`
	Redis   Redis   `json:"redis" yaml:"redis`
	ES      ES      `json:"es" yaml:"es"`
	Upload  Upload  `json:"upload" yaml:"upload"`
	Website Website `json:"website" yaml:"website"`
	Email   Email   `json:"email" yaml:"email"`
	Captcha Captcha `json:"captcha" yaml:"captcha"`
	Gaode   Gaode   `json:"gaode" yaml:"gaode"`
}
