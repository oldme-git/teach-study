package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"golib_test/logic"
	"sync"
	"testing"
	"time"
)

func TestHttpClient(t *testing.T) {
	logic.InitData()
	time.Sleep(1 * time.Second)
	var (
		url = "http://192.168.10.42:30080/"
		ctx = gctx.New()
		wg  = &sync.WaitGroup{}
		num = 3000
	)
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			r, err := g.Client().Get(ctx, url)
			if err != nil {
				panic(err)
			}
			r.Close()
			wg.Done()
		}()
	}
	wg.Wait()
}
