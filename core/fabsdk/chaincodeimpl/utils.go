/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/

package chaincodeimpl

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/wsw365904/wswlog/wlogging"

	"fabapi/core/common/json"
	"fabapi/core/fabsdk/models"
	internalutils "fabapi/core/fabsdk/utils"
	"fabapi/pkg/utils"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/status"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/policydsl"
)

var logger = wlogging.MustGetLoggerWithoutName()

// 链码

// 安装链码////////////////////

// peer安装链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) installCC(packageID string, installCCReq resmgmt.LifecycleInstallCCRequest, orgsInfos []*models.PeerOrgInfo) error {
	logger.Debug("installCC enter")
	if c.isAsy {
		var wg sync.WaitGroup
		var errChan = make(chan error, len(orgsInfos))
		for _, org := range orgsInfos {
			go c.asyInstallCCForOrg(packageID, installCCReq, org, &wg, errChan)
		}
		err := internalutils.WaitAndErgodicErrChan(errChan, len(orgsInfos), &wg)
		if err != nil {
			return utils.ToErr(err, "ergodicErrChan")
		}
	} else {
		for _, org := range orgsInfos {
			err := c.installCCForOrg(packageID, installCCReq, org)
			if err != nil {
				return utils.ToErr(err, "installCCForOrg")
			}
		}
	}
	return nil
}

// 为每个组织的peer安装链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) installCCForOrg(packageID string, installCCReq resmgmt.LifecycleInstallCCRequest, org *models.PeerOrgInfo) error {
	logger.Debug("installCCForOrg enter")
	orgInfo, err := internalutils.GetOrgInfo(org, c.sdk)
	if err != nil {
		return utils.ToErr(err, "getExpectTargetPeers")
	}
	defer models.Release(orgInfo.OrgClient) // 释放资源
	var uninstallPeers []string
	for _, orgPeer := range orgInfo.PeerNames {
		if flag, _ := c.judgeInstalled(packageID, orgPeer, orgInfo.OrgClient.OrgResMgmt); !flag {
			uninstallPeers = append(uninstallPeers, orgPeer)
		} else {
			logger.Debugf("have already installed chaincode label %v", installCCReq.Label)
		}
	}
	if len(uninstallPeers) != 0 {
		res, err := orgInfo.OrgClient.OrgResMgmt.LifecycleInstallCC(installCCReq, resmgmt.WithTargetEndpoints(uninstallPeers...), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
		if err != nil {
			return utils.ToErr(err, "lifecycleInstallCC")
		}
		logger.Debug("install chaincode", res)
		// Get installed cc package
		if err := c.checkInstalledCCPackage(packageID, orgInfo.OrgClient, uninstallPeers); err != nil {
			return utils.ToErr(err, "getInstalledCCPackage")
		}
		// Query installed cc
		if err := c.checkInstalled(packageID, orgInfo.OrgClient, uninstallPeers); err != nil {
			return utils.ToErr(err, "checkInstalled")
		}
		return nil
	}
	return utils.ToErr(nil, "all peers(%v) have already installed", orgInfo.PeerNames)
}

// 为每个组织的peer安装链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) asyInstallCCForOrg(packageID string, installCCReq resmgmt.LifecycleInstallCCRequest, org *models.PeerOrgInfo, wg *sync.WaitGroup, errChan chan error) {
	logger.Debug("asyInstallCCForOrg enter")
	wg.Add(1)
	defer wg.Done()
	errChan <- c.installCCForOrg(packageID, installCCReq, org)
}

// 异步为每个组织的peer安装链码

// 校验每个组织的peer获取安装的链码包
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) checkInstalledCCPackage(packageID string, peerClient *models.PeerClient, orgPeers []string) error {
	logger.Debug("checkInstalledCCPackage enter")
	for _, orgPeer := range orgPeers {
		if _, err := peerClient.OrgResMgmt.LifecycleGetInstalledCCPackage(packageID, resmgmt.WithTargetEndpoints(orgPeer)); err != nil {
			return utils.ToErr(err, "lifecycleGetInstalledCCPackage")
		}
	}
	return nil
}

// 校验每个组织的peer查询安装的链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) checkInstalled(packageID string, peerClient *models.PeerClient, orgPeers []string) error {
	logger.Debug("checkInstalled enter")
	for _, orgPeer := range orgPeers {
		resp, err := peerClient.OrgResMgmt.LifecycleQueryInstalledCC(resmgmt.WithTargetEndpoints(orgPeer))
		if err != nil {
			return utils.ToErr(err, "lifecycleQueryInstalledCC")
		}
		logger.Debug("LifecycleQueryInstalledCC", "resp", resp)
		packageIDRes := models.Empty
		for _, t := range resp {
			if t.PackageID == packageID {
				packageIDRes = t.PackageID
			}
		}
		if !strings.EqualFold(packageID, packageIDRes) {
			return utils.ToErr(nil, "check package id error")
		}
	}
	return nil
}

// 为每个组织的peer校验安装的链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) judgeInstalled(packageID string, peer string, client *resmgmt.Client) (bool, error) {
	logger.Debug("judgeInstalled enter")
	flag := false
	resp, err := client.LifecycleQueryInstalledCC(resmgmt.WithTargetEndpoints(peer))
	if err != nil {
		return false, utils.ToErr(err, "lifecycleQueryInstalledCC")
	}
	for _, t := range resp {
		if t.PackageID == packageID {
			flag = true
		}
	}
	return flag, nil
}

// 安装链码////////////////////

// 批准链码////////////////////

// 批准链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) approveCC(channelID string, orgsInfos []*models.OrgInfo, approveCCReq resmgmt.LifecycleApproveCCRequest) error {
	logger.Debug("approveCC enter")
	if c.isAsy {
		var wg sync.WaitGroup
		var errChan = make(chan error, len(orgsInfos))
		for _, orgInfo := range orgsInfos {
			go c.asyApproveCCForOrg(channelID, orgInfo, approveCCReq, &wg, errChan) // 传入的每个组织批准自家的packageid
		}
		err := internalutils.WaitAndErgodicErrChan(errChan, len(orgsInfos), &wg)
		if err != nil {
			return utils.ToErr(err, "ergodicErrChan")
		}
	} else {
		for _, orgInfo := range orgsInfos {
			err := c.approveCCForOrg(channelID, orgInfo, approveCCReq) // 传入的每个组织批准自家的packageid
			if err != nil {
				return utils.ToErr(err, "approveCCForOrg")
			}
		}
	}
	return nil
}

// 为每个组织的批准链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) approveCCForOrg(channelID string, orgInfo *models.OrgInfo, approveCCReq resmgmt.LifecycleApproveCCRequest) error {
	logger.Debug("approveCCForOrg enter")
	logger.Debugf("hello chaincode approved by %s peers:", orgInfo.OrgName)
	for _, p := range orgInfo.PeerNames {
		logger.Debugf("%s", p)
	}

	res, err := orgInfo.OrgClient.OrgResMgmt.LifecycleApproveCC(channelID, approveCCReq, resmgmt.WithTargetEndpoints(orgInfo.PeerNames...), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return utils.ToErr(err, "lifecycleApproveCC")
	}
	logger.Debug("LifecycleApproveCC", "txid", res)
	return nil
}

// 异步为每个组织的批准链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) asyApproveCCForOrg(channelID string, orgInfo *models.OrgInfo, approveCCReq resmgmt.LifecycleApproveCCRequest, wg *sync.WaitGroup, errChan chan error) {
	logger.Debug("asyApproveCCForOrg enter")
	logger.Debugf("hello chaincode approved by %s peers:", orgInfo.OrgName)
	wg.Add(1)
	defer wg.Done()
	for _, p := range orgInfo.PeerNames {
		logger.Debugf("%s", p)
	}

	res, err := orgInfo.OrgClient.OrgResMgmt.LifecycleApproveCC(channelID, approveCCReq, resmgmt.WithTargetEndpoints(orgInfo.PeerNames...), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err == nil {
		logger.Debug("LifecycleApproveCC", "txid", res)
	}
	errChan <- err
}

// 校验每个组织的查询批准链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) checkApprovedCCForOrg(channelID string, queryApprovedCCReq resmgmt.LifecycleQueryApprovedCCRequest, orgInfo *models.OrgInfo) error {
	// Query approve cc
	logger.Debug("checkApprovedCCForOrg enter")
	for _, p := range orgInfo.PeerNames {
		resp, err := retry.NewInvoker(retry.New(retry.TestRetryOpts)).Invoke(
			func() (interface{}, error) {
				resp, err := orgInfo.OrgClient.OrgResMgmt.LifecycleQueryApprovedCC(channelID, queryApprovedCCReq, resmgmt.WithTargetEndpoints(p))
				if err != nil {
					return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), fmt.Sprintf("LifecycleQueryApprovedCC returned error: %v", err), nil)
				}
				logger.Debug("LifecycleQueryApprovedCC", resp)
				return resp, nil
			},
		)
		if err != nil {
			return utils.ToErr(err, "org %v Peer %v NewInvoker", orgInfo.OrgName, p)
		}
		if resp == nil {
			return utils.ToErr(nil, "org %v Peer %v Got nil invoker", orgInfo.OrgName, p)
		}
	}
	return nil
}

// 校验每个组织的查询批准链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) checkApprovedCC(ccName string, sequence int64, channelID string, orgInfos []*models.OrgInfo) error {
	logger.Debug("checkApprovedCC enter")
	queryApprovedCCReq := resmgmt.LifecycleQueryApprovedCCRequest{
		Name:     ccName,
		Sequence: sequence,
	}

	for _, org := range orgInfos {
		err := c.checkApprovedCCForOrg(channelID, queryApprovedCCReq, org)
		if err != nil {
			return utils.ToErr(err, "checkApprovedCCForOrg")
		}
	}

	return nil
}

// 校验每个组织的链码是否准备好被提交
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) checkCCCommitReadinessForOrg(ccName string, channelID string, req resmgmt.LifecycleCheckCCCommitReadinessRequest, orgInfo *models.OrgInfo) error {
	logger.Debug("checkCCCommitReadinessForOrg enter")
	for _, p := range orgInfo.PeerNames {
		resp, err := retry.NewInvoker(retry.New(retry.TestRetryOpts)).Invoke(
			func() (interface{}, error) {
				resp, err := orgInfo.OrgClient.OrgResMgmt.LifecycleCheckCCCommitReadiness(channelID, req, resmgmt.WithTargetEndpoints(p))
				logger.Debugf("lifecycleCheckCCCommitReadiness cc = %v, = %v\n", ccName, resp)
				if err != nil {
					return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), fmt.Sprintf("LifecycleCheckCCCommitReadiness returned error: %v", err), nil)
				}
				logger.Debug("LifecycleCheckCCCommitReadiness", resp)
				flag := true
				for _, r := range resp.Approvals {
					flag = flag && r
				}
				if !flag {
					return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), fmt.Sprintf("LifecycleCheckCCCommitReadiness returned : %v", resp), nil)
				}
				return resp, nil
			},
		)
		if err != nil {
			return utils.ToErr(err, "newInvoker")
		}
		if resp == nil {
			return utils.ToErr(nil, "got nil invoker response")
		}
	}

	return nil
}

// 校验每个组织的链码是否准备好被提交
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) checkCCCommitReadiness(ccName, ccVersion string, sequence int64, chaincodePolicy string, isInit bool, channelID string, orgInfos []*models.OrgInfo) error {
	logger.Debug("checkCCCommitReadiness enter")
	var mspIds []string
	for _, org := range orgInfos {
		mspIds = append(mspIds, org.OrgMSP)
	}
	ccPolicy, err := c.getCCPolicy(chaincodePolicy, mspIds)
	mspIds = nil
	if err != nil {
		return utils.ToErr(err, "getCCPolicy")
	}
	req := resmgmt.LifecycleCheckCCCommitReadinessRequest{
		Name:              ccName,
		Version:           ccVersion,
		EndorsementPlugin: models.DefaultEndorsementPlugin,
		ValidationPlugin:  models.DefaultValidationPlugin,
		SignaturePolicy:   ccPolicy,
		Sequence:          sequence,
		InitRequired:      isInit,
	}
	for _, orgInfo := range orgInfos {
		err := c.checkCCCommitReadinessForOrg(ccName, channelID, req, orgInfo)
		if err != nil {
			return utils.ToErr(err, "checkCCCommitReadinessForOrg")
		}
	}

	return nil
}

// 为每个组织的查询批准链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) judgeApprovedCCForOrg(channelID string, queryApprovedCCReq resmgmt.LifecycleQueryApprovedCCRequest, orgInfo *models.OrgInfo) error {
	// Query approve cc
	logger.Debug("judgeApprovedCCForOrg enter")
	for _, p := range orgInfo.PeerNames {
		resp, err := orgInfo.OrgClient.OrgResMgmt.LifecycleQueryApprovedCC(channelID, queryApprovedCCReq, resmgmt.WithTargetEndpoints(p))
		if err == nil {
			ret, _ := json.MarshalIndent(&resp, "", "	")
			return utils.ToErr(err, "org(%v) has already approved chaincode(%v) and definition(%v)", orgInfo.OrgName, queryApprovedCCReq.Name, string(ret))
		}
	}
	return nil
}

// 为每个组织的查询批准链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) judgeApprovedCC(ccName string, sequence int64, channelID string, orgInfos []*models.OrgInfo) error {
	logger.Debug("judgeApprovedCC enter")
	queryApprovedCCReq := resmgmt.LifecycleQueryApprovedCCRequest{
		Name:     ccName,
		Sequence: sequence,
	}

	for _, org := range orgInfos {
		err := c.judgeApprovedCCForOrg(channelID, queryApprovedCCReq, org)
		if err != nil {
			return utils.ToErr(err, "judgeApprovedCCForOrg")
		}
	}
	return nil
}

// 批准链码////////////////////

// 提交链码////////////////////

// 校验每个组织的链码是否提交
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) checkCommittedCCForOrg(ccName string, channelID string, sequence int64, req resmgmt.LifecycleQueryCommittedCCRequest, org *models.PeerOrgInfo) error {
	logger.Debug("checkCommittedCCForOrg enter")
	orgInfo, err := internalutils.GetOrgInfo(org, c.sdk)
	if err != nil {
		return utils.ToErr(err, "getExpectTargetPeers")
	}
	defer models.Release(orgInfo.OrgClient)
	for _, p := range orgInfo.PeerNames {
		resp, err := retry.NewInvoker(retry.New(retry.TestRetryOpts)).Invoke(
			func() (interface{}, error) {
				resp, err := orgInfo.OrgClient.OrgResMgmt.LifecycleQueryCommittedCC(channelID, req, resmgmt.WithTargetEndpoints(p))
				if err != nil {
					return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), fmt.Sprintf("LifecycleQueryCommittedCC returned error: %v", err), nil)
				}
				logger.Debugf("lifecycleQueryCommittedCC cc = %v, = %v\n", ccName, resp)
				flag := false
				for _, r := range resp {
					if r.Name == ccName && r.Sequence == sequence {
						flag = true
						break
					}
				}
				if !flag {
					return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), fmt.Sprintf("LifecycleQueryCommittedCC returned : %v", resp), nil)
				}
				return resp, nil
			},
		)
		if err != nil {
			return utils.ToErr(err, "newInvoker")
		}
		if resp == nil {
			return utils.ToErr(nil, "got nil invoker response")
		}
	}
	return nil
}

// 校验每个组织的链码是否提交
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) checkCommittedCC(ccName string, channelID string, sequence int64, orgs []*models.PeerOrgInfo) error {
	logger.Debug("checkCommittedCC enter")
	req := resmgmt.LifecycleQueryCommittedCCRequest{
		Name: ccName,
	}
	for _, org := range orgs {
		err := c.checkCommittedCCForOrg(ccName, channelID, sequence, req, org)
		if err != nil {
			return utils.ToErr(err, "checkCommittedCCForOrg")
		}
	}
	return nil
}

// 为每个组织的链码是否提交
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) judgeCommittedCCForOrg(ccName string, channelID string, sequence int64, req resmgmt.LifecycleQueryCommittedCCRequest, org *models.PeerOrgInfo) error {
	logger.Debug("judgeCommittedCCForOrg enter")
	orgInfo, err := internalutils.GetOrgInfo(org, c.sdk)
	if err != nil {
		return utils.ToErr(err, "getExpectTargetPeers")
	}
	defer models.Release(orgInfo.OrgClient)
	for _, p := range orgInfo.PeerNames {
		resp, err := orgInfo.OrgClient.OrgResMgmt.LifecycleQueryCommittedCC(channelID, req, resmgmt.WithTargetEndpoints(p))
		if err == nil {
			logger.Debugf("lifecycleQueryCommittedCC cc = %v, = %v\n", ccName, resp)
			flag := false
			for _, r := range resp {
				if r.Name == ccName && r.Sequence == sequence {
					flag = true
					break
				}
			}
			if flag {
				return utils.ToErr(nil, "have already commit successfully")
			}
		}
	}
	return nil
}

// 为每个组织的链码是否提交
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) judgeCommittedCC(ccName string, channelID string, sequence int64, orgs []*models.PeerOrgInfo) error {
	logger.Debug("judgeCommittedCC enter")
	req := resmgmt.LifecycleQueryCommittedCCRequest{
		Name: ccName,
	}
	for _, org := range orgs {
		err := c.judgeCommittedCCForOrg(ccName, channelID, sequence, req, org)
		if err != nil {
			return utils.ToErr(err, "judgeCommittedCCForOrg")
		}
	}
	return nil
}

// 提交链码////////////////////

// 获取组织信息
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) getOrgInfos(orgs []*models.PeerOrgInfo) ([]*models.OrgInfo, error) {
	logger.Debug("getOrgInfos enter")
	orgInfos := make([]*models.OrgInfo, 0)
	for _, org := range orgs {
		orgInfo, err := internalutils.GetOrgInfo(org, c.sdk)
		if err != nil {
			return orgInfos, utils.ToErr(err, "getExpectTargetPeers")
		}
		orgInfos = append(orgInfos, orgInfo)
	}
	return orgInfos, nil
}

// 获取cc策略
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) getCCPolicy(chaincodePolicy string, mspids []string) (*common.SignaturePolicyEnvelope, error) {
	logger.Debug("getCCPolicy enter")
	var ccPolicy *common.SignaturePolicyEnvelope
	var err error
	if chaincodePolicy == models.Empty {
		ccPolicy = policydsl.SignedByAnyMember(mspids)
	} else {
		logger.Debugf("ccPolicy:%v", chaincodePolicy)
		ccPolicy, err = policydsl.FromString(chaincodePolicy)
		if err != nil {
			return nil, utils.ToErr(err, "FromString")
		}
	}
	return ccPolicy, nil
}

// 获取链码包
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) getPackageValue(ret interface{}) ([]byte, bool, string, error) {
	logger.Debug("getPackageValue enter")
	var isHaveTmpCcFile bool
	var ccPath string
	ccPkg, ok := ret.([]byte)
	if !ok { // 是否是传入的package的文件地址，是的话，读取文件
		ccPath, ok = ret.(string)
		if ok {
			var err error
			ccPkg, err = ioutil.ReadFile(ccPath)
			if err != nil {
				return nil, isHaveTmpCcFile, models.EmptyReturn, utils.ToErr(err, "read chaincode package")
			}
			isHaveTmpCcFile = true
		} else {
			return nil, isHaveTmpCcFile, models.EmptyReturn, utils.ToErr(nil, "ret para is invalid")
		}
	}
	return ccPkg, isHaveTmpCcFile, ccPath, nil
}

// 移除临时文件
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) removeTmpFile(path string) {
	logger.Debug("removeTmpFile enter")
	logger.Debugf("delete %v", path)
	err := os.Remove(path)
	if err != nil {
		logger.Warnf("delete %s err:%v", path, err)
	}
}

// 写临时文件
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) writeTmpFile(path string, value []byte) (string, error) {
	logger.Debug("writeTmpFile enter")
	err := ioutil.WriteFile(path, value, os.ModePerm) // 存入固定位置
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "write chaincode package")
	}
	return path, nil
}
