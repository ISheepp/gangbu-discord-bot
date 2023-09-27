package handler

import (
	"context"
	"gangbu/pkg/cache"
	"gangbu/pkg/util"
	"gangbu/proto/game"
	"gangbu/proto/player"
	"gangbu/proto/withdraw"
	"github.com/bwmarrin/discordgo"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"os"
	"strconv"
	"time"
)

type botHandler struct {
}

func NewBotHandler() *botHandler {
	return &botHandler{}
}

// AddAllHandlers add all handlers
func (h *botHandler) AddAllHandlers(session *discordgo.Session) {
	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

var (
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"play": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			interactionData := i.Data.(discordgo.ApplicationCommandInteractionData)
			choice := interactionData.Options[0].Value.(float64)
			f := interactionData.Options[1].Value.(float64)
			addr, err := getGrpcAddr()
			if err != nil {
				util.Logger.Error("获取grpc地址失败: %v", err)
				failedContent(s, i, "get grpc address failed")
				return
			}
			conn, err := util.GetGrpcClientConn(addr)
			if err != nil {
				util.Logger.Error("连接服务器失败: %v", err)
				failedContent(s, i, "link to server failed")
				return
			}
			grc := game.NewGameRequestClient(conn)

			go func() {
				txHash, err := grc.CreateGame(context.Background(), &game.GameCreateBo{
					PlayerDiscordUserId: i.Member.User.ID,
					Choice:              uint32(choice),
					BetValue:            util.EtherToWei(f),
					GuildId:             i.GuildID,
					ChannelId:           i.ChannelID,
				})
				if err != nil {
					util.Logger.Error("创建游戏失败!", err)
					// 发送到kafka 由kafka来发送消息
					_, err := grc.SendCallbackMessage(context.Background(), &game.CallbackMessage{
						Message:             "create game failed!" + " reason: " + err.Error(),
						Data:                "",
						Type:                game.CallbackMessageType_ERROR,
						ChannelId:           i.ChannelID,
						PlayerDiscordUserId: i.Member.User.ID,
					})
					if err != nil {
						util.Logger.Error("发送回调消息失败!", err)
					}
					return
				}
				// 发送到kafka
				_, err = grc.SendCallbackMessage(context.Background(), &game.CallbackMessage{
					Message:             "create game success!",
					Data:                txHash.Value,
					Type:                game.CallbackMessageType_SUCCESS,
					ChannelId:           i.ChannelID,
					PlayerDiscordUserId: i.Member.User.ID,
				})
				if err != nil {
					util.Logger.Error("发送回调消息失败!", err)
				}
				err = conn.Close()
				if err != nil {
					util.Logger.Error("关闭连接失败", err)
				}
			}()
			// requestRandomTxUrl := "https://goerli.etherscan.io/tx/" + txHash.String()
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Gaming started! please wait message...",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			if err != nil {
				util.Logger.Error(err)
				return
			}
		},
		"show": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// 展示用户信息
			addr, err := getGrpcAddr()
			if err != nil {
				util.Logger.Error("获取grpc地址失败: %v", err)
				failedContent(s, i, "get grpc address failed")
				return
			}
			conn, err := util.GetGrpcClientConn(addr)
			if err != nil {
				util.Logger.Error("连接服务器失败: %v", err)
				failedContent(s, i, "link to server failed")
				return
			}
			defer func(conn *grpc.ClientConn) {
				err := conn.Close()
				if err != nil {
					util.Logger.Error(err)
				}
			}(conn)
			grc := player.NewPlayerRequestClient(conn)
			playerInfo, err := grc.ShowPlayerInfo(context.Background(), &wrapperspb.StringValue{Value: i.Member.User.ID})
			if err != nil {
				util.Logger.Error("查询玩家信息失败!", err)
				failedContent(s, i, "query player info failed")
				return
			}
			ether := util.WeiToEther(playerInfo.WalletValue)
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Get player info completed!",
					Flags:   discordgo.MessageFlagsEphemeral,
					Embeds: []*discordgo.MessageEmbed{
						{
							Title:     "Player Board",
							Timestamp: time.Now().Format("2006-01-02T15:04:05-07:00"),
							Footer: &discordgo.MessageEmbedFooter{
								Text:         "Developed by ISheep🐑",
								IconURL:      "https://ipfs.xlog.app/ipfs/bafkreihyfxi5fmurms77dv6qv5q6kbosbfshqz6q7kff6uczkfw5hb4qgq",
								ProxyIconURL: "https://isheep.xlog.app/",
							},
							Fields: []*discordgo.MessageEmbedField{
								{
									Name:  "Wallet Address",
									Value: playerInfo.WalletAddress,
								},
								{
									Name:  "Wallet Value",
									Value: strconv.FormatFloat(ether, 'f', -1, 64) + " ETH",
								},
							},
						},
					},
				},
			})

			if err != nil {
				util.Logger.Error("发送游戏结果消息失败", err)
			}
		},
		"history": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// 展示用户信息
			addr, err := getGrpcAddr()
			if err != nil {
				util.Logger.Error("获取grpc地址失败: %v", err)
				failedContent(s, i, "get grpc address failed")
				return
			}
			conn, err := util.GetGrpcClientConn(addr)
			if err != nil {
				util.Logger.Error("连接服务器失败: %v", err)
				failedContent(s, i, "link to server failed")
				return
			}
			defer func(conn *grpc.ClientConn) {
				err := conn.Close()
				if err != nil {
					util.Logger.Error(err)
				}
			}(conn)
			grc := game.NewGameRequestClient(conn)
			historyDtoSlice, err := grc.GetLastFiveGameHistoryByDiscordIdFromTheGraph(context.Background(), &wrapperspb.StringValue{Value: i.Member.User.ID})
			if err != nil {
				util.Logger.Error("获取游戏记录失败", err)
				return
			}
			hisSlice := historyDtoSlice.Histories
			msgEmbeds := make([]*discordgo.MessageEmbed, 0)
			for _, his := range hisSlice {
				finishTime := time.Unix(his.FinishTime.GetSeconds(), 0)
				result := his.GameResult
				value := util.WeiToEther(his.BetValue)
				gameResultStr := ""
				if result {
					gameResultStr = "Win"
				} else {
					gameResultStr = "Lose"
				}
				msg := &discordgo.MessageEmbed{
					Title:     "Final Board",
					Timestamp: time.Now().Format("2006-01-02T15:04:05-07:00"),
					Footer: &discordgo.MessageEmbedFooter{
						Text:         "Developed by ISheep🐑",
						IconURL:      "https://ipfs.xlog.app/ipfs/bafkreihyfxi5fmurms77dv6qv5q6kbosbfshqz6q7kff6uczkfw5hb4qgq",
						ProxyIconURL: "https://isheep.xlog.app/",
					},
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "Game Result",
							Value:  gameResultStr,
							Inline: true,
						},
						{
							Name:   "Bet Value",
							Value:  strconv.FormatFloat(value, 'f', -1, 64) + " ETH",
							Inline: true,
						},
						{
							Name:   "Play Time",
							Value:  finishTime.Format("2006-01-02 15:04:05 MST"),
							Inline: true,
						},
					},
				}
				msgEmbeds = append(msgEmbeds, msg)
			}
			// float64转int64
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Get Game History completed!",
					Flags:   discordgo.MessageFlagsEphemeral,
					Embeds:  msgEmbeds,
				},
			})
			if err != nil {
				util.Logger.Error(err)
			}
		},
		"withdraw": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// 提款
			interactionData := i.Data.(discordgo.ApplicationCommandInteractionData)
			receive := interactionData.Options[0].Value.(string)
			addr, err := getGrpcAddr()
			if err != nil {
				util.Logger.Error("获取grpc地址失败: %v", err)
				failedContent(s, i, "get grpc address failed")
				return
			}
			conn, err := util.GetGrpcClientConn(addr)
			if err != nil {
				util.Logger.Error("连接服务器失败: %v", err)
				failedContent(s, i, "link to server failed")
				return
			}
			defer func(conn *grpc.ClientConn) {
				err := conn.Close()
				if err != nil {
					util.Logger.Error(err)
				}
			}(conn)
			wrc := withdraw.NewWithdrawRequestClient(conn)
			txHash, err := wrc.Withdraw(context.Background(), &withdraw.WithdrawBo{
				DiscordUserId:   i.Member.User.ID,
				WithdrawAddress: receive,
			})
			if err != nil {
				util.Logger.Error(err)
				failedContent(s, i, "withdraw failed, reason: "+err.Error())
				return
			}
			txId := txHash.GetValue()
			txUrl := "https://goerli.etherscan.io/tx/" + txId
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Withdrawing...please wait!",
					Flags:   discordgo.MessageFlagsEphemeral,
					Components: []discordgo.MessageComponent{
						&discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								&discordgo.Button{
									Emoji: discordgo.ComponentEmoji{
										Name: "📜",
									},
									Label: "Etherscan:Withdraw",
									Style: discordgo.LinkButton,
									URL:   txUrl,
								},
							},
						},
					},
				},
			})

			if err != nil {
				util.Logger.Error("发送游戏结果消息失败", err)
			}

		},
	}
)

// failedContent sends an error message to the user
func failedContent(s *discordgo.Session, i *discordgo.InteractionCreate, content string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		util.Logger.Error(err)
	}
	return
}

func getGrpcAddr() (string, error) {
	client := cache.NewRedisCache()
	addr, err := client.GetString(context.Background(), os.Getenv("GRPC_NAME"))
	if err != nil {
		util.Logger.Errorf("getGrpcAddr error: %v", err)
		return "", err
	}
	return addr, nil
}
