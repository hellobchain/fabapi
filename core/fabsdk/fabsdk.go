/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package fabsdk

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

type FabSdk interface {
	SetupFabricSDK(configFile string) (*fabsdk.FabricSDK, error)

	Close()
}
