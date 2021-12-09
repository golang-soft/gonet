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
	}
)

func (this *BaseServer) InitConfig(data interface{}) bool {
	config.Init(system.Args.Env, data)
	return true
}
