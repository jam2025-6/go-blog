package utils

import (
	"net"
	"server/global"
	"server/model/appTypes"
	"server/model/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetAccessToken 从请求头获取Access Token
func GetAccessToken(c *gin.Context) string {
	token := c.Request.Header.Get("x-access-token")
	return token
}

// GetRefreshToken 从cookie获取Refresh Token
func GetRefreshToken(c *gin.Context) string {
	token, _ := c.Cookie("x-refresh-token")
	return token
}

// SetRefreshToken 设置Refresh Token的cookie
func SetRefreshToken(c *gin.Context, token string, maxAge int) {
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	setCookie(c, "x-refresh-token", token, maxAge, host)
}

// ClearRefreshToken 清除Refresh Token的cookie
func ClearRefreshToken(c *gin.Context) {
	// net.SplitHostPort(...)	将 "host:port" 格式的字符串拆分为 host 和 port 两部分
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	setCookie(c, "x-refresh-token", "", -1, host)
}

// setCookie 设置指定名称和值的cookie
func setCookie(c *gin.Context, name, value string, maxAge int, host string) {
	if net.ParseIP(host) != nil {
		c.SetCookie(name, value, maxAge, "/", "", false, true)
	} else {
		c.SetCookie(name, value, maxAge, "/", host, false, true)
	}
}

// 从token中通过token获取claims
func GetClaims(c *gin.Context) (*request.JwtCustomClaims, error) {
	token := GetAccessToken(c)
	j := NewJWT()
	claims, err := j.ParseAccessToken(token)
	if err != nil {
		// 如果解析失败，记录错误日志
		global.Log.Error("Failed to retrieve JWT parsing information from Gin's Context. Please check if the request header contains 'x-access-token' and if the claims structure is correct.", zap.Error(err))
	}
	return claims, nil
}

func GetRoleID(c *gin.Context) appTypes.RoleID {
	if claims, exists := c.Get("claims"); !exists {
		if el, err := GetClaims(c); err != nil {
			return 0
		} else {
			return el.RoleID
		}
	} else {
		waitUser := claims.(*request.JwtCustomClaims)
		return waitUser.RoleID
	}
}
