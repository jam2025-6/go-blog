package database

import "server/global"

type JwtBlacklist struct {
	global.MODEL
	Jwt string `json:"jwt" gorm:"type:text"` // Jwt
}
