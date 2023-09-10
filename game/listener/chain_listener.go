package listener

import (
	"context"
	"gangbu/pkg/models"
	"gangbu/pkg/util"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

func ChainListener() {
	client, err := ethclient.Dial(os.Getenv("ALCHEMY_WEBSOCKET_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 解析合约地址
	contractAddr := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	// 创建一个过滤器以监听事件
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
		Topics:    [][]common.Hash{{}},
	}
	logs := make(chan types.Log)
	_, err = client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	filterer, err := models.NewEvenOddGameFilterer(common.HexToAddress(os.Getenv("CONTRACT_ADDRESS")), client)
	if err != nil {
		log.Fatal(err)
	}
	channel := make(chan *models.EvenOddGameRequestSent)
	sub, err := filterer.WatchRequestSent(&bind.WatchOpts{
		Start:   nil,
		Context: nil,
	}, channel)
	if err != nil {
		log.Fatal(err)
	}
	util.Logger.Info("开始监听事件")
	for {
		select {
		case err := <-sub.Err():
			util.Logger.Error(err)
		case event := <-channel:
			util.Logger.Info(event)
		}
	}

}
