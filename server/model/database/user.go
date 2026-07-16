package database

import (
	"server/global"
	"server/model/appTypes"

	"github.com/gofrs/uuid"
)

type User struct {
	global.MODEL
	UUID      uuid.UUID         `json:"uuid" gorm:"type:char(36);unique"` // 用户唯一标识
	Username  string            `json:"username"`                         // 用户名
	Password  string            `json:"password"`                         // 密码
	Email     string            `json:"email"`                            // 邮箱
	Openid    string            `json:"openid"`                           // openid
	Avatar    string            `json:"avatar" gorm:"size:255"`           // 头像：邮箱注册的头像或 QQ 登录的空间头像
	Address   string            `json:"address"`                          // 地址
	Signature string            `json:"signature"`                        // 签名
	RoleID    appTypes.RoleID   `json:"role_id"`                          // 角色ID
	Register  appTypes.Register `json:"register"`                         // 注册来源
	Freeze    bool              `json:"freeze"`                           // 是否冻结账号
}
