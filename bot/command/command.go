package command

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func InstallCommand(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "play",
		Type:        discordgo.ChatApplicationCommand,
		Description: "play 깐부",
	}

	_, err := s.ApplicationCommandCreate("1140967812983160932", "", command)
	if err != nil {
		log.Fatalln("install command failed!", err)
	}
}

func ShowAllCommands(s *discordgo.Session) []*discordgo.ApplicationCommand {
	cmd, err := s.ApplicationCommands("1140967812983160932", "", nil)
	if err != nil {
		log.Fatalln("show all commands failed!", err)
	}
	return cmd
}
