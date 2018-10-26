package ethereum_service

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ninjadotorg/handshake-wallet/config"
	"github.com/shopspring/decimal"
)

var WeiDecimal = decimal.NewFromBigInt(big.NewInt(1000000000000000000), 0)

type EthereumClient struct {
	client     *ethclient.Client
	address    common.Address
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func GetAddress() string {
	conf := config.GetConfig()
	privateKey, _ := crypto.HexToECDSA(conf.GetString("eth_private_key"))
	tmpPublicKey := privateKey.Public()
	publicKey := tmpPublicKey.(*ecdsa.PublicKey)

	return crypto.PubkeyToAddress(*publicKey).Hex()
}

func (c *EthereumClient) Initialize() (err error) {
	conf := config.GetConfig()
	c.client, err = ethclient.Dial(conf.GetString("eth_network"))
	if err != nil {
		return
	}

	c.privateKey, err = crypto.HexToECDSA(conf.GetString("eth_private_key"))
	if err != nil {
		return
	}

	publicKey := c.privateKey.Public()
	c.publicKey, _ = publicKey.(*ecdsa.PublicKey)

	c.address = crypto.PubkeyToAddress(*c.publicKey)

	return
}

func (c *EthereumClient) InitializeWithKey(key string) (err error) {
	conf := config.GetConfig()
	c.client, err = ethclient.Dial(conf.GetString("eth_network"))
	if err != nil {
		return
	}

	c.privateKey, err = crypto.HexToECDSA(key)
	if err != nil {
		return
	}

	publicKey := c.privateKey.Public()
	c.publicKey, _ = publicKey.(*ecdsa.PublicKey)

	c.address = crypto.PubkeyToAddress(*c.publicKey)

	return
}

func (c *EthereumClient) GetAddress() string {
	return c.address.Hex()
}

func (c *EthereumClient) Close() {
	c.client.Close()
}

func (c *EthereumClient) SendTransaction(address string, amount decimal.Decimal) (string, error) {
	c.Initialize()
	nonce, err := c.client.PendingNonceAt(context.Background(), c.address)
	if err == nil {
		bigAmount := big.NewInt(amount.IntPart())
		bigWei := big.NewInt(WeiDecimal.IntPart())
		value := bigAmount.Mul(bigAmount, bigWei) // in wei
		gasLimit := uint64(21000)
		gasPrice, err := c.client.SuggestGasPrice(context.Background())
		if err == nil {
			toAddress := common.HexToAddress(address)
			tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

			signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, c.privateKey)
			if err == nil {
				err = c.client.SendTransaction(context.Background(), signedTx)
				if err == nil {
					return signedTx.Hash().Hex(), nil
				}
			}
		}
	}

	c.Close()

	return "", err
}

func (c *EthereumClient) GetBalance() (balance decimal.Decimal, err error) {
	c.Initialize()

	intBalance, errCli := c.client.BalanceAt(context.Background(), c.address, nil)

	if errCli == nil {
		balance = decimal.NewFromBigInt(intBalance, 0).Div(WeiDecimal)
	}

	err = errCli
	c.Close()

	return
}

func (c *EthereumClient) GetNonce() (nonce uint64, err error) {
	nonce, err = c.client.PendingNonceAt(context.Background(), c.address)
	if err != nil {
		return
	}
	return
}

func (c *EthereumClient) GetAuth(amount decimal.Decimal) (auth *bind.TransactOpts, err error) {
	nonce, err := c.client.PendingNonceAt(context.Background(), c.address)
	if err != nil {
		return
	}

	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	// Increase gas price to 20%
	increase20 := big.NewInt(decimal.NewFromBigInt(gasPrice, 0).Mul(decimal.NewFromFloat(0.2)).IntPart())
	gasPrice = gasPrice.Add(gasPrice, increase20)

	value := big.NewInt(amount.Mul(WeiDecimal).IntPart()) // in wei
	auth = bind.NewKeyedTransactor(c.privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value              // in wei
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = gasPrice

	return
}

func (c *EthereumClient) GetTransactionReceipt(txHash string) (status bool, isPending bool, amount decimal.Decimal, err error) {
	c.Initialize()

	var tx *types.Transaction
	tx, isPending, err = c.client.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err == nil {
		if !isPending {
			txReceipt, err1 := c.client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
			err = err1
			if err == nil {
				status = txReceipt.Status == 1
				amount = decimal.NewFromBigInt(tx.Value(), 0).Div(WeiDecimal)
			}
		}
	}

	c.Close()

	return
}

func GenerateAddress() (string, string) {
	// Create an account
	key, _ := crypto.GenerateKey()

	// Get the address
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	// 0x8ee3333cDE801ceE9471ADf23370c48b011f82a6

	// Get the private key
	privateKey := hex.EncodeToString(key.D.Bytes())
	// 05b14254a1d0c77a49eae3bdf080f926a2df17d8e2ebdf7af941ea001481e57f

	return address, privateKey
}
