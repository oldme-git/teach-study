// 辅助理解go上下文
package test

import (
	"context"
	"database/sql"
	"testing"
	"time"
)

// 使用ctx控制超时时间
func TestPingCtx(t *testing.T) {
	start := time.Now()
	db, _ := sql.Open("mysql", "root:JCPHqknyy8ATR5ME@tcp(192.168.10.47:3306)/oldme")
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	if err := db.PingContext(ctx); err != nil {
		t.Log("连接失败")
	}
	t.Log("ok")
	elapsed := time.Since(start)
	t.Log(elapsed.Seconds())
}
