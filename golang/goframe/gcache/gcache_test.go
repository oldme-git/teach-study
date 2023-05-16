package gcache

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
	"testing"
)

func TestGcache(t *testing.T) {
	var (
		ctx    = context.Background()
		c      = gcache.New()
		v1, v2 *gvar.Var
	)

	c.Set(ctx, "key", "value", 0)
	c.Set(ctx, "key1", "value1", 0)
	v1, _ = c.Get(ctx, "key")
	v2, _ = c.Get(ctx, "key1")
	fmt.Println(v1, v2)
	c.Clear(ctx)
	v1, _ = c.Get(ctx, "key")
	v2, _ = c.Get(ctx, "key1")
	fmt.Println(v1, v2)
}
