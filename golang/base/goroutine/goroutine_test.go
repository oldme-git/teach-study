package goroutine

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
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

// 测试time.Since是否会受到抢占式调度的影响
func TestTimeSinceWithGMP(t *testing.T) {
	runtime.GOMAXPROCS(1)
	var (
		sMap sync.Map
		wg   = &sync.WaitGroup{}
		num  = 5000000
	)

	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) {
			defer wg.Done()
			start := time.Now()
			// 模拟一个耗时请求
			time.Sleep(3 * time.Second)
			elapsed := time.Since(start).Seconds()
			sMap.Store(strconv.Itoa(i), elapsed)
		}(i)
	}
	wg.Wait()

	j := 0
	sMap.Range(func(key, value any) bool {
		if value.(float64) > 4 {
			j++
		}
		return true
	})
	fmt.Printf("有 %d 个异常", j)
}
