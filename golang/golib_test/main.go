package main

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/guid"
	"golib_test/logic"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		logic.Reserve(r.GetCtx(), guid.S(), grand.Intn(3))
	})
	s.Run()
}
