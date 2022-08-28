package fabapi

import (
	"os"

	"github.com/wsw365904/fabapi/core/fabsdk/chaincodeimpl"
	"github.com/wsw365904/fabapi/core/fabsdk/channelimpl"
	"github.com/wsw365904/fabapi/core/fabsdk/fabsdkimpl"
	"github.com/wsw365904/fabapi/core/fabsdk/ledgerimpl"
	"github.com/wsw365904/fabapi/core/fabsdk/models"
	"github.com/wsw365904/fabapi/internal/app/controller/allcontroller"
	"github.com/wsw365904/fabapi/internal/app/controller/chaincodecontroller"
	"github.com/wsw365904/fabapi/internal/app/controller/channelcontroller"
	"github.com/wsw365904/fabapi/internal/app/controller/ledgercontroller"
	"github.com/wsw365904/fabapi/internal/app/controller/logcontroller"
	"github.com/wsw365904/fabapi/internal/app/service/chaincodeserviceimpl"
	"github.com/wsw365904/fabapi/internal/app/service/channelserviceimpl"
	"github.com/wsw365904/fabapi/internal/app/service/ledgerserviceimpl"
	"github.com/wsw365904/fabapi/internal/app/service/logserviceimpl"
	"github.com/wsw365904/fabapi/internal/pkg/config/fabconfig"
	"github.com/wsw365904/fabapi/internal/router"

	"github.com/urfave/cli"

	"go.uber.org/dig"
)

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func buildContainer() (*dig.Container, error) {
	container := dig.New()

	// new sdk  func NewFabSdk() (FabSdk, *fabsdk.FabricSDK, error) {

	err := container.Provide(fabsdkimpl.NewFabSdk)
	if err != nil {
		return nil, err
	}
	// new config func NewConfig() (Config, *models.Other, error) {

	err = container.Provide(fabconfig.NewConfig)
	if err != nil {
		return nil, err
	}
	// new chain code op func NewChaincodeOp(other *models.Other, sdk *fabsdk.FabricSDK) *ChaincodeOp {

	err = container.Provide(chaincodeimpl.NewChaincodeOp)
	if err != nil {
		return nil, err
	}
	// new channel op func NewChannelOp(other *models.Other, sdk *fabsdk.FabricSDK) *ChannelOp {

	err = container.Provide(channelimpl.NewChannelOp)
	if err != nil {
		return nil, err
	}
	// new ledger op func NewLedgerOp(other *models.Other, sdk *fabsdk.FabricSDK) *LedgerOp {

	err = container.Provide(ledgerimpl.NewLedgerOp)
	if err != nil {
		return nil, err
	}
	// new chaincode service  func NewChaincodeService(cop internalfabsdk.Chaincode) *ChaincodeService {

	err = container.Provide(chaincodeserviceimpl.NewChaincodeService)
	if err != nil {
		return nil, err
	}
	// new channel service func NewChannelService(cop internalfabsdk.Channel) *ChannelService {
	err = container.Provide(channelserviceimpl.NewChannelService)
	if err != nil {
		return nil, err
	}
	// new ledger service func NewLedgerService(lop internalfabsdk.Ledger) *LedgerService {
	err = container.Provide(ledgerserviceimpl.NewLedgerService)
	if err != nil {
		return nil, err
	}
	// new log service func NewLogService() *LogService {

	err = container.Provide(logserviceimpl.NewLogService)
	if err != nil {
		return nil, err
	}
	// new chaincode controller func NewChaincodeController(chaincodeService service.ChaincodeService) *ChaincodeController {

	err = container.Provide(chaincodecontroller.NewChaincodeController)
	if err != nil {
		return nil, err
	}
	// new channel controller func NewChannelController(channelService service.ChannelService) *ChannelController {

	err = container.Provide(channelcontroller.NewChannelController)
	if err != nil {
		return nil, err
	}
	// new ledger controller func NewChannelController(channelService service.ChannelService) *ChannelController {

	err = container.Provide(ledgercontroller.NewLedgerController)
	if err != nil {
		return nil, err
	}
	// new log controller func NewLogController(log service.LogService) *LogController {

	err = container.Provide(logcontroller.NewLogController)
	if err != nil {
		return nil, err
	}

	// new controller  func NewAllController(chaincodeController *chaincode.ChaincodeController,
	// 	channelController *channel.ChannelController,
	// 	ledgerController *ledger.LedgerController,
	// 	logController *log.LogController,
	// 	fabSdk fabsdk.FabSdk) (*Controller, error) {
	err = container.Provide(allcontroller.NewController)
	if err != nil {
		return nil, err
	}
	// newCliApp func newCliApp() *cli.App {

	err = container.Provide(newCliApp)
	if err != nil {
		return nil, err
	}

	return container, nil
}

// StartMain /***************************************************************
func StartMain() {
	logger.Debug("StartMain enter")
	container, err := buildContainer() // 创建
	if err != nil {
		logger.Fatal(err)
	}
	err = container.Invoke(func(app *cli.App, appController *allcontroller.Controller, other *models.Other) error {
		err := app.Run(os.Args)
		if err != nil {
			return err
		}
		router.Start(appController, other)
		return nil
	})
	if err != nil {
		logger.Fatal(err)
	}
}
