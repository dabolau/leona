package main

import (
	"github.com/dabolau/leona/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// 获取应用实例
	app := fiber.New()
	// 使用日志中间件
	app.Use(logger.New())
	// 获取路由
	route.Home(app)
	route.Feedback(app)
	route.Version(app)
	route.User(app)
	// 监听服务
	app.Listen("0.0.0.0:8081")
}
