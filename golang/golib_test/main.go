package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		//logic.Reserve(r.GetCtx(), guid.S(), grand.Intn(1))
	})
	s.Run()
}
