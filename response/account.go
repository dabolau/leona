package response

import (
	"github.com/dabolau/leona/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 响应数据
type ResponseAccount struct {
	Message    string
	StatusCode uint
}

// 响应多条数据
type ResponseMultipleAccount struct {
	Datas      []model.Account
	Message    string
	StatusCode uint
}

// 响应单条数据
type ResponseSingleAccount struct {
	Data       model.Account
	Message    string
	StatusCode uint
}

// 响应编号数据
type ResponseIDAccount struct {
	ID         primitive.ObjectID
	Message    string
	StatusCode uint
}

// 响应编号和令牌数据
type ResponseIDAndBearerTokenAccount struct {
	ID          primitive.ObjectID
	BearerToken string
	Message     string
	StatusCode  uint
}

// 响应错误数据
type ResponseErrorAccount struct {
	Error      error
	Message    string
	StatusCode uint
}
