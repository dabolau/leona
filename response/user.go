package response

import "github.com/dabolau/leona/model"

// 响应数据
type ResponseUser struct {
	Message    string
	StatusCode uint
}

// 响应多条数据
type ResponseMultipleUser struct {
	Datas      []model.User
	Message    string
	StatusCode uint
}

// 响应单条数据
type ResponseSingleUser struct {
	Data       model.User
	Message    string
	StatusCode uint
}
