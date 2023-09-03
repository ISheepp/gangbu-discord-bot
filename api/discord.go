package api

import (
	"fmt"
	"gangbu/pkg/e"
	"gangbu/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// ShowAllCommands show all bot commands
func ShowAllCommands() gin.HandlerFunc {
	return func(c *gin.Context) {
		commands, err := service.ShowAllCommands()
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

// DeleteCommand delete a bot command
func DeleteCommand() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := service.DeleteCommand(id)
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

// CreateCommand create a bot command
func CreateCommand() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := service.CreateCommand()
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
