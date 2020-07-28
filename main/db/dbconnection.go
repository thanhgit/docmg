package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var instance *gorm.DB

func GetInstance() *gorm.DB {
	if instance == nil {
		db, err := gorm.Open("mysql", "docmguser:123456@tcp(172.17.0.2:3306)/docmg?charset=utf8&parseTime=True&loc=Local")

		if err != nil {
			println(err.Error())
			defer db.Close()
		} else {
			println("Connect success to Mysql server")
			instance = db
			instance.AutoMigrate(&Document{}, &Document_user{}, &Event_log{}, &GroupUser{}, &User{}, &User_group{})
		}

	}

	return instance
}