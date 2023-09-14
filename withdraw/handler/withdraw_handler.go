package handler

import (
	"context"
	"errors"
	"gangbu/pkg/app"
	"gangbu/pkg/e"
	"gangbu/pkg/models"
	"gangbu/pkg/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"os"
)

type withdrawHandler struct {
	pu models.PlayerUsecase
}

func (wh *withdrawHandler) withdraw(c *gin.Context) {
	a := app.Gin{C: c}
	discordId := c.Param("id")
	receive := c.Param("receive")
	player, err := wh.pu.GetByDiscordUserID(discordId)
	if err != nil {
		util.Logger.Error("获取用户信息失败", err)
		a.Response(http.StatusInternalServerError, e.ERROR, "获取用户信息失败")
		return
	}
	client, err := ethclient.Dial(os.Getenv("ALCHEMY_URL"))
	if err != nil {
		util.Logger.Error("连接到以太坊节点失败!", err)
		a.Response(http.StatusInternalServerError, e.ERROR, "连接ethereum失败")
		return
	}
	defer client.Close()
	if err != nil {
		util.Logger.Error(err)
		return
	}
	value, err := client.BalanceAt(context.Background(), common.HexToAddress(player.WalletAddress), nil)
	if err != nil {
		util.Logger.Error("获取余额失败", err)
		a.Response(http.StatusInternalServerError, e.ErrorEthereumNodeResponse, "获取余额失败")
		return
	}
	tx, err := sendTransaction(client, player.WalletAddress, receive, player.PrivateKey, value)
	if err != nil {
		util.Logger.Error("发送交易失败", err)
		a.Response(http.StatusInternalServerError, e.ErrorEthereumNodeResponse, "发送交易失败")
		return
	}
	a.Response(http.StatusOK, e.SUCCESS, tx.Hash().Hex())
}

func sendTransaction(cl *ethclient.Client, senderAddress string, toAddress string, privateKey string, amount *big.Int) (*types.Transaction, error) {
	var (
		sk       = crypto.ToECDSAUnsafe(common.FromHex(privateKey))
		to       = common.HexToAddress(toAddress)
		sender   = common.HexToAddress(senderAddress)
		gasLimit = uint64(21000)
	)
	// Get suggested gas price
	tipCap, _ := cl.SuggestGasTipCap(context.Background())
	feeCap, _ := cl.SuggestGasPrice(context.Background())
	// 计算燃气费用
	gasCost := new(big.Int).Mul(feeCap, new(big.Int).SetUint64(gasLimit))

	// 确保发送的价值足够支付燃气费用和实际转账金额
	if amount.Cmp(gasCost) < 0 {
		// 价值不足以支付燃气费用
		util.Logger.Error("Insufficient funds for gas cost")
		return nil, errors.New("insufficient funds for gas cost")
	}

	// 减去燃气费用以获取实际转账金额
	actualValue := new(big.Int).Sub(amount, gasCost)

	// Retrieve the chainid (needed for signer)
	chainid, err := cl.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	// Retrieve the pending nonce
	nonce, err := cl.PendingNonceAt(context.Background(), sender)
	if err != nil {
		return nil, err
	}

	// Create a new transaction
	tx := types.NewTx(
		&types.DynamicFeeTx{
			ChainID:   chainid,
			Nonce:     nonce,
			GasTipCap: tipCap,
			GasFeeCap: feeCap,
			Gas:       gasLimit,
			To:        &to,
			Value:     actualValue,
			Data:      nil,
		})
	// Sign the transaction using our keys
	signedTx, _ := types.SignTx(tx, types.NewLondonSigner(chainid), sk)
	// Send the transaction to our node
	err = cl.SendTransaction(context.Background(), signedTx)
	return signedTx, err
}

func NewWithdrawHandler(c *gin.Engine, playerU models.PlayerUsecase) {
	handler := &withdrawHandler{pu: playerU}
	v1 := c.Group("/v1")
	{
		v1.GET("/withdraw/:id/receive/:receive", handler.withdraw)
	}
}
