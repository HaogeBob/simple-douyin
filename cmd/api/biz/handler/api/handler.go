package api

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/simple/douyin/pkg/errno"
)

type Response struct {
	Code    int64       `json:"status_code"`
	Message string      `json:"status_msg"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, Response{
		Code:    int64(Err.ErrCode),
		Message: Err.ErrMsg,
		Data:    data,
	})
}
