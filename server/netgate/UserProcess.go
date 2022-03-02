package netgate

import (
	"context"
	"gonet/actor"
	"gonet/base"
	"gonet/base/logger"
	"gonet/grpc"
	"gonet/network"
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/rpc"
	"strings"
	"time"
)

var (
	C_A_LoginRequest    = strings.ToLower("C_A_LoginRequest")
	C_A_RegisterRequest = strings.ToLower("C_A_RegisterRequest")
)

type (
	UserPrcoess struct {
		actor.Actor
		m_KeyMap map[uint32]*base.Dh
	}

	IUserProcess interface {
		actor.IActor

		CheckClientEx(uint32, string, rpc.RpcHead) bool
		CheckClient(uint32, string, rpc.RpcHead) *AccountInfo
		SwtichSendToWorld(uint32, string, rpc.RpcHead, []byte)
		SwtichSendToAccount(uint32, string, rpc.RpcHead, []byte)
		SwtichSendToZone(uint32, string, rpc.RpcHead, []byte)
		HandleBasicMessage(socketid uint32, packetId uint32, buff []uint8) bool
		addKey(uint32, *base.Dh)
		delKey(uint32)
	}
)

func (this *UserPrcoess) CheckClientEx(sockId uint32, packetName string, head rpc.RpcHead) bool {
	if IsCheckClient(packetName) {
		return true
	}

	accountId := SERVER.GetPlayerMgr().GetAccount(sockId)
	if accountId <= 0 || accountId != head.Id {
		SERVER.GetLog().Fatalf("Old socket communication or viciousness[%d].", sockId)
		return false
	}
	return true
}

func (this *UserPrcoess) CheckClient(sockId uint32, packetName string, head rpc.RpcHead) *AccountInfo {
	pAccountInfo := SERVER.GetPlayerMgr().GetAccountInfo(sockId)
	if pAccountInfo != nil && (pAccountInfo.AccountId <= 0) {
		SERVER.GetLog().Fatalf("Old socket communication or viciousness[%d].", sockId)
		return nil
	}
	return pAccountInfo
}

func (this *UserPrcoess) SwtichSendToWorld(socketId uint32, packetName string, head rpc.RpcHead, buff []byte) {
	pAccountInfo := this.CheckClient(socketId, packetName, head)
	if pAccountInfo != nil {
		head.DestClusterId = pAccountInfo.WClusterId
		head.DestServerType = rpc.SERVICE_WORLDSERVER
		SERVER.GetCluster().Send(head, buff)
	}
}

func (this *UserPrcoess) SwtichSendToWorldDb(socketId uint32, packetName string, head rpc.RpcHead, buff []byte) {
	//pAccountInfo := this.CheckClient(socketId, packetName, head)
	//if pAccountInfo != nil {
	//head.ClusterId = pAccountInfo.WClusterId
	head.SendType = rpc.SEND_BALANCE
	head.DestServerType = rpc.SERVICE_WORLDDBSERVER
	SERVER.GetCluster().Send(head, buff)
	//}
}

func (this *UserPrcoess) SwtichSendToCenter(socketId uint32, packetName string, head rpc.RpcHead, buff []byte) {
	head.SendType = rpc.SEND_POINT
	head.DestServerType = rpc.SERVICE_CENTERSERVER
	SERVER.GetCluster().Send(head, buff)
}

func (this *UserPrcoess) SwtichSendToLogin(socketId uint32, packetName string, head rpc.RpcHead, buff []byte) {
	head.SendType = rpc.SEND_BALANCE
	head.DestServerType = rpc.SERVICE_LOGINSERVER
	SERVER.GetCluster().Send(head, buff)
}

func (this *UserPrcoess) SwtichSendToGrpc(socketId uint32, packetName string, head rpc.RpcHead, buff []byte) {
	head.SendType = rpc.SEND_BALANCE
	head.DestServerType = rpc.SERVICE_GRPCSERVER
	SERVER.GetCluster().Send(head, buff)
}

func (this *UserPrcoess) SwtichSendToAccount(socketId uint32, packetName string, head rpc.RpcHead, buff []byte) {
	if this.CheckClientEx(socketId, packetName, head) == true {
		head.SendType = rpc.SEND_BALANCE
		head.DestServerType = rpc.SERVICE_ACCOUNTSERVER
		SERVER.GetCluster().Send(head, buff)
	}
}

func (this *UserPrcoess) SwtichSendToZone(socketId uint32, packetName string, head rpc.RpcHead, buff []byte) {
	pAccountInfo := this.CheckClient(socketId, packetName, head)
	if pAccountInfo != nil {
		head.DestClusterId = pAccountInfo.ZClusterId
		head.DestServerType = rpc.SERVICE_ZONESERVER
		SERVER.GetCluster().Send(head, buff)
	}
}

func (this *UserPrcoess) DisconnectClient(stream *base.BitStream, socketid uint32) {
	SERVER.GetPlayerMgr().SendMsg(rpc.RpcHead{}, "DEL_ACCOUNT", uint32(stream.ReadInt(32)))
	this.SendMsg(rpc.RpcHead{}, "DISCONNECT", socketid)
}

func (this *UserPrcoess) UpdateHeardTime(socketid uint32) bool {
	if SERVER.CheckIsWebsocket() {
		if SERVER.WebsocketModeIsGorilla() {
			client := SERVER.GetWebSocketServerG().GetClientById(socketid)
			if client == nil {
				return true
			}
			if client != nil {
				client.SetLastHeardTime(int(time.Now().Unix()) + network.HEART_TIME_OUT)
			}
		} else {
			client := SERVER.GetWebSocketServer().GetClientById(socketid)
			if client == nil {
				return true
			}
			if client != nil {
				client.SetLastHeardTime(int(time.Now().Unix()) + network.HEART_TIME_OUT)
			}
		}

	} else {
		client := SERVER.GetServer().GetClientById(socketid)
		if client == nil {
			return true
		}
		if client != nil {
			client.SetLastHeardTime(int(time.Now().Unix()) + network.HEART_TIME_OUT)
		}
	}
	return true
}

func (this *UserPrcoess) HandleBasicMessage(socketid uint32, packetId uint32, buff []uint8) bool {
	//客户端主动断开
	if packetId == network.DISCONNECTINT {
		//断开客户端的链接
		stream := base.NewBitStream(buff, len(buff))
		stream.ReadInt(32)
		this.DisconnectClient(stream, socketid)
	} else if packetId == network.HEART_PACKET {
		//心跳netsocket做处理，这里不处理
		SERVER.GetLog().Debugf("网关收到心跳包, %d", packetId)

		if !this.UpdateHeardTime(socketid) {
			SERVER.GetLog().Errorf("玩家[%d]更新心跳包失败", socketid)
		}

	} else {
		//未知的消息
		SERVER.GetLog().Errorf("包解析错误, 未知的消息 socket=%d, packetId = %d", socketid, packetId)
	}

	return true
}

func (this *UserPrcoess) PacketFunc(packet1 rpc.Packet) bool {
	buff := packet1.Buff
	socketid := packet1.Id
	packetId, data := common.Decode(buff)
	packet := common.GetPakcet(packetId)
	if packet == nil {
		return this.HandleBasicMessage(socketid, packetId, buff)
	}

	//获取配置的路由地址
	err := common.UnmarshalText(packet, data)
	if err != nil {
		SERVER.GetLog().Printf("包解析错误2  socket=%d", socketid)
		return true
	}
	//destServerType := packet.(common.Packet).GetPacketHead().DestServerType

	packetHead := packet.(common.Packet).GetPacketHead()
	//packetHead.DestServerType = destServerType
	if packetHead == nil /*|| packetHead.Ckx != common.Default_Ipacket_Ckx || packetHead.Stx != common.Default_Ipacket_Stx*/ {
		SERVER.GetLog().Printf("(A)致命的越界包,已经被忽略 socket=%d", socketid)
		return true
	}

	packetName := common.GetMessageName(packet)
	head := rpc.RpcHead{Id: int64(packetHead.Id), SrcClusterId: SERVER.GetCluster().Id()}
	if packetName == C_A_LoginRequest {
		head.DestClusterId = socketid
	} else if packetName == C_A_RegisterRequest {
		head.DestClusterId = socketid
	}
	head.SocketId = socketid
	//解析整个包
	//if packetHead.DestServerType == smessage.SERVICE_WORLDSERVER {
	//	this.SwtichSendToWorld(socketid, packetName, head, rpc.Marshal(head, packetName, packet))
	//} else if packetHead.DestServerType == smessage.SERVICE_ACCOUNTSERVER {
	//	this.SwtichSendToAccount(socketid, packetName, head, rpc.Marshal(head, packetName, packet))
	//} else if packetHead.DestServerType == smessage.SERVICE_ZONESERVER {
	//	this.SwtichSendToZone(socketid, packetName, head, rpc.Marshal(head, packetName, packet))
	//} else {
	//	this.Actor.PacketFunc(rpc.Packet{Id: socketid, Buff: rpc.Marshal(head, packetName, packet)})
	//}
	dest := common.GetPacketByName(packetName)
	switch dest {
	case rpc.SERVICE_NONE:
		{
			this.Actor.PacketFunc(rpc.Packet{Id: socketid, Buff: grpc.Marshal(head, packetName, packet)})
		}
	case rpc.SERVICE_CLIENT:
		{

		}
	case rpc.SERVICE_GATESERVER:
		{
			this.Actor.PacketFunc(rpc.Packet{Id: socketid, Buff: grpc.Marshal(head, packetName, packet)})
		}
	case rpc.SERVICE_ACCOUNTSERVER:
		{
			this.SwtichSendToAccount(socketid, packetName, head, grpc.Marshal(head, packetName, packet))
		}
	case rpc.SERVICE_WORLDSERVER:
		{
			this.SwtichSendToWorld(socketid, packetName, head, grpc.Marshal(head, packetName, packet))
		}
	case rpc.SERVICE_ZONESERVER:
		{
			this.SwtichSendToZone(socketid, packetName, head, grpc.Marshal(head, packetName, packet))
		}
	case rpc.SERVICE_CENTERSERVER:
		{
			this.SwtichSendToCenter(socketid, packetName, head, grpc.Marshal(head, packetName, packet))
		}
	case rpc.SERVICE_LOGINSERVER:
		{
			this.SwtichSendToLogin(socketid, packetName, head, grpc.Marshal(head, packetName, packet))
		}
	case rpc.SERVICE_GRPCSERVER:
		{
			this.SwtichSendToGrpc(socketid, packetName, head, grpc.Marshal(head, packetName, packet))
		}
	case rpc.SERVICE_WORLDDBSERVER:
		{
			//无
		}
	default:
		{
			this.Actor.PacketFunc(rpc.Packet{Id: socketid, Buff: grpc.Marshal(head, packetName, packet)})
		}
	}
	//switch packetName {
	//case base.ToLower("C_W_Game_LoginRequset"):
	//case base.ToLower("C_W_CreatePlayerRequest"):
	//	{
	//		this.SwtichSendToWorld(socketid, packetName, head, grpc.Marshal(head, packetName, packet))
	//	}
	//	break
	//case base.ToLower("C_A_RegisterRequest"):
	//	{
	//		this.SwtichSendToAccount(socketid, packetName, head, grpc.Marshal(head, packetName, packet))
	//	}
	//	break
	//default:
	//	{
	//		this.Actor.PacketFunc(rpc.Packet{Id: socketid, Buff: grpc.Marshal(head, packetName, packet)})
	//	}
	//}

	return true
}

func (this *UserPrcoess) addKey(SocketId uint32, pDh *base.Dh) {
	this.m_KeyMap[SocketId] = pDh
}

func (this *UserPrcoess) delKey(SocketId uint32) {
	delete(this.m_KeyMap, SocketId)
}

func (this *UserPrcoess) Init() {
	this.Actor.Init()
	this.m_KeyMap = map[uint32]*base.Dh{}
	this.RegisterCall("C_G_LogoutRequest", func(ctx context.Context, accountId int, UID int) {
		SERVER.GetLog().Printf("logout Socket:%d Account:%d UID:%d ", this.GetRpcHead(ctx).SocketId, accountId, UID)
		SERVER.GetPlayerMgr().SendMsg(rpc.RpcHead{}, "DEL_ACCOUNT", this.GetRpcHead(ctx).SocketId)
		SendToClient(this.GetRpcHead(ctx).SocketId, &cmessage.C_G_LogoutResponse{PacketHead: common.BuildPacketHead(0, 0)})
	})

	this.RegisterCall("C_G_LoginResquest", func(ctx context.Context, packet *cmessage.C_G_LoginResquest) {
		head := this.GetRpcHead(ctx)
		dh := base.Dh{}
		dh.Init()
		dh.ExchangePubk(packet.GetKey())
		this.addKey(head.SocketId, &dh)
		SendToClient(head.SocketId, &cmessage.G_C_LoginResponse{PacketHead: common.BuildPacketHead(cmessage.MessageID_MSG_G_C_LoginResponse, 0), Key: dh.PubKey()})
	})

	this.RegisterCall("C_A_LoginRequest", func(ctx context.Context, packet *cmessage.C_A_LoginRequest) {
		head := this.GetRpcHead(ctx)
		dh, bEx := this.m_KeyMap[head.SocketId]
		if bEx {
			if dh.ShareKey() == packet.GetKey() {
				this.delKey(head.SocketId)
				//head.Id = int64(base.ToHash(packet.AccountName))
				this.SwtichSendToAccount(head.SocketId, "C_A_LoginRequest", head, grpc.Marshal(head, base.ToLower("C_A_LoginRequest"), packet))
			} else {
				SERVER.GetLog().Println("client key cheat", dh.ShareKey(), packet.GetKey())
			}
		} else {
			logger.Debug("找不到對應的客戶端 socketid : %d", head.SocketId)
		}
	})

	this.RegisterCall("DISCONNECT", func(ctx context.Context, socketid uint32) {
		this.delKey(socketid)
	})

	this.RegisterCall("HeartPacket", func(ctx context.Context, packet *cmessage.W_C_Test) {
		head := this.GetRpcHead(ctx)
		logger.Debug(head)
	})

	this.RegisterCall("W_C_Test", func(ctx context.Context, packet *cmessage.W_C_Test) {
		head := this.GetRpcHead(ctx)
		dh := base.Dh{}
		dh.Init()
		this.addKey(head.SocketId, &dh)
		//SendToClient(head.SocketId, &message.G_C_LoginResponse{PacketHead: message.BuildPacketHead(0, 0), Key: dh.PubKey()})

		this.SwtichSendToWorldDb(head.SocketId, base.ToLower("W_C_Test"), head, grpc.Marshal(head, base.ToLower("W_C_Test"), packet))

	})

	this.RegisterCall("AttackReq", func(ctx context.Context, packet *cmessage.AttackReq) {
		head := this.GetRpcHead(ctx)
		SendToClient(head.SocketId, &cmessage.AttackResp{PacketHead: common.BuildPacketHead(cmessage.MessageID_MSG_AttackResp, rpc.SERVICE_NONE)})
		this.SwtichSendToWorld(head.SocketId, base.ToLower("AttackReq"), head, grpc.Marshal(head, base.ToLower("AttackReq"), packet))
	})

	this.RegisterCall("GameTimeReq", func(ctx context.Context, packet *cmessage.GameTimeReq) {
		head := this.GetRpcHead(ctx)
		SendToClient(head.SocketId, &cmessage.GameTimeResp{PacketHead: common.BuildPacketHead(cmessage.MessageID_MSG_GameTimeResp, rpc.SERVICE_NONE)})
	})

	this.Actor.Start()
}
