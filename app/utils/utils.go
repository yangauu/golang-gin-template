package utils

import "gorm.io/gorm"

func Paginate(page int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * 10).Limit(10)
	}
}
