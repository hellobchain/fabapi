/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package service

import (
	"github.com/wsw365904/fabapi/core/fabsdk/models"
)

type LedgerService interface {
	QueryLedger(channelID string, org *models.PeerOrgInfo) (interface{}, error)

	QueryConfig(channelID string, org *models.PeerOrgInfo) (interface{}, error)

	QueryBlock(channelID string, org *models.PeerOrgInfo, input interface{}) (interface{}, error)

	QueryBlockByHash(channelID string, org *models.PeerOrgInfo, input interface{}) (interface{}, error)

	QueryBlockByTxId(channelID string, org *models.PeerOrgInfo, input interface{}) (interface{}, error)

	QueryTx(channelID string, org *models.PeerOrgInfo, input interface{}) (interface{}, error)

	QueryBlockNum(channelID string, org *models.PeerOrgInfo) (uint64, error)

	Close()
}
