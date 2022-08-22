package controller

import (
	"github.com/dabolau/leona/response"
	"github.com/gofiber/fiber/v2"
)

// 首页信息
// https://docs.gofiber.io/api/app#route-handlers
func HomeHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(response.ResponseHome{
		Message:    "接口程序1.0",
		StatusCode: fiber.StatusOK,
	})
}
