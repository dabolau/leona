package route

import (
	"github.com/dabolau/leona/controller"
	"github.com/gofiber/fiber/v2"
)

// 首页路由
// https://docs.gofiber.io/api/app#group
// https://docs.gofiber.io/api/app#route
func Home(app *fiber.App) {
	// 首页分组
	homeGroup := app.Group("/")
	// 首页信息
	homeGroup.Get("/", controller.HomeHandler)
}
