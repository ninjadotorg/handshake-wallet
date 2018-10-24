package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/handshake-wallet/bean"
	"github.com/ninjadotorg/handshake-wallet/dao"
	"github.com/ninjadotorg/handshake-wallet/utils"
)

type GiftCardApi struct{}

func (api GiftCardApi) CreateOrder(c *gin.Context) {
	var orderForm bean.GiftCardOrderForm
	if err := c.ShouldBindJSON(&orderForm); err != nil {
		resp := CodeMessage[InvalidRequestBody]
		resp.Message = "Invalid order data"
		c.JSON(resp.StatusCode, resp)
		return
	}

	orderDetails := orderForm.OrderDetails
	name := orderForm.Name
	email := orderForm.Email
	shippingType := orderForm.ShippingType
	serviceFee := orderForm.ServiceFee
	additionalFee := orderForm.AdditionalFee
	ethAddress := orderForm.EthAddress

	var errors []string

	if name == "" {
		errors = append(errors, "Name cannot be empty")
	}

	if email == "" {
		errors = append(errors, "Email cannot be empty")
	}

	if len(orderDetails) < 1 {
		errors = append(errors, "Order details cannot be empty")
	}

	if serviceFee < 0 {
		errors = append(errors, "Invalid service fee")
	}

	if ethAddress == "" {
		errors = append(errors, "Invalid eth address")
	}

	if shippingType < 1 || shippingType > 2 {
		errors = append(errors, "Invalid shipping type")
	} else if len(orderDetails) > 0 {
		for i, orderDetail := range orderDetails {
			api.ValidateOrderShippingDetail(&errors, shippingType, orderDetail.ShippingDetail, i)
		}
	}

	if len(errors) > 0 {
		resp := CodeMessage[InvalidRequestBody]
		resp.Data = errors
		c.JSON(resp.StatusCode, resp)
		return
	}

	giftCardDAO := dao.GiftCardDAO{}
	var orderNumber uint
	var success bool

	if success, orderNumber = giftCardDAO.CreateOrder(name, email, serviceFee, shippingType, ethAddress, additionalFee); !success {
		resp := CodeMessage[CreateOrderFailed]
		c.JSON(resp.StatusCode, resp)
		return
	}

	for _, orderDetail := range orderDetails {
		giftCardDAO.CreateOrderDetail(orderNumber, orderDetail.ShippingDetail, orderDetail.Amount, orderDetail.Quantity)
	}

	orderID := utils.CreateOrderID(orderNumber)

	if orderID == "" {
		resp := CodeMessage[CreateOrderFailed]
		resp.Message = "Cannot create order ID"
		c.JSON(resp.StatusCode, resp)
		return
	}

	resp := CodeMessage[Success]
	resp.Data = orderID
	c.JSON(resp.StatusCode, resp)
}

func (api GiftCardApi) UpdateOrder(c *gin.Context) {
	var orderForm bean.GiftCardOrderForm
	if err := c.ShouldBindJSON(&orderForm); err != nil {
		resp := CodeMessage[InvalidRequestBody]
		resp.Message = "Invalid order data"
		c.JSON(resp.StatusCode, resp)
		return
	}

	orderID := orderForm.OrderID
	email := orderForm.Email
	transactionID := orderForm.TransactionID
	contractID := orderForm.ContractID

	orderNumber := -1
	var errors []string

	if orderID == "" {
		errors = append(errors, "Invalid order ID")
	} else if orderNumber = utils.GetOrderNumber(orderID); orderNumber < 0 {
		errors = append(errors, "Invalid order ID")
	}

	if email == "" {
		errors = append(errors, "Invalid email address")
	}

	if transactionID == "" {
		errors = append(errors, "Invalid transaction ID")
	}

	if contractID == "" {
		errors = append(errors, "Invalid contract ID")
	}

	if len(errors) > 0 {
		resp := CodeMessage[InvalidRequestBody]
		resp.Data = errors
		c.JSON(resp.StatusCode, resp)
		return
	}

	giftCardDAO := dao.GiftCardDAO{}
	resp := CodeMessage[Success]
	if success := giftCardDAO.UpdateOrder(uint(orderNumber), email, transactionID, contractID); !success {
		resp = CodeMessage[UpdateOrderFailed]
	}

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

func (api GiftCardApi) ValidateOrderShippingDetail(errors *[]string, shippingType uint, shippingDetail *bean.GiftCardOrderFormShipping, position int) {
	if shippingDetail == nil {
		*errors = append(*errors, fmt.Sprintf("Shipping detail cannot be empty for order [%d]", position))
		return
	}
	switch shippingType {
	case 1:
		if shippingDetail.Name == "" {
			*errors = append(*errors, fmt.Sprintf("Shipping name cannot be empty for order [%d]", position))
		}
		if shippingDetail.Email == "" {
			*errors = append(*errors, fmt.Sprintf("Shipping email cannot be empty for order [%d]", position))
		}
		break
	case 2:
		if shippingDetail.Name == "" {
			*errors = append(*errors, fmt.Sprintf("Shipping name cannot be empty for order [%d]", position))
		}
		if shippingDetail.Email == "" {
			*errors = append(*errors, fmt.Sprintf("Shipping email cannot be empty for order [%d]", position))
		}
		if shippingDetail.Address == "" {
			*errors = append(*errors, fmt.Sprintf("Shipping address cannot be empty for order [%d]", position))
		}
		if shippingDetail.City == "" {
			*errors = append(*errors, fmt.Sprintf("Shipping city cannot be empty for order [%d]", position))
		}
		if shippingDetail.State == "" {
			*errors = append(*errors, fmt.Sprintf("Shipping state cannot be empty for order [%d]", position))
		}
		if shippingDetail.Zip == "" {
			*errors = append(*errors, fmt.Sprintf("Shipping zip cannot be empty for order [%d]", position))
		}
		if shippingDetail.Country == "" {
			*errors = append(*errors, fmt.Sprintf("Shipping country cannot be empty for order [%d]", position))
		}
		if shippingDetail.Phone == "" {
			*errors = append(*errors, fmt.Sprintf("Shipping phone cannot be empty for order [%d]", position))
		}
		break
	}
}
