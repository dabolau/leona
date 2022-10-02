package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// 视频模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type Video struct {
	ID          primitive.ObjectID `json:"ID,omitempty" bson:"_id,omitempty"`    // 编号
	CreatedAt   string             `json:"CreatedAt" bson:"CreatedAt"`           // 创建时间
	UpdatedAt   string             `json:"UpdatedAt" bson:"UpdatedAt"`           // 更新时间
	DeletedAt   string             `json:"DeletedAt" bson:"DeletedAt"`           // 删除时间
	Name        string             `json:"Name" bson:"Name" validate:"required"` // 名称
	Category    string             `json:"Category" bson:"Category"`             // 类别
	Type        string             `json:"Type" bson:"Type"`                     // 类型
	Area        string             `json:"Area" bson:"Area"`                     // 地区
	Language    string             `json:"Language" bson:"Language"`             // 语言
	Director    string             `json:"Director" bson:"Director"`             // 导演
	Actor       string             `json:"Actor" bson:"Actor"`                   // 演员
	Year        string             `json:"Year" bson:"Year"`                     // 年份
	Score       string             `json:"Score" bson:"Score"`                   // 评分
	Status      string             `json:"Status" bson:"Status"`                 // 状态
	Description string             `json:"Description" bson:"Description"`       // 描述
	Picture     string             `json:"Picture" bson:"Picture"`               // 图片
	URLs        []URL              `json:"URLs" bson:"URLs"`                     // 资源
}

// 资源模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type URL struct {
	Name string `json:"Name" bson:"Name"` // 名称
	URL  string `json:"URL" bson:"URL"`   // 资源
}
