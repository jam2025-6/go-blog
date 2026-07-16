package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (r *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	// 创建子路由组，路径前缀为 "base"
	baseRouter := Router.Group("base")
	baseApi := api.ApiGroupApp.BaseApi
	{
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.POST("sendEmailCode", baseApi.SendEmailVerificationCode)
		// baseRouter.GET("qqLoginURL")
	}
}
