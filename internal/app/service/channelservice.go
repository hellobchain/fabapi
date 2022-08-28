// Package service /***************************************************************
package service

import "github.com/wsw365904/fabapi/core/fabsdk/models"

type ChannelService interface {
	CreateChannel(orgsInfo []*models.PeerOrgInfo, ordererOrgInfo *models.OrdererOrgInfo, channelInfo *models.ChannelInfo) error

	JoinChannel(orgsInfo []*models.PeerOrgInfo, ordererOrgInfo *models.OrdererOrgInfo, channelInfo *models.ChannelInfo) error

	Close()
}
