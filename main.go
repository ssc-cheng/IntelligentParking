package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	"intelligent_parking/models"
	_ "intelligent_parking/routers"
	"intelligent_parking/server"
)


func init() {
	//设置logs输出
	log := logs.NewLogger(30000)
	log.SetLogger("console", "")
	logs.EnableFuncCallDepth(true)		//输出文件名
	logs.SetLogFuncCallDepth(3)
	logs.Async()						//异步输出日志
	logs.Async(1e3)			//设置缓冲 chan 的大小
	logs.SetLogger(logs.AdapterFile, `{"filename":"log/hm_project.log"}`)
}

func main() {
	//cors跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	models.Init()
	//初始化repo
	server.NewRepo()
	beego.Run()
}

