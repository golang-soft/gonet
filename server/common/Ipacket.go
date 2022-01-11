package common

import (
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"gonet/base"
	"gonet/server/cmessage"
	"gonet/server/rpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"reflect"
	"strings"
)

var (
	Packet_CreateFactorStringMap map[string]func() proto.Message
	Packet_CreateFactorMap       map[uint32]func() proto.Message
	Packet_CrcNamesMap           map[uint32]string
)

const (
	Default_Ipacket_Stx int32 = 0x27
	Default_Ipacket_Ckx int32 = 0x72
)

type (
	//获取包头
	Packet interface {
		GetPacketHead() *cmessage.Ipacket
	}
)

func BuildPacketHead(id cmessage.MessageID, destservertype rpc.SERVICE) *cmessage.Ipacket {
	ipacket := &cmessage.Ipacket{
		Stx: Default_Ipacket_Stx,
		/*DestServerType: smessage.SERVICE(destservertype),*/
		Ckx: Default_Ipacket_Ckx,
		Id:  id,
	}
	return ipacket
}

func GetMessageName(packet proto.Message) string {
	sType := strings.ToLower(proto.MessageName(packet))
	index := strings.Index(sType, ".")
	if index != -1 {
		sType = sType[index+1:]
	}
	return sType
}

func Encode(packet proto.Message) []byte {
	packetId := base.GetMessageCode1(GetMessageName(packet))
	buff, _ := proto.Marshal(packet)
	data := append(base.IntToBytes(int(packetId)), buff...)
	return data
}

func Decode(buff []byte) (uint32, []byte) {
	packetId := uint32(base.BytesToInt(buff[0:4]))
	return packetId, buff[4:]
}

func RegisterPacket(packet proto.Message) {
	packetName := GetMessageName(packet)
	val := reflect.ValueOf(packet).Elem()
	packetFunc := func() proto.Message {
		packet := reflect.New(val.Type())
		packet.Elem().Field(3).Set(val.Field(3))
		//packet.Elem().Set(val)
		return packet.Interface().(proto.Message)
	}

	Packet_CreateFactorStringMap[packetName] = packetFunc
	Packet_CreateFactorMap[base.GetMessageCode1(packetName)] = packetFunc
	glog.Infof("注册协议 %s %d", packetName, base.GetMessageCode1(packetName))
}

func GetPakcet(packetId uint32) proto.Message {
	packetFunc, exist := Packet_CreateFactorMap[packetId]
	if exist {
		return packetFunc()
	}

	return nil
}

func GetPakcetName(packetId uint32) string {
	return Packet_CrcNamesMap[packetId]
}

func UnmarshalText(packet proto.Message, packetBuf []byte) error {
	return proto.Unmarshal(packetBuf, packet)
}

func init() {
	Packet_CreateFactorStringMap = make(map[string]func() proto.Message)
	Packet_CreateFactorMap = make(map[uint32]func() proto.Message)
	Packet_CrcNamesMap = make(map[uint32]string)
}

//统计crc对应string
func initCrcNames() {
	protoFiles := []protoreflect.MessageDescriptors{}
	//protoFiles = append(protoFiles, File_message_proto.Messages())
	//protoFiles = append(protoFiles, File_client_proto.Messages())
	//protoFiles = append(protoFiles, File_game_proto.Messages())
	for _, v := range protoFiles {
		for i := 0; i < v.Len(); i++ {
			packetName := strings.ToLower(string(v.Get(i).Name()))
			crcVal := base.GetMessageCode1(packetName)
			Packet_CrcNamesMap[crcVal] = packetName
		}
	}
}

//网关防火墙
func Init() {
	initCrcNames()
	//注册消息
	//PacketHead 中的 DestServerType 决定转发到那个服务器
	RegisterPacket(&cmessage.C_A_LoginRequest{PacketHead: BuildPacketHead(0, rpc.SERVICE_GATESERVER)})
	RegisterPacket(&cmessage.AttackReq{PacketHead: BuildPacketHead(0, rpc.SERVICE_GATESERVER)})
	RegisterPacket(&cmessage.C_G_LoginResquest{PacketHead: BuildPacketHead(0, rpc.SERVICE_GATESERVER)})
	RegisterPacket(&cmessage.C_A_RegisterRequest{PacketHead: BuildPacketHead(0, rpc.SERVICE_ACCOUNTSERVER)})
	RegisterPacket(&cmessage.C_G_LogoutResponse{PacketHead: BuildPacketHead(0, rpc.SERVICE_GATESERVER)})
	RegisterPacket(&cmessage.C_W_CreatePlayerRequest{PacketHead: BuildPacketHead(0, rpc.SERVICE_WORLDSERVER)})
	RegisterPacket(&cmessage.C_W_Game_LoginRequset{PacketHead: BuildPacketHead(0, rpc.SERVICE_WORLDSERVER)})
	RegisterPacket(&cmessage.W_C_Test{PacketHead: BuildPacketHead(0, rpc.SERVICE_GATESERVER)})
	RegisterPacket(&cmessage.C_W_ChatMessage{PacketHead: BuildPacketHead(0, rpc.SERVICE_WORLDSERVER)})

	RegisterPacket(&cmessage.C_Z_LoginCopyMap{PacketHead: BuildPacketHead(0, rpc.SERVICE_ZONESERVER)})
	RegisterPacket(&cmessage.C_Z_Move{PacketHead: BuildPacketHead(0, rpc.SERVICE_ZONESERVER)})
	RegisterPacket(&cmessage.C_Z_Skill{PacketHead: BuildPacketHead(0, rpc.SERVICE_ZONESERVER)})
}

//client消息回调
func InitClient() {
	initCrcNames()
	//注册消息
	RegisterPacket(&cmessage.W_C_SelectPlayerResponse{})
	RegisterPacket(&cmessage.W_C_CreatePlayerResponse{})
	RegisterPacket(&cmessage.Z_C_LoginMap{})
	RegisterPacket(&cmessage.Z_C_ENTITY{})
	RegisterPacket(&cmessage.W_C_ChatMessage{})
	RegisterPacket(&cmessage.A_C_LoginResponse{})
	RegisterPacket(&cmessage.A_C_RegisterResponse{})
	RegisterPacket(&cmessage.G_C_LoginResponse{})
}
