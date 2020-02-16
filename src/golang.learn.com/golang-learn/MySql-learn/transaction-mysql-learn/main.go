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
	// db.SetMaxOpenConns(10)
	return
}

type user struct {
	id         int
	name       string
	gender     int
	profession string
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed, err: %v\n", err)
		return
	}
	fmt.Println("连接数据库成功s")
	transactionDemo()
}

func transactionDemo() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update users set name=? where id=?"
	_, err = tx.Exec(sqlStr1, "transactionEditName", 6)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	sqlStr2 := "Update users set profession=? where id=?"
	_, err = tx.Exec(sqlStr2, "transactionEditPro", 5)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	err = tx.Commit() // 提交事务
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("commit failed, err:%v\n", err)
		return
	}
	fmt.Println("exec trans success!")
}
