package middleware

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

// 跨域中间件
// https://docs.gofiber.io/api/middleware/cors
func CORS(app *fiber.App) {
	app.Use(cors.New(cors.ConfigDefault))
}
