/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package logcontroller

import (
	"net/http"

	"github.com/wsw365904/wswlog/wlogging"

	"fabapi/core/common/e"
	"fabapi/core/common/gintool"
	"fabapi/core/common/json"
	"fabapi/core/fabsdk/models"
	"fabapi/internal/app/service"
	"fabapi/pkg/utils"

	"github.com/gin-gonic/gin"
)

var logger = wlogging.MustGetLoggerWithoutName()

type LogController struct {
	log service.LogService
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewLogController(log service.LogService) *LogController {
	return &LogController{
		log: log,
	}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LogController) Close() {
	logger.Debug("LogController enter close")
	if l == nil {
		return
	}
	l.log.Close()
	l = nil
}

//
type setLogLevelRequest struct {
	LogLevel string `json:"loglevel,required" example:"info" binding:"required"`
}

// setLogLevel godoc
// @Summary 设置日志级别
// @Description 功能：设置日志级别
// @Tags 设置日志级别
// @Accept  json
// @Produce  json
// @Param resource body  log.setLogLevelRequest true "设置日志级别"
// @Success 200 {object} gintool.ApiResponse
// @Failure 400 {object} gintool.ApiResponse
// @Router /log/setloglevel [post]
func (l *LogController) setLogLevel(ctx *gin.Context) {
	logger.Debug("setLogLevel enter controller")
	req := new(setLogLevelRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("SetLogLevel req\n", string(ret))
	l.log.SetLogLevel(req.LogLevel)
	gintool.ResultCodeWithData(ctx, nil, nil)
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
// 日志管理路由组
func FabLogRouterGroup(app *LogController) gintool.RouteGroup {
	return gintool.RouteGroup{
		Route: []gintool.Route{
			gintool.NewRoute(http.MethodPost, "/setloglevel", app.setLogLevel).AddComment("设置日志级别"),
		},
		Prefix:  "/log",
		Comment: "日志管理",
		Module:  "log",
	}
}
