package route

import (
	"context"
	"gangbu/pkg/dao"
	"gangbu/player/handler"
	"gangbu/player/repository/mysql"
	"gangbu/player/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	db := dao.NewDBClient(context.Background())
	repository := mysql.NewPlayerRepository(db)
	playerUsecase := usecase.NewPlayerUsecase(repository)
	handler.NewPlayerHandler(r, playerUsecase)
	return r
}
