package dao

import (
	"strconv"

	"github.com/ninjadotorg/handshake-wallet/form"
	"github.com/ninjadotorg/handshake-wallet/service/cache"
)

type OnChainDAO struct {
}

func (dao OnChainDAO) GetGiftCardInitRedeemEventBlock() (t TransferObject) {
	obj := form.OnChainEventBlock{}
	GetCacheObject(GetInitRedeemEndBlockCacheKey(), &t, func(val string) interface{} {
		block, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			block = 0
		}
		obj.LastBlock = uint64(block)

		return obj
	})

	return
}

func (dao OnChainDAO) UpdateGiftCardInitRedeemEventBlock(block form.OnChainEventBlock) {
	key := GetInitRedeemEndBlockCacheKey()
	cache.RedisClient.Set(key, block.LastBlock, 0)
}

func (dao OnChainDAO) GetGiftCardUseRedeemEventBlock() (t TransferObject) {
	obj := form.OnChainEventBlock{}
	GetCacheObject(GetUseRedeemEndBlockCacheKey(), &t, func(val string) interface{} {
		block, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			block = 0
		}
		obj.LastBlock = uint64(block)

		return obj
	})
	return
}

func (dao OnChainDAO) UpdateGiftCardUseRedeemEventBlock(block form.OnChainEventBlock) {
	key := GetUseRedeemEndBlockCacheKey()
	cache.RedisClient.Set(key, block.LastBlock, 0)
}
