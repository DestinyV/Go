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
	// var u1 user
	// sqlstr := `select id, name, gender, profession from users where id=1;`
	// db.QueryRow(sqlstr).Scan(&u1.id, &u1.name, &u1.gender, &u1.profession) // 从连接池中取出一个连接去数据库查询单条记录 函数都是值拷贝
	// fmt.Printf("query user is :%#v", u1)
	query()
	insert()
	// update()
	// delete()
}

type user struct {
	id         int
	name       string
	gender     int
	profession string
}

// 查询多条语句
func query() {
	sqlstr := `select id, name, gender, profession from users where id > ?;`
	rows, err := db.Query(sqlstr, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.gender, &u.profession)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("query user is :%#v", u)
	}
}

func insert() {
	sqlStr := "insert into users(name, gender, profession) values(?,?,?)"
	result, err := db.Exec(sqlStr, "排污口", 0, "线控器")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := result.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

func update() {
	sqlStr := "update users set name=? where id = ?"
	result, err := db.Exec(sqlStr, "newName", 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := result.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

func delete() {
	sqlStr := "delete from users where id = ?"
	result, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := result.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
