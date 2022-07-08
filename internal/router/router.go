/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package router

import (
	"fabapi/core/fabsdk/models"
	"fabapi/internal/app/controller/allcontroller"
)

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func Run(arguments []string) {

	logger.Debug("Run enter")

	logger.Info("input args:", arguments)

	// controller层的东西
	app, other := newApp()
	// 启动
	_start(app, other)
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func Start(app *allcontroller.Controller, other *models.Other) {
	logger.Debug("Start enter")
	// 启动
	_start(app, other)
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func _start(app *allcontroller.Controller, other *models.Other) {
	logger.Debug("_start enter")

	// 监控程序停止信号，以此来处理资源的释放
	handleSignals(app)

	// 设置gin框架mode模式
	setGinMode(other)

	// gin框架路由的设置
	router := setGin()

	// 启动pprof
	startPprof(router, other)

	// swagger doc服务
	startSwaggerServer(router, other)

	// 加入路由到gin框架里
	addMyRouteGroups(router, app)

	// 启动服务
	startServer(router, other)
}
