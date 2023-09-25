package test

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"testing"
)

type Post struct {
	CreatedAt *gtime.Time `json:"createdAt" ` //
}

type MyPost struct {
	Post
	Ok string
}

//func Test_2980A(t *testing.T) {
//	var (
//		ctx = gctx.New()
//	)
//	one, err := dao.Post.Ctx(ctx).Where("id", 1).One()
//	if err != nil {
//		panic(err)
//	}
//
//	g.Dump(one["created_at"].GTime().Format("Y-m-d H:i:s"))
//	g.Dump(one["created_at"].GTime().Unix())
//
//	info := g.Map{
//		"Id":        1,
//		"Content":   "hello",
//		"UserId":    1,
//		"CreatedAt": one["created_at"],
//		"UpdatedAt": "2023-09-22 12:00:00",
//		"DeletedAt": "",
//	}
//
//	data := new(MyPost)
//	err = gconv.Scan(info, data)
//	if err != nil {
//		panic(err)
//	}
//
//	g.Dump(data.CreatedAt.Location())
//	g.Dump(data.CreatedAt.Format("Y-m-d H:i:s"))
//	g.Dump(data.CreatedAt.Unix())
//}

func Test_2980A(t *testing.T) {
	date := gtime.New("2023-09-22 12:00:00").UTC()
	g.Dump(date.Unix())

	info := g.Map{
		"CreatedAt": date,
	}

	data := new(MyPost)
	err := gconv.Scan(info, data)
	if err != nil {
		panic(err)
	}

	g.Dump(data.CreatedAt.Location())
	g.Dump(data.CreatedAt.Unix())
}
