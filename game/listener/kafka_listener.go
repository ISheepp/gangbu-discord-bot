package listener

import (
	"context"
	"encoding/json"
	"fmt"
	"gangbu/pkg/e"
	"gangbu/pkg/models"
	"gangbu/pkg/queue"
	"gangbu/pkg/util"
	"gangbu/proto/game"
	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
	"reflect"
	"strconv"
	"time"
)

type kafkaListener struct {
	pUsecase models.PlayerUsecase
}

func NewKafkaListener(pu models.PlayerUsecase) Listener {
	return &kafkaListener{pUsecase: pu}
}

func (k *kafkaListener) StartListen() {
	go ListenGameDoneHistory(k)
	go ListenCreateGameResult(k)
}

// ListenGameDoneHistory 监听游戏结束
func ListenGameDoneHistory(k *kafkaListener) {
	reader, err := queue.NewKafkaReader("game", "test-1")
	if err != nil {
		util.Logger.Error("创建kafka reader失败", err)
	}
	util.Logger.Info("开始监听kafka, topic: game")
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			util.Logger.Error("kafka reader读取消息失败", err)
			time.Sleep(5 * time.Second)
			continue
		}
		gameHis := &models.GameHistory{}
		err = json.Unmarshal(msg.Value, gameHis)
		if err != nil {
			util.Logger.Error("kafka reader解析消息失败", err)
			continue
		}
		// 发送到discord服务器
		go pushMsgToDiscord(k, gameHis)

	}
}

func pushMsgToDiscord(k *kafkaListener, game *models.GameHistory) {
	discord := util.GetDiscordClient()
	defer discord.Close()
	if game.GameStatus == e.IN_PROGRESS {
		content := fmt.Sprintf("Game run failed! <@%s>", game.PlayerDiscordUserId)
		_, err := discord.ChannelMessageSendComplex(game.ChannelID, &discordgo.MessageSend{
			Content: content,
			AllowedMentions: &discordgo.MessageAllowedMentions{
				Users: []string{game.PlayerDiscordUserId},
			},
		})
		if err != nil {
			util.Logger.Error("发送游戏结果消息失败", err)
		}
	}
	// 发送消息
	requestRandomTxUrl := "https://goerli.etherscan.io/tx/" + game.RequestRandomTxId
	mainBetTxUrl := "https://goerli.etherscan.io/tx/" + game.MainBetTxId
	gameResult := ""
	if game.GameResult {
		gameResult = "Win"
	} else {
		gameResult = "Lose"
	}
	betValue := util.WeiToEther(game.BetValue)
	// 查询钱包余额
	player, err := k.pUsecase.GetByDiscordUserID(game.PlayerDiscordUserId)
	client, err := ethclient.Dial(os.Getenv("ALCHEMY_URL"))
	if err != nil {
		util.Logger.Error("连接到以太坊节点失败!", err)
		return
	}
	defer client.Close()
	if err != nil {
		util.Logger.Error(err)
		return
	}
	value, err := client.BalanceAt(context.Background(), common.HexToAddress(player.WalletAddress), nil)

	if err != nil {
		util.Logger.Error("查询玩家余额失败", err)
		return
	}
	walletValue := util.WeiToEther(value.Int64())
	// 将时间对象格式化为 ISO 8601 格式的字符串
	iso8601String := game.FinishTime.Format("2006-01-02T15:04:05-07:00")
	content := fmt.Sprintf("Game completed! <@%s>", game.PlayerDiscordUserId)
	_, err = discord.ChannelMessageSendComplex(game.ChannelID, &discordgo.MessageSend{
		Content: content,
		Embed: &discordgo.MessageEmbed{
			Title:     "Final Board",
			Timestamp: iso8601String,
			Footer: &discordgo.MessageEmbedFooter{
				Text:         "Developed by ISheep🐑",
				IconURL:      "https://ipfs.xlog.app/ipfs/bafkreihyfxi5fmurms77dv6qv5q6kbosbfshqz6q7kff6uczkfw5hb4qgq",
				ProxyIconURL: "https://isheep.xlog.app/",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Game Result",
					Value:  gameResult,
					Inline: true,
				},
				{
					Name:   "Bet Value",
					Value:  strconv.FormatFloat(betValue, 'f', -1, 64) + " ETH",
					Inline: true,
				},
				{
					Name:   "Wallet Balance",
					Value:  strconv.FormatFloat(walletValue, 'f', -1, 64) + " ETH",
					Inline: true,
				},
			},
		},
		Components: []discordgo.MessageComponent{
			&discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					&discordgo.Button{
						Emoji: discordgo.ComponentEmoji{
							Name: "📜",
						},
						Label: "Etherscan:Request",
						Style: discordgo.LinkButton,
						URL:   requestRandomTxUrl,
					},
					&discordgo.Button{
						Emoji: discordgo.ComponentEmoji{
							Name: "🥊",
						},
						Label: "Etherscan:Bet",
						Style: discordgo.LinkButton,
						URL:   mainBetTxUrl,
					},
				},
			},
		},
		AllowedMentions: &discordgo.MessageAllowedMentions{
			Users: []string{game.PlayerDiscordUserId},
		},
	})
	if err != nil {
		util.Logger.Error("发送游戏结果消息失败", err)
	}
}

func pushGameCreateResultToDiscord(k *kafkaListener, g *game.CallbackMessage) {
	discord := util.GetDiscordClient()
	defer discord.Close()
	// 发送消息
	content := fmt.Sprintf(" <@%s>", g.PlayerDiscordUserId)
	requestRandomTxUrl := "https://goerli.etherscan.io/tx/" + g.GetData()
	components := []discordgo.MessageComponent{
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				&discordgo.Button{
					Emoji: discordgo.ComponentEmoji{
						Name: "📜",
					},
					Label: "Etherscan:Request",
					Style: discordgo.LinkButton,
					URL:   requestRandomTxUrl,
				},
			},
		},
	}
	i := reflect.ValueOf(g.Type).Interface()

	if game.CallbackMessageType_ERROR == i {
		components = []discordgo.MessageComponent{}
	}
	_, err := discord.ChannelMessageSendComplex(g.ChannelId, &discordgo.MessageSend{
		Content:    g.Message + content,
		Components: components,
		AllowedMentions: &discordgo.MessageAllowedMentions{
			Users: []string{g.PlayerDiscordUserId},
		},
	})
	if err != nil {
		util.Logger.Error("发送游戏结果消息失败", err)
	}
}

func ListenCreateGameResult(k *kafkaListener) {
	reader, err := queue.NewKafkaReader("callback_msg", "test-2")
	if err != nil {
		util.Logger.Error("创建kafka reader失败", err)
	}
	util.Logger.Info("开始监听kafka, topic: callback_msg")
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			util.Logger.Error("kafka reader读取消息失败", err)
			time.Sleep(5 * time.Second)
			continue
		}
		g := &game.CallbackMessage{}
		err = json.Unmarshal(msg.Value, g)
		if err != nil {
			util.Logger.Error("kafka reader解析消息失败", err)
			continue
		}
		// 发送到discord服务器
		go pushGameCreateResultToDiscord(k, g)

	}
}
