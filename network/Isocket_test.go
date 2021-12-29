package network_test

import (
	"gonet/grpc"
	"gonet/network"
	"gonet/server/rpc"
	"testing"
)

var (
	m_pInBuffer  []byte
	m_pInBuffer1 []byte
	nTimes       = 1000
)

const (
	TCP_END   = "💞♡"   //解决tpc粘包半包,结束标志
	ARRAY_LEN = 100000 //800kb 100 * 1000 * 8
)

func TestEndFlag(t *testing.T) {
	t.Log("c语言结束标志", []byte(TCP_END))
	buff := []byte{}
	packetParse := network.NewPacketParser(network.PacketConfig{})
	for j := 0; j < 1; j++ {
		buff = append(buff, packetParse.Write(grpc.Marshal(rpc.RpcHead{}, "test1", [ARRAY_LEN]int64{1, 2, 3, 4, 5, 6}))...)
	}
	for i := 0; i < nTimes; i++ {
		packetParse.Read(buff)
	}
}

func SetTcpEnd1(buff []byte) []byte {
	buff = append(buff, []byte(TCP_END)...)
	return buff
}
