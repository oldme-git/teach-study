package goframe

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gcache"
	"testing"
)

func TestGcache(t *testing.T) {
	ctx := context.Background()
	gcache.Set(ctx, "key", "value", 0)
	v, _ := gcache.Get(ctx, "key")
	fmt.Println(v)
}
