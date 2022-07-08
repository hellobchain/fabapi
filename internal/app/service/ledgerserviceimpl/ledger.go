/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package ledgerserviceimpl

import (
	"fmt"

	"github.com/wsw365904/wswlog/wlogging"

	"fabapi/core/common/e"
	internalfabsdk "fabapi/core/fabsdk"
	"fabapi/core/fabsdk/models"
	"fabapi/internal/app/service"
	"fabapi/pkg/utils"
)

var _ service.LedgerService = (*LedgerService)(nil)

var logger = wlogging.MustGetLoggerWithoutName()

type LedgerService struct {
	lop internalfabsdk.Ledger
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewLedgerService(lop internalfabsdk.Ledger) service.LedgerService {
	logger.Debug("NewLedgerService enter")
	return &LedgerService{
		lop: lop,
	}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LedgerService) Close() {
	logger.Debug("LedgerService Close enter")
	if l == nil {
		return
	}
	l = nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LedgerService) QueryLedger(channelID string, org *models.PeerOrgInfo) (interface{}, error) {
	logger.Debug("QueryLedger enter service")
	return l.lop.QueryLedger(channelID, org)
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LedgerService) QueryConfig(channelID string, org *models.PeerOrgInfo) (interface{}, error) {
	logger.Debug("QueryConfig enter service")
	return l.lop.QueryConfig(channelID, org)
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LedgerService) QueryBlock(channelID string, org *models.PeerOrgInfo, input interface{}) (interface{}, error) {
	logger.Debug("QueryBlock enter service")
	blockNum, ok := input.(uint64)
	if !ok {
		return "", utils.NewError(e.ERROR_TYPE_CHANGE_FAILED, fmt.Errorf("input is not uint64, input:%v", input))
	}
	return l.lop.QueryBlock(channelID, org, blockNum)
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LedgerService) QueryBlockByHash(channelID string, org *models.PeerOrgInfo, input interface{}) (interface{}, error) {
	logger.Debug("QueryBlockByHash enter service")
	blockHash, ok := input.(string)
	if !ok {
		return "", utils.NewError(e.ERROR_TYPE_CHANGE_FAILED, fmt.Errorf("input is not string, l.input:%v", input))
	}
	return l.lop.QueryBlockByHash(channelID, org, blockHash)
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LedgerService) QueryBlockByTxId(channelID string, org *models.PeerOrgInfo, input interface{}) (interface{}, error) {
	logger.Debug("QueryBlockByTxId enter service")
	txid, ok := input.(string)
	if !ok {
		return "", utils.NewError(e.ERROR_TYPE_CHANGE_FAILED, fmt.Errorf("input is not string, input:%v", input))
	}
	return l.lop.QueryBlockByTxId(channelID, org, txid)
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LedgerService) QueryTx(channelID string, org *models.PeerOrgInfo, input interface{}) (interface{}, error) {
	logger.Debug("QueryTx enter service")
	txid, ok := input.(string)
	if !ok {
		return "", utils.NewError(e.ERROR_TYPE_CHANGE_FAILED, fmt.Errorf("input is not string, input:%v", input))
	}
	return l.lop.QueryTx(channelID, org, txid)
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LedgerService) QueryBlockNum(channelID string, org *models.PeerOrgInfo) (uint64, error) {
	logger.Debug("QueryBlockNum enter service")
	return l.lop.QueryBlockNum(channelID, org)
}
