package main

import (
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"grpc/app/article/internal/cmd"
	_ "grpc/app/article/internal/packed"
)

func main() {
	grpcx.Resolver.Register(etcd.New("192.168.10.43:2379"))
	cmd.Main.Run(gctx.GetInitCtx())
}
