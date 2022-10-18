package conf

type Paging struct {
	Page      int `json:"page" example:"1"`
	Limit     int `json:"limit" example:"10"`
	Total     int `json:"total" example:"20"`
	Timestamp int `json:"timestamp" example:"1661246492"`
}

type UserCollection struct {
	Id       string `json:"id" example:"id"`
	Username string `json:"username" example:"用户名"`
	Age      int    `json:"age" example:"20"`
	Sex      string `json:"sex" example:"男"`
	Address  string `json:"address" example:"地址示例"`
}
