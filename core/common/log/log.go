package log

import (
	"fmt"
	"strings"

	"github.com/wsw365904/wswlog/wlogging"
)

var logger = wlogging.MustGetLoggerWithoutName()

// 设置程序日志级别
func SetLogLevel(logLevel string) {
	logLevel = fmt.Sprintf("%s", strings.ToLower(logLevel))
	logger.Debug("日志级别", logLevel)
	wlogging.SetGlobalLogLevel(logLevel)
}
