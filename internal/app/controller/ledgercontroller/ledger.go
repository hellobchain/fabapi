/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package ledgercontroller

import (
	"net/http"

	"github.com/wsw365904/wswlog/wlogging"

	"github.com/wsw365904/fabapi/core/common/e"
	"github.com/wsw365904/fabapi/core/common/gintool"
	"github.com/wsw365904/fabapi/core/common/json"
	"github.com/wsw365904/fabapi/core/fabsdk/models"
	"github.com/wsw365904/fabapi/internal/app/service"
	"github.com/wsw365904/fabapi/pkg/utils"

	"github.com/gin-gonic/gin"
)

var logger = wlogging.MustGetLoggerWithoutName()

type LedgerController struct {
	ledgerService service.LedgerService
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewLedgerController(ledgerService service.LedgerService) *LedgerController {
	return &LedgerController{
		ledgerService: ledgerService,
	}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (l *LedgerController) Close() {
	logger.Debug("LedgerController enter close")
	if l == nil {
		return
	}
	l.ledgerService.Close()
	l = nil
}

//
type queryLedgerRequest struct {
	OrgNames  []string   `json:"orgnames,required" example:"org1,org2,org3" binding:"required"`
	OrgMsps   []string   `json:"orgmsps,required" example:"Org1MSP,Org2MSP,Org3MSP" binding:"required"`
	PeerNames [][]string `json:"peernames,required" example:"[peer0.org1,peer1.org1],[peer0.org2,peer1.org2],[peer0.org3,peer1.org3]" binding:"required"`
	ChannelID string     `json:"channelid,required" example:"mychannel" binding:"required"`
}

//
type queryLedgerResponse struct {
	QueryRes interface{} `json:"queryres,required" example:"0" binding:"-"`
}

// queryLedger godoc
// @Summary 查询账本管理
// @Description 功能：查询账本管理
// @Tags 查询账本管理
// @Accept  json
// @Produce  json
// @Param resource body  ledger.queryLedgerRequest true "查询账本管理"
// @Success 200 {object} gintool.ApiResponse{result=ledger.queryLedgerResponse}
// @Failure 400 {object} gintool.ApiResponse
// @Router /ledger/queryledger [post]
func (l *LedgerController) queryLedger(ctx *gin.Context) {
	logger.Debug("queryLedger enter controller")
	req := new(queryLedgerRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("queryLedger req\n", string(ret))
	peerOrgInfos := models.NewDefaultPeerOrgInfos(req.OrgNames, req.OrgMsps, req.PeerNames)
	res, err := l.ledgerService.QueryLedger(req.ChannelID, peerOrgInfos[0])
	if err != nil {
		logger.Error("queryLedger", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.ERROR_QUERY_LEDGER_FAILED, err), nil)
		return
	}
	gintool.ResultCodeWithData(ctx, nil, queryLedgerResponse{
		QueryRes: res,
	})
}

//
type queryConfigRequest struct {
	OrgNames  []string   `json:"orgnames,required" example:"org1,org2,org3" binding:"required"`
	OrgMsps   []string   `json:"orgmsps,required" example:"Org1MSP,Org2MSP,Org3MSP" binding:"required"`
	PeerNames [][]string `json:"peernames,required" example:"[[peer0.org1,peer1.org1],[peer0.org2,peer1.org2],[peer0.org3,peer1.org3]]" binding:"required"`
	ChannelID string     `json:"channelid,required" example:"mychannel" binding:"required"`
}

//
type queryConfigResponse struct {
	QueryRes interface{} `json:"queryres,required" example:"0" binding:"-"`
}

// queryConfig godoc
// @Summary 查询配置块管理
// @Description 功能：查询配置块管理
// @Tags 查询配置块管理
// @Accept  json
// @Produce  json
// @Param resource body  ledger.queryConfigRequest true "查询配置块管理"
// @Success 200 {object} gintool.ApiResponse{result=ledger.queryConfigResponse}
// @Failure 400 {object} gintool.ApiResponse
// @Router /ledger/queryconfig [post]
func (l *LedgerController) queryConfig(ctx *gin.Context) {
	logger.Debug("queryConfig enter controller")
	req := new(queryConfigRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("queryConfig req\n", string(ret))
	peerOrgInfos := models.NewDefaultPeerOrgInfos(req.OrgNames, req.OrgMsps, req.PeerNames)
	res, err := l.ledgerService.QueryConfig(req.ChannelID, peerOrgInfos[0])
	if err != nil {
		logger.Error("queryConfig", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.ERROR_QUERY_CONFIG_FAILED, err), nil)
		return
	}
	gintool.ResultCodeWithData(ctx, nil, queryConfigResponse{
		QueryRes: res,
	})
}

//
type queryBlockRequest struct {
	OrgNames  []string   `json:"orgnames,required" example:"org1,org2,org3" binding:"required"`
	OrgMsps   []string   `json:"orgmsps,required" example:"Org1MSP,Org2MSP,Org3MSP" binding:"required"`
	PeerNames [][]string `json:"peernames,required" example:"[peer0.org1,peer1.org1],[peer0.org2,peer1.org2],[peer0.org3,peer1.org3]" binding:"required"`
	ChannelID string     `json:"channelid,required" example:"mychannel" binding:"required"`
	BlockNum  uint64     `json:"blocknum" example:"0" binding:"-"`
}

//
type queryBlockResponse struct {
	QueryRes interface{} `json:"queryres,required" example:"0" binding:"-"`
}

// queryBlock godoc
// @Summary 查询块管理
// @Description 功能：查询块管理
// @Tags 查询块管理
// @Accept  json
// @Produce  json
// @Param resource body  ledger.queryBlockRequest true "查询块管理"
// @Success 200 {object} gintool.ApiResponse{result=ledger.queryBlockResponse}
// @Failure 400 {object} gintool.ApiResponse
// @Router /ledger/queryblock [post]
func (l *LedgerController) queryBlock(ctx *gin.Context) {
	logger.Debug("queryBlock enter controller")
	req := new(queryBlockRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("queryBlock req\n", string(ret))
	peerOrgInfos := models.NewDefaultPeerOrgInfos(req.OrgNames, req.OrgMsps, req.PeerNames)
	res, err := l.ledgerService.QueryBlock(req.ChannelID, peerOrgInfos[0], req.BlockNum)
	if err != nil {
		logger.Error("queryBlock", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.ERROR_QUERY_BLOCK_BY_BLOCKNUM_FAILED, err), nil)
		return
	}
	gintool.ResultCodeWithData(ctx, nil, queryBlockResponse{
		QueryRes: res,
	})
}

//
type queryBlockByHashRequest struct {
	OrgNames  []string   `json:"orgnames,required" example:"org1,org2,org3" binding:"required"`
	OrgMsps   []string   `json:"orgmsps,required" example:"Org1MSP,Org2MSP,Org3MSP" binding:"required"`
	PeerNames [][]string `json:"peernames,required" example:"[peer0.org1,peer1.org1],[peer0.org2,peer1.org2],[peer0.org3,peer1.org3]" binding:"required"`
	ChannelID string     `json:"channelid,required" example:"mychannel" binding:"required"`
	Hash      string     `json:"hash,required" example:"0xacedababacedababacedababacedababacedababacedababacedababacedabab" binding:"required"`
}

//
type queryBlockByHashResponse struct {
	QueryRes interface{} `json:"queryres,required" example:"0" binding:"-"`
}

// queryBlockByHash godoc
// @Summary 通过哈希查询块管理
// @Description 功能：通过哈希查询块管理
// @Tags 通过哈希查询块管理
// @Accept  json
// @Produce  json
// @Param resource body  ledger.queryBlockByHashRequest true "通过哈希查询块管理"
// @Success 200 {object} gintool.ApiResponse{result=ledger.queryBlockByHashResponse}
// @Failure 400 {object} gintool.ApiResponse
// @Router /ledger/queryblockbyhash [post]
func (l *LedgerController) queryBlockByHash(ctx *gin.Context) {
	logger.Debug("queryBlockByHash enter controller")
	req := new(queryBlockByHashRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("queryBlockByHash req\n", string(ret))
	peerOrgInfos := models.NewDefaultPeerOrgInfos(req.OrgNames, req.OrgMsps, req.PeerNames)
	res, err := l.ledgerService.QueryBlockByHash(req.ChannelID, peerOrgInfos[0], req.Hash)
	if err != nil {
		logger.Error("queryBlockByHash", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.ERROR_QUERY_BLOCK_BY_BLOCKHASH_FAILED, err), nil)
		return
	}
	gintool.ResultCodeWithData(ctx, nil, queryBlockByHashResponse{
		QueryRes: res,
	})
}

//
type queryBlockByTxIdRequest struct {
	OrgNames  []string   `json:"orgnames,required" example:"org1,org2,org3" binding:"required"`
	OrgMsps   []string   `json:"orgmsps,required" example:"Org1MSP,Org2MSP,Org3MSP" binding:"required"`
	PeerNames [][]string `json:"peernames,required" example:"[peer0.org1,peer1.org1],[peer0.org2,peer1.org2],[peer0.org3,peer1.org3]" binding:"required"`
	ChannelID string     `json:"channelid,required" example:"mychannel" binding:"required"`
	Txid      string     `json:"txid,required" example:"0xacedababacedababacedababacedababacedababacedababacedababacedabab" binding:"required"`
}

//
type queryBlockByTxIdResponse struct {
	QueryRes interface{} `json:"queryres,required" example:"0" binding:"-"`
}

// queryBlockByTxId godoc
// @Summary 通过交易id查询块管理
// @Description 功能：通过交易id查询块管理
// @Tags 通过交易id查询块管理
// @Accept  json
// @Produce  json
// @Param resource body  ledger.queryBlockByTxIdRequest true "通过交易id查询块管理"
// @Success 200 {object} gintool.ApiResponse{result=ledger.queryBlockByTxIdResponse}
// @Failure 400 {object} gintool.ApiResponse
// @Router /ledger/queryblockbytxid [post]
func (l *LedgerController) queryBlockByTxId(ctx *gin.Context) {
	logger.Debug("queryBlockByTxId enter controller")
	req := new(queryBlockByTxIdRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("queryBlockByTxId req\n", string(ret))
	peerOrgInfos := models.NewDefaultPeerOrgInfos(req.OrgNames, req.OrgMsps, req.PeerNames)
	res, err := l.ledgerService.QueryBlockByTxId(req.ChannelID, peerOrgInfos[0], req.Txid)
	if err != nil {
		logger.Error("queryBlockByTxId", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.ERROR_QUERY_BLOCK_BY_TXID_FAILED, err), nil)
		return
	}
	gintool.ResultCodeWithData(ctx, nil, queryBlockByTxIdResponse{
		QueryRes: res,
	})
}

//
type queryTxRequest struct {
	OrgNames  []string   `json:"orgnames,required" example:"org1,org2,org3" binding:"required"`
	OrgMsps   []string   `json:"orgmsps,required" example:"Org1MSP,Org2MSP,Org3MSP" binding:"required"`
	PeerNames [][]string `json:"peernames,required" example:"[peer0.org1,peer1.org1],[peer0.org2,peer1.org2],[peer0.org3,peer1.org3]" binding:"required"`
	ChannelID string     `json:"channelid,required" example:"mychannel" binding:"-"`
	Txid      string     `json:"txid,required" example:"0xacedababacedababacedababacedababacedababacedababacedababacedabab" binding:"required"`
}

//
type queryTxResponse struct {
	QueryRes interface{} `json:"queryres,required" example:"0" binding:"-"`
}

// queryTx godoc
// @Summary 查询交易管理
// @Description 功能：查询交易管理
// @Tags 查询交易管理
// @Accept  json
// @Produce  json
// @Param resource body  ledger.queryTxRequest true "查询交易管理"
// @Success 200 {object} gintool.ApiResponse{result=ledger.queryTxResponse}
// @Failure 400 {object} gintool.ApiResponse
// @Router /ledger/querytx [post]
func (l *LedgerController) queryTx(ctx *gin.Context) {
	logger.Debug("queryTx enter controller")
	req := new(queryTxRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("queryTx req\n", string(ret))
	peerOrgInfos := models.NewDefaultPeerOrgInfos(req.OrgNames, req.OrgMsps, req.PeerNames)
	res, err := l.ledgerService.QueryTx(req.ChannelID, peerOrgInfos[0], req.Txid)
	if err != nil {
		logger.Error("queryTx", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.ERROR_QUERY_TX_BY_HASH_FAILED, err), nil)
		return
	}
	gintool.ResultCodeWithData(ctx, nil, queryTxResponse{
		QueryRes: res,
	})
}

//
type queryBlockNumRequest struct {
	OrgNames  []string   `json:"orgnames,required" example:"org1,org2,org3" binding:"required"`
	OrgMsps   []string   `json:"orgmsps,required" example:"Org1MSP,Org2MSP,Org3MSP" binding:"required"`
	PeerNames [][]string `json:"peernames,required" example:"[peer0.org1,peer1.org1],[peer0.org2,peer1.org2],[peer0.org3,peer1.org3]" binding:"required"`
	ChannelID string     `json:"channelid,required" example:"mychannel" binding:"-"`
}

//
type queryBlockNumResponse struct {
	BlockNum uint64 `json:"blocknum,required" example:"0" binding:"-"`
}

// queryBlockNum godoc
// @Summary 查询块高管理
// @Description 功能：查询块高管理
// @Tags 查询块高
// @Accept  json
// @Produce  json
// @Param resource body  ledger.queryBlockNumRequest true "查询块高管理"
// @Success 200 {object} gintool.ApiResponse{result=ledger.queryBlockNumResponse}
// @Failure 400 {object} gintool.ApiResponse
// @Router /ledger/queryblocknum [post]
func (l *LedgerController) queryBlockNum(ctx *gin.Context) {
	logger.Debug("queryblocknum enter controller")
	req := new(queryBlockNumRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("queryblocknum req\n", string(ret))
	peerOrgInfos := models.NewDefaultPeerOrgInfos(req.OrgNames, req.OrgMsps, req.PeerNames)
	res, err := l.ledgerService.QueryBlockNum(req.ChannelID, peerOrgInfos[0])
	if err != nil {
		logger.Error("queryblocknum", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.ERROR_QUERY_BLOCKNUM_FAILED, err), nil)
		return
	}
	gintool.ResultCodeWithData(ctx, nil, queryBlockNumResponse{
		BlockNum: res,
	})
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
// 账本管理路由组
func FabLedgerRouterGroup(app *LedgerController) gintool.RouteGroup {
	return gintool.RouteGroup{
		Route: []gintool.Route{
			gintool.NewRoute(http.MethodPost, "/queryledger", app.queryLedger).AddComment("查询账本管理"),
			gintool.NewRoute(http.MethodPost, "/queryconfig", app.queryConfig).AddComment("查询配置块管理"),
			gintool.NewRoute(http.MethodPost, "/queryblock", app.queryBlock).AddComment("查询块管理"),
			gintool.NewRoute(http.MethodPost, "/queryblockbyhash", app.queryBlockByHash).AddComment("通过哈希查询块管理"),
			gintool.NewRoute(http.MethodPost, "/queryblockbytxid", app.queryBlockByTxId).AddComment("通过交易id查询块管理"),
			gintool.NewRoute(http.MethodPost, "/querytx", app.queryTx).AddComment("查询交易管理"),
			gintool.NewRoute(http.MethodPost, "/queryblocknum", app.queryBlockNum).AddComment("查询块高管理"),
		},
		Prefix:  "/ledger",
		Comment: "区块链账本管理",
		Module:  "ledger",
	}
}
