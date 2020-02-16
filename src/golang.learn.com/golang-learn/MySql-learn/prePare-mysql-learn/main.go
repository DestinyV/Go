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
	prePareQuery(0)
	prePareInsert("琉璃珠", 1, "作家")
}

func prePareQuery(p1 int) {
	sqlStr := "select id, name, gender, profession from users where id > ? "
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(p1)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.gender, &u.profession)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("prePareQuery user is :%#v\n", u)
	}
}

func prePareInsert(name string, gender int, pro string) {
	sqlStr := "insert into users(name, gender, profession) values(?,?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(name, gender, pro)
	if err != nil {
		fmt.Println(err)
		return
	}
	effectID, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("insert success, inset item id is:%d\n", effectID)
}
