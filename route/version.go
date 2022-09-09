package route

import (
	"github.com/dabolau/leona/controller"

	"github.com/gofiber/fiber/v2"
)

// 版本路由
// https://docs.gofiber.io/api/app#group
// https://docs.gofiber.io/api/app#route
func Version(app *fiber.App) {
	// 版本分组
	versionGroup := app.Group("/version")
	// 版本信息
	versionGroup.Get("/", controller.VersionHandler)
	// 版本详情
	versionGroup.Get("/detail", controller.VersionDetailHandler)
	// 版本添加
	versionGroup.Post("/add", controller.VersionAddHandler)
	// 版本编辑
	versionGroup.Put("/change", controller.VersionChangeHandler)
	// 版本删除
	versionGroup.Delete("/delete", controller.VersionDeleteHandler)
}
