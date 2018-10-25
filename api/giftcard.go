package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/handshake-wallet/api_response"
	"github.com/ninjadotorg/handshake-wallet/common"
	"github.com/ninjadotorg/handshake-wallet/form"
	"github.com/ninjadotorg/handshake-wallet/service"
	"github.com/ninjadotorg/handshake-wallet/utils"
)

type GiftCardApi struct{}

func (api GiftCardApi) CreateOrder(c *gin.Context) {
	var orderForm form.GiftCardCreateOrderForm

	if common.ValidateBody(c, &orderForm) != nil {
		return
	}

	order, ce := service.GiftCardServiceInst.CreateOrder(orderForm)

	if ce.ContextValidate(c) {
		return
	}

	api_response.SuccessResponse(c, utils.CreateOrderID(order.ID))
}

func (api GiftCardApi) UpdateOrder(c *gin.Context) {
	var orderForm form.GiftCardUpdateOrderForm

	if common.ValidateBody(c, &orderForm) != nil {
		return
	}

	_, ce := service.GiftCardServiceInst.UpdateOrder(uint(utils.GetOrderNumber(orderForm.OrderID)), orderForm)

	if ce.ContextValidate(c) {
		return
	}

	api_response.SuccessResponse(c, nil)
}

func (api GiftCardApi) CheckCode(c *gin.Context) {
	var checkCodeForm form.GiftCardCheckCodeForm

	if common.ValidateBody(c, &checkCodeForm) != nil {
		return
	}

	giftCode, ce := service.GiftCardServiceInst.CheckCode(checkCodeForm)

	if ce.ContextValidate(c) {
		return
	}

	api_response.SuccessResponse(c, map[string]interface{}{
		"is_redeemed":     giftCode.Status == 1,
		"expiration_date": giftCode.ExpirationDate,
		"amount":          giftCode.Amount,
	})
}
