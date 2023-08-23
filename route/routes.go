package route

import (
	"gangbu/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/commands", api.ShowAllCommands())
		v1.DELETE("/commands/:id", api.DeleteCommand())
		v1.POST("/commands", api.CreateCommand())
	}
	return r
}
