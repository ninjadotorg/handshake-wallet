package main

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/ninjadotorg/handshake-wallet/common"
	"github.com/ninjadotorg/handshake-wallet/config"
	"github.com/ninjadotorg/handshake-wallet/model"
)

func main() {
	log.Print("Start migrate database")
	config.Init()

	var db *gorm.DB
	db = common.Database()
	defer db.Close()

	db.AutoMigrate(&model.GiftCard{})
	db.AutoMigrate(&model.GiftCardOrder{})
	db.AutoMigrate(&model.GiftCardOrderDetail{})

	log.Print("Migrate success")
}
