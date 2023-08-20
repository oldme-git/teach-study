package main

import (
	"context"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/util/guid"
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
		num = 1000
	)
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(i int) {
			logic.Reserve(ctx, guid.S(), i%3)
			wg.Done()
		}(i)
	}
	wg.Wait()
	logic.InitData()
}

func TestInit(t *testing.T) {
	logic.InitData()
}
