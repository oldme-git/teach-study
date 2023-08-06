package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "service/protobuf/goods"
)

type Goods struct {
	pb.UnimplementedGoodsRpcServer
}

func (g *Goods) GetGoods(ctx context.Context, req *pb.GoodsReq) (*pb.GoodsRes, error) {
	var (
		name string
		err  error
	)
	if req.Id == 0 {
		err = errors.New("没有商品")
	} else {
		name = fmt.Sprintf("%d号商品", req.Id)
	}
	return &pb.GoodsRes{Name: name}, err
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGoodsRpcServer(s, &Goods{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
