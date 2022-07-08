/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package main

import (
	"fabapi/cmd/fabapi"
	"github.com/wsw365904/wswlog/wlogging"
)

var logger = wlogging.MustGetLoggerWithoutName()

// @Title fabapi API文档
// @Version 1.1.0
// @Description 描述:fabapi 接口文档.
// @Host localhost:6922
// @BasePath /fab

func main() {
	logger.Debug("main enter")
	fabapi.StartMain()
}
