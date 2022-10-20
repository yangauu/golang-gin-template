package db

import (
	"fmt"
	"go-template/app/conf"
	"go-template/app/utils"
	"time"
)

// 添加
func ExamInsert(user conf.User) error {
	t := time.Now()
	temp := t.Unix()
	creat := t.Format("2006-01-02 15:04:05")
	user.Timestamp = temp
	user.CreateTime = creat

	tx := db.Create(&user)
	if tx.Error != nil {
		fmt.Println("添加失败 ", tx.Error)
		return tx.Error
	}
	// fmt.Println("自增id", user.Id)
	return nil
}

// 删除
func ExamDel(id int64) error {
	tx := db.Delete(&conf.User{}, id)
	if tx.Error != nil {
		fmt.Println("删除失败 ", tx.Error)
		return tx.Error
	}
	return nil
}

// 修改
func ExamEdit(user conf.User) error {
	tx := db.Model(&conf.User{}).Where("id = ?", user.Id).Updates(user)
	if tx.Error != nil {
		fmt.Println("更新失败 ", tx.Error)
		return tx.Error
	}
	return nil
}

// 查询
func ExamQuery(page int) (conf.Paging, error) {
	result := []conf.User{}
	var total int64

	tx := db.Model(conf.User{}).Count(&total).Scopes(utils.Paginate(page)).Find(&result)
	if tx.Error != nil {
		fmt.Println("查询失败", tx.Error)
		return conf.Paging{}, tx.Error
	}

	p := conf.Paging{
		Page:  page,
		Limit: 10,
		Total: total,
		List:  result,
	}
	return p, nil
}

// 连接
func init() {
	connect()
}
