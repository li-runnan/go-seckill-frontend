package main

import (
	"github.com/astaxie/beego"
	_ "myseckill/routers"
	"myseckill/util"
)

func main() {
	// 初始化系统时先把秒杀商品库存放入redis缓存
	util.InitClient()
	util.SetTicketsInfoToRedis()

	// 启动消息队列，持续监听，接收订单请求
	// util.ReceiveMessage()

	go func() {
		util.ReceiveMessage()
	}()

	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.Run()
}

