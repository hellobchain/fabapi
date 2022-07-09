/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package chaincodeserviceimpl

import (
	internalfabsdk "github.com/wsw365904/fabapi/core/fabsdk"
	"github.com/wsw365904/fabapi/core/fabsdk/models"
	"github.com/wsw365904/fabapi/internal/app/service"

	"github.com/wsw365904/wswlog/wlogging"
)

var _ service.ChaincodeService = (*ChaincodeService)(nil)

var logger = wlogging.MustGetLoggerWithoutName()

type ChaincodeService struct {
	cop internalfabsdk.Chaincode
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewChaincodeService(cop internalfabsdk.Chaincode) service.ChaincodeService {
	logger.Debug("NewChaincodeService enter")
	return &ChaincodeService{
		cop: cop,
	}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeService) Close() {
	logger.Debug("ChaincodeService Close enter")
	if c == nil {
		return
	}
	c = nil
}

// package chaincode
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeService) PackageChaincode(chaincodeInfo *models.ChaincodeInfo) (string, interface{}, string, error) {
	logger.Debug("PackageChaincode enter service")
	return c.cop.PackageCC(chaincodeInfo.ChaincodeID, chaincodeInfo.ChaincodeVersion, chaincodeInfo.ChaincodePath, chaincodeInfo.ChaincodeType)
}

// install chaincode
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeService) InstallChaincode(chaincodeInfo *models.ChaincodeInfo, orgs []*models.PeerOrgInfo) (string, error) {
	logger.Debug("InstallChaincode enter service")
	return c.cop.InstallCC(chaincodeInfo.ChaincodeID, chaincodeInfo.ChaincodeVersion, chaincodeInfo.PackagePara, orgs)
}

// approve chaincode
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeService) ApproveChaincode(chaincodeInfo *models.ChaincodeInfo, orgs []*models.PeerOrgInfo, channelID string) error {
	logger.Debug("ApproveChaincode enter service")
	return c.cop.ApproveCC(chaincodeInfo.PackageID, chaincodeInfo.ChaincodeID, chaincodeInfo.ChaincodeVersion, chaincodeInfo.Sequence, chaincodeInfo.ChaincodePolicy, chaincodeInfo.IsInit, channelID, orgs)
}

// upgrade chaincode
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeService) UpgradeChaincode(chaincodeInfo *models.ChaincodeInfo, orgs []*models.PeerOrgInfo, channelID string) error {
	logger.Debug("ApproveChaincode enter service")
	return c.cop.UpgradeCC(chaincodeInfo.PackageID, chaincodeInfo.ChaincodeID, chaincodeInfo.ChaincodeVersion, chaincodeInfo.Sequence, chaincodeInfo.ChaincodePolicy, chaincodeInfo.IsInit, channelID, orgs)
}

// commit chaincode
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeService) CommitChaincode(chaincodeInfo *models.ChaincodeInfo, orgs []*models.PeerOrgInfo, channelID string) error {
	logger.Debug("CommitChaincode enter service")
	return c.cop.CommitCC(chaincodeInfo.ChaincodeID, chaincodeInfo.ChaincodeVersion, chaincodeInfo.Sequence, chaincodeInfo.ChaincodePolicy, chaincodeInfo.IsInit, channelID, orgs)
}

// init chaincode
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeService) InitChaincode(chaincodeInfo *models.ChaincodeInfo, orgs []*models.PeerOrgInfo, channelID string) (string, error) {
	logger.Debug("InitChaincode enter service")
	return c.cop.InitCC(chaincodeInfo.ChaincodeID, channelID, chaincodeInfo.Args, orgs[0])
}

// invoke chaincode
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeService) InvokeChaincode(chaincodeInfo *models.ChaincodeInfo, orgs []*models.PeerOrgInfo, channelID string) (string, error) {
	logger.Debug("InvokeChaincode enter service")
	return c.cop.InvokeCC(chaincodeInfo.ChaincodeID, channelID, chaincodeInfo.Args, orgs[0])
}

// query chaincode
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeService) QueryChaincode(chaincodeInfo *models.ChaincodeInfo, orgs []*models.PeerOrgInfo, channelID string) (string, error) {
	logger.Debug("QueryChaincode enter service")
	return c.cop.QueryCC(chaincodeInfo.ChaincodeID, channelID, chaincodeInfo.Args, orgs[0])
}
