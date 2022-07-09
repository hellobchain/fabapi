/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package router

import (
	"os"
	"syscall"

	"github.com/wsw365904/wswlog/wlogging"

	"github.com/wsw365904/fabapi/core/common/gintool"
	"github.com/wsw365904/fabapi/core/fabsdk/models"
	_ "github.com/wsw365904/fabapi/docs"
	"github.com/wsw365904/fabapi/internal/app/controller/allcontroller"
	"github.com/wsw365904/fabapi/internal/app/controller/chaincodecontroller"
	"github.com/wsw365904/fabapi/internal/app/controller/channelcontroller"
	"github.com/wsw365904/fabapi/internal/app/controller/ledgercontroller"
	"github.com/wsw365904/fabapi/internal/app/controller/logcontroller"
	"github.com/wsw365904/fabapi/pkg/utils"

	"github.com/gin-contrib/pprof"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var logger = wlogging.MustGetLoggerWithoutName()

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func newApp() (*allcontroller.Controller, *models.Other) {
	logger.Debug("newApp enter")
	app, other, err := allcontroller.New()
	if err != nil {
		logger.Errorf("NewController err: %v", err)
		os.Exit(0)
	}
	return app, other
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func addMyRouteGroups(router *gin.Engine, app *allcontroller.Controller) {
	logger.Debug("addMyRouteGroups enter")
	const prefix = "/fab"
	gintool.AddRouteGroups(router, prefix, routGroups(app))
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func setGin() *gin.Engine {
	logger.Debug("setGin enter")
	router := gin.New()
	router.Use(gintool.Logger()) // 设置路由日志
	router.Use(gin.Recovery())
	return router
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func startSwaggerServer(router *gin.Engine, other *models.Other) {
	logger.Debug("startSwaggerServer enter")
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	logger.Debugf("swag server(http://%v:%v/swagger/index.html)", "127.0.0.1", other.Port)
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func setGinMode(other *models.Other) {
	logger.Debug("setGinMode enter")
	var mn string
	switch other.Envs {
	case gin.DebugMode, gin.ReleaseMode, gin.TestMode:
		mn = other.Envs
	default:
		mn = gin.ReleaseMode
	}
	gin.SetMode(mn)
	logger.Debug("gin mode", mn)
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func handleSignals(app *allcontroller.Controller) {
	go utils.HandleSignals(utils.AddPlatformSignals(map[os.Signal]func(){
		syscall.SIGINT: func() {
			logger.Debug("释放资源")
			app.Close()
			os.Exit(0)
		},
		syscall.SIGTERM: func() {
			logger.Debug("释放资源")
			app.Close()
			os.Exit(0)
		},
	}, logger), logger)
}

// 批次路由
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func routGroups(app *allcontroller.Controller) []gintool.RouteGroup {
	logger.Debug("routGroups enter")
	return []gintool.RouteGroup{
		// fab chaincode
		chaincodecontroller.FabChaincodeRouterGroup(app.ChaincodeController),
		// fab channel
		channelcontroller.FabChannelRouterGroup(app.ChannelController),
		// fab ledger
		ledgercontroller.FabLedgerRouterGroup(app.LedgerController),
		// fab log
		logcontroller.FabLogRouterGroup(app.LogController),
	}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func startServer(router *gin.Engine, other *models.Other) {
	logger.Debug("startServer enter")
	err := router.Run(":" + other.Port)
	if err != nil {
		logger.Errorf("Server err: %v", err)
		os.Exit(-1)
	}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func startPprof(r *gin.Engine, other *models.Other) {
	// pprof start
	if other.PprofEnable {
		logger.Debug("startPprof enter")
		logger.Debugf("pprof server(http://localhost:%v/debug/pprof/)", other.Port)
		pprof.Register(r)
	}

}
