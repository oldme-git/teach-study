package main

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	Main = &gcmd.Command{
		Name:        "测试",
		Brief:       "测试brief",
		Description: "测试description",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			fmt.Println(parser.GetArgAll())
			fmt.Println(parser.GetOptAll())
			return
		},
	}
)

// ----------------

type cMain struct {
	g.Meta `name:"main" brief:"start http server"`
}

type cMainHttpInput struct {
	g.Meta `name:"http" brief:"start http server"`
	Name   string `v:"required" name:"NAME" arg:"true" brief:"server name"`
	Port   int    `v:"required" short:"p" name:"port"  brief:"port of http server"`
}
type cMainHttpOutput struct{}

func (c *cMain) Http(ctx context.Context, in cMainHttpInput) (out *cMainHttpOutput, err error) {
	fmt.Printf("start http, name %s, port: %d\n", in.Name, in.Port)
	return
}

func main() {
	cmd, err := gcmd.NewFromObject(cMain{})
	if err != nil {
		panic(err)
	}
	cmd.Run(gctx.New())
}
