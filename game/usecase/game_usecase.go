package usecase

import (
	"context"
	"errors"
	"gangbu/pkg/e"
	"gangbu/pkg/models"
	"gangbu/pkg/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"math/big"
	"os"
	"strconv"
)

type gameHistoryUsecase struct {
	gameRepo   models.GameHistoryRepository
	playerRepo models.PlayerRepository
	db         *gorm.DB
}

func (ghu *gameHistoryUsecase) CreateGame(bo models.GameHistoryBo) error {
	tx := ghu.db.Begin()
	err := ghu.gameRepo.CreateGame(&models.GameHistory{
		PlayerDiscordUserId: bo.PlayerDiscordUserId,
		Choice:              int(bo.Choice),
		GameStatus:          e.IN_PROGRESS,
		BetValue:            bo.BetValue,
		ServerId:            bo.ServerId,
		FinishTime:          nil,
	}, tx)
	if err != nil {
		util.Logger.Error("创建游戏失败!", err)
		tx.Rollback()
		return err
	}
	// 调用合约函数，拿到requestId
	// 连接到以太坊节点
	client, err := ethclient.Dial(os.Getenv("ALCHEMY_URL"))
	if err != nil {
		util.Logger.Error("连接到以太坊节点失败!", err)
		tx.Rollback()
		return err
	}
	defer client.Close()

	user, err := ghu.playerRepo.GetByDiscordUserID(bo.PlayerDiscordUserId)
	if err != nil {
		util.Logger.Error("查询用户信息失败!", err)
		tx.Rollback()
		return err
	}

	// 检查用户游戏内地址的余额是否大于筹码
	// todo 数据库里的数据如何实时更新？？
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(user.WalletAddress), nil)
	if err != nil {
		util.Logger.Error("查询用户余额失败!", err)
		tx.Rollback()
		return err
	}
	if balance.Cmp(big.NewInt(bo.BetValue)) < 0 {
		util.Logger.Error("用户余额不足!")
		tx.Rollback()
		return errors.New("用户余额不足: 当前余额：" + balance.String())
	}

	// 合约地址
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	// 创建合约对象
	game, err := models.NewEvenOddGame(contractAddress, client)
	if err != nil {
		util.Logger.Error("创建合约对象失败!", err)
		tx.Rollback()
		return err
	}
	// 签名对象
	chainId, _ := strconv.Atoi(os.Getenv("CHAIN_ID"))
	// 创建私钥（用于签名交易）
	privateKey, err := crypto.HexToECDSA(user.PrivateKey)
	if err != nil {
		util.Logger.Error("创建私钥失败!", err)
		return err
	}
	// 创建一个session
	// session := &models.EvenOddGameSession{
	// 	Contract: game,
	// 	CallOpts: bind.CallOpts{
	// 		Pending: true,
	// 	},
	// 	TransactOpts: bind.TransactOpts{},
	// }
	// util.StringToPrivateKey(user.PrivateKey)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(chainId)))
	if err != nil {
		util.Logger.Error("创建签名对象失败!", err)
		tx.Rollback()
		return err
	}
	// 筹码
	auth.Value = big.NewInt(bo.BetValue)
	blockTx, err := game.RequestRandomWords(auth, bo.Choice)
	if err != nil {
		util.Logger.Error("创建tx交易失败!", err)
		tx.Rollback()
		return err
	}
	// 等待交易, 打印hash地址
	util.Logger.Info("交易hash:", blockTx.Hash().Hex())
	txReceipt, err := bind.WaitMined(context.Background(), client, blockTx)
	if err != nil {
		util.Logger.Error("等待交易失败!", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	util.Logger.Info("请求随机数完成!", txReceipt)
	return nil
}

func NewGameUsecase(gameRepo models.GameHistoryRepository, playerRepo models.PlayerRepository, db *gorm.DB) models.GameHistoryUsecase {
	return &gameHistoryUsecase{
		gameRepo:   gameRepo,
		playerRepo: playerRepo,
		db:         db,
	}
}
