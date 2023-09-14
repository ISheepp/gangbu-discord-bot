package route

import (
	"context"
	_discordHandler "gangbu/discord/handler"
	_gameHandler "gangbu/game/handler"
	"gangbu/game/listener"
	_gameRepo "gangbu/game/repository/mysql"
	_gameUsecase "gangbu/game/usecase"
	"gangbu/pkg/dao"
	_playerHandler "gangbu/player/handler"
	_playerRepo "gangbu/player/repository/mysql"
	_playerUsecase "gangbu/player/usecase"
	_withdrawHandler "gangbu/withdraw/handler"
	"github.com/gin-gonic/gin"
	"os"
)

func NewRouter() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	db := dao.NewDBClient(context.Background())
	playerRepository := _playerRepo.NewPlayerRepository(db)
	gameHistoryRepository := _gameRepo.NewGameHistoryRepository()
	gameUsecase := _gameUsecase.NewGameUsecase(gameHistoryRepository, playerRepository, db)
	playerUsecase := _playerUsecase.NewPlayerUsecase(playerRepository)

	// router
	_playerHandler.NewPlayerHandler(r, playerUsecase)
	_gameHandler.NewGameHistoryHandler(r, gameUsecase, playerUsecase)
	_discordHandler.NewDiscordHandler(r)
	_withdrawHandler.NewWithdrawHandler(r, playerUsecase)
	// 监听器
	chainListener := listener.NewChainListener(gameUsecase, playerUsecase)
	kafkaListener := listener.NewKafkaListener(playerUsecase)
	chainListener.StartListen()
	kafkaListener.StartListen()
	return r
}
