package conf

type SuccApiExamDel struct {
	Code int    `json:"code" example:"200"`
	Msg  string `json:"msg" example:"删除成功"`
}

type ErrApiExamDel struct {
	Code int    `json:"code" example:"400"`
	Msg  string `json:"msg" example:"删除失败"`
}

type SuccApiDelQiNiuFile struct {
	Code int    `json:"code" example:"200"`
	Msg  string `json:"msg" example:"删除成功"`
	Data string `json:"data" example:"xxx.png"`
}

type ErrApiDelQiNiuFile struct {
	Code int    `json:"code" example:"400"`
	Msg  string `json:"msg" example:"删除失败"`
}

type SuccApiExamQueryData struct {
	Paging
	List []User `json:"list"`
}

type SuccApiExamQuery struct {
	Code int                  `json:"code" example:"200"`
	Msg  string               `json:"msg" example:"获取成功"`
	Data SuccApiExamQueryData `json:"data"`
}

type ErrApiExamQuery struct {
	Code int    `json:"code" example:"400"`
	Msg  string `json:"msg" example:"获取列表出错"`
}
