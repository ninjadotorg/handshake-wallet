package model

import (
	_ "github.com/jinzhu/gorm"
)

type GiftCardOrderDetail struct {
	BaseModel
	OrderID  uint    `gorm:"column:order_id" sql:"not null" json:"order_id"`
	Amount   float64 `gorm:"column:amount;default:0" sql:"not null" json:"amount"`
	Quantity uint    `gorm:"column:quantity;default:0" sql:"not null" json:"quantity"`
}

func (g GiftCardOrderDetail) TableName() string {
	return "gift_card_order_detail"
}
