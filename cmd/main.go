// main is the entry point of the Go program.
//
// It loads the environment variables for the bot token and the OpenWeather token.
// It then starts the bot by setting the bot token and OpenWeather token,
// and calling the Run function.
package main

import (
	"context"
	"gangbu/bot"
	_gameRepo "gangbu/game/repository/mysql"
	"gangbu/game/server"
	_gameUsecase "gangbu/game/usecase"
	"gangbu/pkg/db"
	"gangbu/pkg/util"
	_playerRepo "gangbu/player/repository/mysql"
	_playerUsecase "gangbu/player/usecase"
	_ "github.com/joho/godotenv/autoload"
	"sync"
)

func main() {
	// start grpc server
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		util.Logger.Info("start grpc server")
		dbClient := db.NewDBClient(context.Background())
		playerRepository := _playerRepo.NewPlayerRepository(dbClient)
		gameHistoryRepository := _gameRepo.NewGameHistoryRepository()
		gameUsecase := _gameUsecase.NewGameUsecase(gameHistoryRepository, playerRepository, dbClient)
		playerUsecase := _playerUsecase.NewPlayerUsecase(playerRepository)

		gameServer := server.NewGameServer(gameUsecase, playerUsecase)
		util.StartGrpcServer(wg, gameServer)
		//r := route.NewRouter()
		util.Logger.Info("GRPC server is running at http://localhost:8989")
	}()
	//_ = r.Run(os.Getenv("PORT"))
	go bot.Run(wg)
	select {}
}

func init() {
	util.InitLog()
	db.MySQLInit()
}
