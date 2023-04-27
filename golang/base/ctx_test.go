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

// 父context的协程
func watch1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //取出值即说明是结束信号
			fmt.Println("收到信号，父context的协程退出,time=", time.Now().Unix())
			return
		default:
			fmt.Println("父context的协程监控中,time=", time.Now().Unix())
			time.Sleep(1 * time.Second)
		}
	}
}

// 子context的协程
func watch2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //取出值即说明是结束信号
			fmt.Println("收到信号，子context的协程退出,time=", time.Now().Unix())
			return
		default:
			fmt.Println("子context的协程监控中,time=", time.Now().Unix())
			time.Sleep(1 * time.Second)
		}
	}
}
