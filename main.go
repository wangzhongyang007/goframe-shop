package main

import (
	"github.com/gogf/gf/frame/g"
	"shop/app/middleware"
	"shop/app/system/backend"
	"shop/app/system/frontend"
	_ "shop/boot"
	_ "shop/router"
)

func main() {
	//现在的写法 把app.Run内的方法直接写到main方法中执行。
	s := g.Server()
	s.Use(middleware.Cors.CORS)
	//后台项目
	backend.Init(s)
	//前端项目
	frontend.Init(s)
	s.Run()
}
