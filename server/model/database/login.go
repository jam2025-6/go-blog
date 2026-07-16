package database

import "server/global"

type Login struct {
	global.MODEL
	UserID      uint   `json:"user_id"`                       // 用户ID
	User        User   `json:"user" gorm:"foreignKey:UserID"` // 用户信息
	LoginMethod string `json:"login_method"`                  // 登录方式
	IP          string `json:"ip"`                            // IP地址
	Address     string `json:"address"`                       // 地址
	OS          string `json:"os"`                            // 操作系统
	DeviceInfo  string `json:"device_info"`                   // 设备信息
	BrowserInfo string `json:"browser_info"`                  // 浏览器信息
	Status      int    `json:"status"`                        // 登录状态
}
