package models

import (
	"fabapi/pkg/utils"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// 通道客户端
type ChannelClient struct {
	Sdk                  *fabsdk.FabricSDK        // FabricSDK提供了对由SDK管理的客户端的访问(和上下文)。
	ClientChannelContext *context.ChannelProvider // channnel的上下文
	Client               *channel.Client          // channel的client
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func newChannelClient(sdk *fabsdk.FabricSDK, channelId string, orgName string, orgUser string) (*ChannelClient, error) {
	logger.Debug("NewChannelClient enter")
	clientChannelContext := sdk.ChannelContext(channelId, fabsdk.WithUser(orgUser), fabsdk.WithOrg(orgName))
	client, err := channel.New(clientChannelContext)
	if err != nil {
		closeContext(sdk, clientChannelContext)
		clientChannelContext = nil
		return nil, utils.ToErr(err, "channel.New")
	}
	return &ChannelClient{
		Sdk:                  sdk,
		ClientChannelContext: &clientChannelContext,
		Client:               client,
	}, nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChannelClient) close() {
	logger.Debug("ChannelClient close enter")
	if c == nil {
		return
	}
	closeContext(c.Sdk, c.ClientChannelContext)
	c.ClientChannelContext = nil
	c.Client = nil
	c = nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func getChannelClientFromCache(sdk *fabsdk.FabricSDK, channelId string, orgName string, orgUser string) (*ChannelClient, error) {
	logger.Debug("getChannelClientFromCache enter")
	key := fmt.Sprintf("%s%s%s@CC", channelId, orgName, orgUser)
	ret, ok := cache.Get(key)
	if ok {
		return ret.(*ChannelClient), nil
	}
	channelClient, err := newChannelClient(sdk, channelId, orgName, orgUser)
	if err != nil {
		return nil, utils.ToErr(err, "NewChannelClient")
	}
	cache.Set(key, channelClient, 0)
	return channelClient, nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func getChannelClient(sdk *fabsdk.FabricSDK, channelId string, orgName string, orgUser string) (*ChannelClient, error) {
	logger.Debug("getChannelClient enter")
	channelClient, err := newChannelClient(sdk, channelId, orgName, orgUser)
	if err != nil {
		return nil, err
	}
	return channelClient, nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func GetOneChannelClient(sdk *fabsdk.FabricSDK, channelId string, orgName string, orgUser string, isConcurrent bool) (*ChannelClient, error) {
	logger.Debug("GetOneChannelClient enter")
	var channelClient *ChannelClient
	var err error
	if isConcurrent {
		channelClient, err = getChannelClientFromCache(sdk, channelId, orgName, orgUser)
		if err != nil {
			return nil, utils.ToErr(err, "failed to getChannelClient channel channelClient")
		}
	} else {
		channelClient, err = getChannelClient(sdk, channelId, orgName, orgUser)
		if err != nil {
			return nil, utils.ToErr(err, "failed to create new channel channelClient")
		}
	}
	return channelClient, nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func loopCloseChannelClient(ps []*ChannelClient) {
	logger.Debug("loopCloseChannelClient enter")
	for _, p := range ps {
		p.close()
	}
}
