package model

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type GiftCard struct {
	BaseModel
	OrderID        uint      `gorm:"column:order_id" sql:"not null" json:"order_id"`
	Code           string    `gorm:"column:code" sql:"not null" json:"code"`
	RedeemUserID   uint      `gorm:"column:redeem_user_id" json:"redeem_user_id"`
	Amount         float64   `gorm:"column:amount;default:0" sql:"not null" json:"amount"`
	RedeemDate     time.Time `gorm:"redeem_date" json:"redeem_date"`
	ExpirationDate time.Time `gorm:"expiration_date" json:"expiration_date"`
	Status         uint      `gorm:"column:status;default:0" sql:"not null" json:"status"`
}

func (g GiftCard) TableName() string {
	return "gift_card"
}
