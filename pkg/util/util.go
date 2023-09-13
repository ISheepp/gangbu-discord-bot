package util

import (
	"crypto/ecdsa"
	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
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

// Timer 函数接受一个函数作为参数，并返回一个新的函数，
// 该新函数会测量传递的函数的运行时间并将其打印到日志中。
func Timer(fn func()) func() {
	return func() {
		startTime := time.Now()
		fn()
		endTime := time.Now()
		elapsed := endTime.Sub(startTime)
		Logger.Infof("Function took %s to run\n", elapsed)
	}
}
