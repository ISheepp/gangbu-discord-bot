package route

import (
	"context"
	_gameHandler "gangbu/game/handler"
	"gangbu/game/listener"
	_gameRepo "gangbu/game/repository/mysql"
	_gameUsecase "gangbu/game/usecase"
	"gangbu/pkg/dao"
	_playerHandler "gangbu/player/handler"
	_playerRepo "gangbu/player/repository/mysql"
	_playerUsecase "gangbu/player/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	db := dao.NewDBClient(context.Background())
	playerRepository := _playerRepo.NewPlayerRepository(db)
	gameHistoryRepository := _gameRepo.NewGameHistoryRepository()
	gameUsecase := _gameUsecase.NewGameUsecase(gameHistoryRepository, playerRepository, db)
	playerUsecase := _playerUsecase.NewPlayerUsecase(playerRepository)
	_playerHandler.NewPlayerHandler(r, playerUsecase)
	_gameHandler.NewGameHistoryHandler(r, gameUsecase, playerUsecase)
	chainListener := listener.NewChainListener(gameUsecase, playerUsecase)
	kafkaListener := listener.NewKafkaListener()
	// 监听器
	chainListener.StartListen()
	kafkaListener.StartListen()
	return r
}
