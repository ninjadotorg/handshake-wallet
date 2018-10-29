package dao

import (
	"log"
	"time"

	"github.com/ninjadotorg/handshake-wallet/model"
	"github.com/ninjadotorg/handshake-wallet/utils"
)

type GiftCardDAO struct{}

func (dao GiftCardDAO) CreateOrder(orderModel *model.GiftCardOrder) error {
	if err := GetDB().Save(&orderModel).Error; err != nil {
		log.Print("CreateOrder Error", err.Error())
		return err
	}

	return nil
}

func (dao GiftCardDAO) UpdateOrder(orderModel *model.GiftCardOrder) error {
	if err := GetDB().Save(&orderModel).Error; err != nil {
		log.Print("UpdateOrder Error", err.Error())
		return err
	}

	return nil
}

func (dao GiftCardDAO) GetOrder(orderID uint, email string) (model.GiftCardOrder, error) {
	var order model.GiftCardOrder
	query := GetDB().Where("id = ?", orderID)
	if email != "" {
		query = query.Where("email = ?", email)
	}
	if err := query.First(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (dao GiftCardDAO) CreateOrderDetail(orderDetailModel *model.GiftCardOrderDetail) error {
	if err := GetDB().Save(&orderDetailModel).Error; err != nil {
		log.Print("CreateOrderDetail Error", err.Error())
		return err
	}
	return nil
}

func (dao GiftCardDAO) GetOrderDetails(orderID uint) ([]model.GiftCardOrderDetail, error) {
	var orderDetails []model.GiftCardOrderDetail
	if err := GetDB().Where("order_id = ?", orderID).Find(&orderDetails).Error; err != nil {
		return orderDetails, err
	}

	return orderDetails, nil
}

func (dao GiftCardDAO) CreateCode(orderID uint, amount float64) (string, error) {
	code := utils.GenerateGiftCardCode()
	encodedCode := utils.Md5(code)
	giftCardCode := model.GiftCard{
		OrderID:    orderID,
		Code:       encodedCode,
		Amount:     amount,
		RedeemDate: time.Now(),
	}

	if err := GetDB().Save(&giftCardCode).Error; err != nil {
		log.Print("CreateCode Error", err.Error())
		return "", err
	}
	return code, nil
}

func (dao GiftCardDAO) GetCode(code string) (model.GiftCard, error) {
	var codeModel model.GiftCard
	encodedCode := utils.Md5(code)

	if err := GetDB().Where("code = ?", encodedCode).First(&codeModel).Error; err != nil {
		log.Print("CheckCode Error", err)
		return codeModel, err
	}

	return codeModel, nil
}

func (dao GiftCardDAO) GetCodeWithoutEncrypt(code string) (model.GiftCard, error) {
	var codeModel model.GiftCard

	if err := GetDB().Where("code = ?", code).First(&codeModel).Error; err != nil {
		log.Print("CheckCode Error", err)
		return codeModel, err
	}

	return codeModel, nil
}

func (dao GiftCardDAO) UpdateCode(giftCardModel *model.GiftCard) error {
	if err := GetDB().Save(&giftCardModel).Error; err != nil {
		log.Print("UpdateCode Error", err.Error())
		return err
	}

	return nil
}
