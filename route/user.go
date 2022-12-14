package route

import (
	"github.com/dabolau/leona/controller"
	"github.com/dabolau/leona/middleware"

	"github.com/gofiber/fiber/v2"
)

// 用户路由
// https://docs.gofiber.io/api/app#group
// https://docs.gofiber.io/api/app#route
func User(app *fiber.App) {
	// 用户分组
	userGroup := app.Group("/user", middleware.AuthToken)
	// 用户信息
	userGroup.Get("/", controller.UserHandler)
	// 用户详情
	userGroup.Get("/detail", controller.UserDetailHandler)
	// 用户添加
	userGroup.Post("/add", controller.UserAddHandler)
	// 用户编辑
	userGroup.Put("/change", controller.UserChangeHandler)
	// 用户删除
	userGroup.Delete("/delete", controller.UserDeleteHandler)
}
