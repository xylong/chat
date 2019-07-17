package service

import (
	"chat/model"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var Db *xorm.Engine

func init() {
	driver := "mysql"
	database := "root:root@(127.0.0.1:3306)/chat?charset=utf8"

	err := errors.New("")
	Db, err = xorm.NewEngine(driver, database)
	if nil != err && "" != err.Error() {
		log.Fatal(err.Error())
	}
	//是否显示sql
	Db.ShowSQL(true)
	Db.SetMaxOpenConns(2)
	Db.Sync2(new(model.User))
	fmt.Println("init database")
}
