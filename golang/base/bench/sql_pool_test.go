package bench

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

var db *sql.DB

func BenchmarkInsertPool(b *testing.B) {
	db, _ = sql.Open("mysql", "root:JCPHqknyy8ATR5ME@tcp(192.168.10.47:3306)/oldme")
	db.SetMaxOpenConns(14)                  // 一般设置为CPU核心数+2乘可用磁盘数量
	db.SetMaxIdleConns(7)                   // 一般设置为SetMaxOpenConns的一半
	db.SetConnMaxLifetime(time.Minute * 60) // 小于数据库超时时间

	// 重置计时器
	b.ResetTimer()
	// 并发执行
	b.RunParallel(func(pb *testing.PB) {
		// 迭代次数由b.N控制
		for pb.Next() {
			_, _ = db.Exec("insert into saying (saying) values (?)", "我是gopher")
		}
	})
	// 停止计时器
	b.StopTimer()
}

// 与pool对比
func BenchmarkInsert(b *testing.B) {
	db, _ = sql.Open("mysql", "root:JCPHqknyy8ATR5ME@tcp(192.168.10.47:3306)/oldme")
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 60)
	// 重置计时器
	b.ResetTimer()
	// 并发执行
	b.RunParallel(func(pb *testing.PB) {
		// 迭代次数由b.N控制
		for pb.Next() {
			_, _ = db.Exec("insert into saying (saying) values (?)", "我是gopher")
		}
	})
	// 停止计时器
	b.StopTimer()
}
