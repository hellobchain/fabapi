/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package channelcontroller

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

type ChannelController struct {
	channelService service.ChannelService
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewChannelController(channelService service.ChannelService) *ChannelController {
	return &ChannelController{
		channelService: channelService,
	}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChannelController) Close() {
	logger.Debug("ChannelController enter close")
	if c == nil {
		return
	}
	c.channelService.Close()
	c = nil
}

// 创建通道请求
type createChannelRequest struct {
	OrgNames        []string `json:"orgnames,required" example:"org1,org2,org3" binding:"required"`
	OrgMsps         []string `json:"orgmsps,required" example:"Org1MSP,Org2MSP,Org3MSP" binding:"required"`
	ChannelID       string   `json:"channelid,required" example:"mychannel" binding:"-"`
	OrdererNames    []string `json:"orderernames,required" example:"orderer0,orderer1,orderer2" binding:"required"`
	OrdererOrgName  string   `json:"ordererorgname,omitempty" example:"orderer" binding:"-"`
	OrdererUserName string   `json:"ordererusername,omitempty" example:"Admin" binding:"-"`
}

// createChannel godoc
// @Summary 创建通道管理
// @Description 功能：创建通道管理
// @Tags 创建通道管理
// @Accept  json
// @Produce  json
// @Param resource body  channel.createChannelRequest true "创建通道管理"
// @Success 200 {object} gintool.ApiResponse
// @Failure 400 {object} gintool.ApiResponse
// @Router /channel/createchannel [post]
func (c *ChannelController) createChannel(ctx *gin.Context) {
	logger.Debug("createChannel enter controller")
	req := new(createChannelRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("createChannel req\n", string(ret))
	peerOrgInfos := models.NewDefaultPeerOrgInfos(req.OrgNames, req.OrgMsps, nil)
	ordererOrgInfo := models.NewOrdererOrgInfo(req.OrdererUserName, req.OrdererOrgName, req.OrdererNames)
	channelInfo := models.NewDefaultChannelInfo(req.ChannelID)
	err := c.channelService.CreateChannel(peerOrgInfos, ordererOrgInfo, channelInfo)
	if err != nil {
		logger.Error("createChannel", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.ERROR_CREATE_CHANNEL_FAILED, err), nil)
		return
	}
	gintool.ResultCodeWithData(ctx, nil, nil)
}

// 加入通道请求
type joinChannelRequest struct {
	OrgNames     []string   `json:"orgnames,required" example:"org1,org2,org3" binding:"required"`
	OrgMsps      []string   `json:"orgmsps,required" example:"Org1MSP,Org2MSP,Org3MSP" binding:"required"`
	PeerNames    [][]string `json:"peernames,required" example:"[peer0.org1,peer1.org1],[peer0.org2,peer1.org2],[peer0.org3,peer1.org3]" binding:"required"`
	ChannelID    string     `json:"channelid,required" example:"mychannel" binding:"-"`
	OrdererNames []string   `json:"orderernames,required" example:"orderer0,orderer1,orderer2" binding:"required"`
}

// joinChannel godoc
// @Summary 加入通道管理
// @Description 功能：加入通道管理
// @Tags 加入通道管理
// @Accept  json
// @Produce  json
// @Param resource body  channel.joinChannelRequest true "加入通道管理"
// @Success 200 {object} gintool.ApiResponse
// @Failure 400 {object} gintool.ApiResponse
// @Router /channel/joinchannel [post]
func (c *ChannelController) joinChannel(ctx *gin.Context) {
	logger.Debug("joinChannel enter controller")
	req := new(joinChannelRequest)
	if err := ctx.ShouldBind(req); err != nil {
		logger.Error("params validate error:", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.INVALID_PARAMS, err), nil)
		return
	}
	ret, _ := json.MarshalIndent(req, models.Empty, models.TAB)
	logger.Debug("joinChannel req\n", string(ret))
	peerOrgInfos := models.NewDefaultPeerOrgInfos(req.OrgNames, req.OrgMsps, req.PeerNames)
	ordererOrgInfo := models.NewDefaultOrdererOrgInfo(req.OrdererNames)
	channelInfo := models.NewDefaultChannelInfo(req.ChannelID)
	err := c.channelService.JoinChannel(peerOrgInfos, ordererOrgInfo, channelInfo)
	if err != nil {
		logger.Error("joinChannel", err)
		gintool.ResultCodeWithData(ctx, utils.NewError(e.ERROR_JOIN_CHANNEL_FAILED, err), nil)
		return
	}
	gintool.ResultCodeWithData(ctx, nil, nil)
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
// 通道管理路由组
func FabChannelRouterGroup(app *ChannelController) gintool.RouteGroup {
	return gintool.RouteGroup{
		Route: []gintool.Route{
			gintool.NewRoute(http.MethodPost, "/createchannel", app.createChannel).AddComment("创建通道管理"),
			gintool.NewRoute(http.MethodPost, "/joinchannel", app.joinChannel).AddComment("加入通道管理"),
		},
		Prefix:  "/channel",
		Comment: "区块链通道管理",
		Module:  "channel",
	}
}
