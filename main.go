package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"shop/api/backend"
	"shop/api/frontend"
	"shop/middleware"
)

func main() {
	s := g.Server()
	s.Use(middleware.Cors.CORS)
	//后台项目
	backend.Init(s)
	//前端项目
	frontend.Init(s)
	s.Run()
}
