package goframe

import (
	"context"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"

	"testing"
)

type RedisTest struct {
	id    int
	title string
}

func TestRedisSet(t *testing.T) {
	redis := g.Redis()
	rt := &RedisTest{
		id:    1,
		title: "title",
	}
	redis.Set(context.Background(), "key", rt)
}
