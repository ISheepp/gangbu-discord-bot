// main is the entry point of the Go program.
//
// It loads the environment variables for the bot token and the OpenWeather token.
// It then starts the bot by setting the bot token and OpenWeather token,
// and calling the Run function.
package main

import (
	"bot-demo/bot"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	bot.Run()
}
