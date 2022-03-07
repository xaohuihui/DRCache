package common

import (
	pb "DRCache/rpcpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

// author: xaohuihui
// datetime: 2022/3/7 17:04:51
// software: GoLand


func CreateRPCClient(target string) {
	// 连接服务器
	conn, err := grpc.Dial(target)
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGroupCacheClient(conn)
	// 调用服务端的获取远程节点中的缓存值
	r, err := c.GetRemoteValue(context.Background(), &pb.GetParam{Key: "q1mi"})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
	}
	fmt.Printf("Greeting: %s !\n", r.Value)
}
