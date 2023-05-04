// 辅助理解go上下文
package test

import (
	"context"
	"database/sql"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 使用ctx控制超时时间
func TestPingCtx(t *testing.T) {
	start := time.Now()
	db, _ := sql.Open("mysql", link)
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	if err := db.PingContext(ctx); err != nil {
		t.Log("连接失败")
	}
	t.Log("ok")
	elapsed := time.Since(start)
	t.Log(elapsed.Seconds())
}

// 使用ctx关闭多层goroutine
func TestCtx(m *testing.T) {
	fmt.Printf("开始了，有%d个协程\n", runtime.NumGoroutine())
	// 父context(利用根context得到)
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "key", "value")

	var e func(ctx context.Context, i int)

	e = func(ctx context.Context, i int) {
		if i < 3 {
			c, _ := context.WithCancel(ctx)
			if i == 1 {
				// 第1层后改变键值
				c = context.WithValue(c, "key", "value2")
			}
			go e(c, i+1)
		}
		for {
			select {
			case <-ctx.Done():
				var s string
				// 获取取消原因
				if err := ctx.Err(); err != nil {
					s = err.Error()
				}
				fmt.Printf("收到关闭信号，第%d层退出，退出原因：%s\n", i, s)
				return
			default:
				fmt.Printf("第%d层协程监听中，key值：%s\n", i, ctx.Value("key"))
				time.Sleep(1 * time.Second)
			}
		}
	}
	go e(ctx, 0)

	fmt.Println("等待3秒")
	time.Sleep(3 * time.Second)
	// 调用cancel()
	fmt.Printf("调用cancel()前协程数%d\n", runtime.NumGoroutine())
	cancel()

	// 等待2秒
	time.Sleep(2 * time.Second)
	fmt.Printf("最终结束，有%d个协程\n", runtime.NumGoroutine())
}

// 向上找到最近的上下文值
func TestValue(t *testing.T) {
	ctx := context.Background()
	ctx1 := context.WithValue(ctx, "key", "ctx1")
	ctx2 := context.WithValue(ctx1, "key", "ctx2")
	value := ctx2.Value("key")
	if value != nil {
		fmt.Println(value)
	}
}
