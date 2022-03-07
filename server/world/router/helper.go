package router

import (
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/world/param"
	"gonet/server/world/public"
	"gonet/server/world/socket"
	"time"
)

func HandleHelper(socket socket.Socket) {

	socket.On(public.USER_EVENT.GLOBAL.TIME, func(param param.Param) {
		socket.Emit(public.USER_EVENT.GLOBAL.TIME, &cmessage.GameTimeResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_GameTimeResp), 0),
			Time:       time.Now().Unix(),
		})
	})

}
