package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/golangdb"
	db, err = sqlx.Connect("mysql", dsn) // open 不会校验用户名和密码是否正确
	if err != nil {
		return
	}
	// 设置数据库连接池最大连接数
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

type user struct {
	ID         int
	Name       string
	Gender     int
	Profession string
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	fmt.Println("数据库连接成功")
	// queryRow()
	// queryRows()
	// insertRowDemo()
	// updateRowDemo()
	// deleteRowDemo()
	transactionDemo()
}

// 单行查询
func queryRow() {
	sqlStr := "select id, name, gender, profession from users where id=?"
	var u user
	err := db.Get(&u, sqlStr, 2)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s gender:%d profession:%s\n", u.ID, u.Name, u.Gender, u.Profession)
}

// 多行查询
func queryRows() {
	sqlStr := "select id, name, gender, profession from users where id > ?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into users(name, gender, profession) values (?,?,?)"
	ret, err := db.Exec(sqlStr, "阿林新", 1, "评论员")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo() {
	// sqpStr := "update users set name=? where id=?"
	sqlStr := "update users set gender=? where id = ?"
	ret, err := db.Exec(sqlStr, 1, 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from users where id = ?"
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func transactionDemo() {
	tx, err := db.Beginx() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update users set name=? where id=?"
	tx.MustExec(sqlStr1, "the second", 2)
	sqlStr2 := "Update users set profession=? where id=?"
	tx.MustExec(sqlStr2, "player", 1)
	err = tx.Commit() // 提交事务
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("commit failed, err:%v\n", err)
		return
	}
	fmt.Println("exec trans success!")
}
