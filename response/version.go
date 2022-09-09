package response

import "github.com/dabolau/leona/model"

// 响应数据
type ResponseVersion struct {
	Message    string
	StatusCode uint
}

// 响应多条数据
type ResponseMultipleVersion struct {
	Datas      []model.Version
	Message    string
	StatusCode uint
}

// 响应单条数据
type ResponseSingleVersion struct {
	Data       model.Version
	Message    string
	StatusCode uint
}
