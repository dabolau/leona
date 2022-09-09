package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// 版本模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type Version struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" validate:"required"`
	Version     string             `json:"version" validate:"required"`
	Description string             `json:"description" validate:"required"`
	URL         string             `json:"url" validate:"required"`
}
