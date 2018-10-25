package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/ninjadotorg/handshake-wallet/common"
)

func GetDB() *gorm.DB {
	return common.Database()
}
