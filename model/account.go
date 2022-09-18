package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// 帐号模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type Account struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" validate:"required,min=6,max=32"`
	Email    string             `json:"email" validate:"required,email"`
	Phone    string             `json:"phone" validate:"required,min=11,max=11"`
	Password string             `json:"password" validate:"required,min=6,max=32"`
}

// 登录模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type LoginAccount struct {
	Username string `json:"username" validate:"required,min=6,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}

// 注册模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type RegisterAccount struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" validate:"required,min=6,max=32"`
	Email    string             `json:"email" validate:"required,email"`
	Phone    string             `json:"phone" validate:"required,min=11,max=11"`
	Password string             `json:"password" validate:"required,min=6,max=32"`
}

// 修改密码模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type ChangePasswordAccount struct {
	Username        string `json:"username" validate:"required,min=6,max=32"`
	Password        string `json:"password" validate:"required,min=6,max=32"`
	NewPassword     string `json:"newpassword"  validate:"required,min=6,max=32"`
	ConfirmPassword string `json:"confirmpassword"  validate:"required,eqfield=NewPassword"`
}

// 找回密码模型
// https://docs.gofiber.io/api/ctx#bodyparser
// https://github.com/go-playground/validator
type GetPasswordAccount struct {
	Email string `json:"email" validate:"required,email"`
}
