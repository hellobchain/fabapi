// Package fabsdk /***************************************************************
package fabsdk

import (
	"github.com/wsw365904/fabapi/core/fabsdk/models"
)

type Channel interface {
	// JoinChannel 加入通道
	JoinChannel(orgsInfos []*models.PeerOrgInfo, channelInfo *models.ChannelInfo, ordererOrgInfo *models.OrdererOrgInfo) error

	// CreateChannel 创建通道
	CreateChannel(orgsInfos []*models.PeerOrgInfo, ordererOrgInfo *models.OrdererOrgInfo, channelInfo *models.ChannelInfo) error
}
