package common

import (
	"fmt"
	"gonet/base"
	"gonet/server/rpc"
	"strings"
)

type (
	//集群信息
	ClusterInfo rpc.ClusterInfo

	IClusterInfo interface {
		Id() uint32
		String() string
		ServiceType() rpc.SERVICE
		IpString() string
	}
)

func (this *ClusterInfo) IpString() string {
	return fmt.Sprintf("%s:%d", this.Ip, this.Port)
}

func (this *ClusterInfo) String() string {
	return strings.ToLower(this.Type.String())
}

func (this *ClusterInfo) Id() uint32 {
	return base.ToHash(this.IpString())
}

func (this *ClusterInfo) ServiceType() rpc.SERVICE {
	return this.Type
}
