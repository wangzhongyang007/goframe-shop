package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"shop/app/middleware"
	"shop/app/system/backend"
	"shop/app/system/frontend"
	_ "shop/boot"
	_ "shop/router"
)

func main() {
	//之前的写法 在Linux服务器上无法编译通过 提示gcmd.AutoRun()方法不存在
	//err := gcmd.BindHandle("server", app.Run)
	//if err != nil {
	//	glog.Error("gcmd 启动server报错：", err)
	//	return
	//}
	//err = gcmd.AutoRun()
	//if err != nil {
	//	glog.Error("gcmd AutoRun 报错：", err)
	//	return
	//}

	//现在的写法 把app.Run内的方法直接写到main方法中执行。
	s := g.Server()
	s.Use(middleware.Cors.CORS)
	//后台项目
	backend.Init(s)
	//前端项目
	frontend.Init(s)
	s.Run()
}
