package handler

import (
	"gangbu/pkg/app"
	"gangbu/pkg/e"
	"gangbu/pkg/models"
	"gangbu/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type gameHistoryHandler struct {
	ghUsecase models.GameHistoryUsecase
	pUsecase  models.PlayerUsecase
}

func (ghh *gameHistoryHandler) play(c *gin.Context) {
	ghu := ghh.ghUsecase
	appG := app.Gin{
		C: c,
	}
	bo := &models.GameHistoryBo{}
	err := c.BindJSON(bo)
	if err != nil {
		util.Logger.Error("解析json失败!", err)
		appG.Response(http.StatusInternalServerError, e.ERROR, err)
	}
	// todo blocking
	err = ghu.CreateGame(*bo)
	if err != nil {
		util.Logger.Error("创建游戏失败!", err)
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	util.Logger.Info("创建游戏成功!")
	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func (ghh *gameHistoryHandler) getGameHistoryByDiscordId(c *gin.Context) {
	ghu := ghh.ghUsecase
	appG := app.Gin{
		C: c,
	}
	discordId := c.Param("id")

	gameHistory, err := ghu.GetGameHistoryByDiscordId(discordId)
	if err != nil {
		util.Logger.Error("获取游戏记录失败!", err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, gameHistory)
}

func NewGameHistoryHandler(c *gin.Engine, ghu models.GameHistoryUsecase, pu models.PlayerUsecase) {
	handler := &gameHistoryHandler{
		ghUsecase: ghu,
		pUsecase:  pu,
	}
	v1 := c.Group("/v1")
	{
		v1.POST("/play", handler.play)
		v1.GET("/game-history/:id", handler.getGameHistoryByDiscordId)
	}
}
