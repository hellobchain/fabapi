/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package service

import "fabapi/core/fabsdk/models"

type ChannelService interface {
	CreateChannel(orgsInfo []*models.PeerOrgInfo, ordererOrgInfo *models.OrdererOrgInfo, channelInfo *models.ChannelInfo) error

	JoinChannel(orgsInfo []*models.PeerOrgInfo, ordererOrgInfo *models.OrdererOrgInfo, channelInfo *models.ChannelInfo) error

	Close()
}
