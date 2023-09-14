package usecase

import (
	"gangbu/pkg/util"
	"github.com/bwmarrin/discordgo"
	"os"
)

func ShowAllCommands() ([]*discordgo.ApplicationCommand, error) {
	discord := util.GetDiscordClient()
	err := discord.Open()
	if err != nil {
		return nil, err
	}

	defer discord.Close()

	cmd, err := discord.ApplicationCommands(os.Getenv("APP_ID"), "")
	if err != nil {
		return nil, err
	}
	return cmd, nil
}

func DeleteCommand(id string) error {
	discord := util.GetDiscordClient()
	err := discord.Open()
	defer discord.Close()
	if err != nil {
		return err
	}
	err = discord.ApplicationCommandDelete(os.Getenv("APP_ID"), "", id)
	if err != nil {
		return err
	}
	return nil
}

func CreateCommand() error {
	discord := util.GetDiscordClient()
	err := discord.Open()
	defer discord.Close()
	if err != nil {
		return err
	}
	commands := map[string]*discordgo.ApplicationCommand{
		"show": {
			Name:        "show",
			Description: "show player info",
		},
		"play": {
			Name:        "play",
			Description: "play 깐부",
			// 要有输入
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "choice",
					Description: "Please choose even or odd",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "even",
							Value: 0,
						},
						{
							Name:  "odd",
							Value: 1,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionNumber,
					Name:        "value",
					Description: "Input your bet ETH value! unit is ether",
					Required:    true,
				},
			},
		},
		"history": {
			Name:        "history",
			Description: "show game history",
		},
		"withdraw": {
			Name:        "withdraw",
			Description: "withdraw your ETH",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "address",
					Description: "Please input your withdraw address",
					Required:    true,
				},
			},
		},
	}

	for name, command := range commands {
		_, err = discord.ApplicationCommandCreate(os.Getenv("APP_ID"), "", command)
		if err != nil {
			util.Logger.Errorf("创建discord command: %v 失败, err: %v", name, err)
			continue
		}
	}
	return nil
}
