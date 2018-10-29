package redeemhandshake_service

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ninjadotorg/handshake-wallet/abi"
	"github.com/ninjadotorg/handshake-wallet/config"
	"github.com/ninjadotorg/handshake-wallet/form"
	"github.com/ninjadotorg/handshake-wallet/integration/ethereum_service"
	"github.com/ninjadotorg/handshake-wallet/utils"
	"github.com/shopspring/decimal"
)

var WeiDecimal = decimal.NewFromBigInt(big.NewInt(1000000000000000000), 0)

type RedeemHandshakeClient struct {
	client          *ethclient.Client
	address         common.Address
	redeemHandshake *abi.RedeemHandshake
	writeClient     ethereum_service.EthereumClient
}

func (c *RedeemHandshakeClient) initialize() (err error) {
	conf := config.GetConfig()
	c.client, err = ethclient.Dial(conf.GetString("eth_network"))
	if err != nil {
		return
	}
	c.address = common.HexToAddress(conf.GetString("eth_giftcard_contract_address"))
	c.redeemHandshake, err = abi.NewRedeemHandshake(c.address, c.client)
	if err != nil {
		return
	}

	return
}

func (c *RedeemHandshakeClient) initializeWrite() {
	c.writeClient = ethereum_service.EthereumClient{}
	c.writeClient.Initialize()
}

func (c *RedeemHandshakeClient) initializeWriteWithKey(key string) {
	c.writeClient = ethereum_service.EthereumClient{}
	c.writeClient.InitializeWithKey(key)
}

func (c *RedeemHandshakeClient) close() {
	c.client.Close()
}

func (c *RedeemHandshakeClient) closeWrite() {
	c.writeClient.Close()
}

func (c *RedeemHandshakeClient) GetInitRedeemEvent(startBlock uint64) (orders []form.OnChainInitRedeemBlock, endBlock uint64, err error) {
	c.initialize()

	opt := &bind.FilterOpts{
		Start: startBlock,
	}
	past, errInit := c.redeemHandshake.FilterInitRedeem(opt)
	if errInit != nil {
		err = errInit
		return
	}

	notEmpty := true
	endBlock = startBlock
	for notEmpty {
		notEmpty = past.Next()
		if notEmpty {
			endBlock = past.Event.Raw.BlockNumber
			txHash := past.Event.Raw.TxHash.String()

			orderID := string(bytes.Trim(past.Event.Offchain[:], "\x00"))
			rid := uint(past.Event.Rid.Uint64())
			fee, _ := decimal.NewFromBigInt(past.Event.Fee, 0).Div(WeiDecimal).Float64()
			if orderID != "" {
				orders = append(orders, form.OnChainInitRedeemBlock{
					OrderID:    utils.GetOrderNumber(orderID),
					ContractID: rid,
					Fee:        fee,
					BaseOnChainBlock: form.BaseOnChainBlock{
						BlockNumber: endBlock,
						TxHash:      txHash,
					},
				})
			}
		}
	}
	c.close()

	return
}

func (c *RedeemHandshakeClient) GetUseRedeemEvent(startBlock uint64) (orders []form.OnChainUseRedeemBlock, endBlock uint64, err error) {
	c.initialize()

	opt := &bind.FilterOpts{
		Start: startBlock,
	}
	past, errInit := c.redeemHandshake.FilterUseRedeem(opt)
	if errInit != nil {
		err = errInit
		return
	}

	notEmpty := true
	endBlock = startBlock
	for notEmpty {
		notEmpty = past.Next()
		if notEmpty {
			endBlock = past.Event.Raw.BlockNumber
			txHash := past.Event.Raw.TxHash.String()

			code := string(bytes.Trim(past.Event.Offchain[:], "\x00"))
			rid := uint(past.Event.Rid.Uint64())
			if code != "" {
				orders = append(orders, form.OnChainUseRedeemBlock{
					Code:       code,
					ContractID: rid,
					BaseOnChainBlock: form.BaseOnChainBlock{
						BlockNumber: endBlock,
						TxHash:      txHash,
					},
				})
			}
		}
	}
	c.close()

	return
}

func (c *RedeemHandshakeClient) UseRedeem(giftCardCodeId string, redeemId int, amount decimal.Decimal, address string, key string) (txHash string, err error) {
	c.initialize()
	if key != "" {
		c.initializeWriteWithKey(key)
	} else {
		c.initializeWrite()
	}

	auth, err := c.writeClient.GetAuth(decimal.NewFromFloat(0))

	if err != nil {
		return
	}

	toAddress := common.HexToAddress(address)
	decimalAmount := amount.Sub(amount.Floor())
	intAmount := amount.Sub(decimalAmount)
	weiBigAmount := big.NewInt(WeiDecimal.IntPart())
	intBigAmount := big.NewInt(intAmount.IntPart())
	intWeiAmount := intBigAmount.Mul(intBigAmount, weiBigAmount)
	decimalBigAmount := big.NewInt(decimalAmount.Mul(WeiDecimal).IntPart())
	sendAmount := intWeiAmount.Add(intWeiAmount, decimalBigAmount)
	offerId := big.NewInt(int64(redeemId))

	offChain := [32]byte{}
	copy(offChain[:], []byte(giftCardCodeId))

	tx, err := c.redeemHandshake.UseRedeem(auth, offerId, sendAmount, toAddress, offChain)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()

	c.closeWrite()
	c.close()

	return
}
