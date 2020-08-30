package test

import (
	"Oauth2Server/env"
	"fmt"
)

func DBTest() {
	sql := "select * from teacher"
	results, err := env.DBengine.Query(sql)

	if err != nil {
		panic(err)
	}

	for _, v := range results {
		fmt.Println("教师ID: %s, 姓名: %s\n", string(v["op_id"]), string(v["teacher_name"]))
	}
}
