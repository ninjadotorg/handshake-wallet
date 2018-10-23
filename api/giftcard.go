package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/handshake-wallet/bean"
	"github.com/ninjadotorg/handshake-wallet/dao"
	"github.com/ninjadotorg/handshake-wallet/utils"
)

type GiftCardApi struct{}

func (api GiftCardApi) GenerateOrderID(c *gin.Context) {
	orderID := utils.GenerateGiftCardOrderID()
	resp := CodeMessage[Success]
	resp.Data = orderID
	c.JSON(resp.StatusCode, resp)
}

func (api GiftCardApi) CreateOrder(c *gin.Context) {
	var orderForm bean.CreateOrderFormJSON
	if err := c.ShouldBindJSON(&orderForm); err != nil {
		resp := CodeMessage[InvalidRequestBody]
		resp.Message = "Invalid order data"
		c.JSON(resp.StatusCode, resp)
		return
	}

	orderID := orderForm.OrderID
	orderDetails := orderForm.OrderDetails
	serviceFee := orderForm.ServiceFee
	createdUserID := orderForm.CreatedUserID
	transactionID := orderForm.TransactionID
	contractID := orderForm.ContractID
	deliveryType := orderForm.DeliveryType
	additionalFee := orderForm.AdditionalFee

	var errors []string

	if orderID == "" {
		errors = append(errors, "Invalid Order ID")
	}

	if len(orderDetails) < 1 {
		errors = append(errors, "Invalid Order Details")
	}

	if serviceFee < 0 {
		errors = append(errors, "Invalid Service Fee")
	}

	if createdUserID < 0 {
		errors = append(errors, "Invalid User ID")
	}

	if transactionID == "" {
		errors = append(errors, "Invalid Transaction ID")
	}

	if contractID == "" {
		errors = append(errors, "Invalid Contract ID")
	}

	if deliveryType < 1 || deliveryType > 2 {
		errors = append(errors, "Invalid Delivery Type")
	}

	if len(errors) > 0 {
		resp := CodeMessage[InvalidRequestBody]
		resp.Data = errors
		c.JSON(resp.StatusCode, resp)
		return
	}

	giftCardDAO := dao.GiftCardDAO{}
	var oid uint
	var success bool

	if giftCardDAO.IsOrderExists(orderID) {
		resp := CodeMessage[OrderIDExists]
		c.JSON(resp.StatusCode, resp)
		return
	}

	if success, oid = giftCardDAO.CreateOrder(orderID, serviceFee, createdUserID, transactionID, contractID, deliveryType, additionalFee); !success {
		resp := CodeMessage[CreateOrderFailed]
		c.JSON(resp.StatusCode, resp)
		return
	}

	for _, orderDetail := range orderDetails {
		giftCardDAO.CreateOrderDetail(oid, orderDetail.Amount, orderDetail.Quantity)
	}

	resp := CodeMessage[Success]
	c.JSON(resp.StatusCode, resp)
}

func (api GiftCardApi) CheckCode(c *gin.Context) {
	code := c.PostForm("code")

	if code == "" {
		resp := CodeMessage[InvalidGiftCardCode]
		c.JSON(resp.StatusCode, resp)
		return
	}

	type CodeStatus struct {
		IsRedeemed     bool      `json:"is_redeemed"`
		Amount         float64   `json:"amount"`
		ExpirationDate time.Time `json:"expiration_date"`
	}

	giftCardDAO := dao.GiftCardDAO{}
	isRedeemed, amount, expirationDate := giftCardDAO.CheckCode(code)
	resp := CodeMessage[Success]
	resp.Data = CodeStatus{IsRedeemed: isRedeemed, Amount: amount, ExpirationDate: expirationDate}
	c.JSON(resp.StatusCode, resp)
}
