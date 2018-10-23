package main

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/ninjadotorg/handshake-wallet/config"
	"github.com/ninjadotorg/handshake-wallet/database"
	"github.com/ninjadotorg/handshake-wallet/model"
)

func main() {
	log.Print("Start migrate database")
	config.Init()

	var db *gorm.DB
	db = database.Database()
	defer db.Close()

	db.AutoMigrate(&model.GiftCard{})
}
