package server

import (
	"gonet/base/config"
	"gonet/base/system"
)

type (
	BaseServer struct {
	}

	IBaseServer interface {
		InitConfig(data interface{}) bool
		ConnectCenter(addr string) bool
	}
)

func (this *BaseServer) InitConfig(data interface{}) bool {
	config.Init(system.Args.Env, data)
	return true
}

func (this *BaseServer) ConnectCenter(addr string) bool {

	return true
}

func (this *BaseServer) ListServerByType() []string {

	return nil
}
