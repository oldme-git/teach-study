package client_test

import (
	"context"
	"fmt"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc/app/article/api/article/v1"
	"grpc/consts"
	"log"
	"testing"
	"time"
)

// 原生客户端调用
func TestClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()
	c := pb.NewArticleClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.List(ctx, &pb.ListReq{
		Page: 1,
		Size: 15,
	})
	if err != nil {
		log.Fatalf("无法调用: %v", err)
	}
	for _, v := range r.GetArticle() {
		//fmt.Printf("%+v\n", v)
		fmt.Printf("id: %d\n", v.Id)
		fmt.Printf("title: %s\n", v.Title)
	}
}

// gf客户端调用
func TestGfClient(t *testing.T) {
	grpcx.Resolver.Register(etcd.New(consts.EtcdAddress))
	var (
		ctx  = gctx.New()
		conn = grpcx.Client.MustNewGrpcClientConn("article")
		//conn, err = grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()))
		client = pb.NewArticleClient(conn)
	)

	r, err := client.List(ctx, &pb.ListReq{
		Page: 1,
		Size: 15,
	})
	if err != nil {
		panic(err)
	}
	for _, v := range r.GetArticle() {
		//fmt.Printf("%+v\n", v)
		fmt.Printf("id: %d\n", v.Id)
		fmt.Printf("title: %s\n", v.Title)
	}
}
