package models

import (
	"fmt"

	"fabapi/pkg/utils"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// 账本客户端
type LedgerClient struct {
	Sdk                  *fabsdk.FabricSDK        //
	ClientChannelContext *context.ChannelProvider //
	LedgerCli            *ledger.Client           //
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func newLedgerClient(channelId, userName, orgName string, sdk *fabsdk.FabricSDK) (*LedgerClient, error) {
	logger.Debug("NewLedgerClient enter")
	// mutex.Lock()
	// defer mutex.Unlock()
	clientChannelContext := sdk.ChannelContext(channelId, fabsdk.WithUser(userName), fabsdk.WithOrg(orgName))
	ledgerClient, err := ledger.New(clientChannelContext)
	if err != nil {
		closeContext(sdk, clientChannelContext)
		clientChannelContext = nil
		return nil, utils.ToErr(err, "ledger.New")
	}
	return &LedgerClient{
		Sdk:                  sdk,
		ClientChannelContext: &clientChannelContext,
		LedgerCli:            ledgerClient,
	}, nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LedgerClient) close() {
	logger.Debug("LedgerClient close enter")
	if l == nil {
		return
	}
	closeContext(l.Sdk, l.ClientChannelContext)
	l.ClientChannelContext = nil
	l.LedgerCli = nil
	l = nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func loopCloseLedgerClient(ps []*LedgerClient) {
	logger.Debug("loopCloseLedgerClient enter")
	for _, p := range ps {
		p.close()
	}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func getLedgerClientFromCache(channelId, userName, orgName string, sdk *fabsdk.FabricSDK) (*LedgerClient, error) {
	logger.Debug("getLedgerClientFromCache enter")
	key := fmt.Sprintf("%s%s%s@LC", channelId, orgName, userName)
	ret, ok := cache.Get(key)
	if ok {
		return ret.(*LedgerClient), nil
	}
	ledgerClient, err := newLedgerClient(channelId, userName, orgName, sdk)
	if err != nil {
		return nil, utils.ToErr(err, "NewLedgerClient")
	}
	cache.Set(key, ledgerClient, 0)
	return ledgerClient, nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func getLedgerClient(channelId, userName, orgName string, sdk *fabsdk.FabricSDK) (*LedgerClient, error) {
	logger.Debug("getLedgerClient enter")
	ledgerClient, err := newLedgerClient(channelId, userName, orgName, sdk)
	if err != nil {
		return nil, err
	}
	return ledgerClient, nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func GetOneLedgerClient(channelId, userName, orgName string, sdk *fabsdk.FabricSDK, isConcurrent bool) (*LedgerClient, error) {

	logger.Debug("GetOneLedgerClient enter")
	var ledgerClient *LedgerClient
	var err error
	if isConcurrent {
		ledgerClient, err = getLedgerClientFromCache(channelId, userName, orgName, sdk)
		if err != nil {
			return nil, utils.ToErr(err, "getLedgerClientFromCache")
		}
	} else {
		ledgerClient, err = getLedgerClient(channelId, userName, orgName, sdk)
		if err != nil {
			return nil, utils.ToErr(err, "getLedgerClient")
		}
	}
	return ledgerClient, nil
}
