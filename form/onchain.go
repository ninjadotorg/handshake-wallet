package form

type OnChainEventBlock struct {
	LastBlock uint64
}

type BaseOnChainBlock struct {
	BlockNumber uint64
	TxHash      string
}

type OnChainInitRedeemBlock struct {
	OrderID    uint
	ContractID uint
	Fee        float64
	BaseOnChainBlock
}

type OnChainUseRedeemBlock struct {
	ContractID uint
	Code       string
	BaseOnChainBlock
}
