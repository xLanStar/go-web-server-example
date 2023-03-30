package controllers

import (
	"go-web-server-example/internal/exception"
	"go-web-server-example/internal/services/ai"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {
	var data ai.RequestData

	c.BindJSON(&data)

	response, err := ai.Request(&data)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": exception.UNKNOWN,
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
