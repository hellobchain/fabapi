/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package fabsdkimpl

import (
	"github.com/wsw365904/fabapi/internal/pkg/config/fabconfig"
	"github.com/wsw365904/fabapi/pkg/utils"

	"github.com/wsw365904/wswlog/wlogging"

	internalfabsdk "github.com/wsw365904/fabapi/core/fabsdk"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

var _ internalfabsdk.FabSdk = (*FabSdkOp)(nil)

var logger = wlogging.MustGetLoggerWithoutName()

type FabSdkOp struct {
	sdk *fabsdk.FabricSDK
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewFabSdkOp() internalfabsdk.FabSdk {
	logger.Debug("NewFabSdkOp enter")
	return &FabSdkOp{}
}

// 安装sdk
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (f *FabSdkOp) SetupFabricSDK(configFile string) (*fabsdk.FabricSDK, error) {
	logger.Debug("SetupFabricSDK enter")
	logger.Debug("set up fab sdk enter", "path", configFile)
	sdk, err := fabsdk.New(config.FromFile(configFile)) // 加载network.yaml文件
	if err != nil {
		return nil, utils.ToErr(err, "new fab sdk failed")
	}
	f.sdk = sdk
	return sdk, nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (f *FabSdkOp) Close() {
	logger.Debug("Close enter")
	logger.Debug("close fab sdk")
	if f != nil {
		if f.sdk != nil {
			f.sdk.Close()
		}
	}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewFabSdk() (internalfabsdk.FabSdk, *fabsdk.FabricSDK, error) {
	var fabSdk = NewFabSdkOp()
	sdk, err := fabSdk.SetupFabricSDK(fabconfig.GetNetworkYaml())
	if err != nil {
		logger.Error("SetupFabricSDK err:", err)
		return nil, nil, err
	}
	return fabSdk, sdk, nil
}
