package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = 200
	ERROR   = 500
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

// 成功 且无数据
func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "success", c)
}

// 成功 返回自定义信息， 无数据
func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

// 成功，且返回数据
func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}

// 成功，返回自定义数据、提示
func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

// 失败
func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "failure", c)
}

// 失败 返回自定义信息， 无数据
func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

// 成功，返回自定义数据、提示
func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

// 用于在用户未登录或身份验证失败时返回统一的错误响应
func NoAuth(message string, c *gin.Context) {
	Result(ERROR, gin.H{"reload": true}, message, c)
}

// 自定义的权限不足响应函数，用于在用户已登录但无权访问某个资源时返回 HTTP 403 错误。
func Forbidden(message string, c *gin.Context) {
	c.JSON(http.StatusForbidden, Response{
		Code: ERROR,
		Data: nil,
		Msg:  message,
	})
}
