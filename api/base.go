package api

import "net/http"

const Success = "Success"
const InvalidRequestBody = "InvalidRequestBody"
const InvalidRequestParam = "InvalidRequestParam"
const InvalidGiftCardCode = "InvalidGiftCardCode"
const OrderIDExists = "OrderIDExists"
const CreateOrderFailed = "CreateOrderFailed"
const UpdateOrderFailed = "UpdateOrderFailed"

var CodeMessage = map[string]struct {
	StatusCode int         `json:"status_code"`
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}{
	Success:             {http.StatusOK, 1, "Success", nil},
	InvalidRequestBody:  {http.StatusBadRequest, -1, "Invalid request body", nil},
	InvalidRequestParam: {http.StatusBadRequest, -2, "Invalid request param", nil},
	OrderIDExists:       {http.StatusBadRequest, -200, "Order ID cannot be used", nil},
	CreateOrderFailed:   {http.StatusBadRequest, -201, "Create order failed", nil},
	InvalidGiftCardCode: {http.StatusBadRequest, -202, "Invalid code", nil},
	UpdateOrderFailed:   {http.StatusBadRequest, -203, "Update order failed", nil},
}
