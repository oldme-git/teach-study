package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "service/protobuf/goods"
)

// Goods 定义 Goods, 实现接口
type Goods struct {
	pb.UnimplementedGoodsRpcServer
}

// GetGoods 实现获取商品的功能
func (g *Goods) GetGoods(ctx context.Context, req *pb.GoodsReq) (*pb.GoodsRes, error) {
	// 模拟超时
	//time.Sleep(2 * time.Second)
	var (
		name string
		err  error
	)
	if req.Id == 0 {
		err = errors.New("商品不存在")
	} else {
		name = fmt.Sprintf("%d号商品", req.Id)
	}
	return &pb.GoodsRes{
		Name:  name,
		Price: 20,
	}, err
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":800")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGoodsRpcServer(s, &Goods{})
	log.Printf("正在监听端口： %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("启用服务失败: %v", err)
	}
}
