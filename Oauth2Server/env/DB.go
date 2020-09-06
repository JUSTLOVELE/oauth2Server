package env

import (
	"Oauth2Server/constant"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var DBengine *xorm.Engine

func CreateDB () {
	//dbconf := "root:12345678@(127.0.0.1:3306)/mydb?charset=utf8"
	dbconf, e := Conf.GetValue(constant.Profile, constant.DBConf)

	fmt.Println(dbconf)
	if e != nil {
		Log.Panicln(e)
	}
	//db := "mysql"
	db, _ := Conf.GetValue(constant.Profile, constant.DB)
	var err error
	DBengine, err = xorm.NewEngine(db, dbconf)

	if err != nil {
		Log.Panicln(err)
	}
}

