package route

import (
	"github.com/dabolau/leona/controller"

	"github.com/gofiber/fiber/v2"
)

// 反馈路由
// https://docs.gofiber.io/api/app#group
// https://docs.gofiber.io/api/app#route
func Feedback(app *fiber.App) {
	// 反馈分组
	feedbackGroup := app.Group("/feedback")
	// 反馈信息
	feedbackGroup.Get("/", controller.FeedbackHandler)
	// 反馈详情
	feedbackGroup.Get("/detail", controller.FeedbackDetailHandler)
	// 反馈添加
	feedbackGroup.Post("/add", controller.FeedbackAddHandler)
	// 反馈编辑
	feedbackGroup.Put("/change", controller.FeedbackChangeHandler)
	// 反馈删除
	feedbackGroup.Delete("/delete", controller.FeedbackDeleteHandler)
}
