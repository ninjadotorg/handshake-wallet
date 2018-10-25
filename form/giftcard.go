package form

type GiftCardOrderFormDetail struct {
	Amount          float64                    `json:"amount" validate:"required,numeric"`
	Quantity        uint                       `json:"quantity" validate:"required,numeric"`
	PersonalMessage string                     `json:"personal_message" validate:"omitempty"`
	ShippingDetail  *GiftCardOrderFormShipping `json:"shipping_detail" validate:"required"`
}

type GiftCardOrderFormShipping struct {
	Name    string `json:"name" validate:"omitempty"`
	Address string `json:"address" validate:"omitempty"`
	City    string `json:"city" validate:"omitempty"`
	State   string `json:"state" validate:"omitempty"`
	Country string `json:"country" validate:"omitempty"`
	Zip     string `json:"zip" validate:"omitempty"`
	Phone   string `json:"phone" validate:"omitempty"`
	Email   string `json:"email" validate:"required,email"`
}

type GiftCardCreateOrderForm struct {
	OrderDetails  []GiftCardOrderFormDetail `json:"order_details" validate:"required,gt=0"`
	ServiceFee    float64                   `json:"service_fee" validate:"required,numeric"`
	AdditionalFee string                    `json:"additional_fee" validate:"omitempty"`
	Name          string                    `json:"name" validate:"required"`
	Email         string                    `json:"email" validate:"required,email"`
	ShippingType  uint                      `json:"shipping_type" validate:"required,oneof=1 2"`
	EthAddress    string                    `json:"eth_address" validate:"required,eth_addr"`
}

type GiftCardUpdateOrderForm struct {
	OrderID       string `json:"order_id" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	TransactionID string `json:"transaction_id" validate:"required"`
}

type GiftCardCheckCodeForm struct {
	Code string `json:"code" validate:"required"`
}
