package controller

import (
	"fmt"
	"log"
	"strconv"

	"github.com/dabolau/leona/database"
	"github.com/dabolau/leona/model"
	"github.com/dabolau/leona/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 用户集合
var userCollection = database.GetCollection(database.MongoClient, "users")

// 用户验证器
var userValidate = validator.New()

// 用户信息
// https://docs.gofiber.io/api/app#route-handlers
func UserHandler(c *fiber.Ctx) error {
	// 数据模型
	var users []model.User
	// 分页信息
	var findOptions options.FindOptions
	var text string = ""
	var pageSize int64 = 10
	var page int64 = 1
	// 获取参数
	text = c.Query("text")
	pageSize, _ = strconv.ParseInt(c.Query("size"), 10, 64)
	page, _ = strconv.ParseInt(c.Query("page"), 10, 64)
	if pageSize == 0 {
		pageSize = 1
	}
	if page == 0 {
		page = 1
	}
	if pageSize > 0 {
		findOptions.SetSkip((page - 1) * pageSize)
		findOptions.SetLimit(pageSize)
	}
	// 查询数据
	cursor, err := userCollection.Find(c.Context(), bson.M{"Username": bson.M{"$regex": primitive.Regex{Pattern: fmt.Sprintf(".*%v.*", text), Options: "i"}}}, &findOptions)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseUser{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	defer cursor.Close(c.Context())
	// 获取所有数据
	err = cursor.All(c.Context(), &users)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseUser{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 查询成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseMultipleUser{
		Datas:      users,
		Message:    "查询成功",
		StatusCode: fiber.StatusOK,
	})
}

// 用户详情
// https://docs.gofiber.io/api/app#route-handlers
func UserDetailHandler(c *fiber.Ctx) error {
	// 数据模型
	var user model.User
	// 获取编号
	objectId, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseUser{
			Message:    "编号错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 查询数据
	err = userCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseUser{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 查询成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseSingleUser{
		Data:       user,
		Message:    "查询成功",
		StatusCode: fiber.StatusOK,
	})
}

// 用户添加
// https://docs.gofiber.io/api/app#route-handlers
func UserAddHandler(c *fiber.Ctx) error {
	// 数据模型
	var requestUser model.User
	// 验证请求参数
	err := c.BodyParser(&requestUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseUser{
			Message:    "请求参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 使用验证器验证必填参数
	err = userValidate.Struct(&requestUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseUser{
			Message:    "验证参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 新增编号
	requestUser.ID = primitive.NewObjectID()
	// 添加数据
	result, err := userCollection.InsertOne(c.Context(), &requestUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseUser{
			Message:    "添加失败",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	log.Println(result.InsertedID)
	// 添加成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseSingleUser{
		Data:       requestUser,
		Message:    "添加成功",
		StatusCode: fiber.StatusOK,
	})
}

// 用户编辑
// https://docs.gofiber.io/api/app#route-handlers
func UserChangeHandler(c *fiber.Ctx) error {
	// 数据模型
	var requestUser model.User
	var user model.User
	// 获取编号
	objectId, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseUser{
			Message:    "编号错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 验证请求参数
	err = c.BodyParser(&requestUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseUser{
			Message:    "请求参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 使用验证器验证必填参数
	err = userValidate.Struct(&requestUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseUser{
			Message:    "验证参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 查询数据
	err = userCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseUser{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 更新数据
	result, err := userCollection.UpdateOne(c.Context(), bson.M{"_id": objectId}, bson.M{"$set": requestUser})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseUser{
			Message:    "更新失败",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	log.Println(result)
	// 查询数据
	err = userCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseUser{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 更新成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseSingleUser{
		Data:       user,
		Message:    "更新成功",
		StatusCode: fiber.StatusOK,
	})
}

// 用户删除
// https://docs.gofiber.io/api/app#route-handlers
func UserDeleteHandler(c *fiber.Ctx) error {
	// 数据模型
	var user model.User
	// 获取编号
	objectId, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseUser{
			Message:    "编号错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 查询数据
	err = userCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseUser{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 删除数据
	result, err := userCollection.DeleteOne(c.Context(), bson.M{"_id": objectId})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseUser{
			Message:    "删除失败",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	log.Println(result.DeletedCount)
	// 删除成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseSingleUser{
		Data:       user,
		Message:    "删除成功",
		StatusCode: fiber.StatusOK,
	})
}
