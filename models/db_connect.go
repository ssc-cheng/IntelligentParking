package models

import (
	"github.com/astaxie/beego/logs"
	"intelligent_parking/db"
)

func Init()  {

	var err error
	// 连接mysql
	if err = db.ConnectMysql();err != nil{
		panic(err)
	}
	//db.GetLocalDB().DropTables(Card{},User{},UserCardbag{},UserConsume{},UserVehicle{},UserIntegral{},UserLogin{},UserNews{})
	// 同步表数据
	if err = db.GetLocalDB().Sync2(Card{},User{},UserCardbag{},UserConsume{},UserVehicle{},UserIntegral{},UserLogin{},UserNews{},Vehicle{});err != nil{
		logs.Error(err)
	}
}
