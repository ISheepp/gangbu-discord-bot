package bot

import (
	"gangbu/bot/handler"
	"gangbu/pkg/util"
	"log"
	"os"
	"os/signal"
	"sync"
)

// Run bot running
func Run(wg *sync.WaitGroup) {
	wg.Wait()
	discord := util.GetDiscordClient()

	client, err := util.GetGrpcClient("127.0.0.1:8989")
	// Add event handler
	botHandler := handler.NewBotHandler(client)
	botHandler.AddAllHandlers(discord)

	// Open session
	err = discord.Open()
	if err != nil {
		log.Fatal("Cannot open Discord session: ", err)
	}
	defer discord.Close()

	// Run until code is terminated
	util.Logger.Info("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
