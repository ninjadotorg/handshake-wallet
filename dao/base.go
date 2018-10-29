package dao

import (
	"errors"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/ninjadotorg/handshake-wallet/api_response"
	"github.com/ninjadotorg/handshake-wallet/common"
	"github.com/ninjadotorg/handshake-wallet/form"
	"github.com/ninjadotorg/handshake-wallet/service/cache"
)

type TransferObject struct {
	Object    interface{}
	Objects   []interface{}
	Page      interface{}
	CanMove   bool
	Found     bool
	StatusKey string
	Error     error
}

func (t *TransferObject) SetError(statusKey string, err error) bool {
	// Only set to error and status key if there is really error
	if err != nil {
		t.StatusKey = statusKey
		t.Error = err
		return true
	}

	return false
}

func (t *TransferObject) SetStatusKey(statusKey string) {
	t.SetError(statusKey, errors.New(statusKey))
}

func (t *TransferObject) HasError() bool {
	if !t.Found || t.StatusKey != "" || t.Error != nil {
		return true
	}
	return false
}

func GetDB() *gorm.DB {
	return common.Database()
}

func GetCacheObject(key string, t *TransferObject, f func(string) interface{}) {
	val, err := cache.RedisClient.Get(key).Result()
	if err == nil {
		t.Object = f(val)
		t.Found = true
	} else {
		if err != redis.Nil {
			t.SetError(api_response.GetDataFailed, err)
		}
	}
}

func UpdateLastestBlockToCache(block *form.OnChainEventBlock, lastBlock uint64, f func(form.OnChainEventBlock)) {
	if block.LastBlock != lastBlock {
		block.LastBlock = lastBlock
		f(*block)
	}
}
