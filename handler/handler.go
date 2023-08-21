package handler

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

// AddAllHandlers add all handlers
func AddAllHandlers(session *discordgo.Session) {
	session.AddHandler(newMessage)
	// todo
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
