package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"gonet/actor"
	"gonet/base"
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/game/lmath"
	"gonet/server/rpc"
	"time"
)

type (
	EventProcess struct {
		SuperEventProcess
		//Client      *network.ClientSocket

		AccountId   int64
		PlayerId    int64
		AccountName string
		PassWd      string
		SimId       int64
		Pos         lmath.Point3F
		Rot         lmath.Point3F
		m_Dh        base.Dh
	}

	IEventProcess interface {
		actor.IActor
		LoginGame()
		LoginAccount()
		SendPacket(proto.Message)
	}
)

var (
	id int32
)

func ToSlat(accountName string, pwd string) string {
	return fmt.Sprintf("%s__%s", accountName, pwd)
}

func ToCrc(accountName string, pwd string, buildNo string, nKey int64) uint32 {
	return base.GetMessageCode1(fmt.Sprintf("%s_%s_%s_%d", accountName, pwd, buildNo, nKey))
}

func (this *EventProcess) Init() {
	this.Actor.Init()
	this.Pos = lmath.Point3F{1, 1, 1}
	this.m_Dh.Init()
	//this.RegisterTimer((network.HEART_TIME_OUT/6)*time.Second, this.Update) //定时器

	this.RegisterCall("W_C_SelectPlayerResponse", func(ctx context.Context, packet *cmessage.W_C_SelectPlayerResponse) {
		this.AccountId = packet.GetAccountId()
		nLen := len(packet.GetPlayerData())
		//fmt.Println(len(packet.PlayerData), this.AccountId, packet.PlayerData)
		if nLen == 0 {
			packet1 := &cmessage.C_W_CreatePlayerRequest{PacketHead: common.BuildPacketHead(cmessage.MessageID_MSG_C_W_CreatePlayerRequest, rpc.SERVICE_GATESERVER),
				PlayerName: "我是大坏蛋",
				Sex:        int32(0)}
			this.SendPacket(this, packet1)
		} else {
			this.PlayerId = packet.GetPlayerData()[0].GetPlayerID()
			//this.LoginGame()
			e := eventmanager.GetEvent("LoginEvent")
			if e != nil {
				(*e).DoEvent(this)
			}
		}
	})

	this.RegisterCall("W_C_CreatePlayerResponse", func(ctx context.Context, packet *cmessage.W_C_CreatePlayerResponse) {
		if packet.GetError() == 0 {
			this.PlayerId = packet.GetPlayerId()
			m_Log.Debugf("玩家 %d 登录游戏", this.PlayerId)
			//this.LoginGame()
			e := eventmanager.GetEvent("LoginEvent")
			if e != nil {
				(*e).DoEvent(this)
			}
		} else { //创建失败

		}
	})

	this.RegisterCall("C_W_Game_LoginResponse", func(ctx context.Context, packet *cmessage.C_W_Game_LoginResponse) {
		this.PlayerId = packet.GetPlayerId()
		if this.Robot != nil {
			this.Robot.status = ROBOT_PLAYING
			m_Log.Debugf("玩家 %d 的状态改变 %d", this.PlayerId, ROBOT_PLAYING)
		}
	})

	this.RegisterCall("G_C_LoginResponse", func(ctx context.Context, packet *cmessage.G_C_LoginResponse) {
		this.m_Dh.ExchangePubk(packet.GetKey())
		e := eventmanager.GetEvent("LoginAccountEvent")
		if e != nil {
			(*e).DoEvent(this)
		}
		//this.LoginAccount()
	})

	this.RegisterCall("A_C_LoginResponse", func(ctx context.Context, packet *cmessage.A_C_LoginResponse) {
		if packet.GetError() == base.ACCOUNT_NOEXIST {
			packet1 := &cmessage.C_A_RegisterRequest{PacketHead: common.BuildPacketHead(0, rpc.SERVICE_GATESERVER),
				AccountName: packet.AccountName, Password: this.PassWd}
			this.SendPacket(this, packet1)
		} else if packet.GetError() == base.PASSWORD_ERROR {
			fmt.Println("账号【", packet.GetAccountName(), "】密码错误")
		}
	})

	this.RegisterCall("A_C_RegisterResponse", func(ctx context.Context, packet *cmessage.A_C_RegisterResponse) {
		//注册失败
		if packet.GetError() != 0 {
		}
	})

	this.RegisterCall("W_C_ChatMessage", func(ctx context.Context, packet *cmessage.W_C_ChatMessage) {
		fmt.Println("收到【", packet.GetSenderName(), "】发送的消息[", packet.GetMessage()+"]")
	})

	//map
	this.RegisterCall("Z_C_LoginMap", func(ctx context.Context, packet *cmessage.Z_C_LoginMap) {
		this.SimId = packet.GetId()
		this.Pos = lmath.Point3F{packet.GetPos().GetX(), packet.GetPos().GetY(), packet.GetPos().GetZ()}
		this.Rot = lmath.Point3F{0, 0, packet.GetRotation()}
		//fmt.Println("login map")
	})

	this.RegisterCall("Z_C_ENTITY", func(ctx context.Context, packet *cmessage.Z_C_ENTITY) {
		for _, v := range packet.EntityInfo {
			if v.Data != nil {
				if v.Data.RemoveFlag {
					fmt.Printf("Z_C_ENTITY_DATA  destory:[%d], [%d], [%t]\n", v.GetId(), v.Data.Type, v.Data.RemoveFlag)
					continue
				}
				fmt.Printf("Z_C_ENTITY_DATA :[%d], [%d], [%t]\n", v.GetId(), v.Data.Type, v.Data.RemoveFlag)
			}
			if v.Move != nil {
				if v.Id == this.SimId {
					this.Pos = lmath.Point3F{v.Move.GetPos().GetX(), v.Move.GetPos().GetY(), v.Move.GetPos().GetZ()}
					this.Rot = lmath.Point3F{0, 0, v.Move.GetRotation()}
				}
				fmt.Printf("Z_C_ENTITY_MOVE :[%d], Pos:[x:%f, y:%f, z:%f], Rot[%f]\n", v.GetId(), v.Move.GetPos().GetX(), v.Move.GetPos().GetY(), v.Move.GetPos().GetZ(), v.Move.GetRotation())
			}
		}
	})

	this.RegisterCall("AttackResp", func(ctx context.Context, packet *cmessage.AttackResp) {
		fmt.Println("AttackResp")
		e := eventmanager.GetEvent("AttackEvent")
		if e != nil {
			(*e).HandleEvent(e, packet)
		}
	})

	//链接断开
	this.RegisterCall("DISCONNECT", func(ctx context.Context, socketId uint32) {
		this.Stop()
	})
	this.Actor.Start()
}

func (this *EventProcess) Move(yaw float32, time float32) {
	packet1 := &cmessage.C_Z_Move{PacketHead: common.BuildPacketHead(cmessage.MessageID_MSG_C_Z_Move, rpc.SERVICE_GATESERVER),
		Move: &cmessage.C_Z_Move_Move{Mode: 0, Normal: &cmessage.C_Z_Move_Move_Normal{Pos: &cmessage.Point3F{X: this.Pos.X, Y: this.Pos.Y, Z: this.Pos.Z}, Yaw: yaw, Duration: time}}}
	this.SendPacket(this, packet1)
}

func (this *EventProcess) Update() {
	packet1 := &cmessage.HeartPacket{Time: time.Now().Unix()}
	this.SendPacket(this, packet1)
	m_Log.Debugf("发送心跳包.........")
}
