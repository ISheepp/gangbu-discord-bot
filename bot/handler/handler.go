package handler

import (
	"gangbu/pkg/util"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

// AddAllHandlers add all handlers
func AddAllHandlers(session *discordgo.Session) {
	session.AddHandler(newMessage)
	// todo
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
	discord.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hey there! Congratulations, you just executed your first slash command",
		},
	})

}

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "play",
			Description: "Basic command",
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"play": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Huh. I see, maybe some of these resources might help you?",
					Flags:   discordgo.MessageFlagsEphemeral,
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Emoji: discordgo.ComponentEmoji{
										Name: "📜",
									},
									Label: "Documentation",
									Style: discordgo.LinkButton,
									URL:   "https://discord.com/developers/docs/interactions/message-components#buttons",
								},
								discordgo.Button{
									Emoji: discordgo.ComponentEmoji{
										Name: "🔧",
									},
									Label: "Discord developers",
									Style: discordgo.LinkButton,
									URL:   "https://discord.gg/discord-developers",
								},
							},
						},
					},
				},
				//Type: discordgo.InteractionResponseChannelMessageWithSource,
				//Data: &discordgo.InteractionResponseData{
				//	Content: "CONTENT",
				//},
			})
			if err != nil {
				util.Logger.Error(err)
				return
			}
		},
		"info": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		},
		"history": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		},
		"withdraw": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		},
	}
)

func installCommands() {
	s := util.GetDiscordClient()
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
}
