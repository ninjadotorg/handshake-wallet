package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ninjadotorg/handshake-wallet/config"
	"github.com/segmentio/ksuid"
)

const GIFT_CARD_ORDER_ID_PREFIX = "gift-card-order-"

func GenerateGiftCardCode() string {
	guid := ksuid.New()
	return fmt.Sprintf("NINJA-%s", guid.String())
}

func CreateOrderID(orderNumber uint) string {
	conf := config.GetConfig()
	secretKey := conf.GetString("secret_key")
	orderID := fmt.Sprintf("%s%08d", GIFT_CARD_ORDER_ID_PREFIX, orderNumber)
	orderIDEncoded, err := HashEncrypt([]byte(secretKey), orderID)
	if err != nil {
		orderIDEncoded = ""
	}
	return orderIDEncoded
}

func GetOrderNumber(orderID string) int {
	conf := config.GetConfig()
	secretKey := conf.GetString("secret_key")
	orderID, _ = HashDecrypt([]byte(secretKey), orderID)
	orderID = strings.Replace(orderID, GIFT_CARD_ORDER_ID_PREFIX, "", -1)
	orderNumber, err := strconv.Atoi(orderID)
	if err != nil {
		orderNumber = -1
	}
	return orderNumber
}
