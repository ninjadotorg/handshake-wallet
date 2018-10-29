package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/handshake-wallet/api"
)

type OnChainRouter struct{}

var OnChainRouterInst = OnChainRouter{}

func (router OnChainRouter) Create(routerEngine *gin.Engine) *gin.RouterGroup {
	group := routerEngine.Group("/onchain")
	onChainApi := api.OnChainApi{}

	group.GET("/gift-card/order/scan", onChainApi.ScanOrders)
	group.GET("/gift-card/code/scan", onChainApi.ScanRedeemedCodes)

	return group
}
