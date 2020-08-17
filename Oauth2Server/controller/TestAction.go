package controller

import "github.com/gin-gonic/gin"

func TestAction(c *gin.Context)  {

		c.JSON(200, gin.H{
			"message": "test",
		})
}