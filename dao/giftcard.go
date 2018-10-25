package dao

import (
	"log"
	"time"

	"github.com/ninjadotorg/handshake-wallet/model"
	"github.com/ninjadotorg/handshake-wallet/utils"
)

type GiftCardDAO struct{}

func (dao GiftCardDAO) GetOrder(orderID uint, email string) (model.GiftCardOrder, error) {
	var order model.GiftCardOrder
	if err := GetDB().Where("id = ?", orderID).Where("email = ?", email).First(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (dao GiftCardDAO) GetOrderDetails(orderID uint) ([]model.GiftCardOrderDetail, error) {
	var orderDetails []model.GiftCardOrderDetail
	if err := GetDB().Where("order_id = ?", orderID).Find(&orderDetails).Error; err != nil {
		return orderDetails, err
	}
	return orderDetails, nil
}

func (dao GiftCardDAO) CreateOrder(orderModel *model.GiftCardOrder) error {
	if err := GetDB().Save(&orderModel).Error; err != nil {
		log.Print("CreateOrder Error", err.Error())
		return err
	}

	return nil
}

func (dao GiftCardDAO) UpdateOrder(orderID uint, orderModel *model.GiftCardOrder) error {
	var order model.GiftCardOrder
	var err error

	if order, err = dao.GetOrder(orderID, orderModel.Email); err != nil {
		log.Print("UpdateOrder Error: order not found")
		return err
	}

	order.TransactionID = orderModel.TransactionID
	order.ContractID = orderModel.ContractID

	if err := GetDB().Save(&order).Error; err != nil {
		log.Print("UpdateOrder Error", err.Error())
		return err
	}

	orderModel = &order

	return nil
}

func (dao GiftCardDAO) UpdateOrderStatus(orderID uint, orderModel *model.GiftCardOrder) error {
	var order model.GiftCardOrder
	var err error

	if order, err = dao.GetOrder(orderModel.ID, orderModel.Email); err != nil {
		log.Print("UpdateOrderStatus Error: order not found")
		return err
	}

	order.Status = orderModel.Status
	if err := GetDB().Save(&order).Error; err != nil {
		log.Print("UpdateOrderStatus Error:", err.Error())
		return err
	}

	return nil
}

func (dao GiftCardDAO) CreateOrderDetail(orderDetailModel *model.GiftCardOrderDetail) error {
	if err := GetDB().Save(&orderDetailModel).Error; err != nil {
		log.Print("CreateOrderDetail Error", err.Error())
		return err
	}
	return nil
}

func (dao GiftCardDAO) CreateGiftCardCode(orderID uint, amount float64) (string, error) {
	code := utils.GenerateGiftCardCode()
	encodedCode := utils.Md5(code)
	giftCardCode := model.GiftCard{
		OrderID:        orderID,
		Code:           encodedCode,
		Amount:         amount,
		RedeemDate:     time.Now(),
		ExpirationDate: time.Now().AddDate(0, 6, 0),
	}

	if err := GetDB().Save(&giftCardCode).Error; err != nil {
		log.Print("CreateGiftCardCode Error", err.Error())
		return "", err
	}
	return code, nil
}

func (dao GiftCardDAO) CheckCode(code string) (model.GiftCard, error) {
	var codeModel model.GiftCard
	encodedCode := utils.Md5(code)

	if err := GetDB().Where("code = ?", encodedCode).First(&codeModel).Error; err != nil {
		log.Print("CheckCode Error", err)
		return codeModel, err
	}

	return codeModel, nil
}
