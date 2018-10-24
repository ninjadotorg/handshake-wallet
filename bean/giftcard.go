package bean

type GiftCardOrderFormDetail struct {
	Amount         float64                    `json:"amount" binding:"required"`
	Quantity       uint                       `json:"quantity" binding:"required"`
	ShippingDetail *GiftCardOrderFormShipping `json:"shipping_detail" binding:"required"`
}

type GiftCardOrderFormShipping struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Zip     string `json:"zip"`
	Phone   string `json:"phone"`
	Email   string `json:"email" binding:"required"`
}

type GiftCardOrderForm struct {
	OrderID       string                    `json:"order_id"`
	OrderDetails  []GiftCardOrderFormDetail `json:"order_details"`
	ServiceFee    float64                   `json:"service_fee"`
	AdditionalFee string                    `json:"additional_fee"`
	TransactionID string                    `json:"transaction_id"`
	ContractID    string                    `json:"contract_id"`
	Name          string                    `json:"name"`
	Email         string                    `json:"email"`
	ShippingType  uint                      `json:"shipping_type"`
	EthAddress    string                    `json:"eth_address"`
}
