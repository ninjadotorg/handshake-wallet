package common

import (
	"github.com/gin-gonic/gin"
)

func GetHeaderWithDefault(context *gin.Context, key string, defaultValue string) string {
	value := context.GetHeader(key)
	if value == "" {
		value = defaultValue
	}

	return value
}

func GetLanguage(context *gin.Context) string {
	return GetHeaderWithDefault(context, "Custom-Language", "en-US")
}

func GetUserId(context *gin.Context) string {
	return GetHeaderWithDefault(context, "Uid", "")
}

func GetChainId(context *gin.Context) string {
	return GetHeaderWithDefault(context, "Chainid", "0")
}

func GetFCM(context *gin.Context) string {
	return GetHeaderWithDefault(context, "Fcm-Token", "")
}
