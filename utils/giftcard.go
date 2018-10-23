package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/segmentio/ksuid"
)

func GenerateGiftCardCode() string {
	guid := ksuid.New()
	return fmt.Sprintf("NINJA-%s", guid.String())
}

func GenerateGiftCardOrderID() string {
	guid := ksuid.New()
	orderID := md5.Sum([]byte(fmt.Sprintf("gift-card-order-%s", guid.String())))
	return hex.EncodeToString(orderID[:])
}
