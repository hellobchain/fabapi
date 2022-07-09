/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package fabconfig

import (
	"fmt"
	"os"

	"github.com/wsw365904/fabapi/core/common/json"
	"github.com/wsw365904/fabapi/core/common/log"
	"github.com/wsw365904/fabapi/core/fabsdk/models"
	"github.com/wsw365904/fabapi/internal/pkg/config"

	"github.com/spf13/viper"
)

var _ config.Config = (*fabConfig)(nil)

type fabServerEngine struct {
	Port        string `yaml:"Port"`
	Environment string `yaml:"Environment"`
	LogLevel    string `yaml:"LogLevel"`
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (f *fabServerEngine) validate() error {
	logger.Debug("fabapiEngine validate enter")
	if f.Port == "" {
		return fmt.Errorf("port is empty")
	}
	if f.Environment == "" {
		return fmt.Errorf("environment is empty")
	}
	return nil
}

type otherConfig struct {
	IsFile       bool `yaml:"IsFile"`
	IsAsy        bool `yaml:"IsAsy"`
	IsConcurrent bool `yaml:"IsConcurrent"`
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (o *otherConfig) validate() error {
	logger.Debug("otherConfig validate enter")
	return nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (f *fabConfig) validate() error {
	logger.Debug("fabConfig validate enter")
	err := f.Server.validate()
	if err != nil {
		return err
	}

	err = f.Other.validate()
	if err != nil {
		return err
	}
	return f.PprofConf.validate()
}

type pprofConf struct {
	Enable bool `yaml:"Enable"`
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (p *pprofConf) validate() error {

	return nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (f *fabConfig) loadConfig() error {
	logger.Debug("fabConfig loadConfig enter")
	logger.Debug("load config")
	viper.SetConfigName(fabConfigName)         // name of fab-config file
	path, ok := os.LookupEnv(fabConfigPathEnv) // 优先走环境变量
	if !ok {
		logger.Warn("no set env variable:", fabConfigPathEnv, "default value:", defaultFabConfigPath)
		viper.AddConfigPath(defaultFabConfigPath)
	} else {
		viper.AddConfigPath(path)
	}
	err := viper.ReadInConfig() // Find and read the fab-config.yaml file
	if err != nil {
		return fmt.Errorf("viper read config:%v ", err)
	}
	conf := viper.GetViper()
	err = conf.Unmarshal(f)
	if err != nil {
		return fmt.Errorf("unmarshal yaml data: %v ", err)
	}
	err = f.validate()
	if err != nil {
		return fmt.Errorf("validate yaml data: %v ", err)
	}
	log.SetLogLevel(f.Server.LogLevel) // 设置日志级别
	return nil
}

type fabConfig struct {
	Server    *fabServerEngine `yaml:"Server"`
	Other     *otherConfig     `yaml:"Other"`
	PprofConf *pprofConf       `yaml:"PprofConf"`
}

// 加载配置
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (f *fabConfig) LoadConfig() (*models.Other, error) {
	logger.Debug("fabConfig LoadConfig enter")
	setEnvVariables()
	err := f.loadConfig()
	if err != nil {
		return nil, err
	}
	return &models.Other{
		Envs:         f.Server.Environment,
		Port:         f.Server.Port,
		IsFile:       f.Other.IsFile,
		IsAsy:        f.Other.IsAsy,
		IsConcurrent: f.Other.IsConcurrent,
		PprofEnable:  f.PprofConf.Enable,
	}, nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewFabConfig() config.Config {
	logger.Debug("NewFabConfig enter")
	return &fabConfig{}
}

// 加载配置
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (f *fabConfig) ReloadConfig() error {
	logger.Debug("fabConfig ReloadConfig enter")
	setEnvVariables()
	return f.loadConfig()
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (f *fabConfig) ReleaseConfig() error {
	logger.Debug("fabConfig ReleaseConfig enter")
	return nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewConfig() (*models.Other, error) {
	fabConf := NewFabConfig()
	other, err := fabConf.LoadConfig()
	if err != nil {
		logger.Error("LoadConfig err:", err)
		return nil, err
	}
	res, _ := json.MarshalIndent(fabConf, models.Empty, models.TAB)
	logger.Debug("fab-server engine fab config:", string(res))

	err = fabConf.ReleaseConfig()
	if err != nil {
		logger.Error("ReleaseConfig err:%v", err)
		return nil, err
	}

	return other, nil
}
