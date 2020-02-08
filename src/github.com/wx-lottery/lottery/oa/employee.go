package oa

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:Vankey@2017@tcp(rm-bp1x3aqpn76339a4mio.mysql.rds.aliyuncs.com:3306)/eventoa?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	if gin.Mode() == gin.DebugMode {
		DB.LogMode(true)
	}
	DB.DB().SetMaxIdleConns(25)
	DB.DB().SetMaxOpenConns(50)
}

type Employee struct {
	ID        int    `gorm:"primary_key"`
	Name      string `gorm:"default:NULL"`
	Telephone string
	OpenID    string `gorm:"column:open_id"`
	GroupID   int    `gorm:"column:group_id"`
	Group     *Group `gorm:"foreignKey:GroupID"`
}

func (e Employee) TableName() string {
	return "user"
}

type Group struct {
	ID        int `gorm:"primary_key"`
	GroupName string
}

func (g Group) TableName() string {
	return "department_group"
}
