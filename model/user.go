package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// 用户模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type User struct {
	ID       primitive.ObjectID `json:"ID,omitempty" bson:"_id,omitempty"`            // 编号
	Username string             `json:"Username" bson:"Username" validate:"required"` // 用户
	Password string             `json:"Password" bson:"Password" validate:"required"` // 密码
	Email    string             `json:"Email" bson:"Email" validate:"required,email"` // 邮件
	Phone    string             `json:"Phone" bson:"Phone" validate:"required"`       // 电话
}
