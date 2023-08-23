// main is the entry point of the Go program.
//
// It loads the environment variables for the bot token and the OpenWeather token.
// It then starts the bot by setting the bot token and OpenWeather token,
// and calling the Run function.
package main

import (
	"fmt"
	"gangbu/bot"
	"gangbu/route"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	go bot.Run()
	r := route.NewRouter()
	fmt.Println("web server is running at http://localhost:8989")
	_ = r.Run(":8989")
}
