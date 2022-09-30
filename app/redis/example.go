package redis

import (
	"fmt"
	"time"
)

// 获取
func GetValue() (string, error) {
	val, err := rdb.Get("test").Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// 存储
func SetValue(key string) {
	err := rdb.Set("test", "10秒后不存在", time.Second*10).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}
}

// 连接
func init() {
	connect()
}
