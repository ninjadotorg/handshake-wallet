package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/handshake-wallet/dao"
	"github.com/ninjadotorg/handshake-wallet/utils"
)

type GiftCardApi struct{}

func (api GiftCardApi) Generate(c *gin.Context) {
	numberOfCode, convertErr := strconv.Atoi(c.DefaultPostForm("number_of_code", "1"))

	if convertErr != nil {
		resp := JsonResponse{0, "Please enter valid number of code to be generated", nil}
		c.JSON(http.StatusOK, resp)
	}

	giftCardDao := dao.GiftCardDAO{}
	totalCodeGenerated := 0

	for i := 0; i < numberOfCode; i++ {
		code := utils.GenerateGiftCardCode()
		success := giftCardDao.AddCode(code)
		if success {
			totalCodeGenerated++
		}
	}

	remainingCode := numberOfCode - totalCodeGenerated
	var msg string
	if msg = fmt.Sprintf("Only %d codes generated", totalCodeGenerated); remainingCode == 0 {
		msg = fmt.Sprintf("All of %d code generated", numberOfCode)
	}
	resp := JsonResponse{1, msg, nil}
	c.JSON(http.StatusOK, resp)
}
