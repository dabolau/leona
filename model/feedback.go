package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// 反馈模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type Feedback struct {
	ID          primitive.ObjectID `json:"ID,omitempty" bson:"_id,omitempty"`                  // 编号
	Name        string             `json:"Name" bson:"Name" validate:"required"`               // 名称
	Description string             `json:"Description" bson:"Description" validate:"required"` // 描述
	Email       string             `json:"Email" bson:"Email" validate:"required,email"`       // 邮件
}
