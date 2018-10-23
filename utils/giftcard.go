package utils

import (
	"fmt"

	"github.com/rs/xid"
)

func GenerateGiftCardCode() string {
	guid := xid.New()
	return fmt.Sprintf("NINJA%s", guid.String())
}
