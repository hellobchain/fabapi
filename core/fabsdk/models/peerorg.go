package models

import (
	"fmt"
	"path/filepath"
)

// peer组织信息
type PeerOrgInfo struct {
	OrgAdminUser   string   // like "Admin"    peer组织的管理员
	OrgName        string   // like "Org1"     peer组织名称
	OrgMspId       string   // like "Org1MSP"  peer组织的MSP
	OrgUser        string   // like "User1"    peer组织的普通用户
	OrgPeerDomains []string //                 peer组织下的peer成员
	OrgAnchorFile  string   // like ./channel-artifacts/Org2MSPanchors.tx   更新锚节点的配置文件
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func newPeerOrgInfo(orgAdminUser string, orgName string, orgMspid string, orgUser string, orgPeerDomains []string, orgAnchorFile string) *PeerOrgInfo {
	logger.Debug("NewPeerOrgInfo enter")
	if orgAdminUser == Empty {
		orgAdminUser = DefaultOrgAdminUser
	}
	if orgUser == Empty {
		orgUser = DefaultOrgUser
	}
	return &PeerOrgInfo{
		OrgAdminUser:   orgAdminUser,
		OrgName:        orgName,
		OrgMspId:       orgMspid,
		OrgUser:        orgUser,
		OrgPeerDomains: orgPeerDomains,
		OrgAnchorFile:  orgAnchorFile,
	}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func newDefaultPeerOrgInfo(orgName string, orgMspid string, orgPeerDomains []string, orgAnchorFile string) *PeerOrgInfo {
	logger.Debug("NewDefaultPeerOrgInfo enter")
	return newPeerOrgInfo(DefaultOrgAdminUser, orgName, orgMspid, DefaultOrgUser, orgPeerDomains, orgAnchorFile)
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (p *PeerOrgInfo) close() {
	logger.Debug("close enter")
	p = nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func loopCloseP(ps []*PeerOrgInfo) {
	logger.Debug("loopCloseP enter")
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
func NewDefaultPeerOrgInfos(orgNames []string, orgMspids []string, orgPeerDomains [][]string) []*PeerOrgInfo {
	logger.Debug("NewDefaultPeerOrgInfos enter")
	peerOrgInfos := make([]*PeerOrgInfo, 0)
	var orgPeerDomain []string
	for i := range orgNames {
		if orgPeerDomains == nil {
			orgPeerDomain = nil
		} else {
			orgPeerDomain = orgPeerDomains[i]
		}
		peerOrgInfos = append(peerOrgInfos, newDefaultPeerOrgInfo(orgNames[i], orgMspids[i], orgPeerDomain, filepath.Join(DefaultPath, fmt.Sprintf("%s%s", orgNames[i], DefaultOrgAnchorFile))))
	}
	return peerOrgInfos
}

// 组织信息
type OrgInfo struct {
	OrgClient *PeerClient // peer的client的
	PeerNames []string    // peer的名称
	OrgName   string      // 组织名称
	OrgMSP    string      // 组织msp
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewOrgInfo(orgClient *PeerClient, peerNames []string, orgName string, orgMSP string) *OrgInfo {
	logger.Debug("NewOrgInfo enter")
	return &OrgInfo{
		OrgClient: orgClient,
		PeerNames: peerNames,
		OrgName:   orgName,
		OrgMSP:    orgMSP,
	}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (o *OrgInfo) close() {
	logger.Debug("OrgInfo close enter")
	if o == nil {
		return
	}
	o.OrgClient.close()
	o.OrgClient = nil
	o.PeerNames = nil
	o = nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func loopCloseOrgInfo(ps []*OrgInfo) {
	logger.Debug("loopCloseOrgInfo enter")
	for _, p := range ps {
		p.close()
	}
}
