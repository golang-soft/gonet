package game

import (
	"context"
	"fmt"
	"gonet/server/message"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

var client message.GreeterClient

type GrpcClient struct {
}

func NewGrpcClient() *GrpcClient {
	return &GrpcClient{}
}

func (this *GrpcClient) DialToServer(port int) chan error {
	err1 := make(chan error, 1)
	conn, err := grpc.Dial(":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接服务端失败: %s", err)
		return err1
	}
	//defer conn.Close()
	// 新建一个客户端
	client = message.NewGreeterClient(conn)
	fmt.Printf("连接grpc服务器成功, port: %d", port)
	return nil
}

func (this *GrpcClient) ConnectToServer(port int) bool {
	// 连接服务器
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*3000))
	defer cancel()
	select {
	case <-this.DialToServer(port):
		return true
	case <-ctx.Done():
		fmt.Println("ConnectToServer Timeout")
		return false
	}
	return true
}

func (this *GrpcClient) ReqOnlineUsers() {
	// 调用服务端函数
	res, err := client.ReqOnlineUserCount(context.Background(), &message.OnlineUserRequest{})
	if err != nil {
		fmt.Printf("调用服务端代码失败: %s", err)
		return
	}
	fmt.Printf("调用成功: %d", res.Count)
}

func (this *GrpcClient) ReqServerId() int64 {
	// 调用服务端函数
	res, err := client.ReqServerId(context.Background(), &message.ServerIdRequest{})
	if err != nil {
		fmt.Printf("调用服务端代码失败: %s", err)
		return 0
	}
	fmt.Printf("调用ReqServerId成功: %d", res.Id)
	return res.Id
}
