package dao

import (
	"github.com/ninjadotorg/handshake-wallet/database"
	"github.com/ninjadotorg/handshake-wallet/model"
)

type GiftCardDAO struct{}

func (dao GiftCardDAO) AddCode(code string) bool {
	db := database.Database()
	giftCardModel := model.GiftCard{Code: code}
	errDb := db.Save(&giftCardModel).Error
	return errDb == nil
}
