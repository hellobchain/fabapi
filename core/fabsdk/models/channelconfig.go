package models

import (
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/chconfig"
)

type ChannelConfig struct {
	Id          string           `json:"channelid"`
	BlockNumber uint64           `json:"blocknum"`
	Msps        []*mspConfig     `json:"msps"`
	AnchorPeers []*orgAnchorPeer `json:"anchorpeers"`
	Orderers    []string         `json:"orderers"`
	Versions    *versions        `json:"versions"`
}

// mspConfig collects all the configuration information for
// an MSP. The Config field should be unmarshalled in a way
// that depends on the Type

type mspConfig struct {
	// Type holds the type of the MSP; the default one would
	// be of type FABRIC implementing an X.509 based provider
	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// Config is MSP dependent configuration info
	Config *msp.FabricMSPConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

// orgAnchorPeer contains information about an anchor peer on this channel
type orgAnchorPeer struct {
	Org  string `json:"org"`
	Host string `json:"host"`
	Port int32  `json:"port"`
}

// versions ...
type versions struct {
	ReadSet  *common.ConfigGroup `json:"readset"`
	WriteSet *common.ConfigGroup `json:"writeset"`
	Channel  *common.ConfigGroup `json:"channel"`
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func ChannelCfgToChannelConfig(cfg *chconfig.ChannelCfg) (*ChannelConfig, error) {
	logger.Debug("ChannelCfgToChannelConfig enter")
	conf := new(ChannelConfig)
	conf.Id = cfg.ID()
	conf.BlockNumber = cfg.BlockNumber()
	for _, cfgMsp := range cfg.MSPs() {
		fabMspConfig := new(msp.FabricMSPConfig)
		err := proto.Unmarshal(cfgMsp.Config, fabMspConfig)
		if err != nil {
			conf = nil
			return nil, err
		}
		mspConfig := new(mspConfig)
		mspConfig.Config = fabMspConfig
		mspConfig.Type = cfgMsp.Type
		conf.Msps = append(conf.Msps, mspConfig)
	}
	for _, anchorPeer := range cfg.AnchorPeers() {
		ap := new(orgAnchorPeer)
		ap.Port = anchorPeer.Port
		ap.Host = anchorPeer.Host
		ap.Org = anchorPeer.Org
		conf.AnchorPeers = append(conf.AnchorPeers, ap)
	}
	conf.Orderers = cfg.Orderers()
	if cfg.Versions() != nil {
		conf.Versions = new(versions)
		conf.Versions.ReadSet = cfg.Versions().ReadSet
		conf.Versions.WriteSet = cfg.Versions().WriteSet
		conf.Versions.Channel = cfg.Versions().Channel
	}
	return conf, nil
}
