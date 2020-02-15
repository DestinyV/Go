package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// _ "github.com/go-sql-driver/mysql"

var db *sql.DB // 是一个连接池对象

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/golangdb"
	db, err = sql.Open("mysql", dsn) // open 不会校验用户名和密码是否正确
	if err != nil {
		return
	}
	err = db.Ping() // 尝试连接数据库
	if err != nil {
		return
	}
	// 设置数据库连接池最大连接数
	db.SetMaxOpenConns(10)
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed, err: %v\n", err)
		return
	}
	fmt.Println("连接数据库成功s")
	sqlstr := `select id, name, gender, profession from users where id=1;`
	rowObj := db.QueryRow(sqlstr) // 从连接池中取出一个连接去数据库查询单条记录
	var u1 user
	rowObj.Scan(&u1.id, &u1.name, &u1.gender, &u1.profession) //函数都是值拷贝
	fmt.Printf("query user is :%#v", u1)
}

type user struct {
	id         int
	name       string
	gender     int
	profession string
}

func query() {

}

func insert() {

}
