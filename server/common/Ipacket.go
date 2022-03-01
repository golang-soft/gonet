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
	Packet_CrcDestMap            map[uint32]rpc.SERVICE
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
	//head := packet.(Packet).GetPacketHead()

	val := reflect.ValueOf(packet).Elem()
	packetFunc := func() proto.Message {
		packet := reflect.New(val.Type())
		packet.Elem().Field(3).Set(val.Field(3))
		//packet.Elem().Set(val)
		return packet.Interface().(proto.Message)
	}

	Packet_CreateFactorStringMap[packetName] = packetFunc
	Packet_CreateFactorMap[base.GetMessageCode1(packetName)] = packetFunc
	glog.Infof("注册协议 %s %v", packetName, base.GetMessageCode1(packetName))
}

func RegisterPacket2(packet proto.Message, destservertype rpc.SERVICE) {
	packetName := GetMessageName(packet)
	//head := packet.(Packet).GetPacketHead()

	val := reflect.ValueOf(packet).Elem()
	packetFunc := func() proto.Message {
		packet := reflect.New(val.Type())
		packet.Elem().Field(3).Set(val.Field(3))
		//packet.Elem().Set(val)
		return packet.Interface().(proto.Message)
	}

	Packet_CreateFactorStringMap[packetName] = packetFunc
	Packet_CreateFactorMap[base.GetMessageCode1(packetName)] = packetFunc
	Packet_CrcDestMap[base.GetMessageCode1(packetName)] = destservertype

	glog.Infof("注册协议 %s %v", packetName, base.GetMessageCode1(packetName))
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

func GetPacketByName(packetName string) rpc.SERVICE {
	packetId := base.GetMessageCode1(packetName)
	if packetId > 0 {
		dest, exist := Packet_CrcDestMap[packetId]
		if exist {
			return dest
		}
	}
	return rpc.SERVICE_NONE
}

func UnmarshalText(packet proto.Message, packetBuf []byte) error {
	return proto.Unmarshal(packetBuf, packet)
}

func init() {
	Packet_CreateFactorStringMap = make(map[string]func() proto.Message)
	Packet_CreateFactorMap = make(map[uint32]func() proto.Message)
	Packet_CrcNamesMap = make(map[uint32]string)
	Packet_CrcDestMap = make(map[uint32]rpc.SERVICE)
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
	RegisterPacket2(&cmessage.C_A_LoginRequest{PacketHead: BuildPacketHead(0, rpc.SERVICE_GATESERVER)}, rpc.SERVICE_GATESERVER)
	RegisterPacket2(&cmessage.AttackReq{PacketHead: BuildPacketHead(0, rpc.SERVICE_GATESERVER)}, rpc.SERVICE_GATESERVER)
	RegisterPacket2(&cmessage.C_G_LoginResquest{PacketHead: BuildPacketHead(0, rpc.SERVICE_GATESERVER)}, rpc.SERVICE_GATESERVER)
	RegisterPacket2(&cmessage.C_A_RegisterRequest{PacketHead: BuildPacketHead(0, rpc.SERVICE_ACCOUNTSERVER)}, rpc.SERVICE_ACCOUNTSERVER)
	RegisterPacket2(&cmessage.C_G_LogoutResponse{PacketHead: BuildPacketHead(0, rpc.SERVICE_GATESERVER)}, rpc.SERVICE_GATESERVER)
	RegisterPacket2(&cmessage.C_W_CreatePlayerRequest{PacketHead: BuildPacketHead(0, rpc.SERVICE_WORLDSERVER)}, rpc.SERVICE_WORLDSERVER)
	RegisterPacket2(&cmessage.C_W_Game_LoginRequset{PacketHead: BuildPacketHead(0, rpc.SERVICE_WORLDSERVER)}, rpc.SERVICE_WORLDSERVER)
	RegisterPacket2(&cmessage.W_C_Test{PacketHead: BuildPacketHead(0, rpc.SERVICE_GATESERVER)}, rpc.SERVICE_GATESERVER)
	RegisterPacket2(&cmessage.C_W_ChatMessage{PacketHead: BuildPacketHead(0, rpc.SERVICE_WORLDSERVER)}, rpc.SERVICE_WORLDSERVER)
	RegisterPacket2(&cmessage.AttackReq{PacketHead: BuildPacketHead(0, rpc.SERVICE_WORLDSERVER)}, rpc.SERVICE_WORLDSERVER)
	//RegisterPacket2(&cmessage.HeartPacket{PacketHead: BuildPacketHead(0, rpc.SERVICE_GATESERVER)}, rpc.SERVICE_GATESERVER)

	RegisterPacket2(&cmessage.C_Z_LoginCopyMap{PacketHead: BuildPacketHead(0, rpc.SERVICE_ZONESERVER)}, rpc.SERVICE_ZONESERVER)
	RegisterPacket2(&cmessage.C_Z_Move{PacketHead: BuildPacketHead(0, rpc.SERVICE_ZONESERVER)}, rpc.SERVICE_ZONESERVER)
	RegisterPacket2(&cmessage.C_Z_Skill{PacketHead: BuildPacketHead(0, rpc.SERVICE_ZONESERVER)}, rpc.SERVICE_ZONESERVER)
}

//client消息回调
func InitClient() {
	initCrcNames()
	//注册消息
	RegisterPacket2(&cmessage.W_C_SelectPlayerResponse{}, rpc.SERVICE_GATESERVER)
	RegisterPacket2(&cmessage.W_C_CreatePlayerResponse{}, rpc.SERVICE_GATESERVER)
	RegisterPacket2(&cmessage.Z_C_LoginMap{}, rpc.SERVICE_GATESERVER)
	RegisterPacket2(&cmessage.Z_C_ENTITY{}, rpc.SERVICE_GATESERVER)
	RegisterPacket2(&cmessage.W_C_ChatMessage{}, rpc.SERVICE_GATESERVER)
	RegisterPacket2(&cmessage.A_C_LoginResponse{}, rpc.SERVICE_GATESERVER)
	RegisterPacket2(&cmessage.A_C_RegisterResponse{}, rpc.SERVICE_GATESERVER)
	RegisterPacket2(&cmessage.G_C_LoginResponse{}, rpc.SERVICE_GATESERVER)
	RegisterPacket2(&cmessage.C_W_Game_LoginResponse{}, rpc.SERVICE_GATESERVER)
}
