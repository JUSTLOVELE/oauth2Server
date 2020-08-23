package main

import (
	"Oauth2Server/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	router := registerRoutes()
	router.Run(":12001")
}

func registerRoutes() *gin.Engine {

	router := gin.Default()
	router.GET("getAppId", controller.GetAppId)

	return router
}

