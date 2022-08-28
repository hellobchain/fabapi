// Package channelserviceimpl /***************************************************************
package channelserviceimpl

import (
	internalfabsdk "github.com/wsw365904/fabapi/core/fabsdk"
	"github.com/wsw365904/fabapi/core/fabsdk/models"
	"github.com/wsw365904/fabapi/internal/app/service"

	"github.com/wsw365904/wswlog/wlogging"
)

var _ service.ChannelService = (*ChannelService)(nil)

var logger = wlogging.MustGetLoggerWithoutName()

type ChannelService struct {
	cop internalfabsdk.Channel
}

// NewChannelService /***************************************************************
func NewChannelService(cop internalfabsdk.Channel) service.ChannelService {
	logger.Debug("NewChannelService enter")
	return &ChannelService{
		cop: cop,
	}
}

// Close /***************************************************************
func (c *ChannelService) Close() {
	logger.Debug("ChannelService Close enter")
	if c == nil {
		return
	}
	c = nil
}

// CreateChannel /***************************************************************
func (c *ChannelService) CreateChannel(orgsInfo []*models.PeerOrgInfo, ordererOrgInfo *models.OrdererOrgInfo, channelInfo *models.ChannelInfo) error {
	logger.Debug("CreateChannel enter service")
	return c.cop.CreateChannel(orgsInfo, ordererOrgInfo, channelInfo)
}

// JoinChannel /***************************************************************
func (c *ChannelService) JoinChannel(orgsInfo []*models.PeerOrgInfo, ordererOrgInfo *models.OrdererOrgInfo, channelInfo *models.ChannelInfo) error {
	logger.Debug("JoinChannel enter service")
	return c.cop.JoinChannel(orgsInfo, channelInfo, ordererOrgInfo)
}
