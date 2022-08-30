package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// 用户模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" validate:"required"`
	Password string             `json:"password" validate:"required"`
	Email    string             `json:"email" validate:"required,email"`
	Phone    string             `json:"phone" validate:"required"`
}
