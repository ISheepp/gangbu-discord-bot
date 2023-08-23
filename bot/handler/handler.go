package handler

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

// AddAllHandlers add all handlers
func AddAllHandlers(session *discordgo.Session) {
	session.AddHandler(newMessage)
	// todo
	session.AddHandler(TestPlay)
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

func Play(discord *discordgo.Session, message *discordgo.ApplicationCommand) {

}

func ShowInfo(discord *discordgo.Session, message *discordgo.MessageCreate) {

}

func Withdraw(discord *discordgo.Session, message *discordgo.MessageCreate) {

}

func TestPlay(discord *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()
	switch data.Name {
	case "play":
		// Do something
		// discord.ChannelMessageSend(i.ChannelID, "那我来帮帮你!")
		embed := &discordgo.MessageEmbed{
			Title:       "那我来帮帮你",
			Description: "测试描述",
		}

		_, err := discord.ChannelMessageSendEmbed(i.ChannelID, embed)
		if err != nil {
			log.Println("发送嵌入式消息失败:", err)
		}
	}
}
