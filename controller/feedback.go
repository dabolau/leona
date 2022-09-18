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

// 反馈集合
var feedbackCollection = database.GetCollection(database.MongoClient, "feedbacks")

// 反馈验证器
var feedbackValidate = validator.New()

// 反馈信息
// https://docs.gofiber.io/api/app#route-handlers
func FeedbackHandler(c *fiber.Ctx) error {
	// 数据模型
	var feedbacks []model.Feedback
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
	cursor, err := feedbackCollection.Find(c.Context(), bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: fmt.Sprintf(".*%v.*", text), Options: "i"}}}, &findOptions)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseFeedback{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	defer cursor.Close(c.Context())
	// 获取所有数据
	err = cursor.All(c.Context(), &feedbacks)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseFeedback{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 查询成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseMultipleFeedback{
		Datas:      feedbacks,
		Message:    "查询成功",
		StatusCode: fiber.StatusOK,
	})
}

// 反馈详情
// https://docs.gofiber.io/api/app#route-handlers
func FeedbackDetailHandler(c *fiber.Ctx) error {
	// 数据模型
	var feedback model.Feedback
	// 获取编号
	objectId, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseFeedback{
			Message:    "编号错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 查询数据
	err = feedbackCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&feedback)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseFeedback{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 查询成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseSingleFeedback{
		Data:       feedback,
		Message:    "查询成功",
		StatusCode: fiber.StatusOK,
	})
}

// 反馈添加
// https://docs.gofiber.io/api/app#route-handlers
func FeedbackAddHandler(c *fiber.Ctx) error {
	// 数据模型
	var requestFeedback model.Feedback
	// 验证请求参数
	err := c.BodyParser(&requestFeedback)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseFeedback{
			Message:    "请求参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 使用验证器验证必填参数
	err = feedbackValidate.Struct(&requestFeedback)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseFeedback{
			Message:    "验证参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 新增编号
	requestFeedback.ID = primitive.NewObjectID()
	// 添加数据
	result, err := feedbackCollection.InsertOne(c.Context(), &requestFeedback)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseFeedback{
			Message:    "添加失败",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	log.Println(result.InsertedID)
	// 添加成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseSingleFeedback{
		Data:       requestFeedback,
		Message:    "添加成功",
		StatusCode: fiber.StatusOK,
	})
}

// 反馈编辑
// https://docs.gofiber.io/api/app#route-handlers
func FeedbackChangeHandler(c *fiber.Ctx) error {
	// 数据模型
	var requestFeedback model.Feedback
	var feedback model.Feedback
	// 获取编号
	objectId, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseFeedback{
			Message:    "编号错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 验证请求参数
	err = c.BodyParser(&requestFeedback)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseFeedback{
			Message:    "请求参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 使用验证器验证必填参数
	err = feedbackValidate.Struct(&requestFeedback)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseFeedback{
			Message:    "验证参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 查询数据
	err = feedbackCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&feedback)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseFeedback{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 更新数据
	result, err := feedbackCollection.UpdateOne(c.Context(), bson.M{"_id": objectId}, bson.M{"$set": requestFeedback})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseFeedback{
			Message:    "更新失败",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	log.Println(result)
	// 查询数据
	err = feedbackCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&feedback)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseFeedback{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 更新成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseSingleFeedback{
		Data:       feedback,
		Message:    "更新成功",
		StatusCode: fiber.StatusOK,
	})
}

// 反馈删除
// https://docs.gofiber.io/api/app#route-handlers
func FeedbackDeleteHandler(c *fiber.Ctx) error {
	// 数据模型
	var feedback model.Feedback
	// 获取编号
	objectId, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseFeedback{
			Message:    "编号错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 查询数据
	err = feedbackCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&feedback)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseFeedback{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 删除数据
	result, err := feedbackCollection.DeleteOne(c.Context(), bson.M{"_id": objectId})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseFeedback{
			Message:    "删除失败",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	log.Println(result.DeletedCount)
	// 删除成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseSingleFeedback{
		Data:       feedback,
		Message:    "删除成功",
		StatusCode: fiber.StatusOK,
	})
}
