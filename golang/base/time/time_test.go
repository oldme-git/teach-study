package time

import (
	"fmt"
	"testing"
	"time"
)

// NewTicker 会重复向通道中发送数据
func TestNewTicker(t *testing.T) {
	t1 := time.NewTicker(1 * time.Second)
	exit := make(chan struct{})
	go func() {
		for {
			select {
			case r := <-t1.C:
				fmt.Println(r)
			case <-exit:
				fmt.Println("exit")
				return
			}
		}
	}()
	time.Sleep(3 * time.Second)
	t1.Stop()
	close(exit)
	time.Sleep(5 * time.Second)
}

// After 一段时间后执行
func TestAfter(t *testing.T) {
	t1 := time.After(1 * time.Second)
	<-t1
	fmt.Println("触发")
}

// NewTimer 一段时间后执行，区别After，可以手动stop
func TestNewTimer(t *testing.T) {
	t1 := time.NewTimer(2 * time.Second)
	exit := make(chan struct{})
	go func() {
		for {
			select {
			case r := <-t1.C:
				fmt.Println(r)
			case <-exit:
				fmt.Println("exit")
				return
			}
		}
	}()

	time.Sleep(1 * time.Second)
	t1.Stop()
	close(exit)
	time.Sleep(3 * time.Second)
}
