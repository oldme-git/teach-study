package bench

import (
	"sync"
	"testing"
)

// wg.add 一次添加
func BenchmarkWgOnce(b *testing.B) {
	wg := &sync.WaitGroup{}
	wg.Add(50000000)
	for i := 0; i < 50000000; i++ {
		go func(wg *sync.WaitGroup) {
			wg.Done()
		}(wg)
	}
	wg.Wait()
}

// wg.add 多次添加
func BenchmarkWgMultiple(b *testing.B) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 50000000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			wg.Done()
		}(wg)
	}
	wg.Wait()
}
