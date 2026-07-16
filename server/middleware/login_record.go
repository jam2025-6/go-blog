package middleware

import (
	"server/global"
	"server/model/database"
	"server/service"

	"github.com/gin-gonic/gin"
	"github.com/ua-parser/uap-go/uaparser" // 正确导入
	"go.uber.org/zap"
)

func LoginRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		go func() {
			gaodeService := service.ServiceGroupApp.GaodeService
			var userID uint
			var address string
			ip := ctx.ClientIP()
			loginMethod := ctx.DefaultQuery("flag", "email") // 弱未传递flag参数，则默认为email登录
			userAgent := ctx.Request.UserAgent()

			if value, exists := ctx.Get("user_id"); exists {
				if id, ok := value.(uint); ok {
					userID = id
				}
			}

			// 获取用户Ip的地理位置
			address = getAddressFromIP(ip, gaodeService)

			// 解析用户浏览器、操作系统和设备信息
			os, device, browser := paresUserAgent(userAgent)

			login := database.Login{
				UserID:      userID,
				LoginMethod: loginMethod,
				IP:          ip,
				Address:     address,
				OS:          os,
				DeviceInfo:  device,
				BrowserInfo: browser,
				Status:      ctx.Writer.Status(),
			}

			if err := global.DB.Create(&login).Error; err != nil {
				global.Log.Error("Failed to record login", zap.Error(err))
			}

		}()
	}
}

// 获取IP地址对应的地理位置信息
func getAddressFromIP(ip string, gaodeService service.GaodeService) string {
	res, err := gaodeService.GetLocationByIP(ip)
	if err != nil || res.Province == "" {
		return "未知"
	}
	if res.City != "" && res.Province != res.City {
		return res.Province + "-" + res.City
	}
	return res.Province
}

// 解析用户代理（User-Agent）字符串，提取操作系统、设备信息和浏览器信息
func paresUserAgent(userAgent string) (os, device, browser string) {
	os = userAgent
	device = userAgent
	browser = userAgent
	// 创建一个预加载了完整设备解析规则的解析器实例
	parser := uaparser.NewFromSaved()
	cli := parser.Parse(userAgent)
	os = cli.Os.Family
	device = cli.Device.Family
	browser = cli.UserAgent.Family
	return
}
