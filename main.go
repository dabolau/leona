package main

import (
	"github.com/dabolau/leona/middleware"
	"github.com/dabolau/leona/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// 获取应用实例
	app := fiber.New()
	// 使用日志中间件
	middleware.Logger(app)
	// 使用跨域中间件
	middleware.CORS(app)
	// 使用限制中间件
	// middleware.Limiter(app)
	// 获取路由
	route.Home(app)
	route.Video(app)
	route.Feedback(app)
	route.Version(app)
	route.Account(app)
	route.User(app)
	// 监听服务
	app.Listen("0.0.0.0:8081")
}
