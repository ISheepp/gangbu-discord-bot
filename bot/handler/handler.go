package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gangbu/pkg/app"
	"gangbu/pkg/e"
	"gangbu/pkg/models"
	"gangbu/pkg/util"
	"github.com/bwmarrin/discordgo"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// AddAllHandlers add all handlers
func AddAllHandlers(session *discordgo.Session) {
	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

// newMessage 第二个参数是侦听的事件类型
func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore bot messaage
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// Respond to messages
	switch {
	case strings.Contains(message.Content, "weather"):
		discord.ChannelMessageSend(message.ChannelID, "I can help with that!")
	case strings.Contains(message.Content, "bot"):
		discord.ChannelMessageSend(message.ChannelID, "Hi there!")
	}
}

var (
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"play": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// 调用后端/play接口，传递GuildID和ChannelID
			interactionData := i.Data.(discordgo.ApplicationCommandInteractionData)
			choice := interactionData.Options[0].Value.(float64)
			f := interactionData.Options[1].Value.(float64)

			gameBo := models.GameHistoryBo{
				PlayerDiscordUserId: i.Member.User.ID,
				Choice:              uint8(choice),
				BetValue:            util.EtherToWei(f),
				GuildID:             i.GuildID,
				ChannelID:           i.ChannelID,
			}
			req, err := json.Marshal(gameBo)
			if err != nil {
				util.Logger.Error("json序列化失败", err)
				return
			}
			url := "http://127.0.0.1:8989/v1/play"
			request, err := http.NewRequest("POST", url, bytes.NewBuffer(req))
			if err != nil {
				util.Logger.Error("创建post请求出错：", err)
				return
			}
			request.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			go func() {
				response, err := client.Do(request)
				if err != nil {
					util.Logger.Error("发起请求时出错:", err)
					return
				}
				defer response.Body.Close()
			}()

			// 读取响应内容 todo 不读取返回，后续解决异常处理，因为discord只能响应3s内的数据
			// body, err := io.ReadAll(response.Body)
			// if err != nil {
			// 	fmt.Println("读取响应时出错:", err)
			// 	return
			// }
			// responseData := app.Response{}
			// err = json.Unmarshal(body, &responseData)
			// if err != nil {
			// 	util.Logger.Error("json反序列化失败:", err)
			// 	return
			// }
			// if responseData.Code != e.SUCCESS {
			// 	util.Logger.Error("调用play接口失败！原因：", responseData.Data)
			// 	content := fmt.Sprintf("Invoke game failed! please try again! reason: %v", responseData.Data)
			// 	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			// 		Type: discordgo.InteractionResponseChannelMessageWithSource,
			// 		Data: &discordgo.InteractionResponseData{
			// 			Content: content,
			// 			Flags:   discordgo.MessageFlagsEphemeral,
			// 		},
			// 	})
			// 	if err != nil {
			// 		util.Logger.Error(err)
			// 		return
			// 	}
			// 	return
			// }
			// etherscanUrl := "https://goerli.etherscan.io/tx/" + responseData.Data.(string)
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
			url := "http://127.0.0.1:8989/v1/player/" + i.Member.User.ID
			response, err := http.Get(url)
			if err != nil {
				util.Logger.Error("发起请求时出错:", err)
				return
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				util.Logger.Error("读取响应时出错:", err)
				return
			}
			responseData := app.Response{}
			err = json.Unmarshal(body, &responseData)
			if err != nil {
				util.Logger.Error("json反序列化失败:", err)
				return
			}
			if responseData.Code != e.SUCCESS {
				util.Logger.Error("调用showInfo接口失败！原因：", responseData.Data)
				content := fmt.Sprintf("Invoke game failed! please try again! reason: %v", responseData.Data)
				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: content,
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				})
				if err != nil {
					util.Logger.Error(err)
					return
				}
				return
			}
			result := responseData.Data.(map[string]interface{})
			sourceData := result["WalletValue"].(float64)
			// float64转int64
			ether := util.WeiToEther(int64(sourceData))
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
									Value: result["WalletAddress"].(string),
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
			url := "http://127.0.0.1:8989/v1/game-history-last-five/" + i.Member.User.ID
			response, err := http.Get(url)
			if err != nil {
				util.Logger.Error("发起请求时出错:", err)
				return
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				util.Logger.Error("读取响应时出错:", err)
				return
			}
			responseData := app.Response{}
			err = json.Unmarshal(body, &responseData)
			if err != nil {
				util.Logger.Error("json反序列化失败:", err)
				return
			}
			if responseData.Code != e.SUCCESS {
				util.Logger.Error("调用showInfo接口失败！原因：", responseData.Data)
				content := fmt.Sprintf("Invoke game failed! please try again! reason: %v", responseData.Data)
				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: content,
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				})
				if err != nil {
					util.Logger.Error(err)
					return
				}
				return
			}
			marshal, _ := json.Marshal(responseData.Data)
			var hisSlice []*models.GameHistory
			_ = json.Unmarshal(marshal, &hisSlice)

			msgEmbeds := make([]*discordgo.MessageEmbed, 0)
			for _, his := range hisSlice {
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
							Value:  his.FinishTime.Format("2006-01-02 15:04:05 MST"),
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
			url := fmt.Sprintf("http://127.0.0.1:8989/v1/withdraw/%s/receive/%s", i.Member.User.ID, receive)
			response, err := http.Get(url)
			if err != nil {
				util.Logger.Error("发起请求时出错:", err)
				return
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				util.Logger.Error("读取响应时出错:", err)
				return
			}
			responseData := app.Response{}
			err = json.Unmarshal(body, &responseData)
			if err != nil {
				util.Logger.Error("json反序列化失败:", err)
				return
			}
			if responseData.Code != e.SUCCESS {
				util.Logger.Error("调用withdraw接口失败！原因：", responseData.Data)
				content := fmt.Sprintf("Invoke game failed! please try again! reason: %v", responseData.Data)
				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: content,
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				})
				if err != nil {
					util.Logger.Error(err)
					return
				}
				return
			}
			txId := responseData.Data.(string)
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
