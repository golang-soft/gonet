package center

import (
	"context"
	"gonet/server/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
)

type GrpcServer struct{}

func (s *GrpcServer) ReqOnlineUserCount(ctx context.Context, in *message.OnlineUserRequest) (*message.OnlineUserResponse, error) {
	data := &message.OnlineUserResponse{}
	data.Count = GetUserManager().getOnlineUserCount()
	return data, nil
}

func (s *GrpcServer) ReqServerId(ctx context.Context, in *message.ServerIdRequest) (*message.ServerIdResponse, error) {
	data := &message.ServerIdResponse{}
	data.Id = GetServerIdManager().getNextServerId()
	return data, nil
}

func StartGrpcServer(port int64) {

	// 监听本地端口
	lis, err := net.Listen("tcp", ":"+strconv.FormatInt(port, 10))
	if err != nil {
		SERVER.m_Log.Printf("监听端口失败: %s", err)
		return
	}
	// 创建gRPC服务器
	s := grpc.NewServer()
	// 注册服务
	message.RegisterGreeterServer(s, &GrpcServer{})
	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		SERVER.m_Log.Printf("开启服务失败: %s", err)
		return
	}

	SERVER.m_Log.Debugf("启动grpc服务成功，端口: %d", port)
}
