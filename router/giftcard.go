package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ninjadotorg/handshake-wallet/api"
)

type GiftCardRouter struct{}

func (router GiftCardRouter) Create(routerEngine *gin.Engine) *gin.RouterGroup {
	group := routerEngine.Group("/gift-card")
	giftCardAPI := api.GiftCardApi{}

	group.GET("/order/generate-order-id", giftCardAPI.GenerateOrderID)
	group.POST("/order/create", giftCardAPI.CreateOrder)
	group.POST("/check", giftCardAPI.CheckCode)

	return group
}
