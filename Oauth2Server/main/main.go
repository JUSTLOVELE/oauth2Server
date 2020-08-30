package main

import (
	"Oauth2Server/controller"
	"Oauth2Server/env"
	"Oauth2Server/test"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//创建数据库引擎
	env.CreateDB()
	env.InitConf()
	env.InitLog()
	test.DBTest()

	router := registerRoutes()
	router.Run(":12001")
}

func registerRoutes() *gin.Engine {

	router := gin.Default()
	router.GET("getAppId", controller.GetAppId)

	return router
}

