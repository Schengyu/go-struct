package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义统一的返回格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // omitempty: 当 Data 为空时，JSON 不包含该字段
}

// 成功返回
func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "成功",
		Data:    data,
	})
}

// 失败返回
func Fail(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
	})
}
