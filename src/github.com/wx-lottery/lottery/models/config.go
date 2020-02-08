package models

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:VKT.666666@tcp(db:3306)/lottery?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	if gin.Mode() == gin.DebugMode {
		DB.LogMode(true)
	}
	DB.DB().SetMaxIdleConns(25)
	DB.DB().SetMaxOpenConns(50)
}
