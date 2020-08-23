package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"strings"
)

func GetAppId(c *gin.Context) {

	id := uuid.NewV4().String()
	uuid := strings.ReplaceAll(id, "-", "")

	c.JSON(200, gin.H{
		"appid": uuid,
	})
}