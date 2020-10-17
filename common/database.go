package common

import (
	"fmt"
	"MessageBoard/conf"
	"MessageBoard/model"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func InitDB() *gorm.DB{
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",

		conf.User,
		conf.Password,
		conf.Address,
		conf.Port,
		conf.Databases,
		conf.Chatset,
	)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database err= " + err.Error())
	}
	err=db.AutoMigrate(&model.Message{})
	if err != nil {
		panic("Create Table Error:"+err.Error())
	}
	DB = db
	return db
}
func GetDB() *gorm.DB {
	return DB
}