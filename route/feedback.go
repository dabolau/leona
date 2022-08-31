package route

import (
	"github.com/dabolau/leona/controller"

	"github.com/gofiber/fiber/v2"
)

// 反馈路由
// https://docs.gofiber.io/api/app#group
// https://docs.gofiber.io/api/app#route
func Feedback(app *fiber.App) {
	// 用户分组
	feedbackGroup := app.Group("/feedback")
	// 用户信息
	feedbackGroup.Get("/", controller.FeedbackHandler)
	// 用户详情
	feedbackGroup.Get("/detail", controller.FeedbackDetailHandler)
	// 用户添加
	feedbackGroup.Post("/add", controller.FeedbackAddHandler)
	// 用户编辑
	feedbackGroup.Put("/change", controller.FeedbackChangeHandler)
	// 用户删除
	feedbackGroup.Delete("/delete", controller.FeedbackDeleteHandler)
}
