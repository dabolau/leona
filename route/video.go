package route

import (
	"github.com/dabolau/leona/controller"

	"github.com/gofiber/fiber/v2"
)

// 视频路由
// https://docs.gofiber.io/api/app#group
// https://docs.gofiber.io/api/app#route
func Video(app *fiber.App) {
	// 视频分组
	videoGroup := app.Group("/video")
	// 视频信息
	videoGroup.Get("/", controller.VideoHandler)
	// 视频详情
	videoGroup.Get("/detail", controller.VideoDetailHandler)
	// 视频添加
	videoGroup.Post("/add", controller.VideoAddHandler)
	// 视频编辑
	videoGroup.Put("/change", controller.VideoChangeHandler)
	// 视频删除
	videoGroup.Delete("/delete", controller.VideoDeleteHandler)
}
