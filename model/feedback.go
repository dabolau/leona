package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// 反馈模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type Feedback struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" validate:"required"`
	Description string             `json:"description" validate:"required"`
	Email       string             `json:"email" validate:"required,email"`
}
