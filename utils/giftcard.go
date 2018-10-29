package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/segmentio/ksuid"
)

const GIFT_CARD_ORDER_ID_PREFIX = "GIFTCARDORDER-"

func GenerateGiftCardCode() string {
	guid := ksuid.New()
	return fmt.Sprintf("NINJA-%s", guid.String())
}

func CreateOrderID(orderNumber uint) string {
	return fmt.Sprintf("%s%08d", GIFT_CARD_ORDER_ID_PREFIX, orderNumber)
}

func GetOrderNumber(orderID string) uint {
	orderID = strings.Replace(orderID, GIFT_CARD_ORDER_ID_PREFIX, "", -1)
	orderNumber, err := strconv.ParseUint(orderID, 10, 16)
	if err != nil {
		orderNumber = 0
	}
	return uint(orderNumber)
}
