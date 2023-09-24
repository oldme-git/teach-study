package main

import (
	_ "goframe-feature-test/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe-feature-test/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
