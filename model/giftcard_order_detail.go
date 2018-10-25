package model

import (
	_ "github.com/jinzhu/gorm"
)

type GiftCardOrderDetail struct {
	BaseModel
	OrderID         uint    `gorm:"column:order_id" sql:"not null" json:"order_id"`
	Amount          float64 `gorm:"column:amount;default:0" sql:"not null" json:"amount"`
	Quantity        uint    `gorm:"column:quantity;default:0" sql:"not null" json:"quantity"`
	ShippingName    string  `gorm:"column:shipping_name" json:"shipping_name"`
	ShippingAddress string  `gorm:"column:shipping_address" json:"shipping_address"`
	ShippingCity    string  `gorm:"column:shipping_city" json:"shipping_city"`
	ShippingState   string  `gorm:"column:shipping_state" json:"shipping_state"`
	ShippingZip     string  `gorm:"column:shipping_zip" json:"shipping_zip"`
	ShippingCountry string  `gorm:"column:shipping_country" json:"shipping_country"`
	ShippingPhone   string  `gorm:"column:shipping_phone" json:"shipping_phone"`
	ShippingEmail   string  `gorm:"column:shipping_email" sql:"not null" json:"shipping_email"`
	PersonalMessage string  `gorm:"column:personal_message" sql:"not null" json:"personal_message"`
}

func (g GiftCardOrderDetail) TableName() string {
	return "gift_card_order_detail"
}
