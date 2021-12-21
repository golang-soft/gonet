package server

import (
	"gonet/base/config"
	"gonet/base/system"
	"gonet/server/game"
)

type (
	BaseServer struct {
		id            int64
		M_pGrpcClient *game.GrpcClient
	}

	IBaseServer interface {
		InitConfig(data interface{}) bool
		ConnectCenter(addr string) bool
		SetId(id int64)
		GetId() int64
	}
)

func (this *BaseServer) InitConfig(data interface{}) bool {
	config.Init(system.Args.Env, data)
	return true
}

func (this *BaseServer) ReadTableData(data interface{}) bool {
	config.Init(system.Args.Env, data)
	return true
}

func (this *BaseServer) ConnectCenter(addr string) bool {

	return true
}

func (this *BaseServer) ListServerByType() []string {

	return nil
}

func (this *BaseServer) SetId(id int64) {
	this.id = id
}
func (this *BaseServer) GetId() int64 {
	return this.id
}
