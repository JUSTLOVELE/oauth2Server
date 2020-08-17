package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := registerRoutes()
	router.Run(":12001")

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context){
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//
	//r.Run(":12001")
}

func registerRoutes() *gin.Engine {

	router := gin.Default()
	//router.GET("/test", controller.TestAction)

	return router
}

