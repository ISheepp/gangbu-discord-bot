package usecase

import (
	"crypto/ecdsa"
	"errors"
	"gangbu/pkg/models"
	"gangbu/pkg/util"
	"github.com/ethereum/go-ethereum/crypto"
	"gorm.io/gorm"
)

type playerUsecase struct {
	playerRepo models.PlayerRepository
}

func (pu *playerUsecase) GetByDiscordUserID(discordUserID string) (*models.Player, error) {
	player, err := pu.playerRepo.GetByDiscordUserID(discordUserID)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func NewPlayerUsecase(pr models.PlayerRepository) models.PlayerUsecase {
	return &playerUsecase{playerRepo: pr}
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
	res := &models.PlayerVo{
		ID:                 player.ID,
		DiscordUserId:      player.DiscordUserId,
		WalletAddress:      player.WalletAddress,
		WalletValue:        player.WalletValue, // todo 转为eth单位
		WithDrawHistoryDto: nil,                // todo 查询提款记录
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
