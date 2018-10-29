package config

import (
	"github.com/ninjadotorg/handshake-wallet/service/cache"
)

func InitializeProject() {
	// initialize config
	Init()
	config := GetConfig()

	redisEndpoint := config.GetString("redis_endpoint")
	redisPassword := config.GetString("redis_password")
	cache.InitializeRedisClient(redisEndpoint, redisPassword)
}
