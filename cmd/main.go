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
	_ "github.com/joho/godotenv/autoload"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	// start grpc server
	wg := &sync.WaitGroup{}
	ticker := time.NewTicker(4 * time.Hour)
	wg.Add(1)
	go RunGrpcServer(wg)
	go bot.Run(wg)
	go ListenRenewChan(ticker, wg)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ticker.Stop()
	unRegistryFromRedis()
	util.Logger.Info("stopping server...")
}

func init() {
	util.InitLog()
	db.MySQLInit()
}
