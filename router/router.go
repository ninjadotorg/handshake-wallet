package router

import "github.com/gin-gonic/gin"

func SetupRouters(router *gin.Engine) {
	GiftCardRouterInst.Create(router)
	OnChainRouterInst.Create(router)
}
