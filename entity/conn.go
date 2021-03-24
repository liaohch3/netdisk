package entity

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 这是为了导入mysql驱动程序
)

var db *gorm.DB

// todo 主从读写分离

func InitOrm() {
	// todo 数据库连接配置化
	var err error
	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3329)/file?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
