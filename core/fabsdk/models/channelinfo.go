package models

import (
	"fmt"
	"path/filepath"
)

// 通道信息
type ChannelInfo struct {
	// 通道信息
	ChannelID     string // like "simplecc"
	ChannelConfig string // like os.Getenv("GOPATH") + "/src/github.com/hyperledger/fabric-samples/test-network/channel-artifacts/testchannel.tx"
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func newChannelInfo(channelId string, channelConfig string) *ChannelInfo {
	logger.Debug("NewChannelInfo enter")
	return &ChannelInfo{
		ChannelConfig: channelConfig,
		ChannelID:     channelId,
	}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewDefaultChannelInfo(channelId string) *ChannelInfo {
	logger.Debug("NewDefaultChannelInfo enter")
	return newChannelInfo(channelId, filepath.Join(DefaultPath, fmt.Sprintf("%s.tx", channelId)))
}
