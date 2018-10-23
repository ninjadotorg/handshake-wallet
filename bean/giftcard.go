package bean

type CreateOrderFormDetailJSON struct {
	Amount   float64 `json:"amount" binding:"required"`
	Quantity uint    `json:"quantity" binding:"required"`
}

type CreateOrderFormJSON struct {
	OrderID       string                      `json:"order_id" binding:"required"`
	OrderDetails  []CreateOrderFormDetailJSON `json:"order_details" binding:"required"`
	CreatedUserID uint                        `json:"user_id" binding:"required"`
	ServiceFee    float64                     `json:"service_fee" binding:"required"`
	AdditionalFee string                      `json:"additional_fee"`
	TransactionID string                      `json:"transaction_id" binding:"required"`
	ContractID    string                      `json:"contract_id" binding:"required"`
	DeliveryType  uint                        `json:"delivery_type" binding:"required"`
}
