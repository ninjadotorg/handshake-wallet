package model

import (
	_ "github.com/jinzhu/gorm"
)

type GiftCardOrder struct {
	BaseModel
	OrderID       string  `gorm:"column:order_id" sql:"not null" json:"order_id"`
	BuyerUserID   uint    `gorm:"column:buyer_user_id" sql:"not null" json:"buyer_user_id"`
	ServiceFee    float64 `gorm:"column:service_fee;default:0" sql:"not null" json:"service_fee"`
	AdditionalFee string  `gorm:"column:additional_fee" sql:"type:text" json:"additional_fee"`
	DeliveryType  uint    `gorm:"column:delivery_type" sql:"not null" json:"delivery_type"`
	ContractID    string  `gorm:"column:contract_id" sql:"not null" json:"contract_id"`
	TransactionID string  `gorm:"column:transaction_id" sql:"not null" json:"transaction_id"`
	Status        uint    `gorm:"column:status;default:0" json:"status"`
}

func (g GiftCardOrder) TableName() string {
	return "gift_card_order"
}
