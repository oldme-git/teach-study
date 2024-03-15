package basetest

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
	"testing"
	"time"
)

func TestTimerBase(t *testing.T) {
	var (
		ctx = gctx.New()
		now = time.Now()
	)
	gtimer.AddTimes(ctx, 2*time.Second, 10, func(ctx context.Context) {
		fmt.Println(gtime.Now(), time.Duration(time.Now().UnixNano()-now.UnixNano()))
		now = time.Now()
	})

	select {}
}
