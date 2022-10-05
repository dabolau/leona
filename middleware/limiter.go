package middleware

import (
	"time"

	"github.com/dabolau/leona/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// 速率限制中间件
// https://docs.gofiber.io/api/middleware/limiter
func RateLimiter(app *fiber.App) {
	app.Use(limiter.New(limiter.Config{
		// 定义函数，在返回true时跳过此中间件
		// Next: func(c *fiber.Ctx) bool {
		// 	return c.IP() == "127.0.0.1"
		// },
		// 过期时间内的最大连接数
		Max: 5,
		// 过期时间，指在内存中保留请求记录的时间
		Expiration: 1 * time.Minute,
		// 密钥生成器，允许你生成自定义密钥
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		// 当请求达到限制时调用
		LimitReached: func(c *fiber.Ctx) error {
			// return c.SendStatus(fiber.StatusTooManyRequests)
			return c.Status(fiber.StatusTooManyRequests).JSON(response.ResponseHome{
				Message:    "太多请求",
				StatusCode: fiber.StatusTooManyRequests,
			})
		},
		// 当设置为true时，状态码大于等于400的请求将不计算在内
		SkipFailedRequests: false,
		// 当设置为true时，状态码小于400的请求将不计算在内
		SkipSuccessfulRequests: false,
		// 固定窗口速率限制算法
		// LimiterMiddleware: limiter.FixedWindow{},
		// 滑动窗口速率限制算法
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
}
