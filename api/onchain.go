package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/handshake-wallet/service"
)

type OnChainApi struct{}

// CRON JOB
func (api OnChainApi) ScanOrders(c *gin.Context) {
	service.OnChainServiceInst.ScanGiftCardOrders()
}

// CRON JOB
func (api OnChainApi) ScanRedeemedCodes(c *gin.Context) {
	service.OnChainServiceInst.ScanUsedCodes()
}
