package gamedata

import (
	"gonet/server/common"
	"gonet/server/common/data"
	"gonet/server/common/mredis"
	"gonet/server/world/redisInstnace"
	"math"
	"time"
)

type (
	SPositionCtrl struct {
	}
	ISPositionCtrl interface {
	}
)

var PositionCtrl = &SPositionCtrl{}

func (this *SPositionCtrl) onlineUserMsg(round int) []string {
	roundUser := GetOnlineUsers()
	var users []string
	for _, user := range roundUser {
		keyUser := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(user, round)
		datas := redisInstnace.M_pRedisClient.HGetAll(keyUser)
		datas.Val()
		for _, val := range datas.Val() {
			users = append(users, val)
		}
	}
	return users
}

func (this *SPositionCtrl) toServerDirection(_direction float64) float64 {
	direction := 90 - _direction
	if direction <= 0 {
		direction += 360
	}
	// if (direction < 0) direction += 360
	return direction
}

func (this *SPositionCtrl) distance(a data.UserPositionData, b data.UserPositionData) float64 {
	return math.Sqrt((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y))
}
func (this *SPositionCtrl) angle(a data.UserPositionData, b data.UserPositionData) float64 {
	b.X = b.X - a.X
	b.Y = b.Y - a.Y
	a.X = 0
	a.Y = 0
	if b.X == 0 {
		if b.Y > 0 {
			return 90
		}

		return 270
	}
	value := math.Abs(b.Y / b.X)
	an := math.Atan(value) * (180 / math.Pi)
	if b.X > 0 && b.Y > 0 {
		return an
	}
	if b.X < 0 && b.Y > 0 {
		return 180 - an
	}
	if b.X < 0 && b.Y < 0 {
		return an + 180
	}
	if b.X > 0 && b.Y < 0 {
		return 360 - an
	}
	return an
}

func (this *SPositionCtrl) pointDistance(a data.PointData, b data.PointData) float64 {
	return math.Sqrt((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y))
}

func (this *SPositionCtrl) attackSpendTime(startPoint data.PointData, targetPoint data.PointData) float64 {
	//1m/s
	distance := this.pointDistance(startPoint, targetPoint)
	return distance
}

func (this *SPositionCtrl) GetPos(round int, user string) *data.UserPositionData {
	userData := GGame.GetUserById(round, user)
	if userData != nil {
		//return userData
		var pos = new(data.UserPositionData)

		pos.X = userData.X
		pos.Y = userData.Y
		pos.Speed = userData.Speed
		pos.Direction = userData.Direction
		pos.Barrier = userData.Barrier
		pos.Dizzy = userData.Dizzy
		pos.PosUpdateTs = userData.PosUpdateTs
		pos.ReduceSpeedTs = userData.ReduceSpeedTs
		pos.DizzyTs = userData.DizzyTs

		return pos
	} else {
		key := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(user, round)

		selectAttr := []string{"x", "y", "speed", "direction", "barrier", "dizzy", "posUpdateTs", "reduceSpeedTs", "dizzyTs"}
		posArr := redisInstnace.M_pRedisClient.HMGet(key, selectAttr...)
		if posArr == nil {
			return nil
		}
		var pos = new(data.UserPositionData)

		pos.X = (posArr.Val()[0]).(float64)
		pos.Y = (posArr.Val()[1]).(float64)
		pos.Speed = (posArr.Val()[2]).(int64)
		pos.Direction = (posArr.Val()[3]).(float64)
		pos.Barrier = (posArr.Val()[4]).(int32)
		pos.Dizzy = (posArr.Val()[5]).(int32)
		pos.PosUpdateTs = (posArr.Val()[6]).(int64)
		pos.ReduceSpeedTs = (posArr.Val()[7]).(int64)
		pos.DizzyTs = (posArr.Val()[8]).(int64)

		return pos
	}
	return nil
}

func (this *SPositionCtrl) updataPos(round int, user string, newPos data.UserPositionData) {
	GGame.updateGameUserPos(round, user, newPos)
}

func (this *SPositionCtrl) calPos(obj *data.UserPositionData) *data.UserPositionData {
	//let { x, y, speed, reduceSpeedTs, direction, posUpdateTs, barrier, dizzy, dizzyTs, stopmove, stopmoveTs } = obj
	var speed int64
	if obj.Speed < POSITION_CONFIG.SPEED.Min {
		speed = POSITION_CONFIG.SPEED.Min
	}

	var distance float64 = 0
	barrier := obj.Barrier
	dizzy := obj.Dizzy
	stopmove := obj.StopMove
	posUpdateTs := time.Now().Unix()
	direction := obj.Direction

	if barrier == 1 || dizzy == 1 || speed == 0 || stopmove == 1 {
		// 眩晕  障碍物
		distance = 0
	} else {
		distance = ((float64(time.Now().Unix()) - float64(posUpdateTs)) / float64(1000)) * float64(speed)
	}
	var x, y float64
	x = x + distance*math.Cos((this.toServerDirection(direction)*math.Pi)/float64(180))
	y = y + distance*math.Sin((this.toServerDirection(direction)*math.Pi)/180)
	return &data.UserPositionData{
		Speed:         speed,
		ReduceSpeedTs: obj.ReduceSpeedTs,
		Direction:     direction,
		Barrier:       barrier,
		Dizzy:         dizzy,
		DizzyTs:       obj.DizzyTs,
		X:             x,
		Y:             y,
		StopMove:      stopmove,
		StopMoveTs:    obj.StopMoveTs,
		PosUpdateTs:   posUpdateTs,
	}
}

func (this *SPositionCtrl) skillPos(obj data.UserPositionData) *data.UserPositionData {
	var speed int64
	if obj.Speed < POSITION_CONFIG.SPEED.Min {
		speed = POSITION_CONFIG.SPEED.Min
	}
	if obj.Speed > POSITION_CONFIG.SPEED.Max {
		speed = POSITION_CONFIG.SPEED.Max
	}
	var distance float64 = 0
	barrier := obj.Barrier
	dizzy := obj.Dizzy
	stopmove := obj.StopMove
	posUpdateTs := time.Now().Unix()
	direction := obj.Direction
	if barrier == 1 || dizzy == 1 || speed == 0 || stopmove == 1 {
		// 眩晕  障碍物
		distance = 0
	} else {
		distance = ((float64(time.Now().Unix()) - float64(posUpdateTs)) / float64(1000)) * float64(speed)
	}
	var x, y float64
	x = x + distance*math.Cos((this.toServerDirection(direction)*math.Pi)/float64(180))
	y = y + distance*math.Sin((this.toServerDirection(direction)*math.Pi)/180)
	return &data.UserPositionData{
		Speed:         speed,
		ReduceSpeedTs: obj.ReduceSpeedTs,
		Direction:     direction,
		Barrier:       barrier,
		Dizzy:         dizzy,
		DizzyTs:       obj.DizzyTs,
		X:             x,
		Y:             y,
		StopMove:      stopmove,
		StopMoveTs:    obj.StopMoveTs,
		PosUpdateTs:   posUpdateTs,
	}
}

func (this *SPositionCtrl) updatePosition(user string, round int, obj data.UserPositionData) *data.UserPositionData {
	pos := this.GetPos(round, user)
	if pos == nil {
		return nil
	}

	newPos := this.calPos(pos)
	newPos.ReduceSpeedTs = obj.ReduceSpeedTs
	newPos.Speed = obj.Speed
	newPos.Direction = obj.Direction
	newPos.Barrier = obj.Barrier
	newPos.Dizzy = obj.Dizzy
	newPos.DizzyTs = obj.DizzyTs

	//移除眩晕
	if newPos.DizzyTs < time.Now().Unix() {
		newPos.DizzyTs = 0
		newPos.Dizzy = 0
	}

	//移除不能移动
	if newPos.StopMoveTs < time.Now().Unix() {
		newPos.StopMoveTs = 0
		newPos.StopMove = 0
	}

	//同步缓存内存数据
	// this.updataPos(round, user, newPos)
	GGame.updateGameUserPos(round, user, *newPos)
	return newPos
}

func (this *SPositionCtrl) newUpdatePosition(user string, round int, obj data.UserPositionData) *data.UserPositionData {
	pos := this.GetPos(round, user)
	if pos == nil {
		return nil
	}

	newPos := this.calPos(pos)

	newPos.ReduceSpeedTs = obj.ReduceSpeedTs
	newPos.Speed = obj.Speed
	newPos.Direction = obj.Direction
	newPos.Barrier = obj.Barrier
	newPos.Dizzy = obj.Dizzy
	newPos.DizzyTs = obj.DizzyTs
	newPos.X = obj.X
	newPos.Y = obj.Y

	//移除眩晕
	if newPos.DizzyTs < time.Now().Unix() {
		newPos.DizzyTs = 0
		newPos.Dizzy = 0
	}

	//移除不能移动
	if newPos.StopMoveTs < time.Now().Unix() {
		newPos.StopMoveTs = 0
		newPos.StopMove = 0
	}

	//同步缓存内存数据
	// this.updataPos(round, user, newPos)
	GGame.updateGameUserPos(round, user, *newPos)
	return newPos
}

func (this *SPositionCtrl) skill1502Position(userData interface{}, speed int, direction int, time0 int) {
	//let pos = this.GetPos(userData.round, userData.user)
	//if (!pos) return
	//let newPos = this.newPoint({ x: userData.x, y: userData.y }, speed, direction)
	//newPos = await gGame.updataGameUserSkillPos(userData.round, userData.user, newPos)
	//pos.x = newPos.x
	//pos.y = newPos.x
	//
	//return pos
}

func (this *SPositionCtrl) relivePlayer(userData interface{}, speed int, direction int, time0 int) {
	//let pos = await this.GetPos(userData.round, userData.user)
	//if (!pos) return
	//let newPos = this.newPoint({ x: userData.x, y: userData.y }, speed, direction)
	//newPos = await gGame.updataGameUserSkillPos(userData.round, userData.user, newPos)
	//pos.x = newPos.x
	//pos.y = newPos.x
	//
	//return pos
}

func (this *SPositionCtrl) skillMoving(userData interface{}, speed int, direction int, time0 int) {
	//let newPos = this.newPoint({ x: userData.x, y: userData.y }, speed, direction)
	//newPos = await gGame.updataGameUserSkillPos(userData.round, userData.user, newPos)
	//return newPos
}

func (this *SPositionCtrl) updateNewPosition(user string, round int) *data.UserPositionData {
	pos := this.GetPos(round, user)
	if pos == nil {
		return nil
	}
	newPos := this.calPos(pos)

	return newPos
}

func (this *SPositionCtrl) updateSkillPosition(user string, round int, obj data.UserPositionData) *data.UserPositionData {
	pos := this.GetPos(round, user)
	if pos == nil {
		return nil
	}
	newPos := this.calPos(pos)

	newPos.ReduceSpeedTs = obj.ReduceSpeedTs
	newPos.Speed = obj.Speed
	newPos.Direction = obj.Direction
	newPos.Barrier = obj.Barrier
	newPos.Dizzy = obj.Dizzy
	newPos.DizzyTs = obj.DizzyTs
	newPos.X = obj.X
	newPos.Y = obj.Y

	//同步缓存内存数据
	// this.updataPos(round, user, newPos)
	GGame.updateGameUserPos(round, user, *newPos)
	return newPos
}

//点在圆内
func (this *SPositionCtrl) IsPointInCircle(pos data.UserPositionData, circle data.Pos, r float64) bool {
	x := pos.X
	y := pos.Y
	if r == 0 {
		return false
	}

	var dx = circle.X - x
	var dy = circle.Y - y
	return dx*dx+dy*dy <= r*r
}

func (this *SPositionCtrl) GetCross(p1 data.PointData, p2 data.PointData, p data.PointData) float64 {
	return (p2.X-p1.X)*(p.Y-p1.Y) - (p.X-p1.X)*(p2.Y-p1.Y)
}
func (this *SPositionCtrl) newPoint(point data.PointData, distance float64, direction float64) data.PointData {
	x := point.X
	y := point.Y
	x = x + distance*math.Cos((this.toServerDirection(direction)*math.Pi)/180)
	y = y + distance*math.Sin((this.toServerDirection(direction)*math.Pi)/180)
	return data.PointData{X: x, Y: y}
}

func (this *SPositionCtrl) getMatrixPoint(startPoint data.PointData, direction float64, widthx float64, highy float64) (data.PointData, data.PointData, data.PointData, data.PointData, data.PointData) {
	p1 := this.newPoint(startPoint, widthx/2, float64(int(direction+90)%360))
	p2 := this.newPoint(startPoint, widthx/2, float64(int(direction+270)%360))
	newP := this.newPoint(startPoint, highy, direction)
	point := data.PointData{X: newP.X, Y: newP.Y}
	p3 := this.newPoint(point, widthx/2, float64(int((direction+270))%360))
	p4 := this.newPoint(point, widthx/2, float64(int(direction+90)%360))

	return p1, p2, p3, p4, newP
}

//点在矩形
func (this *SPositionCtrl) IsPointInMatrix(startPoint data.PointData, direction float64, widthx float64, highy float64, p data.PointData) (bool, data.PointData) {
	p1, p2, p3, p4, newPoint := this.getMatrixPoint(startPoint, direction, widthx, highy)
	isPointIn := this.GetCross(p1, p2, p)*this.GetCross(p3, p4, p) >= 0 && this.GetCross(p2, p3, p)*this.GetCross(p4, p1, p) >= 0
	return isPointIn, newPoint
}

//获得人物中心和鼠标坐标连线，与y轴正半轴之间的夹角
func (this *SPositionCtrl) getAngle(startx float64, starty float64, targetx float64, targetxy float64) float64 {
	x := math.Abs(startx - targetx)
	y := math.Abs(starty - targetxy)
	z := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	cos := y / z
	radina := math.Acos(cos)                      //用反三角函数求弧度
	angle := math.Floor(180 / (math.Pi / radina)) //将弧度转换成角度

	if targetx > startx && targetxy > starty { //鼠标在第四象限
		angle = 180 - angle
	}

	if targetx == startx && targetxy > starty { //鼠标在y轴负方向上
		angle = 180
	}

	if targetx > startx && targetxy == starty { //鼠标在x轴正方向上
		angle = 90
	}

	if targetx < startx && targetxy > starty { //鼠标在第三象限
		angle = 180 + angle
	}

	if targetx < startx && targetxy == starty { //鼠标在x轴负方向
		angle = 270
	}

	if targetx < startx && targetxy < starty { //鼠标在第二象限
		angle = 360 - angle
	}
	if targetx == startx && targetx == 0 {
		if targetxy > starty {
			angle = 90
		} else {
			angle = 270
		}
	}
	if targetxy == starty && targetxy == 0 {
		if targetx > startx {
			angle = 180
		} else {
			angle = 360
		}
	}
	return float64(int(angle) % 360)
}

func (this *SPositionCtrl) calNewPointByAngle(startPoint data.PointData, angle float64, distance float64) data.PointData {
	var endPoint data.PointData
	// 角度转弧度
	var radian = (angle * 3.14) / 180
	// 计算新坐标(对于无限接近0的数字，此处没有优化)
	endPoint.X = startPoint.X + distance*math.Sin(radian)
	endPoint.Y = startPoint.Y + distance*math.Cos(radian)
	return endPoint
}

//点在扇形
func (this *SPositionCtrl) IsPointInFan(startPoint data.PointData, direction float64, angle float64, radius float64, targetPoint data.PointData) bool {
	//判断点到点的半径
	direction = float64(int(direction) % 360)
	if targetPoint.X == startPoint.X {
		//同x轴
		if math.Abs(targetPoint.Y-startPoint.Y) > radius {
			return false
		}
	} else if targetPoint.Y == startPoint.Y {
		//同y轴
		if math.Abs(targetPoint.X-startPoint.X) > radius {
			return false
		}
	}

	if (targetPoint.X-startPoint.X)*(targetPoint.X-startPoint.X)+(targetPoint.Y-startPoint.Y)*(targetPoint.Y-startPoint.Y) > radius*radius {
		return false
	}

	// let newPoint = this.calNewPointByAngle(startPoint, direction, radius)
	newPoint := this.newPoint(startPoint, radius, direction)
	//根据技能释放方向判定夹角 （0-90 第四象限，90-180 第一象限， 180-270第二象限，270-360，第3象限）
	//启始角度
	startAngle := this.getAngle(startPoint.X, startPoint.Y, newPoint.X, newPoint.Y)
	//目标角度
	targetAngle := this.getAngle(startPoint.X, startPoint.Y, targetPoint.X, targetPoint.Y)
	//判断目标角度是否 在扇形两边的角度内
	if targetAngle <= (startAngle+angle/2) && targetAngle >= (startAngle-angle/2) {
		return true
	}
	return false
}
