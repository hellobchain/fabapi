/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package fabsdk

import (
	"fabapi/core/fabsdk/models"
)

type Channel interface {
	// 加入通道
	JoinChannel(orgsInfos []*models.PeerOrgInfo, channelInfo *models.ChannelInfo, ordererOrgInfo *models.OrdererOrgInfo) error

	// 创建通道
	CreateChannel(orgsInfos []*models.PeerOrgInfo, ordererOrgInfo *models.OrdererOrgInfo, channelInfo *models.ChannelInfo) error
}
