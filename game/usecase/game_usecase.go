package usecase

import (
	"context"
	"errors"
	"gangbu/pkg/e"
	"gangbu/pkg/models"
	"gangbu/pkg/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

func (ghu *gameHistoryUsecase) GetLastFiveGameHistoryByDiscordId(discordId string) ([]models.GameHistory, error) {
	game, err := ghu.gameRepo.GetLastFiveGameHistoryByDiscordId(discordId, ghu.db)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (ghu *gameHistoryUsecase) GetGameHistoryByDiscordId(discordId string) ([]models.GameHistory, error) {
	game, err := ghu.gameRepo.GetGameHistoryByDiscordId(discordId, ghu.db)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (ghu *gameHistoryUsecase) UpdateGameAfterMainBet(game *models.GameHistory) error {
	err := ghu.gameRepo.UpdateGameAfterMainBet(game, ghu.db)
	if err != nil {
		return err
	}
	return nil
}

func (ghu *gameHistoryUsecase) GetGameHistoryByRequestId(requestId string) (*models.GameHistory, error) {
	game, err := ghu.gameRepo.GetGameHistoryByRequestId(requestId, ghu.db)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (ghu *gameHistoryUsecase) UpdateRequestIdByTxId(txId string, requestID string) error {
	err := ghu.gameRepo.UpdateRequestIdByTxId(txId, requestID, ghu.db)
	if err != nil {
		return err
	}
	return nil
}

func (ghu *gameHistoryUsecase) CreateGame(bo models.GameHistoryBo) (*types.Transaction, error) {
	tx := ghu.db.Begin()
	// 调用合约函数，拿到requestId
	// 连接到以太坊节点
	client, err := ethclient.Dial(os.Getenv("ALCHEMY_URL"))
	if err != nil {
		util.Logger.Error("连接到以太坊节点失败!", err)
		tx.Rollback()
		return nil, err
	}
	defer client.Close()

	user, err := ghu.playerRepo.GetByDiscordUserID(bo.PlayerDiscordUserId)
	if err != nil {
		util.Logger.Error("查询用户信息失败!", err)
		tx.Rollback()
		return nil, err
	}

	// 检查用户游戏内地址的余额是否大于筹码
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(user.WalletAddress), nil)
	if err != nil {
		util.Logger.Error("查询用户余额失败!", err)
		tx.Rollback()
		return nil, err
	}
	if balance.Cmp(big.NewInt(bo.BetValue)) < 0 {
		util.Logger.Error("用户余额不足!")
		tx.Rollback()
		return nil, errors.New("用户余额不足: 当前余额：" + balance.String())
	}

	// 合约地址
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	// 创建合约对象
	game, err := models.NewEvenOddGame(contractAddress, client)
	if err != nil {
		util.Logger.Error("创建合约对象失败!", err)
		tx.Rollback()
		return nil, err
	}
	// bet value 不能大于合约的余额
	contractBalance, err := client.BalanceAt(context.Background(), contractAddress, nil)
	if bo.BetValue > contractBalance.Int64() {
		util.Logger.Error("合约余额不足!")
		tx.Rollback()
		return nil, errors.New("合约余额不足: 余额:" + contractBalance.String() + "; 用户筹码:" + strconv.Itoa(int(bo.BetValue)))
	}

	// 签名对象
	chainId, _ := strconv.Atoi(os.Getenv("CHAIN_ID"))
	// 创建私钥（用于签名交易）
	privateKey, err := crypto.HexToECDSA(user.PrivateKey)
	if err != nil {
		util.Logger.Error("创建私钥失败!", err)
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(chainId)))
	if err != nil {
		util.Logger.Error("创建签名对象失败!", err)
		tx.Rollback()
		return nil, err
	}
	// 筹码
	auth.Value = big.NewInt(bo.BetValue)
	blockTx, err := game.RequestRandomWords(auth, bo.Choice)
	if err != nil {
		util.Logger.Error("创建tx交易失败!", err)
		tx.Rollback()
		return nil, err
	}
	// 等待交易, 打印hash地址
	txId := blockTx.Hash().Hex()
	util.Logger.Info("交易hash:", txId)
	// 保存此次交易信息
	err = ghu.gameRepo.CreateGame(&models.GameHistory{
		PlayerDiscordUserId: bo.PlayerDiscordUserId,
		Choice:              int(bo.Choice),
		GameStatus:          e.IN_PROGRESS,
		BetValue:            bo.BetValue,
		GuildID:             bo.GuildID,
		ChannelID:           bo.ChannelID,
		FinishTime:          nil,
		RequestRandomTxId:   txId,
	}, tx)
	if err != nil {
		util.Logger.Error("创建游戏失败!", err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return blockTx, nil
}

func NewGameUsecase(gameRepo models.GameHistoryRepository, playerRepo models.PlayerRepository, db *gorm.DB) models.GameHistoryUsecase {
	return &gameHistoryUsecase{
		gameRepo:   gameRepo,
		playerRepo: playerRepo,
		db:         db,
	}
}
