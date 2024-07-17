package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "service/protobuf/goods"
)

func main() {
	conn, err := grpc.NewClient("localhost:800", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGoodsRpcClient(conn)

	// 使用 context.WithTimeout 设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetGoods(ctx, &pb.GoodsReq{Id: 1})
	if err != nil {
		log.Fatalf("无法调用: %v", err)
	}
	log.Printf("商品价格: %d", r.Price)
}
