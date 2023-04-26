package test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 打印协程数量
func TestGoNum(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	chin := make(chan bool)
	go func() {
		for {
		}
	}()
	go func() {
		select {
		case <-chin:
			return
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(runtime.NumGoroutine())
	// 通知协程退出退出
	chin <- true
	time.Sleep(1 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}
