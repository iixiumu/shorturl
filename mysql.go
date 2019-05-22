package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MyDB *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:111111@tcp(10.179.32.217:3306)/test?&charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		panic("mysql connect error: " + err.Error())
	}
	db.DB().Ping()
	db.DB().SetMaxIdleConns(128)
	db.DB().SetMaxOpenConns(1024)
	db.LogMode(true)
	db.AutoMigrate(&MyURL{})

	MyDB = db
}
