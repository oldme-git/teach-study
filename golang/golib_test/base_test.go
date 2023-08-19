package main

import (
	"context"
	"golib_test/logic"
	"runtime"
	"sync"
	"testing"
)

func TestReserve(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var (
		ctx = context.Background()
		wg  = &sync.WaitGroup{}
		num = 100000000
	)
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(i int) {
			logic.Reserve(ctx, i, i%5)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
