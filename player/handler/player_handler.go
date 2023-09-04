package handler

import (
	"gangbu/pkg/app"
	"gangbu/pkg/e"
	"gangbu/pkg/models"
	"gangbu/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PlayerHandler struct {
	PUsecase models.PlayerUsecase
}

func (ph *PlayerHandler) ShowPlayerInfo(c *gin.Context) {
	pu := ph.PUsecase
	appG := app.Gin{
		C: c,
	}
	discordId := c.Param("id")
	// todo 校验discordId是否合理
	player, err := pu.GetByDiscordUserID(discordId)
	if err != nil {
		util.Logger.Error("查询玩家信息失败！", err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	}
	appG.Response(http.StatusOK, e.SUCCESS, player)
}

func NewPlayerHandler(c *gin.Engine, pu models.PlayerUsecase) {
	handler := &PlayerHandler{
		PUsecase: pu,
	}
	v1 := c.Group("/v1")
	{

		v1.GET("/player/:id", handler.ShowPlayerInfo)
	}
}
