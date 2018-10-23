package router

import "github.com/gin-gonic/gin"

func SetupRouters(router *gin.Engine) {
	giftCardRouter := GiftCardRouter{}
	giftCardRouter.Create(router)
}
