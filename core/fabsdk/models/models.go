/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package models

import (
	"fabapi/pkg/utils"

	"github.com/wsw365904/wswlog/wlogging"
)

var logger = wlogging.MustGetLoggerWithoutName()

var cache = utils.NewCache() // 缓存 用户channel client peer client orderer client

// 统一释放资源
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func Release(cl interface{}) {
	logger.Debug("Release enter")
	switch cl.(type) {
	case []*LedgerClient:
		loopCloseLedgerClient(cl.([]*LedgerClient))
	case []*ChannelClient:
		loopCloseChannelClient(cl.([]*ChannelClient))
	case []*OrdererClient:
		loopCloseOrdererClient(cl.([]*OrdererClient))
	case []*PeerClient:
		loopClosePeerClient(cl.([]*PeerClient))
	case []*PeerOrgInfo:
		loopCloseP(cl.([]*PeerOrgInfo))
	case []*OrdererOrgInfo:
		loopCloseO(cl.([]*OrdererOrgInfo))
	case []*OrgInfo:
		loopCloseOrgInfo(cl.([]*OrgInfo))
	case []*ChaincodeInfo:
		loopCloseChaincodeInfo(cl.([]*ChaincodeInfo))
	case *LedgerClient:
		cl.(*LedgerClient).close()
	case *ChannelClient:
		cl.(*ChannelClient).close()
	case *OrdererClient:
		cl.(*OrdererClient).close()
	case *PeerClient:
		cl.(*PeerClient).close()
	case *PeerOrgInfo:
		cl.(*PeerOrgInfo).close()
	case *OrdererOrgInfo:
		cl.(*OrdererOrgInfo).close()
	case *OrgInfo:
		cl.(*OrgInfo).close()
	case *ChaincodeInfo:
		cl.(*ChaincodeInfo).close()
	default:
		logger.Warnf("not support %v", cl)
	}
}
