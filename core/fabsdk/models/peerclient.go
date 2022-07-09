package models

import (
	"github.com/wsw365904/fabapi/pkg/utils"

	mspcli "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// peer客户端
type PeerClient struct {
	Sdk                   *fabsdk.FabricSDK       // FabricSDK提供了对由SDK管理的客户端的访问(和上下文)。
	OrgMspClient          *mspcli.Client          // 组织的msp的client
	OrgAdminClientContext *context.ClientProvider // 组织的管理员的client
	OrgResMgmt            *resmgmt.Client         // 资源管理客户端
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewPeerClient(sdk *fabsdk.FabricSDK, orgName string, orgAdminUser string) (*PeerClient, error) {
	logger.Debug("NewPeerClient enter")
	orgMspClient, err := mspcli.New(sdk.Context(), mspcli.WithOrg(orgName))
	if err != nil {
		return nil, utils.ToErr(err, "mspclient.New")
	}
	orgContext := sdk.Context(fabsdk.WithUser(orgAdminUser), fabsdk.WithOrg(orgName))
	resMgmtClient, err := resmgmt.New(orgContext)
	if err != nil {
		closeContext(sdk, orgContext)
		orgContext = nil
		return nil, utils.ToErr(err, "根据指定的资源管理客户端Context创建通道管理客户端失败")
	}
	return &PeerClient{
		Sdk:                   sdk,
		OrgMspClient:          orgMspClient,
		OrgAdminClientContext: &orgContext,
		OrgResMgmt:            resMgmtClient,
	}, nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func closeContext(sdk *fabsdk.FabricSDK, cont interface{}) {
	logger.Debug("closeContext enter")
	var cc fab.ClientContext
	var err error
	switch cont.(type) {
	case *context.ClientProvider:
		tmp := cont.(*context.ClientProvider)
		cc, err = (*tmp)()
		if err != nil {
			return
		}
	case *context.ChannelProvider:
		tmp := cont.(*context.ChannelProvider)
		cc, err = (*tmp)()
		if err != nil {
			return
		}
	case context.ClientProvider:
		tmp := cont.(context.ClientProvider)
		cc, err = tmp()
		if err != nil {
			return
		}
	case context.ChannelProvider:
		tmp := cont.(context.ChannelProvider)
		cc, err = tmp()
		if err != nil {
			return
		}
	}
	sdk.CloseContext(cc)
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (p *PeerClient) close() {
	logger.Debug("PeerClient close enter")
	if p == nil {
		return
	}
	closeContext(p.Sdk, p.OrgAdminClientContext)
	p.OrgAdminClientContext = nil
	p.OrgResMgmt = nil
	p.OrgMspClient = nil
	p = nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func loopClosePeerClient(ps []*PeerClient) {
	logger.Debug("loopClosePeerClient enter")
	for _, p := range ps {
		p.close()
	}
}
