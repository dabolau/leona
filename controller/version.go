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

// 版本集合
var versionCollection = database.GetCollection(database.MongoClient, "versions")

// 验证器
var validateVersion = validator.New()

// 版本信息
// https://docs.gofiber.io/api/app#route-handlers
func VersionHandler(c *fiber.Ctx) error {
	// 数据模型
	var versions []model.Version
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
	cursor, err := versionCollection.Find(c.Context(), bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: fmt.Sprintf(".*%v.*", text), Options: "i"}}}, &findOptions)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseVersion{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	defer cursor.Close(c.Context())
	// 获取所有数据
	err = cursor.All(c.Context(), &versions)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseVersion{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 查询成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseMultipleVersion{
		Datas:      versions,
		Message:    "查询成功",
		StatusCode: fiber.StatusOK,
	})
}

// 版本详情
// https://docs.gofiber.io/api/app#route-handlers
func VersionDetailHandler(c *fiber.Ctx) error {
	// 数据模型
	var version model.Version
	// 获取编号
	objectId, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseVersion{
			Message:    "编号错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 查询数据
	err = versionCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&version)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseVersion{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 查询成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseSingleVersion{
		Data:       version,
		Message:    "查询成功",
		StatusCode: fiber.StatusOK,
	})
}

// 版本添加
// https://docs.gofiber.io/api/app#route-handlers
func VersionAddHandler(c *fiber.Ctx) error {
	// 数据模型
	var requestVersion model.Version
	// 验证请求参数
	err := c.BodyParser(&requestVersion)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseVersion{
			Message:    "请求参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 使用验证器验证必填参数
	err = validateVersion.Struct(&requestVersion)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseVersion{
			Message:    "验证参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 新增编号
	requestVersion.ID = primitive.NewObjectID()
	// 添加数据
	result, err := versionCollection.InsertOne(c.Context(), &requestVersion)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseVersion{
			Message:    "添加失败",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	log.Println(result.InsertedID)
	// 添加成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseSingleVersion{
		Data:       requestVersion,
		Message:    "添加成功",
		StatusCode: fiber.StatusOK,
	})
}

// 版本编辑
// https://docs.gofiber.io/api/app#route-handlers
func VersionChangeHandler(c *fiber.Ctx) error {
	// 数据模型
	var requestVersion model.Version
	var version model.Version
	// 获取编号
	objectId, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseVersion{
			Message:    "编号错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 验证请求参数
	err = c.BodyParser(&requestVersion)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseVersion{
			Message:    "请求参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 使用验证器验证必填参数
	err = validateVersion.Struct(&requestVersion)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseVersion{
			Message:    "验证参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 查询数据
	err = versionCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&version)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseVersion{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 更新数据
	result, err := versionCollection.UpdateOne(c.Context(), bson.M{"_id": objectId}, bson.M{"$set": requestVersion})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseVersion{
			Message:    "更新失败",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	log.Println(result)
	// 查询数据
	err = versionCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&version)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseVersion{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 更新成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseSingleVersion{
		Data:       version,
		Message:    "更新成功",
		StatusCode: fiber.StatusOK,
	})
}

// 版本删除
// https://docs.gofiber.io/api/app#route-handlers
func VersionDeleteHandler(c *fiber.Ctx) error {
	// 数据模型
	var version model.Version
	// 获取编号
	objectId, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseVersion{
			Message:    "编号错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 查询数据
	err = versionCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&version)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseVersion{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 删除数据
	result, err := versionCollection.DeleteOne(c.Context(), bson.M{"_id": objectId})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseVersion{
			Message:    "删除失败",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	log.Println(result.DeletedCount)
	// 删除成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseSingleVersion{
		Data:       version,
		Message:    "删除成功",
		StatusCode: fiber.StatusOK,
	})
}
