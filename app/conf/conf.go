package conf

type Paging struct {
	Page  int         `json:"page"  example:"1"`
	Limit int         `json:"limit" example:"10"`
	Total int64       `json:"total" example:"20"`
	List  interface{} `json:"list,omitempty"`
}

type Model struct {
	Timestamp  int64  `json:"timestamp"  example:"1661233245"`
	CreateTime string `json:"createTime"  example:"2022-12-12 00:00:00"`
}

type User struct {
	Id       int64  `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Username string `json:"username"  example:"用户名"`
	Age      int64  `json:"age" example:"20"`
	Sex      string `json:"sex" example:"男"`
	Address  string `json:"address"  example:"地址示例"`
	Model
}
