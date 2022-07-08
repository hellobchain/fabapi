/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package fabapi

import (
	"fmt"

	"github.com/wsw365904/wswlog/wlogging"

	"fabapi/internal/router"

	"github.com/urfave/cli"
)

var logger = wlogging.MustGetLoggerWithoutName()

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func run(ctx *cli.Context) {
	logger.Debug("run enter")
	router.Run(ctx.Args())
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func newApp() *cli.App {
	logger.Debug("newApp enter")
	app := cli.NewApp()
	app.Name = fmt.Sprintf("\napp name: %s", appName)
	app.Compiled = buildTime
	app.Version = fmt.Sprintf("\nversion:%s \ncommit version: %s \nbuild time: %s", version, commitVersion, buildTime)
	app.Action = run
	logger.Infof("%s%s", app.Name, app.Version)
	return app
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func newCliApp() *cli.App {
	logger.Debug("newCliApp enter")
	app := cli.NewApp()
	app.Name = fmt.Sprintf("\napp name: %s", appName)
	app.Compiled = buildTime
	app.Version = fmt.Sprintf("\nversion:%s \ncommit version: %s \nbuild time: %s", version, commitVersion, buildTime)
	app.Action = action
	logger.Infof("%s%s", app.Name, app.Version)
	return app
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func action(ctx *cli.Context) {
	logger.Info("action enter", ctx.Args())
}
