/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package utils

import (
	"sync"

	"github.com/wsw365904/wswlog/wlogging"

	"fabapi/core/fabsdk/models"
	"fabapi/pkg/utils"

	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

var logger = wlogging.MustGetLoggerWithoutName()

// 获取组织信息
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func GetOrgInfo(org *models.PeerOrgInfo, sdk *fabsdk.FabricSDK) (*models.OrgInfo, error) {
	logger.Debug("GetOrgInfo enter")
	logger.Debugf("获取组织信息（%v）", org.OrgName)
	peerClient, err := models.NewPeerClient(sdk, org.OrgName, org.OrgAdminUser)
	if err != nil {
		return nil, utils.ToErr(err, "newPeerClient")
	}
	return models.NewOrgInfo(peerClient, org.OrgPeerDomains, org.OrgName, org.OrgMspId), nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func WaitAndErgodicErrChan(errChan chan error, loopNum int, wg *sync.WaitGroup) error {
	logger.Debug("deal with asy err")
	wg.Wait()
	for i := 0; i < loopNum; i++ {
		err := <-errChan
		if err != nil {
			return err
		}
	}
	return nil
}
