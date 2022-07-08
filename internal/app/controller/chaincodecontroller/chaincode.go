/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package chaincodecontroller

import (
	"net/http"

	"github.com/wsw365904/wswlog/wlogging"

	"fabapi/core/common/e"
	"fabapi/core/common/gintool"
	"fabapi/core/common/json"
	"fabapi/core/fabsdk/models"
	"fabapi/internal/app/service"
	"fabapi/pkg/utils"

	"github.com/gin-gonic/gin"
)

var logger = wlogging.MustGetLoggerWithoutName()

type ChaincodeController struct {
	chaincodeService service.ChaincodeService
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewChaincodeController(chaincodeService service.ChaincodeService) *ChaincodeController {
	return &ChaincodeController{
		chaincodeService: chaincodeService,
	}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeController) Close() {
	logger.Debug("ChaincodeController enter close")
	if c == nil {
		return
	}
	c.chaincodeService.Close()
	c = nil
}

// 打包链码请求
type packageChaincodeRequest struct {
	ChaincodeType    string `json:"chaincodetype,required" example:"GOLANG" binding:"required"`
	ChaincodeId      string `json:"chaincodeid,required" example:"test" binding:"required"`
	ChaincodeVersion string `json:"chaincodeversion,required" example:"v1.0" binding:"required"`
	ChaincodeSrcPath string `json:"chaincodesrcpath,required" example:"src/github.com/chaincode" binding:"required"`
}

// 打包链码响应
type packageChaincodeResponse struct {
	Label     string      `json:"label,required" binding:"-"`
	CcPack    interface{} `json:"ccpack,required" binding:"-"`
	PackageID string      `json:"packageid,required" binding:"-"`
}

// packageChaincode godoc
// @Summary 打包链码管理
// @Description 功能：打包链码管理
// @Tags 打包链码管理
// @Accept  json
// @Produce  json
// @Param resource body  chaincode.packageChaincodeRequest true "打包链码管理"
// @Success 200 {object} gintool.ApiResponse{result=chaincode.packageChaincodeResponse}
// @Failure 400 {object} gintool.ApiResponse
// @Router /chaincode/packagecc [post]
func (c *ChaincodeController) packageChaincode(ctx *gin.Context) {
	logger.Debug("packageChaincode enter controller")
	req := new(packageChaincodeRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("packageChaincode req\n", string(ret))
	label, ccPkg, packageID, err := c.chaincodeService.PackageChaincode(models.NewChaincodeInfo(true, req.ChaincodeType, models.DefaultChaincodePolicy, req.ChaincodeId, req.ChaincodeSrcPath, req.ChaincodeVersion, models.DefaultPackageId, models.DefaultPackagePara, models.DefaultSequence, models.DefaultArgs))
	if err != nil {
		logger.Error("packageChaincode", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.ERROR_PACKAGE_CHAINCODE_FAILED, err), nil)
		return
	}
	gintool.ResultCodeWithData(ctx, nil, packageChaincodeResponse{
		Label:     label,
		CcPack:    ccPkg,
		PackageID: packageID,
	})
}

// 安装链码请求
type installChaincodeRequest struct {
	ChaincodeId      string      `json:"chaincodeid,required" example:"test" binding:"required"`
	ChaincodeVersion string      `json:"chaincodeversion,required" example:"v1.0" binding:"required"`
	CcPack           interface{} `json:"ccpack,required" example:"" binding:"required"`
	OrgNames         []string    `json:"orgnames,required" example:"org1,org2,org3" binding:"required"`
	OrgMsps          []string    `json:"orgmsps,required" example:"Org1MSP,Org2MSP,Org3MSP" binding:"required"`
	PeerNames        [][]string  `json:"peernames,required" example:"[peer0.org1,peer1.org1],[peer0.org2,peer1.org2],[peer0.org3,peer1.org3]" binding:"required"`
}

// 安装链码响应
type installChaincodeResponse struct {
	PackageID string `json:"packageid,required" binding:"-"`
}

// installChaincode godoc
// @Summary 安装链码管理
// @Description 功能：安装链码管理
// @Tags 安装链码管理
// @Accept  json
// @Produce  json
// @Param resource body  chaincode.installChaincodeRequest true "安装链码管理"
// @Success 200 {object} gintool.ApiResponse{result=chaincode.installChaincodeResponse}
// @Failure 400 {object} gintool.ApiResponse
// @Router /chaincode/installcc [post]
func (c *ChaincodeController) installChaincode(ctx *gin.Context) {
	logger.Debug("installChaincode enter controller")
	req := new(installChaincodeRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("installChaincode req\n", string(ret))
	peerOrgInfos := models.NewDefaultPeerOrgInfos(req.OrgNames, req.OrgMsps, req.PeerNames)
	packageID, err := c.chaincodeService.InstallChaincode(models.NewChaincodeInfo(true, models.DefaultChaincodeType, models.DefaultChaincodePolicy, req.ChaincodeId, models.DefaultChaincodePath, req.ChaincodeVersion, models.DefaultPackageId, req.CcPack, models.DefaultSequence, models.DefaultArgs), peerOrgInfos)
	if err != nil {
		logger.Error("installChaincode", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.ERROR_INSTALL_CHAINCODE_FAILED, err), nil)
		return
	}
	gintool.ResultCodeWithData(ctx, nil, installChaincodeResponse{
		PackageID: packageID,
	})
}

// 批准链码请求
type approveChaincodeRequest struct {
	PackageID        string     `json:"packageid,required" example:"" binding:"required"`
	Sequence         int64      `json:"sequence,required" example:"1" binding:"required"`
	ChaincodeId      string     `json:"chaincodeid,required" example:"test" binding:"required"`
	ChaincodeVersion string     `json:"chaincodeversion,required" example:"v1.0" binding:"required"`
	ChannelID        string     `json:"channelid,required" example:"mychanel" binding:"required"`
	OrgNames         []string   `json:"orgnames,required" example:"org1,org2,org3" binding:"required"`
	OrgMsps          []string   `json:"orgmsps,required" example:"Org1MSP,Org2MSP,Org3MSP" binding:"required"`
	PeerNames        [][]string `json:"peernames,required" example:"[peer0.org1,peer1.org1],[peer0.org2,peer1.org2],[peer0.org3,peer1.org3]" binding:"required"`
	ChaincodePolicy  string     `json:"chaincodepolicy,omitempty" example:"" binding:"-"`
	IsInit           bool       `json:"isinit,omitempty" example:"" binding:"-"`
}

// approveChaincode godoc
// @Summary 批准链码管理
// @Description 功能：批准链码管理
// @Tags 批准链码管理
// @Accept  json
// @Produce  json
// @Param resource body  chaincode.approveChaincodeRequest true "批准链码管理"
// @Success 200 {object} gintool.ApiResponse
// @Failure 400 {object} gintool.ApiResponse
// @Router /chaincode/approvecc [post]
func (c *ChaincodeController) approveChaincode(ctx *gin.Context) {
	logger.Debug("approveChaincode enter controller")
	req := new(approveChaincodeRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("approveChaincode req\n", string(ret))
	peerOrgInfos := models.NewDefaultPeerOrgInfos(req.OrgNames, req.OrgMsps, req.PeerNames)
	err := c.chaincodeService.ApproveChaincode(models.NewChaincodeInfo(req.IsInit, models.DefaultChaincodeType, req.ChaincodePolicy, req.ChaincodeId, models.DefaultChaincodePath, req.ChaincodeVersion, req.PackageID, models.DefaultPackagePara, req.Sequence, models.DefaultArgs), peerOrgInfos, req.ChannelID)
	if err != nil {
		logger.Error("approveChaincode", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.ERROR_APPROVE_CHAINCODE_FAILED, err), nil)
		return
	}
	gintool.ResultCodeWithData(ctx, nil, nil)
}

// 升级链码请求
type upgradeChaincodeRequest struct {
	PackageID        string     `json:"packageid,required" example:"" binding:"required"`
	Sequence         int64      `json:"sequence,required" example:"1" binding:"required"`
	ChaincodeId      string     `json:"chaincodeid,required" example:"test" binding:"required"`
	ChaincodeVersion string     `json:"chaincodeversion,required" example:"v1.0" binding:"required"`
	ChannelID        string     `json:"channelid,required" example:"mychanel" binding:"required"`
	OrgNames         []string   `json:"orgnames,required" example:"org1,org2,org3" binding:"required"`
	OrgMsps          []string   `json:"orgmsps,required" example:"Org1MSP,Org2MSP,Org3MSP" binding:"required"`
	PeerNames        [][]string `json:"peernames,required" example:"[peer0.org1,peer1.org1],[peer0.org2,peer1.org2],[peer0.org3,peer1.org3]" binding:"required"`
	ChaincodePolicy  string     `json:"chaincodepolicy,omitempty" example:"" binding:"-"`
	IsInit           bool       `json:"isinit,omitempty" example:"" binding:"-"`
}

// upgradeChaincode godoc
// @Summary 升级链码管理
// @Description 功能：升级链码管理
// @Tags 升级链码管理
// @Accept  json
// @Produce  json
// @Param resource body  chaincode.upgradeChaincodeRequest true "升级链码管理"
// @Success 200 {object} gintool.ApiResponse
// @Failure 400 {object} gintool.ApiResponse
// @Router /chaincode/upgradecc [post]
func (c *ChaincodeController) upgradeChaincode(ctx *gin.Context) {
	logger.Debug("upgradeChaincode enter controller")
	req := new(upgradeChaincodeRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("upgradeChaincode req\n", string(ret))
	peerOrgInfos := models.NewDefaultPeerOrgInfos(req.OrgNames, req.OrgMsps, req.PeerNames)
	err := c.chaincodeService.UpgradeChaincode(models.NewChaincodeInfo(req.IsInit, models.DefaultChaincodeType, req.ChaincodePolicy, req.ChaincodeId, models.DefaultChaincodePath, req.ChaincodeVersion, req.PackageID, models.DefaultPackagePara, req.Sequence, models.DefaultArgs), peerOrgInfos, req.ChannelID)
	if err != nil {
		logger.Error("upgradeChaincode", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.ERROR_UPGRADE_CHAINCODE_FAILED, err), nil)
		return
	}
	gintool.ResultCodeWithData(ctx, nil, nil)
}

// 提交链码请求
type commitChaincodeRequest struct {
	PackageID        string     `json:"packageid,required" example:"" binding:"required"`
	Sequence         int64      `json:"sequence,required" example:"1" binding:"required"`
	ChaincodeId      string     `json:"chaincodeid,required" example:"test" binding:"required"`
	ChaincodeVersion string     `json:"chaincodeversion,required" example:"v1.0" binding:"required"`
	ChannelID        string     `json:"channelid,required" example:"mychannel" binding:"required"`
	OrgNames         []string   `json:"orgnames,required" example:"org1,org2,org3" binding:"required"`
	OrgMsps          []string   `json:"orgmsps,required" example:"Org1MSP,Org2MSP,Org3MSP" binding:"required"`
	PeerNames        [][]string `json:"peernames,required" example:"[peer0.org1,peer1.org1],[peer0.org2,peer1.org2],[peer0.org3,peer1.org3]" binding:"required"`
	ChaincodePolicy  string     `json:"chaincodepolicy,omitempty" example:"" binding:"-"`
	IsInit           bool       `json:"isinit,omitempty" example:"" binding:"-"`
}

// CommitChaincode godoc
// @Summary 提交链码管理
// @Description 功能：提交链码管理
// @Tags 提交链码管理
// @Accept  json
// @Produce  json
// @Param resource body  chaincode.commitChaincodeRequest true "提交链码管理"
// @Success 200 {object} gintool.ApiResponse
// @Failure 400 {object} gintool.ApiResponse
// @Router /chaincode/commitcc [post]
func (c *ChaincodeController) commitChaincode(ctx *gin.Context) {
	logger.Debug("commitChaincode enter controller")
	req := new(commitChaincodeRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("CommitChaincode req\n", string(ret))
	peerOrgInfos := models.NewDefaultPeerOrgInfos(req.OrgNames, req.OrgMsps, req.PeerNames)
	err := c.chaincodeService.CommitChaincode(models.NewChaincodeInfo(req.IsInit, models.DefaultChaincodeType, req.ChaincodePolicy, req.ChaincodeId, models.DefaultChaincodePath, req.ChaincodeVersion, req.PackageID, models.DefaultPackagePara, req.Sequence, models.DefaultArgs), peerOrgInfos, req.ChannelID)
	if err != nil {
		logger.Error("CommitChaincode", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.ERROR_COMMIT_CHAINCODE_FAILED, err), nil)
		return
	}
	gintool.ResultCodeWithData(ctx, nil, nil)
}

// 初始化链码请求
type initChaincodeRequest struct {
	ChaincodeId string     `json:"chaincodeid,required" example:"test" binding:"required"`
	Args        []string   `json:"args,omitempty" example:"save,a,b" binding:"-"`
	ChannelID   string     `json:"channelid,required" example:"mychannel" binding:"required"`
	OrgNames    []string   `json:"orgnames,required" example:"org1,org2,org3" binding:"required"`
	OrgMsps     []string   `json:"orgmsps,required" example:"Org1MSP,Org2MSP,Org3MSP" binding:"required"`
	PeerNames   [][]string `json:"peernames,required" example:"[peer0.org1,peer1.org1],[peer0.org2,peer1.org2],[peer0.org3,peer1.org3]" binding:"required"`
}

// 初始化链码响应
type initChaincodeResponse struct {
	Txid string `json:"txid,required" example:"0" binding:"-"`
}

// initChaincode godoc
// @Summary 初始化链码管理
// @Description 功能：初始化链码管理
// @Tags 初始化链码管理
// @Accept  json
// @Produce  json
// @Param resource body  chaincode.initChaincodeRequest true "初始化链码管理"
// @Success 200 {object} gintool.ApiResponse{result=chaincode.initChaincodeResponse}
// @Failure 400 {object} gintool.ApiResponse
// @Router /chaincode/initcc [post]
func (c *ChaincodeController) initChaincode(ctx *gin.Context) {
	logger.Debug("initChaincode enter controller")
	req := new(initChaincodeRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("initChaincode req\n", string(ret))
	peerOrgInfos := models.NewDefaultPeerOrgInfos(req.OrgNames, req.OrgMsps, req.PeerNames)
	txid, err := c.chaincodeService.InitChaincode(models.NewChaincodeInfo(true, models.DefaultChaincodeType, models.DefaultChaincodePolicy, req.ChaincodeId, models.DefaultChaincodePath, models.DefaultChaincodeVersion, models.DefaultPackageId, models.DefaultPackagePara, models.DefaultSequence, req.Args), peerOrgInfos, req.ChannelID)
	if err != nil {
		logger.Error("initChaincode", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.ERROR_INIT_CHAINCODE_FAILED, err), nil)
		return
	}
	gintool.ResultCodeWithData(ctx, nil, initChaincodeResponse{
		Txid: txid,
	})
}

// 调用链码请求
type invokeChaincodeRequest struct {
	ChaincodeId string     `json:"chaincodeid,required" example:"test" binding:"required"`
	ChannelID   string     `json:"channelid,required" example:"mychannel" binding:"required"`
	Args        []string   `json:"args,required" example:"save,a,b" binding:"required"`
	OrgNames    []string   `json:"orgnames,required" example:"org1,org2,org3" binding:"required"`
	OrgMsps     []string   `json:"orgmsps,required" example:"Org1MSP,Org2MSP,Org3MSP" binding:"required"`
	PeerNames   [][]string `json:"peernames,required" example:"[peer0.org1,peer1.org1],[peer0.org2,peer1.org2],[peer0.org3,peer1.org3]" binding:"required"`
}

// 调用链码响应
type invokeChaincodeResponse struct {
	Txid string `json:"txid,required" example:"0" binding:"-"`
}

// invokeChaincode godoc
// @Summary invoke链码管理
// @Description 功能：invoke链码管理
// @Tags invoke链码管理
// @Accept  json
// @Produce  json
// @Param resource body  chaincode.invokeChaincodeRequest true "invoke链码管理"
// @Success 200 {object} gintool.ApiResponse{result=chaincode.invokeChaincodeResponse}
// @Failure 400 {object} gintool.ApiResponse
// @Router /chaincode/invokecc [post]
func (c *ChaincodeController) invokeChaincode(ctx *gin.Context) {
	logger.Debug("invokeChaincode enter controller")
	req := new(invokeChaincodeRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("invokeChaincode req\n", string(ret))
	peerOrgInfos := models.NewDefaultPeerOrgInfos(req.OrgNames, req.OrgMsps, req.PeerNames)
	res, err := c.chaincodeService.InvokeChaincode(models.NewChaincodeInfo(true, models.DefaultChaincodeType, models.DefaultChaincodePolicy, req.ChaincodeId, models.DefaultChaincodePath, models.DefaultChaincodeVersion, models.DefaultPackageId, models.DefaultPackagePara, models.DefaultSequence, req.Args), peerOrgInfos, req.ChannelID)
	if err != nil {
		logger.Error("invokeChaincode", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.ERROR_INVOKE_CHAINCODE_FAILED, err), nil)
		return
	}
	gintool.ResultCodeWithData(ctx, nil, invokeChaincodeResponse{
		Txid: res,
	})
}

// 查询链码请求
type queryChaincodeRequest struct {
	ChaincodeId string     `json:"chaincodeid,required" example:"test" binding:"required"`
	ChannelID   string     `json:"channelid,required" example:"mychannel" binding:"required"`
	Args        []string   `json:"args,required" example:"query,a" binding:"required"`
	OrgNames    []string   `json:"orgnames,required" example:"org1,org2,org3" binding:"required"`
	OrgMsps     []string   `json:"orgmsps,required" example:"Org1MSP,Org2MSP,Org3MSP" binding:"required"`
	PeerNames   [][]string `json:"peernames,required" example:"[peer0.org1,peer1.org1],[peer0.org2,peer1.org2],[peer0.org3,peer1.org3]" binding:"required"`
}

// 查询链码响应
type queryChaincodeResponse struct {
	QueryRes string `json:"queryres,required" example:"0" binding:"-"`
}

// queryChaincode godoc
// @Summary query链码管理
// @Description 功能：query链码管理
// @Tags query链码管理
// @Accept  json
// @Produce  json
// @Param resource body  chaincode.queryChaincodeRequest true "query链码管理"
// @Success 200 {object} gintool.ApiResponse{result=chaincode.queryChaincodeResponse}
// @Failure 400 {object} gintool.ApiResponse
// @Router /chaincode/querycc [post]
func (c *ChaincodeController) queryChaincode(ctx *gin.Context) {
	logger.Debug("queryChaincode enter controller")
	req := new(queryChaincodeRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("queryChaincode req\n", string(ret))
	peerOrgInfos := models.NewDefaultPeerOrgInfos(req.OrgNames, req.OrgMsps, req.PeerNames)
	res, err := c.chaincodeService.QueryChaincode(models.NewChaincodeInfo(true, models.DefaultChaincodeType, models.DefaultChaincodePolicy, req.ChaincodeId, models.DefaultChaincodePath, models.DefaultChaincodeVersion, models.DefaultPackageId, models.DefaultPackagePara, models.DefaultSequence, req.Args), peerOrgInfos, req.ChannelID)
	if err != nil {
		logger.Error("queryChaincode", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.ERROR_QUERY_CHAINCODE_FAILED, err), nil)
		return
	}
	gintool.ResultCodeWithData(ctx, nil, queryChaincodeResponse{
		QueryRes: res,
	})
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
// 链码管理路由组
func FabChaincodeRouterGroup(app *ChaincodeController) gintool.RouteGroup {
	return gintool.RouteGroup{
		Route: []gintool.Route{
			gintool.NewRoute(http.MethodPost, "/packagecc", app.packageChaincode).AddComment("打包链码管理"),
			gintool.NewRoute(http.MethodPost, "/installcc", app.installChaincode).AddComment("安装链码管理"),
			gintool.NewRoute(http.MethodPost, "/approvecc", app.approveChaincode).AddComment("批准链码管理"),
			gintool.NewRoute(http.MethodPost, "/commitcc", app.commitChaincode).AddComment("提交链码管理"),
			gintool.NewRoute(http.MethodPost, "/initcc", app.initChaincode).AddComment("初始化链码管理"),
			gintool.NewRoute(http.MethodPost, "/invokecc", app.invokeChaincode).AddComment("invoke链码管理"),
			gintool.NewRoute(http.MethodPost, "/querycc", app.queryChaincode).AddComment("query链码管理"),
			gintool.NewRoute(http.MethodPost, "/upgradecc", app.upgradeChaincode).AddComment("升级链码管理"),
		},
		Prefix:  "/chaincode",
		Comment: "区块链链码管理",
		Module:  "chaincode",
	}
}
