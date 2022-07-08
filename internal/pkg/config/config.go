/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package config

import (
	"fabapi/core/fabsdk/models"
)

type Config interface {
	LoadConfig() (*models.Other, error) // 加载配置文件
	ReloadConfig() error                // 重新加载配置文件
	ReleaseConfig() error               // 释放配置文件
}
