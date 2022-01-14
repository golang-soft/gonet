package player

import (
	"context"
	"database/sql"
	"fmt"
	"gonet/actor"
	"gonet/base"
	"gonet/common"
	"gonet/common/cluster"
	"gonet/db"
	"gonet/server/cmessage"
	common2 "gonet/server/common"
	"gonet/server/rpc"
	"gonet/server/world"
	"gonet/server/world/wcluster"
	"time"

	"github.com/golang/protobuf/proto"
)

type (
	Player struct {
		actor.Actor

		PlayerData
		m_ItemMgr      IItemMgr
		m_db           *sql.DB
		m_Log          *base.CLog
		m_offlineTimer *common.SimpleTimer
	}
)

func (this *Player) Init() {
	this.Actor.Init()
	this.PlayerData.Init()
	this.RegisterTimer((cluster.OFFLINE_TIME/3)*time.Second, this.UpdateLease) //定时器
	this.m_offlineTimer = common.NewSimpleTimer(5 * 60)
	this.m_db = world.SERVER.GetDB()
	this.m_Log = world.SERVER.GetLog()
	this.m_ItemMgr = &ItemMgr{}
	this.m_ItemMgr.Init(this)
	actor.MGR.BindActor(this)

	//玩家登录
	this.RegisterCall("Login", func(ctx context.Context, gateClusterId uint32, clusterInfo rpc.PlayerClusterInfo) {
		head := this.GetRpcHead(ctx)

		PlayerSimpleList := LoadSimplePlayerDatas(this.AccountId)
		this.PlayerSimpleDataList = PlayerSimpleList

		PlayerDataList := make([]*cmessage.PlayerData, len(PlayerSimpleList))
		this.PlayerIdList = []int64{}
		for i, v := range PlayerSimpleList {
			PlayerDataList[i] = &cmessage.PlayerData{PlayerID: v.PlayerId, PlayerName: v.PlayerName, PlayerGold: int32(v.Gold)}
			this.PlayerIdList = append(this.PlayerIdList, v.PlayerId)
		}

		this.m_Log.Println("玩家[%d]登录成功", this.AccountId)
		this.SetGateClusterId(gateClusterId)
		this.m_PlayerRaft = clusterInfo
		this.SendToClient(head.SocketId, &cmessage.W_C_SelectPlayerResponse{PacketHead: common2.BuildPacketHead(cmessage.MessageID_MSG_W_C_SelectPlayerResponse, rpc.SERVICE_GATESERVER),
			AccountId:  this.AccountId,
			PlayerData: PlayerDataList,
		})
	})

	//玩家登录到游戏
	this.RegisterCall("C_W_Game_LoginRequset", func(ctx context.Context, packet *cmessage.C_W_Game_LoginRequset) {
		head := this.GetRpcHead(ctx)

		nPlayerId := packet.GetPlayerId()
		if !this.SetPlayerId(nPlayerId) {
			this.m_Log.Printf("帐号[%d]登入的玩家[%d]不存在", this.AccountId, nPlayerId)
			return
		}

		//读取玩家数据
		this.LoadPlayerData()
		//加载到地图
		this.AddMap()
		//添加到世界频道
		actor.MGR.SendMsg(rpc.RpcHead{ActorName: "chatmgr"}, "AddPlayerToChannel", this.AccountId, this.GetPlayerId(), int64(-3000), this.GetPlayerName(), this.GetGateClusterId())

		//返回客户端
		this.SendToClient(head.SocketId, &cmessage.C_W_Game_LoginResponse{PacketHead: common2.BuildPacketHead(cmessage.MessageID_MSG_C_W_Game_LoginResponse, rpc.SERVICE_GATESERVER),
			PlayerId: nPlayerId,
		})
	})

	//创建玩家
	this.RegisterCall("C_W_CreatePlayerRequest", func(ctx context.Context, packet *cmessage.C_W_CreatePlayerRequest) {
		head := this.GetRpcHead(ctx)
		fmt.Sprintf("%d", head.SocketId)
		rows, err := this.m_db.Query(fmt.Sprintf("select count(player_id) as player_count from tbl_player where account_id = %d", this.AccountId))
		if err == nil {
			rs := db.Query(rows, err)
			if rs.Next() {
				player_count := rs.Row().Int("player_count")
				if player_count >= 1 {
					this.m_Log.Printf("账号[%d]创建玩家上限", this.AccountId)
					wcluster.SendToClient(this.GetRpcHead(ctx).SrcClusterId, &cmessage.W_C_CreatePlayerResponse{
						PacketHead: common2.BuildPacketHead(cmessage.MessageID_MSG_W_C_CreatePlayerResponse, 0),
						Error:      int32(1),
						PlayerId:   0,
					}, this.GetRpcHead(ctx).SocketId)
				} else {
					wcluster.SendToAccount("W_A_CreatePlayer", this.AccountId, packet.GetPlayerName(), packet.GetSex(), this.GetRpcHead(ctx).SrcClusterId)
				}
			}
		}
	})

	//account创建玩家反馈
	this.RegisterCall("CreatePlayer", func(ctx context.Context, playerId int64, gClusterId uint32, err int) {
		head := this.GetRpcHead(ctx)
		//创建成功
		if err == 0 {
			this.PlayerIdList = []int64{}
			playerSimpleData := LoadSimplePlayerData(playerId)
			this.PlayerSimpleDataList = append(this.PlayerSimpleDataList, playerSimpleData)
			this.PlayerIdList = append(this.PlayerIdList, playerId)
		}

		wcluster.SendToClient(gClusterId, &cmessage.W_C_CreatePlayerResponse{
			PacketHead: common2.BuildPacketHead(cmessage.MessageID_MSG_W_C_CreatePlayerResponse, 0),
			Error:      int32(err),
			PlayerId:   playerId,
		}, head.SocketId)
	})

	//玩家断开链接
	this.RegisterCall("Logout", func(ctx context.Context, accountId int64) {
		world.SERVER.GetLog().Printf("[%d] 断开链接", accountId)
		this.SetGateClusterId(0)
		this.Stop()
		this.LeaveMap()
	})

	this.Actor.Start()
}

func (this *Player) GetDB() *sql.DB {
	return this.m_db
}

func (this *Player) GetLog() *base.CLog {
	return this.m_Log
}

func (this *Player) SendToClient(socketId uint32, packet proto.Message) {
	wcluster.SendToClient(this.GetGateClusterId(), packet, socketId)
}

func (this *Player) UpdateLease() {
	world.SERVER.GetPlayerRaft().Lease(this.m_PlayerRaft.LeaseId)
}
