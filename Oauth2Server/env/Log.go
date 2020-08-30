package env

import (
	"Oauth2Server/constant"
	"fmt"
	"io"
	"log"
	"os"
)

var Log *log.Logger

func InitLog() {

	//address := "/Users/yangzuliang/Documents/develop/github/Oauth2Server/Oauth2Server/log.log"
	//address := "./log.log"
	address, _ := Conf.GetValue(constant.Profile, constant.Log)
	infoFile,err:=os.OpenFile(address,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)

	if err!=nil{
		fmt.Println("失败")
		log.Fatalln("打开日志文件失败：",err)
		return
	}

	Log = log.New(os.Stdout,"Info:",log.Ldate | log.Ltime | log.Lshortfile)
	Log = log.New(io.MultiWriter(os.Stderr,infoFile),"Info:",log.Ldate | log.Ltime | log.Lshortfile)
	Log.Println("创建日志成功!")
}