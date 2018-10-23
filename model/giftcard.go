package model

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type GiftCard struct {
	BaseModel
	Code         string    `gorm:"column:code" json:"code"`
	BuyerUserID  uint      `gorm:"column:buyer_user_id" json:"buyer_user_id"`
	RedeemUserID uint      `gorm:"column:redeem_user_id" json:"redeem_user_id"`
	Amount       float64   `gorm:"column:amount;default:0" json:"amount"`
	Fee          float64   `gorm:"column:fee;default:0" json:"fee"`
	ContractID   uint      `gorm:"column:contract_id;default:0" json:"contract_id"`
	IsRedeemed   uint      `gorm:"column:is_redeemed;default:0" json:"is_redeemed"`
	RedeemDate   time.Time `gorm:"redeem_date" json:"redeem_date"`
	Status       uint      `gorm:"column:status;default:0" json:"status"`
}

func (g GiftCard) TableName() string {
	return "gift_card"
}
