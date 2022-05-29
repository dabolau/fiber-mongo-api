package responses

import "fiber-mongo-api/models"

// 响应数据
type ResponseUser struct {
	Message    string
	StatusCode uint
}

// 响应多条数据
type ResponseMultipleUser struct {
	Datas      []models.User
	Message    string
	StatusCode uint
}

// 响应单条数据
type ResponseSingleUser struct {
	Data       models.User
	Message    string
	StatusCode uint
}
