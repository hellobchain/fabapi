/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package logserviceimpl

import (
	"github.com/wsw365904/fabapi/core/common/log"
	"github.com/wsw365904/fabapi/internal/app/service"

	"github.com/wsw365904/wswlog/wlogging"
)

var _ service.LogService = (*LogService)(nil)

var logger = wlogging.MustGetLoggerWithoutName()

type LogService struct{}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewLogService() service.LogService {
	logger.Debug("NewLogService enter")
	return &LogService{}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LogService) Close() {
	logger.Debug("LogService Close enter")
	if l == nil {
		return
	}
	l = nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LogService) SetLogLevel(logLevel string) {
	logger.Debug("setLogLevel enter service")
	log.SetLogLevel(logLevel)
}
