package article

import (
	"context"
	v1 "grpc/app/article/api/article/v1"
	"grpc/app/article/api/pbentity"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedArticleServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterArticleServer(s.Server, &Controller{})
}

func (*Controller) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	var list []*pbentity.Article
	list = append(list, &pbentity.Article{
		Id:          1,
		GrpId:       1,
		Title:       "title",
		Author:      "author",
		Thumb:       "thumb",
		Tags:        "tags",
		Description: "des",
		Content:     "content",
	}, &pbentity.Article{
		Id:          2,
		GrpId:       1,
		Title:       "title2",
		Author:      "author2",
		Thumb:       "thumb2",
		Tags:        "tags2",
		Description: "des2",
		Content:     "content2",
	})
	res = &v1.ListRes{Article: list}
	return res, nil
}
