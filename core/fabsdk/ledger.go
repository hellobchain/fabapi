/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package fabsdk

import (
	"github.com/wsw365904/fabapi/core/fabsdk/models"
)

type Ledger interface {

	// 查询账本
	QueryLedger(channelID string, orgInfo *models.PeerOrgInfo) (interface{}, error)

	// 查询配置块
	QueryConfig(channelID string, orgInfo *models.PeerOrgInfo) (interface{}, error)

	// 查询最新块
	QueryBlock(channelID string, orgInfo *models.PeerOrgInfo, blockNum uint64) (interface{}, error)

	// 根据块高查询块
	QueryBlockByHash(channelID string, orgInfo *models.PeerOrgInfo, hash string) (interface{}, error)

	// 根据交易id查询块
	QueryBlockByTxId(channelID string, orgInfo *models.PeerOrgInfo, txid string) (interface{}, error)

	// 根据交易id查询交易
	QueryTx(channelID string, orgInfo *models.PeerOrgInfo, txID string) (interface{}, error)

	// 查询块高
	QueryBlockNum(channelID string, orgInfo *models.PeerOrgInfo) (uint64, error)
}
