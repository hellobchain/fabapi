/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package fabconfig

import (
	"os"
	"strings"

	"github.com/wsw365904/wswlog/wlogging"

	"github.com/spf13/viper"
)

// 环境变量开头
var logger = wlogging.MustGetLoggerWithoutName()

// 设置环境变量

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func setEnvVariables() {
	logger.Debug("setEnvVariables enter")
	// For environment variables.
	viper.SetEnvPrefix(cmdPre)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func GetNetworkYaml() string {
	logger.Debug("GetNetworkYaml enter")
	path, ok := os.LookupEnv(networkPathEnv)
	if !ok {
		path = defaultNetworkPath
		logger.Warn("no set env variable:", networkPathEnv, "default value:", path)
	} else {
		logger.Debug("network path", path)
	}
	return path
}
