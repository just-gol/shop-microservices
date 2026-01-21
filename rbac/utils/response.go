package utils

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

type Response struct {
	Code int         `json:"code"`           // 状态码：200 成功，500 错误
	Msg  string      `json:"msg"`            // 提示信息
	Data interface{} `json:"data,omitempty"` // 响应数据，可为空时省略
}

// Success 成功响应（可带 data，可不带）
func Success(ctx *gin.Context, msg string, data interface{}) {
	resp := Response{
		Code: 200,
		Msg:  msg,
	}
	if data != "" {
		resp.Data = data
	}
	ctx.JSON(http.StatusOK, resp)
}

// Error 错误响应
func Error(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, Response{
		Code: 500,
		Msg:  msg,
	})
}
