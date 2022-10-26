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

// 视频集合
var videoCollection = database.GetCollection(database.MongoClient, "videos")

// 视频验证器
var videoValidate = validator.New()

// 视频信息
// https://docs.gofiber.io/api/app#route-handlers
func VideoHandler(c *fiber.Ctx) error {
	// 数据模型
	var videos []model.Video
	// 分页信息
	var findOptions options.FindOptions
	var text string = ""
	var videoCategory string = ""
	var videoType string = ""
	var videoArea string = ""
	var videoLanguage string = ""
	var videoYear string = ""
	var pageSize int64 = 10
	var page int64 = 1
	// 获取参数
	text = c.Query("text")
	videoCategory = c.Query("category")
	videoType = c.Query("type")
	videoArea = c.Query("area")
	videoLanguage = c.Query("language")
	videoYear = c.Query("year")
	pageSize, _ = strconv.ParseInt(c.Query("size"), 10, 64)
	page, _ = strconv.ParseInt(c.Query("page"), 10, 64)
	if pageSize == 0 {
		pageSize = 1
	}
	if page == 0 {
		page = 1
	}
	if pageSize > 0 {
		findOptions.SetSort(bson.M{"UpdatedAt": -1}) // 排序
		findOptions.SetSkip((page - 1) * pageSize)   // 分页
		findOptions.SetLimit(pageSize)               // 分页
	}
	// 查询数据
	cursor, err := videoCollection.Find(c.Context(), bson.M{
		"Name":     bson.M{"$regex": primitive.Regex{Pattern: fmt.Sprintf(".*%v.*", text), Options: "i"}},
		"Category": bson.M{"$regex": primitive.Regex{Pattern: fmt.Sprintf(".*%v.*", videoCategory), Options: "i"}},
		"Type":     bson.M{"$regex": primitive.Regex{Pattern: fmt.Sprintf(".*%v.*", videoType), Options: "i"}},
		"Area":     bson.M{"$regex": primitive.Regex{Pattern: fmt.Sprintf(".*%v.*", videoArea), Options: "i"}},
		"Language": bson.M{"$regex": primitive.Regex{Pattern: fmt.Sprintf(".*%v.*", videoLanguage), Options: "i"}},
		"Year":     bson.M{"$regex": primitive.Regex{Pattern: fmt.Sprintf(".*%v.*", videoYear), Options: "i"}},
	}, &findOptions)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseVideo{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	defer cursor.Close(c.Context())
	// 获取所有数据
	err = cursor.All(c.Context(), &videos)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseVideo{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 查询成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseMultipleVideo{
		Datas:      videos,
		Message:    "查询成功",
		StatusCode: fiber.StatusOK,
	})
}

// 视频详情
// https://docs.gofiber.io/api/app#route-handlers
func VideoDetailHandler(c *fiber.Ctx) error {
	// 数据模型
	var video model.Video
	// 获取编号
	objectId, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseVideo{
			Message:    "编号错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 查询数据
	err = videoCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&video)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseVideo{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 查询成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseSingleVideo{
		Data:       video,
		Message:    "查询成功",
		StatusCode: fiber.StatusOK,
	})
}

// 视频添加
// https://docs.gofiber.io/api/app#route-handlers
func VideoAddHandler(c *fiber.Ctx) error {
	// 数据模型
	var requestVideo model.Video
	// 验证请求参数
	err := c.BodyParser(&requestVideo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseVideo{
			Message:    "请求参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 使用验证器验证必填参数
	err = videoValidate.Struct(&requestVideo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseVideo{
			Message:    "验证参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 新增编号
	requestVideo.ID = primitive.NewObjectID()
	// 添加数据
	result, err := videoCollection.InsertOne(c.Context(), &requestVideo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseVideo{
			Message:    "添加失败",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	log.Println(result.InsertedID)
	// 添加成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseSingleVideo{
		Data:       requestVideo,
		Message:    "添加成功",
		StatusCode: fiber.StatusOK,
	})
}

// 视频编辑
// https://docs.gofiber.io/api/app#route-handlers
func VideoChangeHandler(c *fiber.Ctx) error {
	// 数据模型
	var requestVideo model.Video
	var video model.Video
	// 获取编号
	objectId, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseVideo{
			Message:    "编号错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 验证请求参数
	err = c.BodyParser(&requestVideo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseVideo{
			Message:    "请求参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 使用验证器验证必填参数
	err = videoValidate.Struct(&requestVideo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseVideo{
			Message:    "验证参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 查询数据
	err = videoCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&video)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseVideo{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 更新数据
	result, err := videoCollection.UpdateOne(c.Context(), bson.M{"_id": objectId}, bson.M{"$set": requestVideo})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseVideo{
			Message:    "更新失败",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	log.Println(result)
	// 查询数据
	err = videoCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&video)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseVideo{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 更新成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseSingleVideo{
		Data:       video,
		Message:    "更新成功",
		StatusCode: fiber.StatusOK,
	})
}

// 视频删除
// https://docs.gofiber.io/api/app#route-handlers
func VideoDeleteHandler(c *fiber.Ctx) error {
	// 数据模型
	var video model.Video
	// 获取编号
	objectId, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseVideo{
			Message:    "编号错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 查询数据
	err = videoCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&video)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseVideo{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 删除数据
	result, err := videoCollection.DeleteOne(c.Context(), bson.M{"_id": objectId})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseVideo{
			Message:    "删除失败",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	log.Println(result.DeletedCount)
	// 删除成功
	return c.Status(fiber.StatusOK).JSON(response.ResponseSingleVideo{
		Data:       video,
		Message:    "删除成功",
		StatusCode: fiber.StatusOK,
	})
}
