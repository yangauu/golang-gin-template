package conf

type UserCollection struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
	Address  string `json:"address"`
}
