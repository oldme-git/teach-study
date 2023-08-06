package main

import (
	"flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "service/protobuf/goods"
	"time"
)

var (
	addr = flag.String("addr", "localhost:10001", "the address to connect to")
)

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()
	c := pb.NewGoodsRpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetGoods(ctx, &pb.GoodsReq{Id: 1})
	if err != nil {
		log.Fatalf("无法调用: %v", err)
	}
	log.Printf("商品名称: %s", r.GetName())
	log.Printf("商品价格: %d", r.GetPrice())
}
