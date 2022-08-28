// Package channelimpl /***************************************************************
package channelimpl

import (
	"os"

	internalfabsdk "github.com/wsw365904/fabapi/core/fabsdk"
	"github.com/wsw365904/fabapi/core/fabsdk/models"
	"github.com/wsw365904/fabapi/pkg/utils"

	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// 通道操作

/*

1. 创建通道
2. 加入通道

*/
var _ internalfabsdk.Channel = (*ChannelOp)(nil)

type ChannelOp struct {
	isAsy bool
	sdk   *fabsdk.FabricSDK
}

// NewChannelOp /***************************************************************
func NewChannelOp(other *models.Other, sdk *fabsdk.FabricSDK) internalfabsdk.Channel {
	return &ChannelOp{
		isAsy: other.IsAsy,
		sdk:   sdk,
	}
}

// /////////////////////////////// 通道管理 开始///////////////////////////////////

// JoinChannel 加入通道
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChannelOp) JoinChannel(orgsInfos []*models.PeerOrgInfo, channelInfo *models.ChannelInfo, ordererOrgInfo *models.OrdererOrgInfo) error {
	logger.Debug("JoinChannel enter")
	logger.Debug("hello 加入通道 waiting......")
	if len(orgsInfos) == 0 {
		return utils.ToErr(nil, "通道组织不能为空，请提供组织信息")
	}
	err := c.joinChannel(orgsInfos, channelInfo, ordererOrgInfo)
	if err != nil {
		return utils.ToErr(err, "joinChannel")
	}
	logger.Debug("hello 加入通道成功")
	return nil
}

// CreateChannel 创建通道
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChannelOp) CreateChannel(orgsInfos []*models.PeerOrgInfo, ordererOrgInfo *models.OrdererOrgInfo, channelInfo *models.ChannelInfo) error {
	logger.Debug("CreateChannel enter")
	logger.Debug("hello 开始创建通道 waiting......")
	if len(orgsInfos) == 0 {
		return utils.ToErr(nil, "通道组织不能为空，请提供组织信息")
	}
	peerClients, signIds, err := c.getPeerClientsAndSignIds(orgsInfos) // 获得所有组织的签名信息
	if err != nil {
		return utils.ToErr(err, "获得所有组织的签名信息")
	}
	defer models.Release(peerClients) // 释放资源
	err = c.createChannelForOrg(signIds, ordererOrgInfo, channelInfo)
	if err != nil {
		return utils.ToErr(err, "创建通道失败")
	}
	logger.Debug("hello 创建通道完成")
	logger.Debug("hello 使用每个org的管理员身份更新锚节点配置...")
	for i, orgInfo := range orgsInfos {
		_, err := os.Stat(orgInfo.OrgAnchorFile) // 预先判断锚节点文件是否存在，不存在啥也不做
		if err != nil {
			logger.Warn(orgInfo.OrgAnchorFile, "is not exist")
			continue
		}
		logger.Debugf("%v update anchor channel(%v) for org %v", orgInfo.OrgPeerDomains, channelInfo.ChannelID, orgInfo.OrgName)
		err = c.updateAnchorForOrg(orgInfo, ordererOrgInfo, channelInfo, peerClients[i], orgInfo.OrgPeerDomains, signIds[i]) // 单个组织设置锚节点
		if err != nil {
			return utils.ToErr(err, "updateAnchorForOrg")
		}
	}
	logger.Debug("hello 使用每个org的管理员身份更新锚节点配置完成")
	return nil
}

// /////////////////////////////// 通道管理 结束///////////////////////////////////
