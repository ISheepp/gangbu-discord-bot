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
	// todo 耗时太长了，超出了3s，后续优化
	// start := time.Now()
	// // validate user id
	// discord := util.GetDiscordClient()
	// err := discord.Open()
	// if err != nil {
	// 	util.Logger.Error("打开discord失败!", err)
	// 	appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	// 	return
	// }
	// user, err := discord.User(discordId)
	// if err != nil {
	// 	util.Logger.Error("查询discord用户信息失败!", err)
	// 	appG.Response(http.StatusInternalServerError, e.ErrorDiscordId, nil)
	// 	return
	// }
	// end := time.Now()
	// util.Logger.Info("查询discord用户信息耗时:", end.Sub(start))
	// if user == nil {
	// 	util.Logger.Error("discord用户信息不存在!：", discordId)
	// 	appG.Response(http.StatusInternalServerError, e.DiscordUserNotFound, nil)
	// 	return
	// }
	//
	// defer func() {
	// 	err = discord.Close()
	// 	if err != nil {
	// 		util.Logger.Error("关闭discord失败!", err)
	// 		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	// 		return
	// 	}
	// }()
	// marshal, err := json.Marshal(user)
	// util.Logger.Info("查询discord接口信息成功!：", string(marshal))
	player, err := pu.GetByDiscordUserIDOrCreate(discordId)
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
