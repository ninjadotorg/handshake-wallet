package service

import (
	"time"

	"github.com/ninjadotorg/handshake-wallet/api_response"
	"github.com/ninjadotorg/handshake-wallet/dao"
	"github.com/ninjadotorg/handshake-wallet/form"
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
	orderModel := model.GiftCardOrder{
		TransactionID: updateOrderForm.TransactionID,
		Email:         updateOrderForm.Email,
	}

	err := service.dao.UpdateOrder(orderID, &orderModel)

	if ce.SetError(api_response.UpdateDataFailed, err) {
		return
	}

	order = orderModel

	return
}

func (service GiftCardService) CheckCode(checkCodeForm form.GiftCardCheckCodeForm) (giftCard model.GiftCard, ce SimpleContextError) {
	code := checkCodeForm.Code

	giftCardModel, err := service.dao.CheckCode(code)

	if err != nil {
		giftCard = model.GiftCard{
			ExpirationDate: time.Now(),
		}
	} else {
		giftCard = giftCardModel
	}
	return
}
