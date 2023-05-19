package goroutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 打印协程数量
func TestGoNum(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	ch := make(chan bool)
	go func() {
		for {
		}
	}()
	go func() {
		select {
		case <-ch:
			return
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(runtime.NumGoroutine())
	// 通知协程退出退出
	ch <- true
	time.Sleep(1 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}
