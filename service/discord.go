package service

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
	_, err = discord.ApplicationCommandCreate(os.Getenv("APP_ID"), "", &discordgo.ApplicationCommand{
		Name:        "test",
		Description: "test command",
	})
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}
