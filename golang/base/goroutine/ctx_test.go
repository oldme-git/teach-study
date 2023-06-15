// 辅助理解go上下文
package goroutine

import (
	"base"
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
	db, _ := sql.Open("mysql", base.Link)
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

func TestCtxWithCancel(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		for {
			select {
			// 还记得前文提到的Done的方法吗
			// 当 ctx 取消时，ctx.Done()对应的通道就会关闭，case也就会被执行
			case <-ctx.Done():
				// ctx.Err() 会获取到关闭原因哦
				fmt.Println("协程关闭", ctx.Err())
				return
			default:
				fmt.Println("继续运行")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// 等待一秒后关闭
	time.Sleep(1 * time.Second)
	cancel()
	// 等待一秒，让子协程有时间打印出协程关闭的原因
	time.Sleep(1 * time.Second)
}

// 向上找到最近的上下文值
func TestCtxWithValue(t *testing.T) {
	ctx := context.Background()
	ctx1 := context.WithValue(ctx, "key", "ok")
	ctx2, _ := context.WithCancel(ctx1)
	// Value 会一直向上追溯到根节点，获取当前上下文携带的值，
	value := ctx2.Value("key")
	if value != nil {
		fmt.Println(value)
	}
}

// WithDeadline
func TestCtxWithDeadline(t *testing.T) {
	ctx := context.Background()
	// 调用cancel可以手动关闭，否则等待2秒后自动关闭
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(2*time.Second))
	defer cancel()
	go func() {
		select {
		case <-ctx.Done():
			// 手动关闭 context canceled
			// 自动关闭 context deadline exceeded
			fmt.Println("协程关闭", ctx.Err())
			return
		}
	}()

	time.Sleep(3 * time.Second)
}

func TestCtxWithTimeout(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("协程关闭", ctx.Err())
			return
		}
	}()

	time.Sleep(3 * time.Second)
}
