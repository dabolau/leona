package route

import (
	"github.com/dabolau/leona/controller"

	"github.com/gofiber/fiber/v2"
)

// 帐号路由
// https://docs.gofiber.io/api/app#group
// https://docs.gofiber.io/api/app#route
func Account(app *fiber.App) {
	// 帐号分组
	accountGroup := app.Group("/account")
	// 帐号信息
	accountGroup.Get("/", controller.AccountHandler)
	// 帐号登录
	accountGroup.Post("/login", controller.LoginHandler)
	// 帐号注册
	accountGroup.Post("/register", controller.RegisterHandler)
	// 修改密码
	accountGroup.Post("/changepassword", controller.ChangepasswordHandler)
	// 找回密码
	accountGroup.Post("/getpassword", controller.GetpasswordHandler)
	// 帐号注销
	accountGroup.Get("/logout", controller.LogoutHandler)
}
