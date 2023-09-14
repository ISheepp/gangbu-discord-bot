package util

import (
	"crypto/ecdsa"
	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"os"
)

func GetDiscordClient() *discordgo.Session {
	// Create new Discord Session
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// add proxy if exist
	if os.Getenv("ENV") == "dev" {
		discord.Dialer.Proxy = func(request *http.Request) (*url.URL, error) {
			u, _ := url.Parse("http://127.0.0.1:7890")
			return u, nil
		}
	}
	return discord
}

func StringToPrivateKey(privateKeyStr string) (*ecdsa.PrivateKey, error) {
	privateKeyByte, err := hexutil.Decode(privateKeyStr)
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.ToECDSA(privateKeyByte)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func EtherToWei(etherAmount float64) int64 {
	// 1 ether = 10^18 wei
	weiAmount := new(big.Float).Mul(big.NewFloat(etherAmount), big.NewFloat(1e18))
	weiInt, _ := new(big.Int).SetString(weiAmount.Text('f', 0), 10)
	return weiInt.Int64()
}

func WeiToEther(weiAmount int64) float64 {
	// 1 ether = 10^18 wei
	weiBigInt := big.NewInt(weiAmount)
	etherFloat := new(big.Float).Quo(new(big.Float).SetInt(weiBigInt), new(big.Float).SetInt64(1e18))
	ether, _ := etherFloat.Float64()
	return ether
}
