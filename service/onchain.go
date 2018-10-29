package service

import (
	"fmt"
	"log"

	"github.com/ninjadotorg/handshake-wallet/common"
	"github.com/ninjadotorg/handshake-wallet/dao"
	"github.com/ninjadotorg/handshake-wallet/form"
	"github.com/ninjadotorg/handshake-wallet/integration/redeemhandshake_service"
)

type OnChainService struct {
	dao *dao.OnChainDAO
}

func (service OnChainService) ScanGiftCardOrders() {
	t := service.dao.GetGiftCardInitRedeemEventBlock()
	var block form.OnChainEventBlock
	if common.LogOnChainCronError(0, t.Error) {
		return
	}
	block = t.Object.(form.OnChainEventBlock)
	handshakeClient := redeemhandshake_service.RedeemHandshakeClient{}
	orders, endBlock, err := handshakeClient.GetInitRedeemEvent(block.LastBlock + 1)

	if common.LogOnChainCronError(block.LastBlock, err) {
		return
	}

	for _, order := range orders {
		// get order
		giftCardOrder, err := dao.GiftCardDaoInst.GetOrder(order.OrderID, "")
		if common.LogOnChainCronError(order.BlockNumber, err) {
			continue
		}

		if giftCardOrder.Status != 0 {
			continue
		}

		giftCardOrderDetails, err := dao.GiftCardDaoInst.GetOrderDetails(order.OrderID)
		if common.LogOnChainCronError(order.BlockNumber, err) {
			continue
		}

		for _, giftCardOrderDetail := range giftCardOrderDetails {
			var generatedCodes []string
			var i uint
			for i = 0; i < giftCardOrderDetail.Quantity; i++ {
				code, err := dao.GiftCardDaoInst.CreateCode(giftCardOrder.ID, giftCardOrderDetail.Amount)
				if common.LogOnChainCronError(order.BlockNumber, err) {
					continue
				}
				generatedCodes = append(generatedCodes, code)
			}

			// create order and send email
			log.Print(generatedCodes)
		}

		// update order with rid from contract
		giftCardOrder.Status = 1
		giftCardOrder.ContractID = fmt.Sprint(order.ContractID)
		giftCardOrder.TransactionID = order.TxHash
		dao.GiftCardDaoInst.UpdateOrder(&giftCardOrder)
	}

	dao.UpdateLastestBlockToCache(&block, endBlock, service.dao.UpdateGiftCardInitRedeemEventBlock)
}

func (service OnChainService) ScanUsedCodes() {
	t := service.dao.GetGiftCardUseRedeemEventBlock()
	var block form.OnChainEventBlock
	if common.LogOnChainCronError(0, t.Error) {
		return
	}

	block = t.Object.(form.OnChainEventBlock)
	handshakeClient := redeemhandshake_service.RedeemHandshakeClient{}
	codes, endBlock, err := handshakeClient.GetUseRedeemEvent(block.LastBlock + 1)

	if common.LogOnChainCronError(block.LastBlock, err) {
		return
	}

	for _, code := range codes {
		codeModel, err := dao.GiftCardDaoInst.GetCodeWithoutEncrypt(code.Code)
		if common.LogOnChainCronError(code.BlockNumber, err) {
			continue
		}

		if codeModel.Status != 2 {
			continue
		}

		// update code status & send email
		codeModel.Status = 1
		codeModel.TransactionHash = code.TxHash
		if err := dao.GiftCardDaoInst.UpdateCode(&codeModel); common.LogOnChainCronError(code.BlockNumber, err) {
			continue
		}
	}

	dao.UpdateLastestBlockToCache(&block, endBlock, service.dao.UpdateGiftCardUseRedeemEventBlock)
}
