package main

import (
	"grpc/app/article/internal/cmd"
	_ "grpc/app/article/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
