package gamedata

import (
	"fmt"
	"github.com/nickalie/go-queue"
	"gonet/server/common"
	"gonet/server/common/data"
	"reflect"
)

func Produce(message data.AttackData) {
	queue.Put("taskAttack", message)
}

func handleTaskAttack(message data.AttackData) {
	switch message.Mode {
	//TODO:游戏
	case common.Mode.Game:
		{
			GameCtrl.CheckGame()
		}
		break
		//TODO:地图
	case common.Mode.Map:
		{
			UserCtrl.flag(
				&data.FlagReqData{
					Round: message.Round,
					From:  message.From,
					Part:  message.Part,
					Role:  message.Role,
				},
			)
		}
		break
		//TODO:玩家
	case common.Mode.User:
		{
			UserCtrl.relivePlayer(message.Round, message.From, 0)
		}
		break
	//TODO:技能
	case common.Mode.Skill:
		{
			if UserCtrl.canAttack(message) {
				UserCtrl.attack(message)
			}
		}
		break
	//TODO:道具
	case common.Mode.Item:
		{
			UserCtrl.item(
				&data.ItemData{
					From:   message.From,
					Round:  message.Round,
					Part:   message.Part,
					ItemId: message.ItemId,
					Count:  message.Count,
					Cd:     message.Cd,
				})
		}
		break
	//TODO:定时器
	case common.Mode.Timer:
		break
	default:
		{
			break
		}
	}
}

func Consume() {
	go ConsumeAttack()
	go ConsumeRoom()
	go ConsumeLog()
}

func ConsumeAttack() {
	for {
		var message data.AttackData
		queue.Get("taskAttack", &message)

		fmt.Printf("Consumer got a message: %v\n", message)
		handleTaskAttack(message)
	}
}

func ConsumeRoom() {
	for {
		var message1 data.RoomData
		queue.Get("taskRoom", &message1)
		HandleRoomTask(message1)
	}
}

func ConsumeLog() {
	for {
		var message2 data.LogData
		queue.Get("taskLog", &message2)
		handleLogTask(message2)
	}
}

func handleLogTask(message data.LogData) {
	switch message.Mode {
	case common.LogMode.Login:
		SaveCtrl.SaveLoginLog(message)
		break
	case common.LogMode.Battle_round:
		SaveCtrl.SaveBattleLog(message)
		break
	case common.LogMode.Match:
		SaveCtrl.SaveMatchLog(message)
		break
	default:
		//console.error("error log");
		break
	}
}

type service struct {
	servers map[string]reflect.Method
	rcvr    reflect.Value
	typ     reflect.Type
}
