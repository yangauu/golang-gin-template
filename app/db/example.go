package db

import (
	"fmt"
	"go-template/app/conf"
	"log"

	"github.com/gofrs/uuid"
)

// 添加
func ExamInsert(u conf.UserCollection) error {
	uuid, _ := uuid.NewV4()
	u.Id = uuid.String()

	sql := "insert into user (id,username,age,sex,address) value(?,?,?,?,?)"
	res, err := db.Exec(sql, u.Id, u.Username, u.Age, u.Sex, u.Address)
	if err != nil {
		fmt.Println("添加失败 ", err)
		return err
	}
	id, _ := res.LastInsertId() // 获取新增行id

	fmt.Println("添加成功", id)
	return nil
}

// 删除
func ExamDel(id string) error {
	sql := "delete from user where id=?"
	res, err := db.Exec(sql, id)
	if err != nil {
		log.Println("删除失败", err)
	}

	num, err := res.RowsAffected() // 受影响行数

	if err != nil {
		log.Println("删除失败", err)
	}
	if num == 0 {
		return fmt.Errorf("没有找到数据")
	}
	return nil
}

// 修改
func ExamEdit(u conf.UserCollection) error {
	sql := "update user set username=?,age=?,sex=?,address=? where id = ?"
	res, err := db.Exec(sql, u.Username, u.Age, u.Sex, u.Address, u.Id)
	if err != nil {
		log.Println("更新失败", err)
	}

	_, err = res.RowsAffected() //受影响行数
	if err != nil {
		log.Println("更新失败", err)
	}
	return nil
}

// 查询
func ExamQuery() ([]conf.UserCollection, error) {
	r := []conf.UserCollection{}
	e := []conf.UserCollection{}

	sql := "select id,username,age,sex,address from user"
	res, err := db.Query(sql)
	if err != nil {
		fmt.Println("查询失败 ", res, err)
		return e, err
	}

	// 会释放数据库连接
	defer func() {
		res.Close()
	}()

	// 解析查询结果
	for res.Next() {
		c := conf.UserCollection{}
		err := res.Scan(&c.Id, &c.Username, &c.Age, &c.Sex, &c.Address)
		if err != nil {
			fmt.Println("编译结果失败", err)
			return e, err
		}

		r = append(r, c)
	}

	return r, nil
}

// 连接
func init() {
	connect()
}
