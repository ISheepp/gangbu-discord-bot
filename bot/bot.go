package bot

import (
	"bot-demo/handler"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
)

// Run bot running
func Run() {
	// Create new Discord Session
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// add proxy if exist
	if os.Getenv("ENV") == "dev" {
		discord.Dialer.Proxy = func(request *http.Request) (*url.URL, error) {
			u, _ := url.Parse("http://127.0.0.1:7890")
			return u, nil
		}
	}

	// Add event handler
	handler.AddAllHandlers(discord)

	// Open session
	err = discord.Open()
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
