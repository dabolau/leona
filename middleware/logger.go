package middleware

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

// 日志中间件
// https://docs.gofiber.io/api/middleware/logger
func Logger(app *fiber.App) {
	app.Use(logger.New(logger.ConfigDefault))
}
