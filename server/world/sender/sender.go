package sender

import (
	"github.com/nickalie/go-queue"
	"gonet/server/common/data"
)

//hero
var taskAttack = "taskAttack"

func AddAttackTask(message data.AttackData) {
	queue.Put(taskAttack, message)
}

//room
var taskRoom = "taskRoom"

func AddRoomTask(message data.RoomData) {
	queue.Put(taskRoom, message)
}

//log
var taskLog = "taskLog"

func AddLogTask(message data.LogData) {
	queue.Put(taskLog, message)
}
