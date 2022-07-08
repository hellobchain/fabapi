/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package channelserviceimpl

import (
	internalfabsdk "fabapi/core/fabsdk"
	"fabapi/core/fabsdk/models"
	"fabapi/internal/app/service"

	"github.com/wsw365904/wswlog/wlogging"
)

var _ service.ChannelService = (*ChannelService)(nil)

var logger = wlogging.MustGetLoggerWithoutName()

type ChannelService struct {
	cop internalfabsdk.Channel
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewChannelService(cop internalfabsdk.Channel) service.ChannelService {
	logger.Debug("NewChannelService enter")
	return &ChannelService{
		cop: cop,
	}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChannelService) Close() {
	logger.Debug("ChannelService Close enter")
	if c == nil {
		return
	}
	c = nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChannelService) CreateChannel(orgsInfo []*models.PeerOrgInfo, ordererOrgInfo *models.OrdererOrgInfo, channelInfo *models.ChannelInfo) error {
	logger.Debug("CreateChannel enter service")
	return c.cop.CreateChannel(orgsInfo, ordererOrgInfo, channelInfo)
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChannelService) JoinChannel(orgsInfo []*models.PeerOrgInfo, ordererOrgInfo *models.OrdererOrgInfo, channelInfo *models.ChannelInfo) error {
	logger.Debug("JoinChannel enter service")
	return c.cop.JoinChannel(orgsInfo, channelInfo, ordererOrgInfo)
}
