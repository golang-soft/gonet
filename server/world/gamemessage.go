package world

import (
	"context"
	"gonet/actor"
	"gonet/base/logger"
	"gonet/network"
	"gonet/server/cmessage"
)

type (
	GameProcess struct {
		actor.Actor

		Client *network.ClientSocket
	}

	IGameProcess interface {
		actor.IActor
	}
)

func (this *GameProcess) AttackReq(ctx context.Context, packet *cmessage.AttackReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("AttackReq %v", head)
}

func (this *GameProcess) GameStartReq(ctx context.Context, packet *cmessage.GameStartReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("GameStartReq %v", head)
}
func (this *GameProcess) GameEndReq(ctx context.Context, packet *cmessage.GameEndReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("GameEndReq %v", head)
}

func (this *GameProcess) GetRoomAllDataReq(ctx context.Context, packet *cmessage.GetRoomAllDataReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("GetRoomAllDataReq %v", head)
}

func (this *GameProcess) CreateRoomReq(ctx context.Context, packet *cmessage.CreateRoomReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("CreateRoomReq %v", head)
}

func (this *GameProcess) JoinRoomReq(ctx context.Context, packet *cmessage.JoinRoomReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("JoinRoomReq %v", head)
}
func (this *GameProcess) JoinRoomQuickReq(ctx context.Context, packet *cmessage.JoinRoomQuickReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("JoinRoomQuickReq %v", head)
}
func (this *GameProcess) LeaveRoomReq(ctx context.Context, packet *cmessage.LeaveRoomReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("LeaveRoomReq %v", head)
}
func (this *GameProcess) MatchRoomReq(ctx context.Context, packet *cmessage.MatchRoomReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("MatchRoomReq %v", head)
}
func (this *GameProcess) RenameRoomReq(ctx context.Context, packet *cmessage.RenameRoomReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("RenameRoomReq %v", head)
}
func (this *GameProcess) ChangepwdRoomReq(ctx context.Context, packet *cmessage.ChangepwdRoomReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("ChangepwdRoomReq %v", head)
}
func (this *GameProcess) RoomKickOffReq(ctx context.Context, packet *cmessage.RoomKickOffReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("RoomKickOffReq %v", head)
}
func (this *GameProcess) DelRoomReq(ctx context.Context, packet *cmessage.DelRoomReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("DelRoomReq %v", head)
}

func (this *GameProcess) LoginReq(ctx context.Context, packet *cmessage.LoginReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("LoginReq %v", head)
}
func (this *GameProcess) MovingReq(ctx context.Context, packet *cmessage.MovingReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("MovingReq %v", head)
}
func (this *GameProcess) JumpReq(ctx context.Context, packet *cmessage.JumpReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("JumpReq %v", head)
}
func (this *GameProcess) FlagReq(ctx context.Context, packet *cmessage.FlagReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("FlagReq %v", head)
}
func (this *GameProcess) UseItemReq(ctx context.Context, packet *cmessage.UseItemReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("UseItemReq %v", head)
}
func (this *GameProcess) RewardReq(ctx context.Context, packet *cmessage.RewardReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("RewardReq %v", head)
}
func (this *GameProcess) UpdateDevtReq(ctx context.Context, packet *cmessage.UpdateDevtReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("UpdateDevtReq %v", head)
}
func (this *GameProcess) AddHpReq(ctx context.Context, packet *cmessage.AddHpReq) {
	head := this.GetRpcHead(ctx)

	logger.Debug("AddHpReq %v", head)
}

func (this *GameProcess) Init() {
	this.Actor.Init()

	this.RegisterCall("GameStartReq", this.GameStartReq)
	this.RegisterCall("GameEndReq", this.GameEndReq)

	//--------room
	this.RegisterCall("GetRoomAllDataReq", this.GetRoomAllDataReq)
	this.RegisterCall("CreateRoomReq", this.CreateRoomReq)
	this.RegisterCall("JoinRoomReq", this.JoinRoomReq)
	this.RegisterCall("JoinRoomQuickReq", this.JoinRoomQuickReq)
	this.RegisterCall("LeaveRoomReq", this.LeaveRoomReq)
	this.RegisterCall("MatchRoomReq", this.MatchRoomReq)
	this.RegisterCall("RenameRoomReq", this.RenameRoomReq)
	this.RegisterCall("ChangepwdRoomReq", this.ChangepwdRoomReq)
	this.RegisterCall("RoomKickOffReq", this.RoomKickOffReq)
	this.RegisterCall("DelRoomReq", this.DelRoomReq)

	//------player
	this.RegisterCall("LoginReq", this.LoginReq)
	this.RegisterCall("MovingReq", this.MovingReq)
	this.RegisterCall("AttackReq", this.AttackReq)
	this.RegisterCall("JumpReq", this.JumpReq)
	this.RegisterCall("FlagReq", this.FlagReq)
	this.RegisterCall("UseItemReq", this.UseItemReq)
	this.RegisterCall("RewardReq", this.RewardReq)
	this.RegisterCall("UpdateDevtReq", this.UpdateDevtReq)
	this.RegisterCall("AddHpReq", this.AddHpReq)

	this.Actor.Start()
}
