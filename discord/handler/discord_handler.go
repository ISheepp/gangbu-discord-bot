package handler

import (
	"fmt"
	"gangbu/discord/usecase"
	"gangbu/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type discordHandler struct {
}

func NewDiscordHandler(c *gin.Engine) {
	handler := &discordHandler{}
	botGroup := c.Group("/v1/bot")
	botGroup.GET("/commands", handler.showAllCommands())
	botGroup.DELETE("/commands/:id", handler.deleteCommand())
	botGroup.POST("/commands", handler.createCommand())
}

// showAllCommands show all bot commands
func (dh *discordHandler) showAllCommands() gin.HandlerFunc {
	return func(c *gin.Context) {
		commands, err := usecase.ShowAllCommands()
		if err != nil {
			log.Println("show all commands failed!", err)
			c.JSON(http.StatusInternalServerError, &e.ResponseData{
				Message: err.Error(),
				Status:  false,
				Data:    nil,
			})
		}
		c.JSON(http.StatusOK, &e.ResponseData{
			Message: "show all bot commands successfully!",
			Status:  true,
			Data:    commands,
		})
	}
}

// deleteCommand delete a bot command
func (dh *discordHandler) deleteCommand() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := usecase.DeleteCommand(id)
		if err != nil {
			log.Println("delete command failed!", err)
			c.JSON(http.StatusInternalServerError, &e.ResponseData{
				Message: err.Error(),
				Status:  false,
				Data:    nil,
			})
		}
		message := fmt.Sprintf("delete bot command 「%s」 successfully!", id)
		c.JSON(http.StatusOK, &e.ResponseData{
			Message: message,
			Status:  true,
			Data:    nil,
		})
	}
}

// createCommand create a bot command
func (dh *discordHandler) createCommand() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := usecase.CreateCommand()
		if err != nil {
			log.Println("create command failed!", err)
			c.JSON(http.StatusInternalServerError, &e.ResponseData{
				Message: err.Error(),
				Status:  false,
				Data:    nil,
			})
		}
		message := "create bot command successfully!"
		c.JSON(http.StatusOK, &e.ResponseData{
			Message: message,
			Status:  true,
			Data:    nil,
		})
	}
}
