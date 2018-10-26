package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ninjadotorg/handshake-wallet/api_response"
	"github.com/ninjadotorg/handshake-wallet/common"
	"github.com/ninjadotorg/handshake-wallet/dao"
	"github.com/ninjadotorg/handshake-wallet/form"
	"github.com/ninjadotorg/handshake-wallet/integration/redeemhandshake_service"
	"github.com/ninjadotorg/handshake-wallet/model"
)

type GiftCardService struct {
	dao *dao.GiftCardDAO
}

func (service GiftCardService) CreateOrder(orderForm form.GiftCardCreateOrderForm) (order model.GiftCardOrder, ce SimpleContextError) {
	orderModel := model.GiftCardOrder{
		Name:          orderForm.Name,
		Email:         orderForm.Email,
		ServiceFee:    orderForm.ServiceFee,
		ShippingType:  orderForm.ShippingType,
		EthAddress:    orderForm.EthAddress,
		AdditionalFee: orderForm.AdditionalFee,
	}

	err := service.dao.CreateOrder(&orderModel)

	if ce.SetError(api_response.AddDataFailed, err) {
		return
	}

	// create order details
	for _, orderDetailForm := range orderForm.OrderDetails {
		orderDetailModel := model.GiftCardOrderDetail{
			OrderID:         orderModel.ID,
			Amount:          orderDetailForm.Amount,
			Quantity:        orderDetailForm.Quantity,
			ShippingName:    orderDetailForm.ShippingDetail.Name,
			ShippingAddress: orderDetailForm.ShippingDetail.Address,
			ShippingCity:    orderDetailForm.ShippingDetail.City,
			ShippingState:   orderDetailForm.ShippingDetail.State,
			ShippingZip:     orderDetailForm.ShippingDetail.Zip,
			ShippingCountry: orderDetailForm.ShippingDetail.Country,
			ShippingEmail:   orderDetailForm.ShippingDetail.Email,
			ShippingPhone:   orderDetailForm.ShippingDetail.Phone,
			PersonalMessage: orderDetailForm.PersonalMessage,
		}

		err := service.dao.CreateOrderDetail(&orderDetailModel)

		if ce.SetError(api_response.AddDataFailed, err) {
			return
		}
	}

	order = orderModel

	return
}

func (service GiftCardService) UpdateOrder(orderID uint, updateOrderForm form.GiftCardUpdateOrderForm) (order model.GiftCardOrder, ce SimpleContextError) {
	orderModel, err := service.dao.GetOrder(orderID, updateOrderForm.Email)

	if ce.SetError(api_response.GetDataFailed, err) {
		return
	}

	orderModel.TransactionID = updateOrderForm.TransactionID
	err = service.dao.UpdateOrder(&orderModel)

	if ce.SetError(api_response.UpdateDataFailed, err) {
		return
	}

	order = orderModel

	return
}

func (service GiftCardService) CheckCode(checkCodeForm form.GiftCardCheckCodeForm) (giftCard model.GiftCard, ce SimpleContextError) {
	giftCardModel, err := service.dao.GetCode(checkCodeForm.Code)

	if err != nil {
		giftCard = model.GiftCard{
			ExpirationDate: time.Now(),
		}
	} else {
		giftCard = giftCardModel
	}
	return
}

func (service GiftCardService) RedeemCode(redeemForm form.GiftCardRedeemForm) (giftCard model.GiftCard, ce SimpleContextError) {
	giftCardModel, err := service.dao.GetCode(redeemForm.Code)

	if ce.SetError(api_response.InvalidGiftCardCode, err) {
		return
	}

	if giftCardModel.Status != 0 {
		ce.SetStatusKey(api_response.GiftCardCodeRedeemed)
		return
	}

	if giftCardModel.IsExpired() {
		ce.SetStatusKey(api_response.ExpiredGiftCardCode)
		return
	}

	orderModel, err := service.dao.GetOrder(giftCardModel.OrderID, "")
	if ce.SetError(api_response.GetDataFailed, err) {
		return
	}

	if orderModel.Status != 1 {
		ce.SetStatusKey(api_response.RedeemGiftCardCodeFailed)
		return
	}

	redeemID, err := strconv.Atoi(orderModel.ContractID)

	if ce.SetError(api_response.RedeemGiftCardCodeFailed, err) {
		return
	}

	// set status to processing
	giftCardModel.Status = 2
	err = service.dao.UpdateCode(&giftCardModel)

	if ce.SetError(api_response.UpdateDataFailed, err) {
		return
	}

	contractClient := redeemhandshake_service.RedeemHandshakeClient{}
	amount := common.Float64ToDecimal(giftCardModel.Amount)
	giftCardID := fmt.Sprint(giftCardModel.ID)

	address := redeemForm.ToEthAddress
	txHash, err := contractClient.UseRedeem(giftCardID, redeemID, amount, address, "")

	if ce.SetError(api_response.RedeemGiftCardCodeFailed, err) {
		return
	}

	giftCardModel.TransactionHash = txHash
	err = service.dao.UpdateCode(&giftCardModel)
	if ce.SetError(api_response.UpdateDataFailed, err) {
		return
	}

	giftCard = giftCardModel
	return
}
