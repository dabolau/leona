package middleware

import (
	"log"
	"strings"

	"github.com/dabolau/leona/common"
	"github.com/dabolau/leona/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// 请求头结构体
type Headers struct {
	Authorization string `reqHeader:"authorization"`
}

// 中间件验证器
var MiddlewareValidate = validator.New()

// 用户授权中间件
// https://docs.gofiber.io/guide/routing#middleware
// https://github.com/golang-jwt/jwt/v4
// 登录成功时生成令牌
// token, err := common.GenerateToken("username")
// 用户授权时解析令牌
// claims, err := common.ParseToken("token")
func AuthToken(c *fiber.Ctx) error {
	// 数据模型
	var headers = new(Headers)
	// 解析请求头
	err := c.ReqHeaderParser(headers)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ResponseHome{
			Message:    "授权失败",
			StatusCode: fiber.StatusUnauthorized,
		})
	}
	// 获取请求头中的完整令牌
	var bearerToken = headers.Authorization
	// 判断令牌是否为空
	if bearerToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ResponseAccount{
			Message:    "授权失败",
			StatusCode: fiber.StatusUnauthorized,
		})
	}
	// 判断令牌格式是否正确
	if !strings.Contains(bearerToken, "Bearer") {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ResponseAccount{
			Message:    "授权失败",
			StatusCode: fiber.StatusUnauthorized,
		})
	}
	// 获取令牌
	var token = bearerToken[7:]
	// 使用验证器验证验证令牌是否正确
	err = MiddlewareValidate.Var(token, "jwt")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ResponseAccount{
			Message:    "授权失败",
			StatusCode: fiber.StatusUnauthorized,
		})
	}
	// 解析令牌
	claims, err := common.ParseToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ResponseAccount{
			Message:    "授权失败",
			StatusCode: fiber.StatusUnauthorized,
		})
	}
	// 授权成功
	log.Println(claims.Username)
	log.Println(claims.ExpiresAt)
	// 授权成功，往下执行
	return c.Next()
}
