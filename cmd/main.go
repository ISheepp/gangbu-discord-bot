// main is the entry point of the Go program.
//
// It loads the environment variables for the bot token and the OpenWeather token.
// It then starts the bot by setting the bot token and OpenWeather token,
// and calling the Run function.
package main

import (
	"gangbu/bot"
	"gangbu/pkg/db"
	"gangbu/pkg/util"
	"gangbu/route"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	go bot.Run()
	r := route.NewRouter()
	util.Logger.Info("web server is running at http://localhost:8989")
	_ = r.Run(os.Getenv("PORT"))
}

func init() {
	util.InitLog()
	db.MySQLInit()
}
