package db

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/go-xorm/xorm"
	"time"
)

/*
连接mysql数据库
*/

var (
	localEngin *xorm.Engine //全局数据库连接
)

func ConnectMysql() (err error) {

	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	localEngin, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		logs.Error(err.Error())
		return
	}

	localEngin.DB().SetMaxIdleConns(10)
	localEngin.DB().SetMaxOpenConns(100)
	localEngin.DatabaseTZ = time.Local // 必须
	localEngin.TZLocation = time.Local // 必须

	return

}


func GetLocalDB() *xorm.Engine {
	return localEngin
}