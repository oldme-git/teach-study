package test

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
	"time"
)

var ctx = gctx.New()
var r = g.Redis()

func TestRedis(t *testing.T) {
	// 循环次数，代表并发量
	num := 10000
	// 步长，代表重复多少次并发
	num2 := 10
	for i := 0; i < num; i++ {
		go insert(i*num2, i*num2+num2)
	}

	var input string
	fmt.Scanln(&input)
}

func insert(s int, j int) {
	for i := s; i < j; i++ {
		r.Do(ctx, "set", i, time.Now().UnixNano()/1e6)
	}
}
