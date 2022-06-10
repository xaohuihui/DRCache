package common

import (
	pb "DRCache/rpcpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

// author: xaohuihui
// datetime: 2022/3/7 17:04:51
// software: GoLand


func CreateRPCClient(target, key string) ([]byte, error) {
	// 连接服务器
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGroupCacheClient(conn)

	// 初始化上下文，设置请求超时时间为1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	// 延迟关闭请求会话
	defer cancel()

	// 调用服务端的获取远程节点中的缓存值
	r, err := c.GetRemoteValue(ctx, &pb.GetParam{Key: key})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return nil, err
	}
	fmt.Printf("Greeting: %s !\n", r.Value)
	return r.Value, nil
}
