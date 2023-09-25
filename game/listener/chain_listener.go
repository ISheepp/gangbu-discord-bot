package listener

import (
	"context"
	"encoding/json"
	"gangbu/pkg/e"
	"gangbu/pkg/models"
	"gangbu/pkg/queue"
	"gangbu/pkg/util"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/segmentio/kafka-go"
	"math/big"
	"os"
	"strconv"
	"time"
)

type chainListener struct {
	gameHistoryUCase models.GameHistoryUsecase
	playerUCase      models.PlayerUsecase
}

func NewChainListener(ghu models.GameHistoryUsecase, pu models.PlayerUsecase) Listener {
	return &chainListener{gameHistoryUCase: ghu, playerUCase: pu}
}

func (cl *chainListener) StartListen() {
	go cl.RequestRandomListener()
	go cl.ChainLinkGenerateRandomListener()
}

// RequestRandomListener 监听RequestSent事件
func (cl *chainListener) RequestRandomListener() {
	client, err := ethclient.Dial(os.Getenv("ALCHEMY_WEBSOCKET_URL"))
	if err != nil {
		util.Logger.Fatal(err)
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
		util.Logger.Fatal(err)
	}

	filterer, err := models.NewEvenOddGameFilterer(common.HexToAddress(os.Getenv("CONTRACT_ADDRESS")), client)
	if err != nil {
		util.Logger.Fatal(err)
	}
	channel := make(chan *models.EvenOddGameRequestSent)
	sub, err := filterer.WatchRequestSent(&bind.WatchOpts{
		Start:   nil,
		Context: nil,
	}, channel)
	if err != nil {
		util.Logger.Fatal(err)
	}
	util.Logger.Info("开始监听主合约事件")
	for {
		select {
		case err = <-sub.Err():
			util.Logger.Error("订阅主合约失败", err)
			time.Sleep(5 * time.Second)
			// 重试创建连接
			sub, err = filterer.WatchRequestSent(&bind.WatchOpts{
				Start:   nil,
				Context: nil,
			}, channel)
			if err != nil {
				util.Logger.Error("重试订阅主合约失败", err)
			} else {
				util.Logger.Info("重试监听主合约事件成功")
			}
		case event := <-channel:
			util.Logger.Infof("接收到请求随机数完成事件RequestId:%+v", event.RequestId)
			requestId := event.RequestId.String()
			txId := event.Raw.TxHash.Hex()
			err = cl.gameHistoryUCase.UpdateRequestIdByTxId(txId, requestId)
			if err != nil {
				util.Logger.Error("更新requestId失败", err)
				continue
			}

		}
	}

}

// ChainLinkGenerateRandomListener 监听chain link生成随机数的事件
func (cl *chainListener) ChainLinkGenerateRandomListener() {
	client, err := ethclient.Dial(os.Getenv("ALCHEMY_WEBSOCKET_URL"))
	if err != nil {
		util.Logger.Fatal(err)
	}
	defer client.Close()

	// 解析合约地址
	contractAddr := common.HexToAddress(os.Getenv("CHAIN_LINK_VRF_ADDRESS"))

	// 创建一个过滤器以监听事件
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
		Topics:    [][]common.Hash{},
	}
	logs := make(chan types.Log)
	_, err = client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		util.Logger.Fatal(err)
	}

	filterer, err := models.NewVRFCoordinatorV2Filterer(common.HexToAddress(os.Getenv("CONTRACT_ADDRESS")), client)
	if err != nil {
		util.Logger.Fatal(err)
	}
	channel := make(chan *models.VRFCoordinatorV2RequestFulfilled)
	sub, err := filterer.WatchRequestFulfilled(&bind.WatchOpts{
		Start:   nil,
		Context: nil,
	}, channel)
	if err != nil {
		util.Logger.Fatal(err)
	}
	util.Logger.Info("开始监听VRF事件")
	for {
		select {
		case err = <-sub.Err():
			util.Logger.Error("订阅VRF失败", err)
			time.Sleep(5 * time.Second)
			sub, err = filterer.WatchRequestFulfilled(&bind.WatchOpts{
				Start:   nil,
				Context: nil,
			}, channel)
			if err != nil {
				util.Logger.Error("重试订阅VRF失败", err)
			} else {
				util.Logger.Info("重试监听VRF事件成功")
			}
		case event := <-channel:
			util.Logger.Infof("接收到VRF生成随机数事件:%+v", event.RandomWords)
			go invokeMainBet(event, cl)
		}
	}
}

func invokeMainBet(event *models.VRFCoordinatorV2RequestFulfilled, cl *chainListener) {
	requestId := event.RequestId
	game, err := cl.gameHistoryUCase.GetGameHistoryByRequestId(requestId.String())
	if err != nil {
		util.Logger.Error("查询游戏失败", err)
		return
	}
	// 根据discordId查到用户
	user, err := cl.playerUCase.GetByDiscordUserID(game.PlayerDiscordUserId)
	if err != nil {
		util.Logger.Error("查询用户失败", err)
		return
	}
	client, err := ethclient.Dial(os.Getenv("ALCHEMY_URL"))
	if err != nil {
		util.Logger.Error("连接到以太坊节点失败!", err)
		return
	}
	defer client.Close()
	// 合约地址
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	// 创建合约对象
	contract, err := models.NewEvenOddGame(contractAddress, client)
	if err != nil {
		util.Logger.Error("创建合约对象失败!", err)
		return
	}
	// 签名对象
	chainId, _ := strconv.Atoi(os.Getenv("CHAIN_ID"))
	// 创建私钥（用于签名交易）
	privateKey, err := crypto.HexToECDSA(user.PrivateKey)
	if err != nil {
		util.Logger.Error("创建私钥失败!", err)
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(chainId)))
	if err != nil {
		util.Logger.Error("创建签名对象失败!", err)
		return
	}
	blockTx, err := contract.MainBet(auth, requestId)
	if err != nil {
		util.Logger.Error("调用MainBet合约失败!", err)
		return
	}
	txReceipt, err := bind.WaitMined(context.Background(), client, blockTx)
	if err != nil {
		util.Logger.Error("等待交易失败!", err)
		return
	}
	util.Logger.Info("调用MainBet成功! txHash:", txReceipt.TxHash.Hex())
	requests, err := contract.SRequests(nil, requestId)
	if err != nil {
		util.Logger.Error("查询request失败!", err)
		return
	}
	// util.Logger.Info("查询request成功!", requests)
	game.RandomNumber = requests.RandomNum.String()
	game.GameStatus = e.FINISHED
	game.MainBetTxId = blockTx.Hash().Hex()
	finishTime := time.Unix(requests.GameDoneTimestamp.Int64(), 0)
	game.FinishTime = &finishTime
	// 判断randomNumber是否与choice相同(0)
	isGuessCorrect := (requests.RandomNum.Int64()%2 == 0 && game.Choice == e.EVEN) ||
		(requests.RandomNum.Int64()%2 != 0 && game.Choice == e.ODD)
	game.GameResult = isGuessCorrect
	err = cl.gameHistoryUCase.UpdateGameAfterMainBet(game)
	if err != nil {
		util.Logger.Error("更新game失败!", err)
	}
	// 推送结果到kafka
	writer, err := queue.NewKafkaWriter("game")
	if err != nil {
		util.Logger.Error("创建kafka writer失败!", err)
		return
	}
	marshal, err := json.Marshal(game)
	if err != nil {
		util.Logger.Error("序列化失败!", err)
		return
	}
	util.Logger.Info("序列化完成，结果：", string(marshal))
	err = writer.WriteMessages(context.Background(), kafka.Message{Value: marshal})
	if err != nil {
		util.Logger.Error("写入kafka失败!", err)
		return
	}
	util.Logger.Info("写入kafka成功!")
	err = writer.Close()
	if err != nil {
		return
	}
}
