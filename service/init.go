package service

import "github.com/ninjadotorg/handshake-wallet/dao"

var GiftCardServiceInst = GiftCardService{
	dao: &dao.GiftCardDaoInst,
}
