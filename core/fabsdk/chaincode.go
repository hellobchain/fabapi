/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package fabsdk

import (
	"github.com/wsw365904/fabapi/core/fabsdk/models"
)

type Chaincode interface {
	// 打包链码
	PackageCC(cCName, cCVersion, cCpath string, chaincodeType string) (string, interface{}, string, error)

	// 安装链码
	InstallCC(ccName, ccVersion string, ret interface{}, orgsInfos []*models.PeerOrgInfo) (string, error)

	// 批准链码 "OutOf('1', 'AMSP.member', 'BMSP.member')"
	ApproveCC(packageID string, ccName, ccVersion string, sequence int64, chaincodePolicy string, isInit bool, channelID string, orgsInfos []*models.PeerOrgInfo) error

	// 提交链码 "OutOf('1', 'AMSP.member', 'BMSP.member')"
	CommitCC(ccName, ccVersion string, sequence int64, chaincodePolicy string, isInit bool, channelID string, orgsInfos []*models.PeerOrgInfo) error

	// 升级链码
	UpgradeCC(packageID string, ccName, ccVersion string, sequence int64, chaincodePolicy string, isInit bool, channelID string, orgsInfos []*models.PeerOrgInfo) error

	// 初始化链码
	InitCC(ccName string, channelID string, args []string, orgInfo *models.PeerOrgInfo) (string, error)

	// 调用链码
	InvokeCC(ccName string, channelID string, args []string, orgInfo *models.PeerOrgInfo) (string, error)

	// 查询链码
	QueryCC(ccName string, channelID string, args []string, orgInfo *models.PeerOrgInfo) (string, error)
}
