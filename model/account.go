package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// 帐号模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type Account struct {
	ID       primitive.ObjectID `json:"ID,omitempty" bson:"_id,omitempty"`                         // 编号
	Username string             `json:"Username" bson:"Username" validate:"required,min=6,max=32"` // 用户
	Password string             `json:"Password" bson:"Password" validate:"required,min=6,max=32"` // 密码
	Email    string             `json:"Email" bson:"Email" validate:"required,email"`              // 邮件
	Phone    string             `json:"Phone" bson:"Phone" validate:"required,min=11,max=11"`      // 电话
}

// 登录模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type LoginAccount struct {
	Username string `json:"Username" bson:"Username" validate:"required,min=6,max=32"` // 用户
	Password string `json:"Password" bson:"Password" validate:"required,min=6,max=32"` // 密码
}

// 注册模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type RegisterAccount struct {
	ID       primitive.ObjectID `json:"ID,omitempty" bson:"_id,omitempty"`                         // 编号
	Username string             `json:"Username" bson:"Username" validate:"required,min=6,max=32"` // 用户
	Password string             `json:"Password" bson:"Password" validate:"required,min=6,max=32"` // 密码
	Email    string             `json:"Email" bson:"Email" validate:"required,email"`              // 邮件
	Phone    string             `json:"Phone" bson:"Phone" validate:"required,min=11,max=11"`      // 电话
}

// 修改密码模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type ChangePasswordAccount struct {
	Username        string `json:"Username" bson:"Username" validate:"required,min=6,max=32"`                       // 用户
	Password        string `json:"Password" bson:"Password" validate:"required,min=6,max=32"`                       // 密码
	NewPassword     string `json:"NewPassword" bson:"NewPassword"  validate:"required,min=6,max=32"`                // 新的密码
	ConfirmPassword string `json:"ConfirmPassword" bson:"ConfirmPassword"  validate:"required,eqfield=NewPassword"` // 验证密码
}

// 找回密码模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type GetPasswordAccount struct {
	Email string `json:"Email" bson:"Email" validate:"required,email"` // 邮件

}
