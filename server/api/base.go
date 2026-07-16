package api

import (
	"server/global"
	"server/model/request"
	"server/model/response"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type BaseApi struct{}

var store = base64Captcha.DefaultMemStore

func (baseApi *BaseApi) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(
		global.Config.Captcha.Height,   //图片高度
		global.Config.Captcha.Width,    // 图片宽度
		global.Config.Captcha.Length,   // 验证码数字个数
		global.Config.Captcha.MaxSkew,  // 最大倾斜角度
		global.Config.Captcha.DotCount, // 干扰点数量
	)
	// 创建验证码对象
	captcha := base64Captcha.NewCaptcha(driver, store)
	// 生成验证码
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		global.Log.Error("生成验证码失败:", zap.Error(err))
		response.FailWithMessage("生成验证码失败", c)
		return
	}
	response.OkWithData(response.Captcha{
		CaptchaID: id,
		PicPath:   b64s,
	}, c)
}

func (baseApi *BaseApi) SendEmailVerificationCode(c *gin.Context) {
	var req request.SendEmailVerificationCode
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if store.Verify(req.CaptchaID, req.Captcha, true) {
		err = baseService.SendEmailVerificationCode(c, req.Email)
		if err != nil {
			global.Log.Error("邮件发送失败：", zap.Error(err))
			response.FailWithMessage("邮件发送失败", c)
			return
		}
		response.OkWithMessage("邮件发送成功", c)
		return
	}
	response.FailWithMessage("验证码错误", c)
}
