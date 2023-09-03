package api

import (
	"crypto/ecdsa"
	"errors"
	"gangbu/pkg/dao"
	"gangbu/pkg/e"
	"gangbu/pkg/models"
	"gangbu/pkg/util"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func ShowPlayerInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		discordUserId := c.Param("id")
		db := dao.NewDBClient(c)
		var player = models.Player{DiscordUserId: discordUserId}
		tx := db.First(&player)
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			// create and show
			// 根据椭圆曲线生成私钥、公钥、地址
			player = models.Player{
				DiscordUserId: discordUserId,
				WalletAddress: "",
			}
			c.JSON(http.StatusOK, &e.ResponseData{
				Message: "player created successfully!",
				Status:  true,
				Data:    nil,
			})
		}
		// show
		c.JSON(http.StatusOK, &e.ResponseData{
			Message: "player exist",
			Status:  true,
			Data:    player,
		})

	}
}

func generatePkAndAddress() (string, string) {
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

	util.Logger.Info("生成的私钥:", privateKey.D.Text(16))
	util.Logger.Info("生成的地址:", address)
	return "", ""
}
