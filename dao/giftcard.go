package dao

import (
	"log"
	"time"

	"github.com/ninjadotorg/handshake-wallet/bean"
	"github.com/ninjadotorg/handshake-wallet/database"
	"github.com/ninjadotorg/handshake-wallet/model"
	"github.com/ninjadotorg/handshake-wallet/utils"
)

type GiftCardDAO struct{}

func (dao GiftCardDAO) GetOrder(orderID uint, email string) (bool, model.GiftCardOrder) {
	db := database.Database()
	var order model.GiftCardOrder
	if err := db.Where("id = ?", orderID).Where("email = ?", email).First(&order).Error; err != nil {
		return false, order
	}
	return true, order
}

func (dao GiftCardDAO) GetOrderDetail(orderID uint) (bool, []model.GiftCardOrderDetail) {
	db := database.Database()
	var orderDetails []model.GiftCardOrderDetail
	if err := db.Where("order_id = ?", orderID).Find(&orderDetails).Error; err != nil {
		return false, orderDetails
	}
	return true, orderDetails
}

func (dao GiftCardDAO) CreateOrder(name string, email string, serviceFee float64, shippingType uint, ethAddress string, additionalFee string) (bool, uint) {
	db := database.Database()
	giftCardOrder := model.GiftCardOrder{
		Name:          name,
		Email:         email,
		ServiceFee:    serviceFee,
		ShippingType:  shippingType,
		EthAddress:    ethAddress,
		AdditionalFee: additionalFee,
	}

	if err := db.Save(&giftCardOrder).Error; err != nil {
		log.Print("CreateOrder Error", err.Error())
		return false, 0
	}
	return true, giftCardOrder.ID
}

func (dao GiftCardDAO) UpdateOrder(orderID uint, email string, transactionID string, contracID string) bool {
	db := database.Database()

	var order model.GiftCardOrder
	var success bool

	if success, order = dao.GetOrder(orderID, email); !success {
		log.Print("UpdateOrder Error: order not found")
		return false
	}

	order.TransactionID = transactionID
	order.ContractID = contracID

	if err := db.Save(&order).Error; err != nil {
		log.Print("UpdateOrder Error", err.Error())
		return false
	}

	return true
}

func (dao GiftCardDAO) UpdateOrderStatus(orderID uint, email string, status uint) bool {
	db := database.Database()

	var order model.GiftCardOrder
	var success bool

	if success, order = dao.GetOrder(orderID, email); !success {
		log.Print("UpdateOrderStatus Error: order not found")
		return false
	}

	order.Status = status
	if err := db.Save(&order).Error; err != nil {
		log.Print("UpdateOrderStatus Error:", err.Error())
		return false
	}

	return true
}

func (dao GiftCardDAO) CreateOrderDetail(orderID uint, shippingDetail *bean.GiftCardOrderFormShipping, amount float64, quantity uint) (bool, uint) {
	db := database.Database()
	giftCardOrderDetail := model.GiftCardOrderDetail{
		OrderID:         orderID,
		Amount:          amount,
		Quantity:        quantity,
		ShippingName:    shippingDetail.Name,
		ShippingAddress: shippingDetail.Address,
		ShippingCity:    shippingDetail.City,
		ShippingState:   shippingDetail.State,
		ShippingZip:     shippingDetail.Zip,
		ShippingCountry: shippingDetail.Country,
		ShippingPhone:   shippingDetail.Phone,
		ShippingEmail:   shippingDetail.Email,
	}

	if err := db.Save(&giftCardOrderDetail).Error; err != nil {
		log.Print("CreateOrderDetail Error", err.Error())
		return false, 0
	}
	return true, giftCardOrderDetail.ID
}

func (dao GiftCardDAO) CreateGiftCardCode(orderID uint, amount float64) (bool, string) {
	db := database.Database()
	code := utils.GenerateGiftCardCode()
	encodedCode := utils.Md5(code)
	giftCardCode := model.GiftCard{OrderID: orderID, Code: encodedCode, Amount: amount, RedeemDate: time.Now(), ExpirationDate: time.Now().AddDate(0, 6, 0)}

	if err := db.Save(&giftCardCode).Error; err != nil {
		log.Print("CreateGiftCardCode Error", err.Error())
		return false, ""
	}
	return true, code
}

func (dao GiftCardDAO) CheckCode(code string) (bool, float64, time.Time) {
	db := database.Database()
	var codeModel model.GiftCard
	encodedCode := utils.Md5(code)
	if err := db.Where("code = ?", encodedCode).First(&codeModel).Error; err != nil {
		log.Print(encodedCode)
		log.Print("CheckCode Error", err)
		return false, 0, time.Now()
	}

	return codeModel.Status == 1, codeModel.Amount, codeModel.ExpirationDate
}
