package gintool

import (
	"time"

	"github.com/wsw365904/wswlog/wlogging"

	"github.com/wsw365904/fabapi/pkg/uuid"

	"github.com/gin-gonic/gin"
)

var logClient = wlogging.MustGetLoggerWithoutName()

//请求日志汇总信息
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := uuid.GetUUIDInt()
		// 开始时间
		start := time.Now()
		// path
		path := c.Request.URL.Path
		// ip
		clientIP := c.ClientIP()
		// 方法
		method := c.Request.Method
		logClient.Debugf("| %10d | %3v | %13v | %15s | %s  %s |",
			requestId,
			"",
			"",
			clientIP,
			method,
			path,
		)

		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		// 执行时间
		latency := end.Sub(start)
		// 状态
		statusCode := c.Writer.Status()
		logClient.Infof("| %10d | %3d | %13v | %15s | %s  %s |",
			requestId,
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)
	}
}
