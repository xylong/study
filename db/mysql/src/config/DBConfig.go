package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const DSN = "root:123456@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() error {
	db, err := gorm.Open("mysql", DSN)
	if err != nil {
		return err
	}
	DB = db
	return nil
}
