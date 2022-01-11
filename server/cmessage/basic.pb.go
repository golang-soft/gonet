// Code generated by protoc-gen-go. DO NOT EDIT.
// source: basic.proto

package cmessage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 游戏开始
type GameStartReq struct {
	PacketHead           *Ipacket `protobuf:"bytes,1,opt,name=PacketHead,proto3" json:"PacketHead,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameStartReq) Reset()         { *m = GameStartReq{} }
func (m *GameStartReq) String() string { return proto.CompactTextString(m) }
func (*GameStartReq) ProtoMessage()    {}
func (*GameStartReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc45f6ac745dd88, []int{0}
}

func (m *GameStartReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameStartReq.Unmarshal(m, b)
}
func (m *GameStartReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameStartReq.Marshal(b, m, deterministic)
}
func (m *GameStartReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameStartReq.Merge(m, src)
}
func (m *GameStartReq) XXX_Size() int {
	return xxx_messageInfo_GameStartReq.Size(m)
}
func (m *GameStartReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GameStartReq.DiscardUnknown(m)
}

var xxx_messageInfo_GameStartReq proto.InternalMessageInfo

func (m *GameStartReq) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

type GameData struct {
	Round                int64    `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	StartTs              int64    `protobuf:"varint,2,opt,name=startTs,proto3" json:"startTs,omitempty"`
	EndTs                int64    `protobuf:"varint,3,opt,name=endTs,proto3" json:"endTs,omitempty"`
	FlagOwner            int64    `protobuf:"varint,4,opt,name=flagOwner,proto3" json:"flagOwner,omitempty"`
	FlagUPdateTs         int64    `protobuf:"varint,5,opt,name=flagUPdateTs,proto3" json:"flagUPdateTs,omitempty"`
	Part1Score           float64  `protobuf:"fixed64,6,opt,name=part1Score,proto3" json:"part1Score,omitempty"`
	Part1Ts              int64    `protobuf:"varint,7,opt,name=part1Ts,proto3" json:"part1Ts,omitempty"`
	Part2Ts              int64    `protobuf:"varint,8,opt,name=part2Ts,proto3" json:"part2Ts,omitempty"`
	Part2Score           float64  `protobuf:"fixed64,9,opt,name=part2Score,proto3" json:"part2Score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameData) Reset()         { *m = GameData{} }
func (m *GameData) String() string { return proto.CompactTextString(m) }
func (*GameData) ProtoMessage()    {}
func (*GameData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc45f6ac745dd88, []int{1}
}

func (m *GameData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameData.Unmarshal(m, b)
}
func (m *GameData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameData.Marshal(b, m, deterministic)
}
func (m *GameData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameData.Merge(m, src)
}
func (m *GameData) XXX_Size() int {
	return xxx_messageInfo_GameData.Size(m)
}
func (m *GameData) XXX_DiscardUnknown() {
	xxx_messageInfo_GameData.DiscardUnknown(m)
}

var xxx_messageInfo_GameData proto.InternalMessageInfo

func (m *GameData) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *GameData) GetStartTs() int64 {
	if m != nil {
		return m.StartTs
	}
	return 0
}

func (m *GameData) GetEndTs() int64 {
	if m != nil {
		return m.EndTs
	}
	return 0
}

func (m *GameData) GetFlagOwner() int64 {
	if m != nil {
		return m.FlagOwner
	}
	return 0
}

func (m *GameData) GetFlagUPdateTs() int64 {
	if m != nil {
		return m.FlagUPdateTs
	}
	return 0
}

func (m *GameData) GetPart1Score() float64 {
	if m != nil {
		return m.Part1Score
	}
	return 0
}

func (m *GameData) GetPart1Ts() int64 {
	if m != nil {
		return m.Part1Ts
	}
	return 0
}

func (m *GameData) GetPart2Ts() int64 {
	if m != nil {
		return m.Part2Ts
	}
	return 0
}

func (m *GameData) GetPart2Score() float64 {
	if m != nil {
		return m.Part2Score
	}
	return 0
}

type UserData struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Round                int32    `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Part                 int32    `protobuf:"varint,3,opt,name=part,proto3" json:"part,omitempty"`
	Hid                  string   `protobuf:"bytes,4,opt,name=hid,proto3" json:"hid,omitempty"`
	Type                 int32    `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	DefPercent           int32    `protobuf:"varint,6,opt,name=defPercent,proto3" json:"defPercent,omitempty"`
	Hp                   int32    `protobuf:"varint,7,opt,name=hp,proto3" json:"hp,omitempty"`
	UpdateTs             int64    `protobuf:"varint,8,opt,name=updateTs,proto3" json:"updateTs,omitempty"`
	X                    int32    `protobuf:"varint,9,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,10,opt,name=y,proto3" json:"y,omitempty"`
	Speed                int32    `protobuf:"varint,11,opt,name=speed,proto3" json:"speed,omitempty"`
	ReduceSpeedTs        int32    `protobuf:"varint,12,opt,name=reduceSpeedTs,proto3" json:"reduceSpeedTs,omitempty"`
	Direction            int32    `protobuf:"varint,13,opt,name=direction,proto3" json:"direction,omitempty"`
	Barrier              int32    `protobuf:"varint,14,opt,name=barrier,proto3" json:"barrier,omitempty"`
	Dizzy                int32    `protobuf:"varint,15,opt,name=dizzy,proto3" json:"dizzy,omitempty"`
	DizzyTs              int64    `protobuf:"varint,16,opt,name=dizzyTs,proto3" json:"dizzyTs,omitempty"`
	Shield               int32    `protobuf:"varint,17,opt,name=shield,proto3" json:"shield,omitempty"`
	ShieldTs             int64    `protobuf:"varint,18,opt,name=shieldTs,proto3" json:"shieldTs,omitempty"`
	Immune               int32    `protobuf:"varint,19,opt,name=immune,proto3" json:"immune,omitempty"`
	ImmuneTs             int32    `protobuf:"varint,20,opt,name=immuneTs,proto3" json:"immuneTs,omitempty"`
	Thorns               int32    `protobuf:"varint,21,opt,name=thorns,proto3" json:"thorns,omitempty"`
	ThornsTs             int32    `protobuf:"varint,22,opt,name=thornsTs,proto3" json:"thornsTs,omitempty"`
	StopMoveTs           int32    `protobuf:"varint,23,opt,name=stopMoveTs,proto3" json:"stopMoveTs,omitempty"`
	StopMove             int32    `protobuf:"varint,24,opt,name=stopMove,proto3" json:"stopMove,omitempty"`
	AddDef               int32    `protobuf:"varint,25,opt,name=addDef,proto3" json:"addDef,omitempty"`
	AddDefTs             int32    `protobuf:"varint,26,opt,name=addDefTs,proto3" json:"addDefTs,omitempty"`
	AddAtk               int32    `protobuf:"varint,27,opt,name=addAtk,proto3" json:"addAtk,omitempty"`
	PosUpdateTs          int64    `protobuf:"varint,28,opt,name=posUpdateTs,proto3" json:"posUpdateTs,omitempty"`
	DieTs                int64    `protobuf:"varint,29,opt,name=dieTs,proto3" json:"dieTs,omitempty"`
	AllAttr              int32    `protobuf:"varint,30,opt,name=allAttr,proto3" json:"allAttr,omitempty"`
	Dvt                  int32    `protobuf:"varint,31,opt,name=dvt,proto3" json:"dvt,omitempty"`
	GetDvt_              int32    `protobuf:"varint,32,opt,name=getDvt,proto3" json:"getDvt,omitempty"`
	DesDvt               int32    `protobuf:"varint,33,opt,name=desDvt,proto3" json:"desDvt,omitempty"`
	Kill                 int32    `protobuf:"varint,34,opt,name=kill,proto3" json:"kill,omitempty"`
	Die                  int32    `protobuf:"varint,35,opt,name=die,proto3" json:"die,omitempty"`
	Dps                  int32    `protobuf:"varint,36,opt,name=dps,proto3" json:"dps,omitempty"`
	Skill_1101Cd         int32    `protobuf:"varint,37,opt,name=skill_1101_cd,json=skill1101Cd,proto3" json:"skill_1101_cd,omitempty"`
	Skill_1102Cd         int32    `protobuf:"varint,38,opt,name=skill_1102_cd,json=skill1102Cd,proto3" json:"skill_1102_cd,omitempty"`
	Skill_1103Cd         int32    `protobuf:"varint,39,opt,name=skill_1103_cd,json=skill1103Cd,proto3" json:"skill_1103_cd,omitempty"`
	Skill_1104Cd         int32    `protobuf:"varint,40,opt,name=skill_1104_cd,json=skill1104Cd,proto3" json:"skill_1104_cd,omitempty"`
	Item_1002Cd          int32    `protobuf:"varint,41,opt,name=item_1002_cd,json=item1002Cd,proto3" json:"item_1002_cd,omitempty"`
	Item_1003Cd          int32    `protobuf:"varint,42,opt,name=item_1003_cd,json=item1003Cd,proto3" json:"item_1003_cd,omitempty"`
	Item_1004Cd          int32    `protobuf:"varint,43,opt,name=item_1004_cd,json=item1004Cd,proto3" json:"item_1004_cd,omitempty"`
	Equip_1              int32    `protobuf:"varint,44,opt,name=equip_1,json=equip1,proto3" json:"equip_1,omitempty"`
	Equip_3              int32    `protobuf:"varint,45,opt,name=equip_3,json=equip3,proto3" json:"equip_3,omitempty"`
	Equip                []int32  `protobuf:"varint,46,rep,packed,name=equip,proto3" json:"equip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserData) Reset()         { *m = UserData{} }
func (m *UserData) String() string { return proto.CompactTextString(m) }
func (*UserData) ProtoMessage()    {}
func (*UserData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc45f6ac745dd88, []int{2}
}

func (m *UserData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserData.Unmarshal(m, b)
}
func (m *UserData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserData.Marshal(b, m, deterministic)
}
func (m *UserData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserData.Merge(m, src)
}
func (m *UserData) XXX_Size() int {
	return xxx_messageInfo_UserData.Size(m)
}
func (m *UserData) XXX_DiscardUnknown() {
	xxx_messageInfo_UserData.DiscardUnknown(m)
}

var xxx_messageInfo_UserData proto.InternalMessageInfo

func (m *UserData) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserData) GetRound() int32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *UserData) GetPart() int32 {
	if m != nil {
		return m.Part
	}
	return 0
}

func (m *UserData) GetHid() string {
	if m != nil {
		return m.Hid
	}
	return ""
}

func (m *UserData) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *UserData) GetDefPercent() int32 {
	if m != nil {
		return m.DefPercent
	}
	return 0
}

func (m *UserData) GetHp() int32 {
	if m != nil {
		return m.Hp
	}
	return 0
}

func (m *UserData) GetUpdateTs() int64 {
	if m != nil {
		return m.UpdateTs
	}
	return 0
}

func (m *UserData) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *UserData) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *UserData) GetSpeed() int32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *UserData) GetReduceSpeedTs() int32 {
	if m != nil {
		return m.ReduceSpeedTs
	}
	return 0
}

func (m *UserData) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *UserData) GetBarrier() int32 {
	if m != nil {
		return m.Barrier
	}
	return 0
}

func (m *UserData) GetDizzy() int32 {
	if m != nil {
		return m.Dizzy
	}
	return 0
}

func (m *UserData) GetDizzyTs() int64 {
	if m != nil {
		return m.DizzyTs
	}
	return 0
}

func (m *UserData) GetShield() int32 {
	if m != nil {
		return m.Shield
	}
	return 0
}

func (m *UserData) GetShieldTs() int64 {
	if m != nil {
		return m.ShieldTs
	}
	return 0
}

func (m *UserData) GetImmune() int32 {
	if m != nil {
		return m.Immune
	}
	return 0
}

func (m *UserData) GetImmuneTs() int32 {
	if m != nil {
		return m.ImmuneTs
	}
	return 0
}

func (m *UserData) GetThorns() int32 {
	if m != nil {
		return m.Thorns
	}
	return 0
}

func (m *UserData) GetThornsTs() int32 {
	if m != nil {
		return m.ThornsTs
	}
	return 0
}

func (m *UserData) GetStopMoveTs() int32 {
	if m != nil {
		return m.StopMoveTs
	}
	return 0
}

func (m *UserData) GetStopMove() int32 {
	if m != nil {
		return m.StopMove
	}
	return 0
}

func (m *UserData) GetAddDef() int32 {
	if m != nil {
		return m.AddDef
	}
	return 0
}

func (m *UserData) GetAddDefTs() int32 {
	if m != nil {
		return m.AddDefTs
	}
	return 0
}

func (m *UserData) GetAddAtk() int32 {
	if m != nil {
		return m.AddAtk
	}
	return 0
}

func (m *UserData) GetPosUpdateTs() int64 {
	if m != nil {
		return m.PosUpdateTs
	}
	return 0
}

func (m *UserData) GetDieTs() int64 {
	if m != nil {
		return m.DieTs
	}
	return 0
}

func (m *UserData) GetAllAttr() int32 {
	if m != nil {
		return m.AllAttr
	}
	return 0
}

func (m *UserData) GetDvt() int32 {
	if m != nil {
		return m.Dvt
	}
	return 0
}

func (m *UserData) GetGetDvt_() int32 {
	if m != nil {
		return m.GetDvt_
	}
	return 0
}

func (m *UserData) GetDesDvt() int32 {
	if m != nil {
		return m.DesDvt
	}
	return 0
}

func (m *UserData) GetKill() int32 {
	if m != nil {
		return m.Kill
	}
	return 0
}

func (m *UserData) GetDie() int32 {
	if m != nil {
		return m.Die
	}
	return 0
}

func (m *UserData) GetDps() int32 {
	if m != nil {
		return m.Dps
	}
	return 0
}

func (m *UserData) GetSkill_1101Cd() int32 {
	if m != nil {
		return m.Skill_1101Cd
	}
	return 0
}

func (m *UserData) GetSkill_1102Cd() int32 {
	if m != nil {
		return m.Skill_1102Cd
	}
	return 0
}

func (m *UserData) GetSkill_1103Cd() int32 {
	if m != nil {
		return m.Skill_1103Cd
	}
	return 0
}

func (m *UserData) GetSkill_1104Cd() int32 {
	if m != nil {
		return m.Skill_1104Cd
	}
	return 0
}

func (m *UserData) GetItem_1002Cd() int32 {
	if m != nil {
		return m.Item_1002Cd
	}
	return 0
}

func (m *UserData) GetItem_1003Cd() int32 {
	if m != nil {
		return m.Item_1003Cd
	}
	return 0
}

func (m *UserData) GetItem_1004Cd() int32 {
	if m != nil {
		return m.Item_1004Cd
	}
	return 0
}

func (m *UserData) GetEquip_1() int32 {
	if m != nil {
		return m.Equip_1
	}
	return 0
}

func (m *UserData) GetEquip_3() int32 {
	if m != nil {
		return m.Equip_3
	}
	return 0
}

func (m *UserData) GetEquip() []int32 {
	if m != nil {
		return m.Equip
	}
	return nil
}

type GameStartResp struct {
	PacketHead           *Ipacket    `protobuf:"bytes,1,opt,name=PacketHead,proto3" json:"PacketHead,omitempty"`
	Game                 *GameData   `protobuf:"bytes,2,opt,name=game,proto3" json:"game,omitempty"`
	Users                []*UserData `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GameStartResp) Reset()         { *m = GameStartResp{} }
func (m *GameStartResp) String() string { return proto.CompactTextString(m) }
func (*GameStartResp) ProtoMessage()    {}
func (*GameStartResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc45f6ac745dd88, []int{3}
}

func (m *GameStartResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameStartResp.Unmarshal(m, b)
}
func (m *GameStartResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameStartResp.Marshal(b, m, deterministic)
}
func (m *GameStartResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameStartResp.Merge(m, src)
}
func (m *GameStartResp) XXX_Size() int {
	return xxx_messageInfo_GameStartResp.Size(m)
}
func (m *GameStartResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GameStartResp.DiscardUnknown(m)
}

var xxx_messageInfo_GameStartResp proto.InternalMessageInfo

func (m *GameStartResp) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *GameStartResp) GetGame() *GameData {
	if m != nil {
		return m.Game
	}
	return nil
}

func (m *GameStartResp) GetUsers() []*UserData {
	if m != nil {
		return m.Users
	}
	return nil
}

// 游戏结束
type GameEndReq struct {
	PacketHead           *Ipacket `protobuf:"bytes,1,opt,name=PacketHead,proto3" json:"PacketHead,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameEndReq) Reset()         { *m = GameEndReq{} }
func (m *GameEndReq) String() string { return proto.CompactTextString(m) }
func (*GameEndReq) ProtoMessage()    {}
func (*GameEndReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc45f6ac745dd88, []int{4}
}

func (m *GameEndReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameEndReq.Unmarshal(m, b)
}
func (m *GameEndReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameEndReq.Marshal(b, m, deterministic)
}
func (m *GameEndReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameEndReq.Merge(m, src)
}
func (m *GameEndReq) XXX_Size() int {
	return xxx_messageInfo_GameEndReq.Size(m)
}
func (m *GameEndReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GameEndReq.DiscardUnknown(m)
}

var xxx_messageInfo_GameEndReq proto.InternalMessageInfo

func (m *GameEndReq) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

type GameEndResp struct {
	PacketHead           *Ipacket  `protobuf:"bytes,1,opt,name=PacketHead,proto3" json:"PacketHead,omitempty"`
	Data                 *GameData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GameEndResp) Reset()         { *m = GameEndResp{} }
func (m *GameEndResp) String() string { return proto.CompactTextString(m) }
func (*GameEndResp) ProtoMessage()    {}
func (*GameEndResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc45f6ac745dd88, []int{5}
}

func (m *GameEndResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameEndResp.Unmarshal(m, b)
}
func (m *GameEndResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameEndResp.Marshal(b, m, deterministic)
}
func (m *GameEndResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameEndResp.Merge(m, src)
}
func (m *GameEndResp) XXX_Size() int {
	return xxx_messageInfo_GameEndResp.Size(m)
}
func (m *GameEndResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GameEndResp.DiscardUnknown(m)
}

var xxx_messageInfo_GameEndResp proto.InternalMessageInfo

func (m *GameEndResp) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *GameEndResp) GetData() *GameData {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GameStartReq)(nil), "cmessage.GameStartReq")
	proto.RegisterType((*GameData)(nil), "cmessage.GameData")
	proto.RegisterType((*UserData)(nil), "cmessage.UserData")
	proto.RegisterType((*GameStartResp)(nil), "cmessage.GameStartResp")
	proto.RegisterType((*GameEndReq)(nil), "cmessage.GameEndReq")
	proto.RegisterType((*GameEndResp)(nil), "cmessage.GameEndResp")
}

func init() { proto.RegisterFile("basic.proto", fileDescriptor_0cc45f6ac745dd88) }

var fileDescriptor_0cc45f6ac745dd88 = []byte{
	// 828 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x72, 0x1b, 0x45,
	0x10, 0x2e, 0x59, 0x5e, 0x5b, 0x1a, 0x49, 0x21, 0x19, 0x42, 0xd2, 0x98, 0x10, 0x84, 0x08, 0xc1,
	0xfc, 0x89, 0x48, 0xca, 0x8d, 0x03, 0x65, 0x62, 0x0a, 0x38, 0x50, 0xb8, 0xd6, 0xeb, 0x0b, 0x17,
	0xd7, 0x58, 0xd3, 0xb6, 0xb6, 0x2c, 0xed, 0x6e, 0x66, 0x46, 0x26, 0xce, 0x03, 0xf0, 0x04, 0x3c,
	0x04, 0x8f, 0x49, 0x75, 0xf7, 0xae, 0x56, 0x9b, 0x43, 0x0e, 0xbe, 0xf5, 0xf7, 0xf5, 0xf7, 0x75,
	0x4f, 0xf7, 0xac, 0x46, 0xaa, 0x77, 0x61, 0x7c, 0x3a, 0x1f, 0x17, 0x2e, 0x0f, 0xb9, 0xee, 0xcc,
	0x57, 0xe8, 0xbd, 0xb9, 0xc2, 0x83, 0x6e, 0xb1, 0xbe, 0x10, 0x72, 0x74, 0xa4, 0xfa, 0xbf, 0x9a,
	0x15, 0x9e, 0x06, 0xe3, 0x42, 0x8c, 0xaf, 0xf5, 0x44, 0xa9, 0x13, 0x33, 0xbf, 0xc6, 0xf0, 0x1b,
	0x1a, 0x0b, 0xad, 0x61, 0xeb, 0xb0, 0x37, 0x7d, 0x30, 0xae, 0x9c, 0xe3, 0xdf, 0x0b, 0x4e, 0xc6,
	0x5b, 0xa2, 0xd1, 0x3f, 0x3b, 0xaa, 0x43, 0x35, 0x8e, 0x4d, 0x30, 0xfa, 0xa1, 0x8a, 0x5c, 0xbe,
	0xce, 0xc4, 0xda, 0x8e, 0x05, 0x68, 0x50, 0xfb, 0x9e, 0x3a, 0x24, 0x1e, 0x76, 0x98, 0xaf, 0x20,
	0xe9, 0x31, 0xb3, 0x89, 0x87, 0xb6, 0xe8, 0x19, 0xe8, 0x27, 0xaa, 0x7b, 0xb9, 0x34, 0x57, 0x7f,
	0xfe, 0x9d, 0xa1, 0x83, 0x5d, 0xce, 0xd4, 0x84, 0x1e, 0xa9, 0x3e, 0x81, 0xb3, 0x13, 0x6b, 0x02,
	0x26, 0x1e, 0x22, 0x16, 0x34, 0x38, 0xfd, 0x54, 0xa9, 0xc2, 0xb8, 0x30, 0x39, 0x9d, 0xe7, 0x0e,
	0x61, 0x6f, 0xd8, 0x3a, 0x6c, 0xc5, 0x5b, 0x0c, 0x9d, 0x88, 0x51, 0xe2, 0x61, 0x5f, 0x4e, 0x54,
	0xc2, 0x2a, 0x33, 0x4d, 0x3c, 0x74, 0xea, 0xcc, 0xb4, 0xae, 0x39, 0x95, 0x9a, 0xdd, 0xba, 0xa6,
	0x30, 0xa3, 0xff, 0xba, 0xaa, 0x73, 0xe6, 0xd1, 0xf1, 0x22, 0xb4, 0xda, 0x5d, 0x7b, 0x74, 0xbc,
	0x87, 0x6e, 0xcc, 0x71, 0xbd, 0x1c, 0x5a, 0x42, 0x54, 0x2d, 0x47, 0xab, 0x5d, 0x2a, 0xc2, 0x1b,
	0x88, 0x62, 0x8e, 0xf5, 0x7d, 0xd5, 0x5e, 0xa4, 0x96, 0x47, 0xef, 0xc6, 0x14, 0x92, 0x2a, 0xdc,
	0x16, 0xc8, 0xc3, 0x46, 0x31, 0xc7, 0x74, 0x20, 0x8b, 0x97, 0x27, 0xe8, 0xe6, 0x98, 0x05, 0x1e,
	0x32, 0x8a, 0xb7, 0x18, 0x7d, 0x4f, 0xed, 0x2c, 0x0a, 0x9e, 0x2f, 0x8a, 0x77, 0x16, 0x85, 0x3e,
	0x50, 0x9d, 0x75, 0x51, 0x2e, 0x4d, 0x66, 0xdb, 0x60, 0xdd, 0x57, 0xad, 0x37, 0x3c, 0x53, 0x14,
	0xb7, 0xde, 0x10, 0xba, 0x05, 0x25, 0xe8, 0x96, 0xce, 0xed, 0x0b, 0x44, 0x0b, 0x3d, 0x39, 0x37,
	0x03, 0xfd, 0x4c, 0x0d, 0x1c, 0xda, 0xf5, 0x1c, 0x4f, 0x09, 0x26, 0x1e, 0xfa, 0x9c, 0x6d, 0x92,
	0x74, 0x95, 0x36, 0x75, 0x38, 0x0f, 0x69, 0x9e, 0xc1, 0x80, 0x15, 0x35, 0x41, 0xcb, 0xbe, 0x30,
	0xce, 0xa5, 0xe8, 0xe0, 0x1e, 0xe7, 0x2a, 0x48, 0x3d, 0x6d, 0xfa, 0xf6, 0xed, 0x2d, 0x7c, 0x20,
	0x3d, 0x19, 0x90, 0x9e, 0x83, 0xc4, 0xc3, 0x7d, 0xb9, 0x9c, 0x12, 0xea, 0x47, 0x6a, 0xcf, 0x2f,
	0x52, 0x5c, 0x5a, 0x78, 0xc0, 0x86, 0x12, 0xd1, 0xcc, 0x12, 0x25, 0x1e, 0xb4, 0xcc, 0x5c, 0x61,
	0xf2, 0xa4, 0xab, 0xd5, 0x3a, 0x43, 0xf8, 0x50, 0x3c, 0x82, 0xc8, 0x23, 0x51, 0xe2, 0xe1, 0x21,
	0x67, 0x36, 0x98, 0x3c, 0x61, 0x91, 0xbb, 0xcc, 0xc3, 0x47, 0xe2, 0x11, 0x44, 0x1e, 0x89, 0x12,
	0x0f, 0x8f, 0xc4, 0x53, 0x61, 0xba, 0x27, 0x1f, 0xf2, 0xe2, 0x8f, 0xfc, 0x86, 0x2a, 0x3e, 0x96,
	0x7b, 0xaa, 0x19, 0x3e, 0x63, 0x89, 0x00, 0xc4, 0x5b, 0x61, 0xea, 0x67, 0xac, 0x3d, 0xc6, 0x4b,
	0xf8, 0x58, 0xfa, 0x09, 0x22, 0x8f, 0x44, 0x89, 0x87, 0x03, 0xf1, 0x54, 0xb8, 0xf4, 0x1c, 0x85,
	0x6b, 0xf8, 0x64, 0xe3, 0x39, 0x0a, 0xd7, 0x7a, 0xa8, 0x7a, 0x45, 0xee, 0xcf, 0xaa, 0x4f, 0xe0,
	0x09, 0xaf, 0x63, 0x9b, 0x92, 0xad, 0x53, 0xee, 0x53, 0xf9, 0x39, 0x32, 0xa0, 0xad, 0x9b, 0xe5,
	0xf2, 0x28, 0x04, 0x07, 0x4f, 0xe5, 0x96, 0x4a, 0x48, 0xdf, 0xa9, 0xbd, 0x09, 0xf0, 0x19, 0xb3,
	0x14, 0x52, 0xef, 0x2b, 0x0c, 0xc7, 0x37, 0x01, 0x86, 0xd2, 0x5b, 0x10, 0xf1, 0x16, 0x3d, 0xf1,
	0x9f, 0x0b, 0x2f, 0x88, 0xbe, 0xeb, 0xeb, 0x74, 0xb9, 0x84, 0x91, 0x7c, 0xd7, 0x14, 0x73, 0xd5,
	0x14, 0xe1, 0x8b, 0xb2, 0x6a, 0x8a, 0xcc, 0x14, 0x1e, 0x9e, 0x95, 0x4c, 0xe1, 0xf5, 0x48, 0x0d,
	0x3c, 0x89, 0xcf, 0x27, 0x93, 0x17, 0x93, 0xf3, 0xb9, 0x85, 0x2f, 0x39, 0xd7, 0x63, 0x92, 0xb8,
	0x57, 0xb6, 0xa1, 0x99, 0x92, 0xe6, 0x79, 0x53, 0x33, 0x7d, 0x47, 0x33, 0x23, 0xcd, 0x57, 0x4d,
	0xcd, 0xec, 0x1d, 0xcd, 0x4b, 0xd2, 0x1c, 0x36, 0x35, 0x2f, 0x5f, 0x59, 0x3d, 0x54, 0xfd, 0x34,
	0xe0, 0xea, 0x7c, 0xf2, 0x42, 0x5a, 0x7d, 0x2d, 0xb7, 0x4c, 0x1c, 0x51, 0x4d, 0x05, 0x37, 0xfa,
	0xa6, 0xa1, 0x98, 0x35, 0x15, 0xdc, 0xe6, 0xdb, 0x86, 0x82, 0xba, 0x3c, 0x56, 0xfb, 0xf8, 0x7a,
	0x9d, 0x16, 0xe7, 0x13, 0xf8, 0x4e, 0xd6, 0xc8, 0x70, 0x52, 0x27, 0x66, 0xf0, 0xfd, 0x56, 0x62,
	0xc6, 0x0f, 0x2c, 0x45, 0x30, 0x1e, 0xb6, 0xe9, 0x77, 0xc4, 0x60, 0xf4, 0x6f, 0x4b, 0x0d, 0xb6,
	0xde, 0x7d, 0x5f, 0xdc, 0xe1, 0xe1, 0xd7, 0xcf, 0xd5, 0xee, 0x95, 0x59, 0x21, 0xbf, 0x66, 0xbd,
	0xa9, 0xae, 0xc5, 0xd5, 0xbf, 0x41, 0xcc, 0x79, 0x7d, 0xa8, 0x22, 0x7a, 0xfe, 0xe8, 0x8d, 0x6f,
	0x37, 0x85, 0xd5, 0x6b, 0x19, 0x8b, 0x60, 0xf4, 0x93, 0x52, 0xe4, 0xfd, 0x25, 0xb3, 0x77, 0xfc,
	0x2f, 0x5a, 0xa8, 0xde, 0xa6, 0xc0, 0x9d, 0x87, 0xb2, 0x26, 0x98, 0xf7, 0x0d, 0x45, 0xf9, 0x9f,
	0x07, 0x7f, 0xf5, 0xc6, 0x3f, 0xfc, 0x58, 0x65, 0x2f, 0xf6, 0xf8, 0xef, 0x74, 0xf6, 0x7f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x63, 0x7e, 0x81, 0x01, 0x72, 0x07, 0x00, 0x00,
}
