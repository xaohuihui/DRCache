package initialize

import (
	. "DRCache/common"
	pb "DRCache/rpcpb"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

// author: xaohuihui
// datetime: 2022/2/18 14:57:27
// software: GoLand
var r = RedisController{}

type GroupCacheServer struct{}

func (p *GroupCacheServer) SetRemoteValue(ctx context.Context, in *pb.SetParam) (*pb.SetRes, error) {
	err := r.SetVal(in.Key, in.Value, time.Duration(in.Timeout) * time.Second)
	if err != nil {
		return nil , err
	}
	return &pb.SetRes{Ok: true, Mes: []byte("set success")}, nil
}

func (p *GroupCacheServer) GetRemoteValue(ctx context.Context, in *pb.GetParam) (*pb.GetRes, error) {
	res, err := r.GetVal(in.Key)
	if err != nil {
		return nil, err
	}
	return &pb.GetRes{Value: res}, nil
}

func InitRPCServer(network, address string) {
	lis, err := net.Listen(network, address)
	if err != nil {
		fmt.Printf("fieled to listen: %v", err)
	}
	s := grpc.NewServer()                               // 创建gRPC服务器
	pb.RegisterGroupCacheServer(s, &GroupCacheServer{}) // 在gRPC服务端注册服务
	reflection.Register(s)                              //在给定的gRPC服务器上注册服务器反射服务
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	err = s.Serve(lis)
	if err != nil {
		zap.L().Error(fmt.Sprintf("[InitRPCServer] 初始化RPC服务失败[%v]", err.Error()))
	}
}
