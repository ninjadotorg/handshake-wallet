package model

import (
	_ "github.com/jinzhu/gorm"
)

type GiftCardOrder struct {
	BaseModel
	Name          string  `gorm:"name" sql:"not null" json:"name"`
	Email         string  `gorm:"column:email" sql:"not null" json:"email"`
	ServiceFee    float64 `gorm:"column:service_fee;default:0" sql:"not null" json:"service_fee"`
	AdditionalFee string  `gorm:"column:additional_fee" sql:"type:text" json:"additional_fee"`
	ShippingType  uint    `gorm:"column:shipping_type" sql:"not null" json:"shipping_type"`
	ContractID    string  `gorm:"column:contract_id" sql:"not null" json:"contract_id"`
	TransactionID string  `gorm:"column:transaction_id" sql:"not null" json:"transaction_id"`
	EthAddress    string  `gorm:"column:eth_address" json:"eth_address"`
	Status        uint    `gorm:"column:status;default:0" json:"status"`
}

func (g GiftCardOrder) TableName() string {
	return "gift_card_order"
}
