package controller

import (
	"fmt"
	"log"

	"github.com/dabolau/leona/common"
	"github.com/dabolau/leona/database"
	"github.com/dabolau/leona/model"
	"github.com/dabolau/leona/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 账号集合
var accountCollection = database.GetCollection(database.MongoClient, "users")

// 帐号验证器
var accountValidate = validator.New()

// 账号信息
// https://docs.gofiber.io/api/app#route-handlers
func AccountHandler(c *fiber.Ctx) error {
	// 账号信息
	return c.Status(fiber.StatusOK).JSON(response.ResponseAccount{
		Message:    "帐号信息",
		StatusCode: fiber.StatusOK,
	})
}

// 账号登录
// https://docs.gofiber.io/api/app#route-handlers
func LoginHandler(c *fiber.Ctx) error {
	// 数据模型
	var account model.Account
	var requestAccount model.LoginAccount
	// 验证请求参数
	err := c.BodyParser(&requestAccount)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseAccount{
			Message:    "请求参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 使用验证器验证必填参数
	err = accountValidate.Struct(&requestAccount)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseAccount{
			Message:    "验证参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 查询数据
	err = accountCollection.FindOne(c.Context(), bson.M{"$or": bson.A{bson.M{"username": requestAccount.Username}, bson.M{"email": requestAccount.Username}, bson.M{"phone": requestAccount.Username}}, "password": requestAccount.Password}).Decode(&account)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseAccount{
			Message:    "登录失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 生成令牌
	token, err := common.GenerateToken(account.Username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseAccount{
			Message:    "生成失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 登录成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseIDAndBearerTokenAccount{
		ID:          account.ID,
		BearerToken: token,
		Message:     "登录成功",
		StatusCode:  fiber.StatusOK,
	})
}

// 账号注册
// https://docs.gofiber.io/api/app#route-handlers
func RegisterHandler(c *fiber.Ctx) error {
	// 数据模型
	var account model.Account
	var requestAccount model.RegisterAccount
	// 验证请求参数
	err := c.BodyParser(&requestAccount)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseAccount{
			Message:    "请求参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 使用验证器验证必填参数
	err = accountValidate.Struct(&requestAccount)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseAccount{
			Message:    "验证参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 查询数据
	err = accountCollection.FindOne(c.Context(), bson.M{"$or": bson.A{bson.M{"username": requestAccount.Username}, bson.M{"email": requestAccount.Email}, bson.M{"phone": requestAccount.Phone}}}).Decode(&account)
	if err == nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseAccount{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 新增编号
	requestAccount.ID = primitive.NewObjectID()
	// 添加数据
	result, err := userCollection.InsertOne(c.Context(), &requestAccount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseUser{
			Message:    "注册失败",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	log.Println(result.InsertedID)
	// 注册成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseIDAccount{
		ID:         requestAccount.ID,
		Message:    "注册成功",
		StatusCode: fiber.StatusOK,
	})
}

// 修改密码
// https://docs.gofiber.io/api/app#route-handlers
func ChangepasswordHandler(c *fiber.Ctx) error {
	// 数据模型
	var account model.Account
	var requestAccount model.ChangePasswordAccount
	// 验证请求参数
	err := c.BodyParser(&requestAccount)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseAccount{
			Message:    "请求参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 使用验证器验证必填参数
	err = accountValidate.Struct(&requestAccount)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseAccount{
			Message:    "验证参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 查询数据
	err = accountCollection.FindOne(c.Context(), bson.M{"username": requestAccount.Username, "password": requestAccount.Password}).Decode(&account)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseAccount{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 更新数据
	result, err := accountCollection.UpdateOne(c.Context(), bson.M{"username": requestAccount.Username, "password": requestAccount.Password}, bson.M{"$set": bson.M{"password": requestAccount.NewPassword}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseUser{
			Message:    "修改失败",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	log.Println(result.MatchedCount)
	// 修改成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseIDAccount{
		ID:         account.ID,
		Message:    "修改成功",
		StatusCode: fiber.StatusOK,
	})
}

// 找回密码
// https://docs.gofiber.io/api/app#route-handlers
func GetpasswordHandler(c *fiber.Ctx) error {
	// 数据模型
	var account model.Account
	var requestAccount model.GetPasswordAccount
	// 验证请求参数
	err := c.BodyParser(&requestAccount)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseAccount{
			Message:    "请求参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 使用验证器验证必填参数
	err = accountValidate.Struct(&requestAccount)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseAccount{
			Message:    "验证参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 查询数据
	err = accountCollection.FindOne(c.Context(), bson.M{"email": requestAccount.Email}).Decode(&account)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseAccount{
			Message:    "找回失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 发送文本邮件
	// err = common.SendTextEMail(account.Email, "找到您的信息", fmt.Sprintf(common.TextEMailTemplate, account.Username, account.Password))
	// 发送超文本标记邮件
	err = common.SendHTMLEMail(account.Email, "找到您的信息", fmt.Sprintf(common.HTMLEMailTemplate, account.Username, account.Password))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseErrorAccount{
			Error:      err,
			Message:    "发送失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 找回成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseIDAccount{
		ID:         account.ID,
		Message:    "找回成功",
		StatusCode: fiber.StatusOK,
	})
}

// 账号注销
// https://docs.gofiber.io/api/app#route-handlers
func LogoutHandler(c *fiber.Ctx) error {
	// 账号注销
	return c.Status(fiber.StatusOK).JSON(response.ResponseAccount{
		Message:    "注销成功",
		StatusCode: fiber.StatusOK,
	})
}
