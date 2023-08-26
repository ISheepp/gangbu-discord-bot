package route

import (
	"gangbu/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1Bot := r.Group("/api/v1/bot")
	{
		v1Bot.GET("/commands", api.ShowAllCommands())
		v1Bot.DELETE("/commands/:id", api.DeleteCommand())
		v1Bot.POST("/commands", api.CreateCommand())
	}

	v1 := r.Group("/api/v1")
	{
		v1.GET("/discord", api.ShowAllCommands())
	}
	return r
}
