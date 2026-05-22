package initialize

import (
	"net/http"
	"server/global"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	// 设置gin模式
	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()
	// 创建一个将 Session 数据完全存储在客户端 Cookie 中的存储引擎
	var store = cookie.NewStore([]byte(global.Config.System.SessionsSecret))
	Router.Use(sessions.Sessions("session", store))
	// 将指定目录下的文件提供给客户端
	// "uploads" 是URL路径前缀，http.Dir("files")是实际文件系统中存储文件的目录
	Router.StaticFS("files", http.Dir(global.Config.Upload.Path))

	return Router
}
