package dao

import (
	"log"
	"time"

	"github.com/ninjadotorg/handshake-wallet/database"
	"github.com/ninjadotorg/handshake-wallet/model"
	"github.com/ninjadotorg/handshake-wallet/utils"
)

type GiftCardDAO struct{}

func (dao GiftCardDAO) IsOrderExists(orderID string) bool {
	db := database.Database()
	var order model.GiftCardOrder
	return db.Where("order_id = ?", orderID).First(&order).Error == nil
}

func (dao GiftCardDAO) CreateOrder(orderID string, serviceFee float64, createdUserID uint, transactionID string, contractID string, deliveryType uint, additionalFee string) (bool, uint) {
	db := database.Database()
	giftCardOrder := model.GiftCardOrder{OrderID: orderID, ServiceFee: serviceFee, BuyerUserID: createdUserID, TransactionID: transactionID, ContractID: contractID, DeliveryType: deliveryType, AdditionalFee: additionalFee}

	if err := db.Save(&giftCardOrder).Error; err != nil {
		log.Print("CreateOrder Error", err.Error())
		return false, 0
	}
	return true, giftCardOrder.ID
}

func (dao GiftCardDAO) CreateOrderDetail(orderID uint, amount float64, quantity uint) (bool, uint) {
	db := database.Database()
	giftCardOrderDetail := model.GiftCardOrderDetail{OrderID: orderID, Amount: amount, Quantity: quantity}

	if err := db.Save(&giftCardOrderDetail).Error; err != nil {
		log.Print("CreateOrderDetail Error", err.Error())
		return false, 0
	}

	var i uint
	for i = 0; i < quantity; i++ {
		dao.CreateGiftCardCode(orderID, amount)
	}
	return true, giftCardOrderDetail.ID
}

func (dao GiftCardDAO) CreateGiftCardCode(orderID uint, amount float64) (bool, string) {
	db := database.Database()
	code := utils.GenerateGiftCardCode()
	giftCardCode := model.GiftCard{OrderID: orderID, Code: code, Amount: amount, RedeemDate: time.Now(), ExpirationDate: time.Now().AddDate(0, 6, 0)}

	if err := db.Save(&giftCardCode).Error; err != nil {
		log.Print("CreateGiftCardCode Error", err.Error())
		return false, ""
	}
	return true, code
}

func (dao GiftCardDAO) CheckCode(code string) (bool, float64, time.Time) {
	db := database.Database()
	var codeModel model.GiftCard

	if err := db.Where("code = ?", code).First(&codeModel).Error; err != nil {
		return false, 0, time.Now()
	}

	return codeModel.Status == 1, codeModel.Amount, codeModel.ExpirationDate
}
