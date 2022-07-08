/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package ledgerimpl

import (
	"encoding/hex"

	"github.com/wsw365904/wswlog/wlogging"

	"fabapi/core/common/json"

	internalfabsdk "fabapi/core/fabsdk"
	"fabapi/core/fabsdk/models"
	"fabapi/pkg/utils"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/chconfig"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

var _ internalfabsdk.Ledger = (*LedgerOp)(nil)

var logger = wlogging.MustGetLoggerWithoutName()

// 账本操作

/*

1. 查询账本
2. 查询配置块
3. 查询最新块
4. 根据块高查询块
5. 根据交易id查询块
6. 根据交易id查询交易
7. 查询块高

*/

type LedgerOp struct {
	isConcurrent bool
	sdk          *fabsdk.FabricSDK
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewLedgerOp(other *models.Other, sdk *fabsdk.FabricSDK) internalfabsdk.Ledger {
	logger.Debug("NewLedgerOp enter")
	return &LedgerOp{
		sdk:          sdk,
		isConcurrent: other.IsConcurrent,
	}
}

// /////////////////////////////// 账本管理 开始///////////////////////////////////

// 查询账本
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LedgerOp) QueryLedger(channelID string, orgInfo *models.PeerOrgInfo) (interface{}, error) {
	logger.Debug("QueryLedger enter")
	ledgerClient, err := models.GetOneLedgerClient(channelID, orgInfo.OrgUser, orgInfo.OrgName, l.sdk, l.isConcurrent)
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "GetOneLedgerClient")
	}
	if !l.isConcurrent {
		defer models.Release(ledgerClient)
	}
	blockInfo, err := ledgerClient.LedgerCli.QueryInfo(ledger.WithTargetEndpoints(orgInfo.OrgPeerDomains...))
	if err != nil { // 不是并发才释放
		return models.EmptyReturn, utils.ToErr(err, "ledgerClient.LedgerCli.QueryInfo")
	}
	if blockInfo.Status != models.SUCCESS {
		return models.EmptyReturn, utils.ToErr(nil, "queryInfo failed")
	}
	ret, err := json.MarshalIndent(blockInfo.BCI, models.Empty, models.TAB)
	if err != nil {
		logger.Warn(err)
	}
	logger.Debug("status", blockInfo.Status, "endorser", blockInfo.Endorser, "ret", string(ret))
	return blockInfo.BCI, nil
}

// 查询配置块
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LedgerOp) QueryConfig(channelID string, orgInfo *models.PeerOrgInfo) (interface{}, error) {
	logger.Debug("QueryConfig enter")
	ledgerClient, err := models.GetOneLedgerClient(channelID, orgInfo.OrgUser, orgInfo.OrgName, l.sdk, l.isConcurrent)
	if err != nil {
		return nil, utils.ToErr(err, "getOneLedgerClient")
	}
	if !l.isConcurrent { // 不是并发才释放
		defer models.Release(ledgerClient)
	}
	configBlockInfo, err := ledgerClient.LedgerCli.QueryConfig(ledger.WithTargetEndpoints(orgInfo.OrgPeerDomains...))

	if err != nil {
		return nil, utils.ToErr(err, "ledgerClient.LedgerCli.QueryConfig")
	}
	cfg, err := models.ChannelCfgToChannelConfig(configBlockInfo.(*chconfig.ChannelCfg))
	if err != nil {
		return nil, utils.ToErr(err, "ChannelCfgToChannelConfig")
	}
	ret, _ := json.MarshalIndent(cfg, models.Empty, models.TAB)
	logger.Debug("config value:", string(ret))
	return cfg, nil
}

// 查询最新块
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LedgerOp) QueryBlock(channelID string, orgInfo *models.PeerOrgInfo, blockNum uint64) (interface{}, error) {
	logger.Debug("QueryBlock enter")
	ledgerClient, err := models.GetOneLedgerClient(channelID, orgInfo.OrgUser, orgInfo.OrgName, l.sdk, l.isConcurrent)
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "getOneLedgerClient")
	}
	if !l.isConcurrent { // 不是并发才释放
		defer models.Release(ledgerClient)
	}
	blockInfo, err := ledgerClient.LedgerCli.QueryBlock(blockNum, ledger.WithTargetEndpoints(orgInfo.OrgPeerDomains...))
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "ledgerClient.LedgerCli.QueryBlock")
	}
	ret, err := json.MarshalIndent(blockInfo, models.Empty, models.TAB)
	if err != nil {
		logger.Warn(err)
	}
	logger.Debug("ret", string(ret))
	return blockInfo, nil
}

// 根据块高查询块
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LedgerOp) QueryBlockByHash(channelID string, orgInfo *models.PeerOrgInfo, hash string) (interface{}, error) {
	logger.Debug("QueryBlockByHash enter")
	ledgerClient, err := models.GetOneLedgerClient(channelID, orgInfo.OrgUser, orgInfo.OrgName, l.sdk, l.isConcurrent)
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "getOneLedgerClient")
	}
	if !l.isConcurrent { // 不是并发才释放
		defer models.Release(ledgerClient)
	}
	hashByte, _ := hex.DecodeString(hash)
	blockInfo, err := ledgerClient.LedgerCli.QueryBlockByHash(hashByte, ledger.WithTargetEndpoints(orgInfo.OrgPeerDomains...))
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "ledgerClient.LedgerCli.QueryBlockByHash")
	}
	ret, err := json.MarshalIndent(blockInfo, models.Empty, models.TAB)
	if err != nil {
		logger.Warn(err)
	}
	logger.Debug("ret", string(ret))
	return blockInfo, nil
}

// 根据交易id查询块
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LedgerOp) QueryBlockByTxId(channelID string, orgInfo *models.PeerOrgInfo, txid string) (interface{}, error) {
	logger.Debug("QueryBlockByTxId enter")
	ledgerClient, err := models.GetOneLedgerClient(channelID, orgInfo.OrgUser, orgInfo.OrgName, l.sdk, l.isConcurrent)
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "getOneLedgerClient")
	}
	if !l.isConcurrent { // 不是并发才释放
		defer models.Release(ledgerClient)
	}
	blockInfo, err := ledgerClient.LedgerCli.QueryBlockByTxID(fab.TransactionID(txid), ledger.WithTargetEndpoints(orgInfo.OrgPeerDomains...))
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "ledgerClient.LedgerCli.QueryBlockByTxID")
	}
	ret, err := json.MarshalIndent(blockInfo, models.Empty, models.TAB)
	if err != nil {
		logger.Warn(err)
	}
	logger.Debug("ret", string(ret))
	return blockInfo, nil
}

// 根据交易id查询交易
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LedgerOp) QueryTx(channelID string, orgInfo *models.PeerOrgInfo, txID string) (interface{}, error) {
	logger.Debug("QueryTx enter")
	ledgerClient, err := models.GetOneLedgerClient(channelID, orgInfo.OrgUser, orgInfo.OrgName, l.sdk, l.isConcurrent)
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "getOneLedgerClient")
	}
	if !l.isConcurrent { // 不是并发才释放
		defer models.Release(ledgerClient)
	}
	txInfo, err := ledgerClient.LedgerCli.QueryTransaction(fab.TransactionID(txID), ledger.WithTargetEndpoints(orgInfo.OrgPeerDomains...))
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "ledgerClient.LedgerCli.QueryTransaction")
	}
	ret, err := json.MarshalIndent(txInfo, models.Empty, models.TAB)
	if err != nil {
		logger.Warn(err)
	}
	logger.Debug("ret", string(ret))
	return txInfo, nil
}

// 查询块高
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LedgerOp) QueryBlockNum(channelID string, orgInfo *models.PeerOrgInfo) (uint64, error) {
	logger.Debug("QueryBlockNum enter")
	ledgerClient, err := models.GetOneLedgerClient(channelID, orgInfo.OrgUser, orgInfo.OrgName, l.sdk, l.isConcurrent)
	if err != nil {
		return 0, utils.ToErr(err, "getOneLedgerClient")
	}
	if !l.isConcurrent { // 不是并发才释放
		defer models.Release(ledgerClient)
	}
	txInfo, err := ledgerClient.LedgerCli.QueryInfo(ledger.WithTargetEndpoints(orgInfo.OrgPeerDomains...))
	if err != nil {
		return 0, utils.ToErr(err, "ledgerClient.LedgerCli.QueryInfo")
	}
	return txInfo.BCI.Height, nil
}

// /////////////////////////////// 账本管理 结束///////////////////////////////////
