package hello

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-feature-test/api/hello/v1"
	"goframe-feature-test/internal/model/entity"
	"log"
)

type PostWithUser struct {
	g.Meta    `orm:"table:post"`
	Id        int64        `json:"id"        ` //
	Content   string       `json:"content"   ` //
	UserId    int64        `json:"userId"    ` //
	CreatedAt *gtime.Time  `json:"createdAt" ` //
	UpdatedAt *gtime.Time  `json:"updatedAt" ` //
	DeletedAt *gtime.Time  `json:"deletedAt" ` //
	User      *entity.User `json:"user" orm:"with:id=userId"`
}

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	var p PostWithUser
	err = g.Model(p).WithAll().Where("id", 1).Scan(&p)
	log.Printf("%v is not equals to %v", p.CreatedAt.Unix(), p.User.CreatedAt.Unix())
	return
}
