package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/ninjadotorg/handshake-wallet/config"
	"github.com/ninjadotorg/handshake-wallet/router"
)

func main() {
	log.Print("Start Wallet Service")

	config.InitializeProject()
	conf := config.GetConfig()

	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/wallet.log",
		MaxSize:    10, // megabytes
		MaxBackups: 10,
		MaxAge:     30,   //days
		Compress:   true, // disabled by default
	})
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	routerEngine := gin.New()
	routerEngine.Use(gin.Logger())

	// setup router
	router.SetupRouters(routerEngine)

	// start server
	routerEngine.Run(fmt.Sprintf(":%d", conf.GetInt("port")))
}
