// 学习database/sql的使用
package test

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

// 测试两个基础函数
func TestFunc(t *testing.T) {
	// 返回已经注册的驱动
	fmt.Println(sql.Drivers())
	// 注册一个驱动，driver需要实现driver.Driver
	sql.Register("me", nil)
}

// 测试type Db 的基本方法
func TestTypeDb(t *testing.T) {
	ctx := context.Background()
	// Open 打开一个db
	db, _ := sql.Open("mysql", "root:JCPHqknyy8ATR5ME@tcp(192.168.10.47:3306)/oldme")
	// 测试db是否能通
	err := db.Ping()
	if err != nil {
		t.Log(err)
	}
	// Conn 建立一个单独的数据库连接，不使用数据库连接池查询数据
	_, _ = db.Conn(ctx)
	// Driver 返回当前db驱动
	db.Driver()
}

// 测试查询语句
func TestQuery(t *testing.T) {
	type saying struct {
		Id     uint
		Saying string
	}

	// Open 打开一个db
	db, _ := sql.Open("mysql", "root:JCPHqknyy8ATR5ME@tcp(192.168.10.47:3306)/oldme")

	// 查询单条数据
	var row saying
	db.QueryRow("select * from saying where id = ?", 2).Scan(&row.Id, &row.Saying)
	t.Log(row)

	rows, _ := db.Query("select * from saying")
	// 查询多条数据，返回一个rows
	defer rows.Close()
	data := make([]saying, 0)

	// 获取字段信息
	columns, _ := rows.Columns()
	t.Log(columns)

	// 从rows中读取数据出来
	for rows.Next() {
		if err := rows.Scan(&row.Id, &row.Saying); err != nil {
			log.Fatal(err)
		}
		data = append(data, row)
	}
	t.Log(data)
}

// 测试写入语句
func TestWrite(t *testing.T) {
	// Open 打开一个db
	db, _ := sql.Open("mysql", "root:JCPHqknyy8ATR5ME@tcp(192.168.10.47:3306)/oldme")
	res, _ := db.Exec("insert into saying (saying) values (?)", "我是gopher")
	t.Log(res.LastInsertId()) // 最后插入数据的id
	t.Log(res.RowsAffected()) // 成功行数
}

// SQL准备语句
func TestPrepare(t *testing.T) {
	db, _ := sql.Open("mysql", "root:JCPHqknyy8ATR5ME@tcp(192.168.10.47:3306)/oldme")
	// 准备sql语句
	stmt, err := db.Prepare("update saying set saying = ? where id = ?")
	if err != nil {
		panic(err)
	}
	// 使用后释放资源
	defer stmt.Close()

	_, err = stmt.Exec("我也是gopher", 7)
	if err != nil {
		panic(err)
	}
}

// 打印database信息
func TestStats(t *testing.T) {
	db, _ := sql.Open("mysql", "root:JCPHqknyy8ATR5ME@tcp(192.168.10.47:3306)/oldme")
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(15)
	stats := db.Stats()
	fmt.Printf("%+v\n", stats)
}
