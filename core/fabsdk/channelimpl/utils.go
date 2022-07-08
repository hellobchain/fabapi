/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package channelimpl

import (
	"sync"

	"github.com/wsw365904/wswlog/wlogging"

	"fabapi/core/fabsdk/models"
	internalutils "fabapi/core/fabsdk/utils"
	"fabapi/pkg/utils"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
)

var logger = wlogging.MustGetLoggerWithoutName()

// 通道

// 获取要加入通道的peer节点
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChannelOp) getJoinChannelPeers(orgInfo *models.OrgInfo, channelId string) ([]string, error) {
	logger.Debug("getJoinChannelPeers enter")
	joinChannelPeers := make([]string, 0)
	for _, peerName := range orgInfo.PeerNames {
		find := false
		channelsRes, err := orgInfo.OrgClient.OrgResMgmt.QueryChannels(resmgmt.WithTargetEndpoints(peerName))
		if err != nil {
			joinChannelPeers = nil
			return nil, utils.ToErr(err, "%s peers failed to queryChannel", orgInfo.OrgName)
		}
		logger.Debugf("peer name:%v channel:%v", peerName, channelsRes)
		for _, channelRes := range channelsRes.Channels {
			if channelId == channelRes.ChannelId {
				find = true
				break
			}
		}
		if !find {
			joinChannelPeers = append(joinChannelPeers, peerName)
		}
	}
	if len(joinChannelPeers) == 0 {
		joinChannelPeers = nil
		return nil, utils.ToErr(nil, "QueryChannels")
	}
	return joinChannelPeers, nil
}

// 为peer加入通道
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChannelOp) joinChannel(orgsInfos []*models.PeerOrgInfo, channelInfo *models.ChannelInfo, ordererOrgInfo *models.OrdererOrgInfo) error {
	logger.Debug("joinChannel enter")
	if c.isAsy {
		var wg sync.WaitGroup
		var errChan = make(chan error, len(orgsInfos))
		for _, org := range orgsInfos { // 循环组织
			logger.Debugf("%v join channel(%v) for org %v", org.OrgPeerDomains, channelInfo.ChannelID, org.OrgName)
			go c.asyJoinChannelForOrg(org, channelInfo, ordererOrgInfo, &wg, errChan) // 单个组织的成员加入通道
		}
		err := internalutils.WaitAndErgodicErrChan(errChan, len(orgsInfos), &wg)
		if err != nil {
			return utils.ToErr(err, "ergodicErrChan")
		}
	} else {
		for _, org := range orgsInfos { // 循环组织
			logger.Debugf("%v join channel(%v) for org %v", org.OrgPeerDomains, channelInfo.ChannelID, org.OrgName)
			err := c.joinChannelForOrg(org, channelInfo, ordererOrgInfo) // 单个组织的成员加入通道
			if err != nil {
				return utils.ToErr(err, "joinChannelForOrg")
			}
		}
	}
	return nil
}

// 为每个组织的peer加入通道
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChannelOp) joinChannelForOrg(org *models.PeerOrgInfo, channelInfo *models.ChannelInfo, ordererOrgInfo *models.OrdererOrgInfo) error {
	logger.Debug("joinChannelForOrg enter")
	orgInfo, err := internalutils.GetOrgInfo(org, c.sdk)
	if err != nil {
		return utils.ToErr(err, "getExpectTargetPeers")
	}
	defer models.Release(orgInfo.OrgClient)
	joinChannelPeers, err := c.getJoinChannelPeers(orgInfo, channelInfo.ChannelID)
	if err != nil {
		return utils.ToErr(err, "all peers(%v) have already join channel(%v)", orgInfo.PeerNames, channelInfo.ChannelID)
	}
	err = orgInfo.OrgClient.OrgResMgmt.JoinChannel(channelInfo.ChannelID, resmgmt.WithTargetEndpoints(joinChannelPeers...), resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(ordererOrgInfo.OrdererOrgEndpoint[0]))
	if err != nil {
		return utils.ToErr(err, "%s peers failed to JoinChannel", org.OrgName)
	}
	return nil
}

// 异步为每个组织的peer加入通道
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChannelOp) asyJoinChannelForOrg(org *models.PeerOrgInfo, channelInfo *models.ChannelInfo, ordererOrgInfo *models.OrdererOrgInfo, wg *sync.WaitGroup, errChan chan error) {
	logger.Debug("asyJoinChannelForOrg enter")
	wg.Add(1)
	defer wg.Done()
	orgInfo, err := internalutils.GetOrgInfo(org, c.sdk)
	if err != nil {
		errChan <- err
		return
	}
	defer models.Release(orgInfo.OrgClient)
	joinChannelPeers, err := c.getJoinChannelPeers(orgInfo, channelInfo.ChannelID)
	if err != nil {
		errChan <- utils.ToErr(err, "all peers(%v) have already join channel(%v)", orgInfo.PeerNames, channelInfo.ChannelID)
		return
	}
	err = orgInfo.OrgClient.OrgResMgmt.JoinChannel(channelInfo.ChannelID, resmgmt.WithTargetEndpoints(joinChannelPeers...), resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(ordererOrgInfo.OrdererOrgEndpoint[0]))
	if err != nil {
		errChan <- utils.ToErr(err, "%s peers failed to JoinChannel", org.OrgName)
		return
	}
	errChan <- nil
	return
}

// 为每个组织的peer更新锚节点
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChannelOp) updateAnchorForOrg(org *models.PeerOrgInfo, ordererOrgInfo *models.OrdererOrgInfo, channelInfo *models.ChannelInfo, peerClient *models.PeerClient, orgPeers []string, signId msp.SigningIdentity) error {
	logger.Debug("updateAnchorForOrg enter")
	req := resmgmt.SaveChannelRequest{
		ChannelID:         channelInfo.ChannelID,
		ChannelConfigPath: org.OrgAnchorFile,
		SigningIdentities: []msp.SigningIdentity{signId},
	}

	res, err := peerClient.OrgResMgmt.SaveChannel(req, resmgmt.WithTargetEndpoints(orgPeers...), resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(ordererOrgInfo.OrdererOrgEndpoint[0]))
	if err != nil {
		return utils.ToErr(err, "saveChannel for anchor org %s", org.OrgName)
	}
	logger.Debug("updateAnchorForOrg", "txId", res.TransactionID)
	return nil
}

// 异步为每个组织的peer更新锚节点
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChannelOp) asyUpdateAnchorForOrg(org *models.PeerOrgInfo, ordererOrgInfo *models.OrdererOrgInfo, channelInfo *models.ChannelInfo, peerClient *models.PeerClient, orgPeers []string, signId msp.SigningIdentity, wg *sync.WaitGroup, errChan chan error) {
	logger.Debug("asyUpdateAnchorForOrg enter")
	defer wg.Done()
	req := resmgmt.SaveChannelRequest{
		ChannelID:         channelInfo.ChannelID,
		ChannelConfigPath: org.OrgAnchorFile,
		SigningIdentities: []msp.SigningIdentity{signId},
	}

	res, err := peerClient.OrgResMgmt.SaveChannel(req, resmgmt.WithTargetEndpoints(orgPeers...), resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(ordererOrgInfo.OrdererOrgEndpoint[0]))
	if err != nil {
		errChan <- utils.ToErr(err, "saveChannel for anchor org %s", org.OrgName)
		return
	} else {
		logger.Debug("asyUpdateAnchorForOrg", "txId", res.TransactionID)
	}
	errChan <- nil
	return
}

// 获取peer clients
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChannelOp) getPeerClientsAndSignIds(orgsInfo []*models.PeerOrgInfo) ([]*models.PeerClient, []msp.SigningIdentity, error) {
	logger.Debug("getPeerClientsAndSignIds enter")
	var signIds []msp.SigningIdentity
	peerClients := make([]*models.PeerClient, 0)
	for _, org := range orgsInfo {
		peerClient, err := models.NewPeerClient(c.sdk, org.OrgName, org.OrgAdminUser)
		if err != nil {
			models.Release(peerClients)
			return nil, nil, utils.ToErr(err, "NewPeerClient error: %v", err)
		}
		peerClients = append(peerClients, peerClient)
		// Get signing identity that is used to sign create channel request
		orgSignId, err := peerClient.OrgMspClient.GetSigningIdentity(org.OrgAdminUser)
		if err != nil {
			models.Release(peerClients)
			return nil, nil, utils.ToErr(err, "GetSigningIdentity")
		}
		signIds = append(signIds, orgSignId)
	}
	return peerClients, signIds, nil
}

// 为组织创建通道
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChannelOp) createChannelForOrg(signIds []msp.SigningIdentity, ordererOrgInfo *models.OrdererOrgInfo, channelInfo *models.ChannelInfo) error {
	logger.Debug("createChannelForOrg enter")
	chMgmtClient, err := models.NewOrdererClient(c.sdk, ordererOrgInfo.OrdererOrgAdminUser, ordererOrgInfo.OrdererOrgName) // 某个orderer节点建立client
	if err != nil {
		return utils.ToErr(err, "NewOrdererClient")
	}
	defer models.Release(chMgmtClient)
	req := resmgmt.SaveChannelRequest{
		ChannelID:         channelInfo.ChannelID,
		ChannelConfigPath: channelInfo.ChannelConfig,
		SigningIdentities: signIds,
	}
	_, err = chMgmtClient.OrdererClient.SaveChannel(req, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(ordererOrgInfo.OrdererOrgEndpoint[0])) // 创建通道
	if err != nil {
		return utils.ToErr(err, "error should be nil for SaveChannel of orgchannel")
	}
	return nil
}

// 通道
