package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/handshake-wallet/api_response"
	"github.com/ninjadotorg/handshake-wallet/service"
)

type OnChainApi struct{}

// CRON JOB
func (api OnChainApi) ScanOrders(c *gin.Context) {
	service.OnChainServiceInst.ScanGiftCardOrders()
	api_response.SuccessResponse(c, nil)
}

// CRON JOB
func (api OnChainApi) ScanRedeemedCodes(c *gin.Context) {
	service.OnChainServiceInst.ScanUsedCodes()
	api_response.SuccessResponse(c, nil)
}
