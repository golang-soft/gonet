package helper

import (
	"gonet/server/common/data"
	"math"
	"math/rand"
)

//export const sleep = (s = 1) = > new Promise<number>(r = > setTimeout(() = > r(0), s * 1000))

func isInside(x float64, y float64) bool {
	return x*x+y*y <= 25*25
}

func MaybeSuccess(number float64) bool {
	return math.Floor(float64(rand.Float64()*2000)) < number
}

func MaybeSuccessPercent(number float64) bool {
	return randomMin2Max(0, 100) <= number
}

func randomMin2Max(min float64, max float64) float64 {
	return math.Floor(rand.Float64()*(max-min)) + min
}

//export function obj2Number(obj: any) {
//	for (let key of Object.keys(obj)) {
//	if (typeof obj[key] == 'object') {
//	obj2Number(obj[key])
//	} else {
//	const v = Number(obj[key])
//	if (!Number.isNaN(v)) {
//	obj[key] = v
//	}
//	}
//	}
//	return obj
//}
//
//
//export function IsPointInCircle(pos: any, circle: any, r: number) {
//	let x = pos.x
//	let y = pos.y
//	if (r == = 0) return false
//	var dx = circle[0] - x
//	var dy = circle[1] - y
//	return dx * dx + dy * dy <= r * r
//}

func GetCross(p1 data.PointData, p2 data.PointData, p data.PointData) float64 {
	return (p2.X-p1.X)*(p.Y-p1.Y) - (p.X-p1.X)*(p2.Y-p1.Y)
}
func IsPointInMatrix(p1 data.PointData, p2 data.PointData, p3 data.PointData, p4 data.PointData, p data.PointData) bool {
	isPointIn := GetCross(p1, p2, p)*GetCross(p3, p4, p) >= 0 && GetCross(p2, p3, p)*GetCross(p4, p1, p) >= 0
	return isPointIn
}
