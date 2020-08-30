package env

import (
	"github.com/Unknwon/goconfig"
)

var Conf *goconfig.ConfigFile

func InitConf()  {

	Conf, _ = goconfig.LoadConfigFile("conf.conf")
	//value, _ := cfg.GetValue("dev", "log")
	//fmt.Println(value)
}
