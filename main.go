package main

import (
	"log"

	"github.com/dabolau/leona/common"
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
	route.Account(app)
	route.User(app)

	jwt, err := common.GenerateToken("dabolau")
	if err != nil {
		log.Println(err)
	}
	log.Println(jwt)
	t := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRhYm9sYXUiLCJleHAiOjE2NjYwOTIyNDV9.2opsLKFsKWEYxs0S3M5Y3P58K2lX0nCEDKiggzBgAbM"
	claims, err := common.ParseToken(t)
	if err != nil {
		log.Println(err)
	}
	log.Println(claims)

	// 监听服务
	app.Listen("0.0.0.0:8081")
}
