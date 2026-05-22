package router

import (
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (r *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	// 创建子路由组，路径前缀为 "base"
	// baseRouter := Router.Group("base")
}
