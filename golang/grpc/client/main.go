package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"service/protobuf/user"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:49961", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()
	//c := pb.NewGoodsRpcClient(conn)
	c := user.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.Create(ctx, &user.CreateReq{
		Passport: "1",
		Password: "2",
		Nickname: "3",
	})
	if err != nil {
		log.Fatalf("无法调用: %v", err)
	}
	//log.Printf("商品价格: %d", r)
}
