package main

import (
	"Oauth2Server/controller"
	"Oauth2Server/env"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//第一步:初始化配置文件
	env.InitConf()
	//第二步初始化日志
	env.InitLog()
	//第三步初始化数据库
	env.CreateDB()
	//测试数据库连接的时候取消注释
	//test.DBTest()
	//第四步启动服务
	router := registerRoutes()
	router.Run(":12001")
}

func registerRoutes() *gin.Engine {

	router := gin.Default()
	router.GET("getAppId", controller.GetAppId)

	return router
}

