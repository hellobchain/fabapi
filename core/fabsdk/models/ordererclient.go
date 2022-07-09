package models

import (
	"github.com/wsw365904/fabapi/pkg/utils"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// orderer客户端
type OrdererClient struct {
	Sdk                  *fabsdk.FabricSDK
	OrdererClientContext *context.ClientProvider
	OrdererClient        *resmgmt.Client
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewOrdererClient(sdk *fabsdk.FabricSDK, ordererAdminUser string, ordererOrgName string) (*OrdererClient, error) {
	logger.Debug("NewOrdererClient enter")
	ordererClientContext := sdk.Context(fabsdk.WithUser(ordererAdminUser), fabsdk.WithOrg(ordererOrgName))
	// Channel management client is responsible for managing channels (create/update channel)
	chMgmtClient, err := resmgmt.New(ordererClientContext)
	if err != nil {
		closeContext(sdk, ordererClientContext)
		ordererClientContext = nil
		return nil, utils.ToErr(err, "channel management client create")
	}
	return &OrdererClient{
		Sdk:                  sdk,
		OrdererClientContext: &ordererClientContext,
		OrdererClient:        chMgmtClient,
	}, nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (o *OrdererClient) close() {
	logger.Debug("OrdererClient close enter")
	if o == nil {
		return
	}
	closeContext(o.Sdk, o.OrdererClientContext)
	o.OrdererClient = nil
	o.OrdererClientContext = nil
	o = nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func loopCloseOrdererClient(ps []*OrdererClient) {
	logger.Debug("loopCloseOrdererClient enter")
	for _, p := range ps {
		p.close()
	}
}
