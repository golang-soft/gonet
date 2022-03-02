package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"gonet/base"
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/rpc"
	"log"
	"sync/atomic"
)

type (
	LoginAccountEvent struct {
		BaseEvent
	}
	ILoginAccountEvent interface {
		IBaseEvent
	}
)

func NewLoginAccountEvent() *LoginAccountEvent {
	return &LoginAccountEvent{}
}

func (this *LoginAccountEvent) DoEvent(process *EventProcess) {
	log.Printf("LoginAccountEvent doEvent.......")
	this.LoginAccount(process)
}

func (this *LoginAccountEvent) Name() string {
	return "LoginAccountEvent"
}

func (this *LoginAccountEvent) HandleEvent(event *IBaseEvent, message proto.Message) {

}

func (this *LoginAccountEvent) EventID() int {
	return this.eventid
}

func (this *LoginAccountEvent) SendEvent(event *IBaseEvent, process *EventProcess) {
	this.LoginAccount(process)
}

func (this *LoginAccountEvent) LoginAccount(process *EventProcess) {
	id := atomic.AddInt32(&id, 1)
	this.AccountName = fmt.Sprintf("test321%d", id)
	this.PassWd = base.MD5(ToSlat(this.AccountName, "123456"))
	//this.AccountName = fmt.Sprintf("test%d", base.RAND.RandI(0, 7000))
	packet1 := &cmessage.C_A_LoginRequest{PacketHead: common.BuildPacketHead(cmessage.MessageID_MSG_C_A_LoginRequest, rpc.SERVICE_GATESERVER),
		AccountName: this.AccountName, Password: this.PassWd, BuildNo: base.BUILD_NO, Key: process.m_Dh.ShareKey()}
	this.SendPacket(process, packet1)
	m_Log.Debugf("玩家 %d 登录账号 %s", this.PlayerId, this.AccountName)
}
