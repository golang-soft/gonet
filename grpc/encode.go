package grpc

import (
	"bytes"
	"encoding/gob"
	"github.com/golang/protobuf/proto"
	"gonet/base"
	"gonet/server/rpc"
	"strings"
)

//rpc  Marshal
func Marshal(head rpc.RpcHead, funcName string, params ...interface{}) []byte {
	data, _ := marshal(head, funcName, params...)
	return data
}

//rpc  marshal
func marshal(head rpc.RpcHead, funcName string, params ...interface{}) ([]byte, *rpc.RpcPacket) {
	defer func() {
		if err := recover(); err != nil {
			base.TraceCode(err)
		}
	}()

	rpcPacket := &rpc.RpcPacket{FuncName: strings.ToLower(funcName), ArgLen: int32(len(params)), RpcHead: (*rpc.RpcHead)(&head)}
	buf := bytes.NewBuffer([]byte{})
	enc := gob.NewEncoder(buf)
	for _, param := range params {
		enc.Encode(param)
	}
	rpcPacket.RpcBody = buf.Bytes()
	dat, _ := proto.Marshal(rpcPacket)
	return dat, rpcPacket
}

//rpc  MarshalPB
func marshalPB(bitstream *base.BitStream, packet proto.Message) {
	bitstream.WriteString(proto.MessageName(packet))
	buf, _ := proto.Marshal(packet)
	nLen := len(buf)
	bitstream.WriteInt(nLen, 32)
	bitstream.WriteBits(buf, nLen<<3)
}
