package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/handshake-wallet/config"
	"github.com/ninjadotorg/handshake-wallet/router"
)

func main() {
	log.Print("Start Wallet Service")

	logFile, err := os.OpenFile("logs/log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(gin.DefaultWriter) // You may need this
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	routerEngine := gin.New()
	routerEngine.Use(gin.Logger())

	// setup router
	router.SetupRouters(routerEngine)

	// initialize config
	config.Init()
	// get current config
	config := config.GetConfig()

	// start server
	routerEngine.Run(fmt.Sprintf(":%d", config.GetInt("port")))
}
