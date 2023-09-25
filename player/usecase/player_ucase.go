package usecase

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"gangbu/pkg/models"
	"gangbu/pkg/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"math/big"
	"os"
	"time"
)

type playerUsecase struct {
	playerRepo models.PlayerRepository
}

func NewPlayerUsecase(pr models.PlayerRepository) models.PlayerUsecase {
	return &playerUsecase{playerRepo: pr}
}

func (pu *playerUsecase) GetByDiscordUserID(discordUserID string) (*models.Player, error) {
	player, err := pu.playerRepo.GetByDiscordUserID(discordUserID)
	if err != nil {
		util.Logger.Error("获取玩家失败!", err)
		return nil, err
	}
	walletValue, err := getWalletValueFromChain(player)
	if err != nil {
		util.Logger.Error("获取钱包余额失败!", err)
		return nil, err
	}
	player.WalletValue = walletValue.Int64()
	return player, nil
}

func (pu *playerUsecase) GetByDiscordUserIDOrCreate(discordUserID string) (*models.PlayerVo, error) {
	player, err := pu.playerRepo.GetByDiscordUserID(discordUserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = pu.CreatePlayer(models.PlayerCreateBo{DiscordUserId: discordUserID})
			if err != nil {
				util.Logger.Error("创建玩家失败！")
				return nil, err
			}
			player, err = pu.playerRepo.GetByDiscordUserID(discordUserID)
		} else {
			return nil, err
		}
	}
	// 查询一下链上的余额
	walletValue, err := getWalletValueFromChain(player)
	if err != nil {
		return nil, err
	}
	// 更新数据库记录
	err = pu.playerRepo.UpdateWalletValue(player.ID, walletValue.Int64())
	res := &models.PlayerVo{
		ID:                 player.ID,
		DiscordUserId:      player.DiscordUserId,
		WalletAddress:      player.WalletAddress,
		WalletValue:        walletValue.Int64(), // todo 转为eth单位
		WithDrawHistoryDto: nil,                 // todo 查询提款记录
		CreateAt:           player.CreatedAt,
	}
	return res, nil
}

func (pu *playerUsecase) CreatePlayer(bo models.PlayerCreateBo) error {
	// 生成随机私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		util.Logger.Error("generate private key failed!", err)
	}

	// 获取公钥
	publicKey := privateKey.Public()
	// 转换公钥为ecdsa.PublicKey类型
	ecdsaPublicKey := publicKey.(*ecdsa.PublicKey)

	// 获取eth地址
	address := crypto.PubkeyToAddress(*ecdsaPublicKey).Hex()

	// util.Logger.Info("生成的私钥:", privateKey.D.Text(16))
	util.Logger.Info("生成的地址:", address)
	player := &models.Player{
		DiscordUserId: bo.DiscordUserId,
		WalletAddress: address,
		PrivateKey:    privateKey.D.Text(16),
		WalletValue:   0,
	}
	if e := pu.playerRepo.CreatePlayer(player); e != nil {
		return e
	}
	return nil
}

func getWalletValueFromChain(player *models.Player) (*big.Int, error) {
	start := time.Now()
	eth, err := ethclient.Dial(os.Getenv("ALCHEMY_URL"))
	if err != nil {
		util.Logger.Error("连接到以太坊节点失败!", err)
		return nil, err
	}
	defer eth.Close()
	// 查询地址的余额
	balanceAt, err := eth.BalanceAt(context.Background(), common.HexToAddress(player.WalletAddress), nil)
	if err != nil {
		util.Logger.Error("查询余额失败!", err)
		return nil, err
	}
	end := time.Now()
	util.Logger.Info("查询余额耗时:", end.Sub(start))
	return balanceAt, nil
}
