package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// 版本模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type Version struct {
	ID          primitive.ObjectID `json:"ID,omitempty" bson:"_id,omitempty"`                  // 编号
	Name        string             `json:"Name" bson:"Name" validate:"required"`               // 名称
	Version     string             `json:"Version" bson:"Version" validate:"required"`         // 版本
	Description string             `json:"Description" bson:"Description" validate:"required"` // 描述
	URL         string             `json:"URL" bson:"URL" validate:"required"`                 // 资源
}
