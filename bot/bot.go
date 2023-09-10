package bot

import (
	"fmt"
	"gangbu/bot/handler"
	"gangbu/pkg/util"
	"log"
	"os"
	"os/signal"
)

// Run bot running
func Run() {
	discord := util.GetDiscordClient()

	// Add event handler
	handler.AddAllHandlers(discord)

	// Open session
	err := discord.Open()
	if err != nil {
		log.Fatal("Cannot open Discord session: ", err)
	}
	defer discord.Close()

	// Run until code is terminated
	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
