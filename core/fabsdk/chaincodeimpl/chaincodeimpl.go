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
	"path/filepath"
	"strings"

	"fabapi/core/fabsdk/models"
	internalutils "fabapi/core/fabsdk/utils"
	"fabapi/pkg/utils"

	internalfabsdk "fabapi/core/fabsdk"

	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/lifecycle"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// 链码操作

/*

1. 打包链码
2. 安装链码
3. 批准链码
4. 提交链码
5. 升级链码
6. 初始化链码
7. 调用链码
8. 查询链码

*/
var _ internalfabsdk.Chaincode = (*ChaincodeOp)(nil)

type ChaincodeOp struct {
	isAsy        bool
	isConcurrent bool
	sdk          *fabsdk.FabricSDK
	isFile       bool
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewChaincodeOp(other *models.Other, sdk *fabsdk.FabricSDK) internalfabsdk.Chaincode {
	return &ChaincodeOp{
		isAsy:        other.IsAsy,
		sdk:          sdk,
		isConcurrent: other.IsConcurrent,
		isFile:       other.IsFile,
	}
}

// /////////////////////////////// 链码管理 开始///////////////////////////////////

// 打包链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) PackageCC(cCName, cCVersion, cCpath string, chaincodeType string) (string, interface{}, string, error) {
	logger.Debug("PackageCC enter")
	logger.Debug("hello 开始打包链码 waiting......")
	var ret interface{}
	label := cCName + "_" + cCVersion
	ccType, ok := peer.ChaincodeSpec_Type_value[strings.ToUpper(chaincodeType)]
	if !ok {
		return models.EmptyReturn, nil, models.EmptyReturn, utils.ToErr(nil, fmt.Sprintf("package chaincode chaincode type is wrong %v only support golang、node、java、car", chaincodeType))
	}
	desc := &lifecycle.Descriptor{
		Path:  cCpath,
		Type:  (peer.ChaincodeSpec_Type)(ccType),
		Label: label,
	}
	ccPkg, err := lifecycle.NewCCPackage(desc) // 打包链码
	if err != nil {
		return models.EmptyReturn, nil, models.EmptyReturn, utils.ToErr(err, "package chaincode source")
	}
	ret = ccPkg
	packageID := lifecycle.ComputePackageID(label, ccPkg)
	if c.isFile { // 是否需要存文件
		ret, err = c.writeTmpFile(filepath.Join(models.DefaultCcpackagePath, packageID), ccPkg)
		if err != nil {
			return models.EmptyReturn, nil, models.EmptyReturn, utils.ToErr(err, "writeTmpFile")
		}
	}
	logger.Debug("hello 打包链码成功")
	return desc.Label, ret, packageID, nil
}

// 安装链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) InstallCC(ccName, ccVersion string, ret interface{}, orgsInfos []*models.PeerOrgInfo) (string, error) {
	logger.Debug("InstallCC enter")
	logger.Debug("hello 开始安装链码 waiting......")
	label := ccName + "_" + ccVersion
	ccPkg, isHaveTmpCcFile, ccPath, err := c.getPackageValue(ret)
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "getPackageValue")
	}
	installCCReq := resmgmt.LifecycleInstallCCRequest{
		Label:   label,
		Package: ccPkg,
	}
	packageID := lifecycle.ComputePackageID(installCCReq.Label, installCCReq.Package)
	err = c.installCC(packageID, installCCReq, orgsInfos)
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "installCC")
	}
	if isHaveTmpCcFile { // 有文件，将文件删除
		c.removeTmpFile(ccPath)
	}
	logger.Debug("hello 安装链码成功")
	return packageID, nil
}

// 批准链码 "OutOf('1', 'AMSP.member', 'BMSP.member')"
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) ApproveCC(packageID string, ccName, ccVersion string, sequence int64, chaincodePolicy string, isInit bool, channelID string, orgsInfos []*models.PeerOrgInfo) error {
	logger.Debug("ApproveCC enter")
	logger.Debug("hello 组织认可智能合约定义 waiting......")
	var mspIDs []string
	for _, org := range orgsInfos {
		mspIDs = append(mspIDs, org.OrgMspId)
	}
	ccPolicy, err := c.getCCPolicy(chaincodePolicy, mspIDs) // 转化安装的链码policy
	if err != nil {
		return utils.ToErr(err, "getCCPolicy")
	}
	approveCCReq := resmgmt.LifecycleApproveCCRequest{
		Name:              ccName,
		Version:           ccVersion,
		PackageID:         packageID,
		Sequence:          sequence,
		EndorsementPlugin: models.DefaultEndorsementPlugin,
		ValidationPlugin:  models.DefaultValidationPlugin,
		SignaturePolicy:   ccPolicy,
		InitRequired:      isInit,
	}
	orgInfos, err := c.getOrgInfos(orgsInfos)
	defer models.Release(orgInfos) // 释放资源
	if err != nil {
		return utils.ToErr(err, "getOrgInfos")
	}
	err = c.judgeApprovedCC(ccName, sequence, channelID, orgInfos)
	if err != nil {
		return utils.ToErr(err, "judgeApprovedCC")
	}
	err = c.approveCC(channelID, orgInfos, approveCCReq)
	if err != nil {
		return utils.ToErr(err, "approveCC")
	}
	logger.Debug("hello 查询组织认可智能合约定义")
	// Query approve cc
	if err := c.checkApprovedCC(ccName, sequence, channelID, orgInfos); err != nil { // 校验组织们是否批准
		return utils.ToErr(err, "queryApprovedCC")
	}
	logger.Debug("hello 组织认可智能合约定义完成")

	// Check commit readiness
	logger.Debug("hello 检查智能合约是否就绪 waiting......")
	if err := c.checkCCCommitReadiness(ccName, ccVersion, sequence, chaincodePolicy, isInit, channelID, orgInfos); err != nil { // 检查链码是否准备好
		return utils.ToErr(err, "checkCCCommitReadiness")
	}
	logger.Debug("hello 智能合约已经就绪")
	return nil
}

// 提交链码 "OutOf('1', 'AMSP.member', 'BMSP.member')"
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) CommitCC(ccName, ccVersion string, sequence int64, chaincodePolicy string, isInit bool, channelID string, orgsInfos []*models.PeerOrgInfo) error {
	logger.Debug("CommitCC enter")
	logger.Debug("hello 提交智能合约定义 waiting......")
	var mspIDs []string
	for _, org := range orgsInfos {
		mspIDs = append(mspIDs, org.OrgMspId)
	}
	ccPolicy, err := c.getCCPolicy(chaincodePolicy, mspIDs) // 转化传入的policy
	if err != nil {
		return utils.ToErr(err, "getCCPolicy")
	}
	req := resmgmt.LifecycleCommitCCRequest{
		Name:              ccName,
		Version:           ccVersion,
		Sequence:          sequence,
		EndorsementPlugin: models.DefaultEndorsementPlugin,
		ValidationPlugin:  models.DefaultValidationPlugin,
		SignaturePolicy:   ccPolicy,
		InitRequired:      isInit,
	}
	orgInfo, err := internalutils.GetOrgInfo(orgsInfos[0], c.sdk)
	if err != nil {
		return utils.ToErr(err, "GetOrgInfo")
	}
	defer models.Release(orgInfo.OrgClient)
	err = c.judgeCommittedCC(ccName, channelID, sequence, orgsInfos)
	if err != nil {
		return utils.ToErr(err, "queryCommittedCC")
	}
	_, err = orgInfo.OrgClient.OrgResMgmt.LifecycleCommitCC(channelID, req, resmgmt.WithTargetEndpoints(orgInfo.PeerNames...), resmgmt.WithRetry(retry.DefaultResMgmtOpts)) // 使用其中一个组织提交链码
	if err != nil {
		return utils.ToErr(err, "LifecycleCommitCC")
	}
	// query committed cc
	if err := c.checkCommittedCC(ccName, channelID, sequence, orgsInfos); err != nil { // 校验是否提交成功
		return utils.ToErr(err, "checkCommittedCC")
	}
	logger.Debug("hello 智能合约定义提交完成")
	return nil
}

// 升级链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) UpgradeCC(packageID string, ccName, ccVersion string, sequence int64, chaincodePolicy string, isInit bool, channelID string, orgsInfos []*models.PeerOrgInfo) error {
	logger.Debug("UpgradeCC enter")
	logger.Debug("hello 组织升级智能合约定义 waiting......")
	var mspIDs []string
	for _, org := range orgsInfos {
		mspIDs = append(mspIDs, org.OrgMspId)
	}
	ccPolicy, err := c.getCCPolicy(chaincodePolicy, mspIDs) // 转化安装的链码policy
	if err != nil {
		return utils.ToErr(err, "getCCPolicy")
	}
	approveCCReq := resmgmt.LifecycleApproveCCRequest{
		Name:              ccName,
		Version:           ccVersion,
		PackageID:         packageID,
		Sequence:          sequence,
		EndorsementPlugin: models.DefaultEndorsementPlugin,
		ValidationPlugin:  models.DefaultValidationPlugin,
		SignaturePolicy:   ccPolicy,
		InitRequired:      isInit,
	}
	orgInfos, err := c.getOrgInfos(orgsInfos)
	defer models.Release(orgInfos) // 释放资源
	if err != nil {
		return utils.ToErr(err, "getOrgInfos")
	}
	err = c.judgeApprovedCC(ccName, sequence, channelID, orgInfos)
	if err != nil {
		return utils.ToErr(err, "judgeApprovedCC")
	}
	err = c.approveCC(channelID, orgInfos, approveCCReq)
	if err != nil {
		return utils.ToErr(err, "approveCC")
	}
	logger.Debug("hello 查询组织认可智能合约定义")
	// Query approve cc
	if err := c.checkApprovedCC(ccName, sequence, channelID, orgInfos); err != nil { // 校验组织们是否批准
		return utils.ToErr(err, "checkUpgradeCC")
	}
	logger.Debug("hello 组织认可智能合约定义完成")

	// Check commit readiness
	logger.Debug("hello 检查智能合约是否就绪 waiting......")
	if err := c.checkCCCommitReadiness(ccName, ccVersion, sequence, chaincodePolicy, isInit, channelID, orgInfos); err != nil { // 检查链码是否准备好
		return utils.ToErr(err, "checkCCCommitReadiness")
	}
	logger.Debug("hello 智能合约已经就绪")
	return nil
}

// 初始化链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) InitCC(ccName string, channelID string, args []string, orgInfo *models.PeerOrgInfo) (string, error) {
	logger.Debug("InitCC enter")
	channelClient, err := models.GetOneChannelClient(c.sdk, channelID, orgInfo.OrgName, orgInfo.OrgUser, false)
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "failed to create new channel channelClient")
	}

	defer models.Release(channelClient)

	var paras [][]byte
	for _, arg := range args {
		paras = append(paras, []byte(arg))
	}
	request := channel.Request{ChaincodeID: ccName, Fcn: "init", Args: paras, IsInit: true}
	// init
	response, err := channelClient.Client.Execute(request, channel.WithTargetEndpoints(orgInfo.OrgPeerDomains...), channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "failed to init")
	}
	return string(response.TransactionID), nil
}

// 调用链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) InvokeCC(ccName string, channelID string, args []string, orgInfo *models.PeerOrgInfo) (string, error) {
	logger.Debug("InvokeCC enter")
	if len(args) == 0 {
		return models.EmptyReturn, utils.ToErr(nil, "args len is 0")
	}
	channelClient, err := models.GetOneChannelClient(c.sdk, channelID, orgInfo.OrgName, orgInfo.OrgUser, c.isConcurrent)
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "GetOneChannelClient")
	}
	if !c.isConcurrent {
		defer models.Release(channelClient)
	}

	var paras [][]byte
	for _, arg := range args[1:] {
		paras = append(paras, []byte(arg))
	}
	request := channel.Request{ChaincodeID: ccName, Fcn: args[0], Args: paras}
	response, err := channelClient.Client.Execute(request, channel.WithTargetEndpoints(orgInfo.OrgPeerDomains...), channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "channelClient.Client.Execute")
	}
	return string(response.TransactionID), nil
}

// 查询链码
/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeOp) QueryCC(ccName string, channelID string, args []string, orgInfo *models.PeerOrgInfo) (string, error) {
	logger.Debug("QueryCC enter")
	if len(args) == 0 {
		return models.EmptyReturn, utils.ToErr(nil, "args len is 0")
	}
	channelClient, err := models.GetOneChannelClient(c.sdk, channelID, orgInfo.OrgName, orgInfo.OrgUser, c.isConcurrent)
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "GetOneChannelClient")
	}
	if !c.isConcurrent { // 不是并发才释放
		defer models.Release(channelClient)
	}
	var paras [][]byte
	for _, arg := range args[1:] {
		paras = append(paras, []byte(arg))
	}
	request := channel.Request{ChaincodeID: ccName, Fcn: args[0], Args: paras}
	response, err := channelClient.Client.Query(request, channel.WithTargetEndpoints(orgInfo.OrgPeerDomains...), channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		return models.EmptyReturn, utils.ToErr(err, "channelClient.Client.Query")
	}
	return string(response.Payload), nil
}

// /////////////////////////////// 链码管理 结束///////////////////////////////////
