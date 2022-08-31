package response

import "github.com/dabolau/leona/model"

// 响应数据
type ResponseFeedback struct {
	Message    string
	StatusCode uint
}

// 响应多条数据
type ResponseMultipleFeedback struct {
	Datas      []model.Feedback
	Message    string
	StatusCode uint
}

// 响应单条数据
type ResponseSingleFeedback struct {
	Data       model.Feedback
	Message    string
	StatusCode uint
}
