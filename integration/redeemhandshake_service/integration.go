package redeemhandshake_service

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ninjadotorg/handshake-wallet/abi"
	"github.com/ninjadotorg/handshake-wallet/config"
	"github.com/ninjadotorg/handshake-wallet/integration/ethereum_service"
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
