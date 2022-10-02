package response

import "github.com/dabolau/leona/model"

// 响应数据
type ResponseVideo struct {
	Message    string
	StatusCode uint
}

// 响应多条数据
type ResponseMultipleVideo struct {
	Datas      []model.Video
	Message    string
	StatusCode uint
}

// 响应单条数据
type ResponseSingleVideo struct {
	Data       model.Video
	Message    string
	StatusCode uint
}
