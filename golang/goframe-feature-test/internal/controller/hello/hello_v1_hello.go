package hello

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-feature-test/api/hello/v1"
)

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	one, err := g.Model("user").RightJoinOnFields("user_detail", "id", "=", "uid").Fields("user.*,user_detail.address").All()
	if err != nil {
		return nil, err
	}
	g.Dump(one)
	return
}
