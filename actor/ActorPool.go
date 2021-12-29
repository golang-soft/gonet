package actor

import (
	"gonet/grpc"
	"gonet/server/rpc"
	"sync"
)

//********************************************************
// actorpool 管理,这里的actor可以动态添加,携程池
//********************************************************
type (
	ActorPool struct {
		Actor
		m_ActorMap  map[int64]IActor
		m_ActorLock *sync.RWMutex
		//m_ActorMap sync.Map
	}

	IActorPool interface {
		GetActor(Id int64) IActor                         //获取actor
		AddActor(Id int64, pActor IActor)                 //添加actor
		DelActor(Id int64)                                //删除actor
		BoardCast(funcName string, params ...interface{}) //广播actor
		GetActorNum() int
	}
)

func (this *ActorPool) Init() {
	this.m_ActorMap = make(map[int64]IActor)
	this.m_ActorLock = &sync.RWMutex{}
	this.Actor.Init()
}

func (this *ActorPool) GetActor(Id int64) IActor {
	//v, bOk := this.m_ActorMap.Load(Id)
	this.m_ActorLock.RLock()
	pActor, bEx := this.m_ActorMap[Id]
	this.m_ActorLock.RUnlock()
	if bEx {
		return pActor
	}
	return nil
}

func (this *ActorPool) AddActor(Id int64, pActor IActor) {
	this.m_ActorLock.Lock()
	this.m_ActorMap[Id] = pActor
	this.m_ActorLock.Unlock()
	//this.m_ActorMap.Store(Id, pActor)
}

func (this *ActorPool) DelActor(Id int64) {
	this.m_ActorLock.Lock()
	delete(this.m_ActorMap, Id)
	this.m_ActorLock.Unlock()
	//this.m_ActorMap.Delete(Id)
}

func (this *ActorPool) GetActorNum() int {
	nLen := 0
	this.m_ActorLock.RLock()
	nLen = len(this.m_ActorMap)
	this.m_ActorLock.RUnlock()
	return nLen
}

func (this *ActorPool) BoardCast(funcName string, params ...interface{}) {
	this.m_ActorLock.RLock()
	for _, v := range this.m_ActorMap {
		v.SendMsg(rpc.RpcHead{}, funcName, params...)
	}
	this.m_ActorLock.RUnlock()
}

func (this *ActorPool) SendMsg(head rpc.RpcHead, funcName string, params ...interface{}) {
	buff := grpc.Marshal(head, funcName, params...)
	head.SocketId = 0
	if head.Id != 0 {
		pActor := this.GetActor(int64(head.Id))
		if pActor != nil && pActor.FindCall(funcName) != nil {
			pActor.Send(head, buff)
			return
		}
	}
	this.Send(head, buff)
}

//actor pool must rewrite PacketFunc
func (this *ActorPool) PacketFunc(packet rpc.Packet) bool {
	rpcPacket, head := grpc.UnmarshalHead(packet.Buff)
	if this.FindCall(rpcPacket.FuncName) != nil {
		head.SocketId = packet.Id
		head.Reply = packet.Reply
		this.Send(head, packet.Buff)
		return true
	} else {
		pActor := this.GetActor(int64(rpcPacket.RpcHead.Id))
		if pActor != nil && pActor.FindCall(rpcPacket.FuncName) != nil {
			head.SocketId = packet.Id
			head.Reply = packet.Reply
			pActor.Send(head, packet.Buff)
			return true
		}
	}
	return false
}
