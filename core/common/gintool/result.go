package gintool

import (
	"net/http"

	"github.com/wsw365904/wswlog/wlogging"

	"github.com/wsw365904/fabapi/core/common/e"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	Success      = 0
	Unauthorized = 401
	Fail         = 400
)

var logger = wlogging.MustGetLoggerWithoutName()

type DataList struct {
	List      interface{} `json:"list"` //列表
	SubList   interface{} `json:"sub_list,omitempty"`
	PageNum   int         `json:"page_num"`       //当前页码
	PageSize  int         `json:"page_size"`      //单页条数
	PageCount int         `json:"page_count"`     //总页数
	Total     int64       `json:"total"`          //总数
	Cols      interface{} `json:"cols"`           //检索字段
	Note      interface{} `json:"note,omitempty"` //注解
}

type ApiResponse struct {
	Code    e.ErrCode   `json:"code"`              // 状态码
	Message string      `json:"message,omitempty"` // 状态短语
	Result  interface{} `json:"result,omitempty"`  // 数据结果集
	ErrMsg  string      `json:"err_msg,omitempty"` // 内部错误详情
}

func responseOutput(c *gin.Context, code e.ErrCode, errMsg string, result interface{}) {
	var httpCode = http.StatusOK
	if code != e.SUCCESS {
		httpCode = http.StatusBadRequest
	}
	c.JSON(httpCode, ApiResponse{
		Code:    code,
		Message: code.Error(),
		Result:  result,
		ErrMsg:  errMsg,
	})
}

func ResultCodeWithData(ctx *gin.Context, err error, data interface{}) {
	var code e.ErrCode
	var errMsg string
	if err == nil {
		code = e.SUCCESS
	} else {
		errMsg = err.Error()
		err1 := errors.Cause(err)
		switch err1.(type) {
		case e.ErrCode:
			code = err1.(e.ErrCode)
		default:
			logger.Error("http 返回状态code 类型错误")
			code = e.ERROR
		}
	}

	responseOutput(ctx, code, errMsg, data)
}
