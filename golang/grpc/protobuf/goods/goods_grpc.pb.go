// 使用proto3

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0--rc3
// source: protobuf/goods/goods.proto

package goods

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GoodsRpc_GetGoods_FullMethodName = "/goods.GoodsRpc/GetGoods"
)

// GoodsRpcClient is the client API for GoodsRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsRpcClient interface {
	GetGoods(ctx context.Context, in *GoodsReq, opts ...grpc.CallOption) (*GoodsRes, error)
}

type goodsRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsRpcClient(cc grpc.ClientConnInterface) GoodsRpcClient {
	return &goodsRpcClient{cc}
}

func (c *goodsRpcClient) GetGoods(ctx context.Context, in *GoodsReq, opts ...grpc.CallOption) (*GoodsRes, error) {
	out := new(GoodsRes)
	err := c.cc.Invoke(ctx, GoodsRpc_GetGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsRpcServer is the server API for GoodsRpc service.
// All implementations must embed UnimplementedGoodsRpcServer
// for forward compatibility
type GoodsRpcServer interface {
	GetGoods(context.Context, *GoodsReq) (*GoodsRes, error)
	mustEmbedUnimplementedGoodsRpcServer()
}

// UnimplementedGoodsRpcServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsRpcServer struct {
}

func (UnimplementedGoodsRpcServer) GetGoods(context.Context, *GoodsReq) (*GoodsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoods not implemented")
}
func (UnimplementedGoodsRpcServer) mustEmbedUnimplementedGoodsRpcServer() {}

// UnsafeGoodsRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsRpcServer will
// result in compilation errors.
type UnsafeGoodsRpcServer interface {
	mustEmbedUnimplementedGoodsRpcServer()
}

func RegisterGoodsRpcServer(s grpc.ServiceRegistrar, srv GoodsRpcServer) {
	s.RegisterService(&GoodsRpc_ServiceDesc, srv)
}

func _GoodsRpc_GetGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsRpcServer).GetGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoodsRpc_GetGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsRpcServer).GetGoods(ctx, req.(*GoodsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GoodsRpc_ServiceDesc is the grpc.ServiceDesc for GoodsRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoodsRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goods.GoodsRpc",
	HandlerType: (*GoodsRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGoods",
			Handler:    _GoodsRpc_GetGoods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/goods/goods.proto",
}
