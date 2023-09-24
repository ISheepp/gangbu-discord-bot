package util

import (
	"crypto/ecdsa"
	"fmt"
	"gangbu/proto"
	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"os"
	"sync"
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

func StartGrpcServer(wg *sync.WaitGroup, srv proto.GameRequestServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// 注册服务
	proto.RegisterGameRequestServer(s, srv)
	// 启动RPC并监听
	log.Printf("server listening at %v", lis.Addr())
	wg.Done()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func GetGrpcClient(addr string) (proto.GameRequestClient, error) {
	// grpc client
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// todo 记得要关闭conn
	if err != nil {
		log.Fatalf("连接服务器失败: %v", err)
	}
	log.Printf("建立连接成功: %s", addr)

	// 创建客户端
	client := proto.NewGameRequestClient(conn)
	return client, err
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
