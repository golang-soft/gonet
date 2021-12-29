// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

package message

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

//point3F
type Point3F struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=X,proto3" json:"X,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=Z,proto3" json:"Z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point3F) Reset()         { *m = Point3F{} }
func (m *Point3F) String() string { return proto.CompactTextString(m) }
func (*Point3F) ProtoMessage()    {}
func (*Point3F) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{0}
}

func (m *Point3F) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point3F.Unmarshal(m, b)
}
func (m *Point3F) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point3F.Marshal(b, m, deterministic)
}
func (m *Point3F) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point3F.Merge(m, src)
}
func (m *Point3F) XXX_Size() int {
	return xxx_messageInfo_Point3F.Size(m)
}
func (m *Point3F) XXX_DiscardUnknown() {
	xxx_messageInfo_Point3F.DiscardUnknown(m)
}

var xxx_messageInfo_Point3F proto.InternalMessageInfo

func (m *Point3F) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Point3F) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Point3F) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

//移动包
type C_Z_Move struct {
	PacketHead           *Ipacket       `protobuf:"bytes,1,opt,name=PacketHead,proto3" json:"PacketHead,omitempty"`
	Move                 *C_Z_Move_Move `protobuf:"bytes,2,opt,name=move,proto3" json:"move,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *C_Z_Move) Reset()         { *m = C_Z_Move{} }
func (m *C_Z_Move) String() string { return proto.CompactTextString(m) }
func (*C_Z_Move) ProtoMessage()    {}
func (*C_Z_Move) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{1}
}

func (m *C_Z_Move) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C_Z_Move.Unmarshal(m, b)
}
func (m *C_Z_Move) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C_Z_Move.Marshal(b, m, deterministic)
}
func (m *C_Z_Move) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C_Z_Move.Merge(m, src)
}
func (m *C_Z_Move) XXX_Size() int {
	return xxx_messageInfo_C_Z_Move.Size(m)
}
func (m *C_Z_Move) XXX_DiscardUnknown() {
	xxx_messageInfo_C_Z_Move.DiscardUnknown(m)
}

var xxx_messageInfo_C_Z_Move proto.InternalMessageInfo

func (m *C_Z_Move) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *C_Z_Move) GetMove() *C_Z_Move_Move {
	if m != nil {
		return m.Move
	}
	return nil
}

type C_Z_Move_Move struct {
	Mode                 int32                 `protobuf:"varint,1,opt,name=Mode,proto3" json:"Mode,omitempty"`
	Normal               *C_Z_Move_Move_Normal `protobuf:"bytes,2,opt,name=normal,proto3" json:"normal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *C_Z_Move_Move) Reset()         { *m = C_Z_Move_Move{} }
func (m *C_Z_Move_Move) String() string { return proto.CompactTextString(m) }
func (*C_Z_Move_Move) ProtoMessage()    {}
func (*C_Z_Move_Move) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{1, 0}
}

func (m *C_Z_Move_Move) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C_Z_Move_Move.Unmarshal(m, b)
}
func (m *C_Z_Move_Move) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C_Z_Move_Move.Marshal(b, m, deterministic)
}
func (m *C_Z_Move_Move) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C_Z_Move_Move.Merge(m, src)
}
func (m *C_Z_Move_Move) XXX_Size() int {
	return xxx_messageInfo_C_Z_Move_Move.Size(m)
}
func (m *C_Z_Move_Move) XXX_DiscardUnknown() {
	xxx_messageInfo_C_Z_Move_Move.DiscardUnknown(m)
}

var xxx_messageInfo_C_Z_Move_Move proto.InternalMessageInfo

func (m *C_Z_Move_Move) GetMode() int32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *C_Z_Move_Move) GetNormal() *C_Z_Move_Move_Normal {
	if m != nil {
		return m.Normal
	}
	return nil
}

type C_Z_Move_Move_Normal struct {
	Pos                  *Point3F `protobuf:"bytes,1,opt,name=Pos,proto3" json:"Pos,omitempty"`
	Yaw                  float32  `protobuf:"fixed32,2,opt,name=Yaw,proto3" json:"Yaw,omitempty"`
	Duration             float32  `protobuf:"fixed32,3,opt,name=Duration,proto3" json:"Duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C_Z_Move_Move_Normal) Reset()         { *m = C_Z_Move_Move_Normal{} }
func (m *C_Z_Move_Move_Normal) String() string { return proto.CompactTextString(m) }
func (*C_Z_Move_Move_Normal) ProtoMessage()    {}
func (*C_Z_Move_Move_Normal) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{1, 0, 0}
}

func (m *C_Z_Move_Move_Normal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C_Z_Move_Move_Normal.Unmarshal(m, b)
}
func (m *C_Z_Move_Move_Normal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C_Z_Move_Move_Normal.Marshal(b, m, deterministic)
}
func (m *C_Z_Move_Move_Normal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C_Z_Move_Move_Normal.Merge(m, src)
}
func (m *C_Z_Move_Move_Normal) XXX_Size() int {
	return xxx_messageInfo_C_Z_Move_Move_Normal.Size(m)
}
func (m *C_Z_Move_Move_Normal) XXX_DiscardUnknown() {
	xxx_messageInfo_C_Z_Move_Move_Normal.DiscardUnknown(m)
}

var xxx_messageInfo_C_Z_Move_Move_Normal proto.InternalMessageInfo

func (m *C_Z_Move_Move_Normal) GetPos() *Point3F {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *C_Z_Move_Move_Normal) GetYaw() float32 {
	if m != nil {
		return m.Yaw
	}
	return 0
}

func (m *C_Z_Move_Move_Normal) GetDuration() float32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

//skill
type C_Z_Skill struct {
	PacketHead           *Ipacket `protobuf:"bytes,1,opt,name=PacketHead,proto3" json:"PacketHead,omitempty"`
	SkillId              int32    `protobuf:"varint,2,opt,name=SkillId,proto3" json:"SkillId,omitempty"`
	TargetId             int64    `protobuf:"varint,3,opt,name=TargetId,proto3" json:"TargetId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C_Z_Skill) Reset()         { *m = C_Z_Skill{} }
func (m *C_Z_Skill) String() string { return proto.CompactTextString(m) }
func (*C_Z_Skill) ProtoMessage()    {}
func (*C_Z_Skill) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{2}
}

func (m *C_Z_Skill) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C_Z_Skill.Unmarshal(m, b)
}
func (m *C_Z_Skill) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C_Z_Skill.Marshal(b, m, deterministic)
}
func (m *C_Z_Skill) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C_Z_Skill.Merge(m, src)
}
func (m *C_Z_Skill) XXX_Size() int {
	return xxx_messageInfo_C_Z_Skill.Size(m)
}
func (m *C_Z_Skill) XXX_DiscardUnknown() {
	xxx_messageInfo_C_Z_Skill.DiscardUnknown(m)
}

var xxx_messageInfo_C_Z_Skill proto.InternalMessageInfo

func (m *C_Z_Skill) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *C_Z_Skill) GetSkillId() int32 {
	if m != nil {
		return m.SkillId
	}
	return 0
}

func (m *C_Z_Skill) GetTargetId() int64 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

type Z_C_LoginMap struct {
	PacketHead           *Ipacket `protobuf:"bytes,1,opt,name=PacketHead,proto3" json:"PacketHead,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	Pos                  *Point3F `protobuf:"bytes,3,opt,name=Pos,proto3" json:"Pos,omitempty"`
	Rotation             float32  `protobuf:"fixed32,4,opt,name=Rotation,proto3" json:"Rotation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Z_C_LoginMap) Reset()         { *m = Z_C_LoginMap{} }
func (m *Z_C_LoginMap) String() string { return proto.CompactTextString(m) }
func (*Z_C_LoginMap) ProtoMessage()    {}
func (*Z_C_LoginMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{3}
}

func (m *Z_C_LoginMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Z_C_LoginMap.Unmarshal(m, b)
}
func (m *Z_C_LoginMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Z_C_LoginMap.Marshal(b, m, deterministic)
}
func (m *Z_C_LoginMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Z_C_LoginMap.Merge(m, src)
}
func (m *Z_C_LoginMap) XXX_Size() int {
	return xxx_messageInfo_Z_C_LoginMap.Size(m)
}
func (m *Z_C_LoginMap) XXX_DiscardUnknown() {
	xxx_messageInfo_Z_C_LoginMap.DiscardUnknown(m)
}

var xxx_messageInfo_Z_C_LoginMap proto.InternalMessageInfo

func (m *Z_C_LoginMap) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *Z_C_LoginMap) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Z_C_LoginMap) GetPos() *Point3F {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *Z_C_LoginMap) GetRotation() float32 {
	if m != nil {
		return m.Rotation
	}
	return 0
}

type Z_C_ENTITY struct {
	PacketHead           *Ipacket             `protobuf:"bytes,1,opt,name=PacketHead,proto3" json:"PacketHead,omitempty"`
	EntityInfo           []*Z_C_ENTITY_Entity `protobuf:"bytes,2,rep,name=EntityInfo,proto3" json:"EntityInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Z_C_ENTITY) Reset()         { *m = Z_C_ENTITY{} }
func (m *Z_C_ENTITY) String() string { return proto.CompactTextString(m) }
func (*Z_C_ENTITY) ProtoMessage()    {}
func (*Z_C_ENTITY) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{4}
}

func (m *Z_C_ENTITY) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Z_C_ENTITY.Unmarshal(m, b)
}
func (m *Z_C_ENTITY) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Z_C_ENTITY.Marshal(b, m, deterministic)
}
func (m *Z_C_ENTITY) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Z_C_ENTITY.Merge(m, src)
}
func (m *Z_C_ENTITY) XXX_Size() int {
	return xxx_messageInfo_Z_C_ENTITY.Size(m)
}
func (m *Z_C_ENTITY) XXX_DiscardUnknown() {
	xxx_messageInfo_Z_C_ENTITY.DiscardUnknown(m)
}

var xxx_messageInfo_Z_C_ENTITY proto.InternalMessageInfo

func (m *Z_C_ENTITY) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *Z_C_ENTITY) GetEntityInfo() []*Z_C_ENTITY_Entity {
	if m != nil {
		return m.EntityInfo
	}
	return nil
}

type Z_C_ENTITY_Entity struct {
	Id                   int64                        `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Data                 *Z_C_ENTITY_Entity_DataMask  `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	Move                 *Z_C_ENTITY_Entity_MoveMask  `protobuf:"bytes,3,opt,name=Move,proto3" json:"Move,omitempty"`
	Stats                *Z_C_ENTITY_Entity_StatsMask `protobuf:"bytes,4,opt,name=Stats,proto3" json:"Stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Z_C_ENTITY_Entity) Reset()         { *m = Z_C_ENTITY_Entity{} }
func (m *Z_C_ENTITY_Entity) String() string { return proto.CompactTextString(m) }
func (*Z_C_ENTITY_Entity) ProtoMessage()    {}
func (*Z_C_ENTITY_Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{4, 0}
}

func (m *Z_C_ENTITY_Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Z_C_ENTITY_Entity.Unmarshal(m, b)
}
func (m *Z_C_ENTITY_Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Z_C_ENTITY_Entity.Marshal(b, m, deterministic)
}
func (m *Z_C_ENTITY_Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Z_C_ENTITY_Entity.Merge(m, src)
}
func (m *Z_C_ENTITY_Entity) XXX_Size() int {
	return xxx_messageInfo_Z_C_ENTITY_Entity.Size(m)
}
func (m *Z_C_ENTITY_Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Z_C_ENTITY_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Z_C_ENTITY_Entity proto.InternalMessageInfo

func (m *Z_C_ENTITY_Entity) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Z_C_ENTITY_Entity) GetData() *Z_C_ENTITY_Entity_DataMask {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Z_C_ENTITY_Entity) GetMove() *Z_C_ENTITY_Entity_MoveMask {
	if m != nil {
		return m.Move
	}
	return nil
}

func (m *Z_C_ENTITY_Entity) GetStats() *Z_C_ENTITY_Entity_StatsMask {
	if m != nil {
		return m.Stats
	}
	return nil
}

type Z_C_ENTITY_Entity_DataMask struct {
	Type                 int32                                     `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	RemoveFlag           bool                                      `protobuf:"varint,2,opt,name=RemoveFlag,proto3" json:"RemoveFlag,omitempty"`
	NpcData              *Z_C_ENTITY_Entity_DataMask_NpcDataMask   `protobuf:"bytes,3,opt,name=NpcData,proto3" json:"NpcData,omitempty"`
	SpellData            *Z_C_ENTITY_Entity_DataMask_SpellDataMask `protobuf:"bytes,4,opt,name=SpellData,proto3" json:"SpellData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *Z_C_ENTITY_Entity_DataMask) Reset()         { *m = Z_C_ENTITY_Entity_DataMask{} }
func (m *Z_C_ENTITY_Entity_DataMask) String() string { return proto.CompactTextString(m) }
func (*Z_C_ENTITY_Entity_DataMask) ProtoMessage()    {}
func (*Z_C_ENTITY_Entity_DataMask) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{4, 0, 0}
}

func (m *Z_C_ENTITY_Entity_DataMask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Z_C_ENTITY_Entity_DataMask.Unmarshal(m, b)
}
func (m *Z_C_ENTITY_Entity_DataMask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Z_C_ENTITY_Entity_DataMask.Marshal(b, m, deterministic)
}
func (m *Z_C_ENTITY_Entity_DataMask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Z_C_ENTITY_Entity_DataMask.Merge(m, src)
}
func (m *Z_C_ENTITY_Entity_DataMask) XXX_Size() int {
	return xxx_messageInfo_Z_C_ENTITY_Entity_DataMask.Size(m)
}
func (m *Z_C_ENTITY_Entity_DataMask) XXX_DiscardUnknown() {
	xxx_messageInfo_Z_C_ENTITY_Entity_DataMask.DiscardUnknown(m)
}

var xxx_messageInfo_Z_C_ENTITY_Entity_DataMask proto.InternalMessageInfo

func (m *Z_C_ENTITY_Entity_DataMask) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_DataMask) GetRemoveFlag() bool {
	if m != nil {
		return m.RemoveFlag
	}
	return false
}

func (m *Z_C_ENTITY_Entity_DataMask) GetNpcData() *Z_C_ENTITY_Entity_DataMask_NpcDataMask {
	if m != nil {
		return m.NpcData
	}
	return nil
}

func (m *Z_C_ENTITY_Entity_DataMask) GetSpellData() *Z_C_ENTITY_Entity_DataMask_SpellDataMask {
	if m != nil {
		return m.SpellData
	}
	return nil
}

type Z_C_ENTITY_Entity_DataMask_NpcDataMask struct {
	DataId               int32    `protobuf:"varint,1,opt,name=DataId,proto3" json:"DataId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Z_C_ENTITY_Entity_DataMask_NpcDataMask) Reset() {
	*m = Z_C_ENTITY_Entity_DataMask_NpcDataMask{}
}
func (m *Z_C_ENTITY_Entity_DataMask_NpcDataMask) String() string { return proto.CompactTextString(m) }
func (*Z_C_ENTITY_Entity_DataMask_NpcDataMask) ProtoMessage()    {}
func (*Z_C_ENTITY_Entity_DataMask_NpcDataMask) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{4, 0, 0, 0}
}

func (m *Z_C_ENTITY_Entity_DataMask_NpcDataMask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Z_C_ENTITY_Entity_DataMask_NpcDataMask.Unmarshal(m, b)
}
func (m *Z_C_ENTITY_Entity_DataMask_NpcDataMask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Z_C_ENTITY_Entity_DataMask_NpcDataMask.Marshal(b, m, deterministic)
}
func (m *Z_C_ENTITY_Entity_DataMask_NpcDataMask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Z_C_ENTITY_Entity_DataMask_NpcDataMask.Merge(m, src)
}
func (m *Z_C_ENTITY_Entity_DataMask_NpcDataMask) XXX_Size() int {
	return xxx_messageInfo_Z_C_ENTITY_Entity_DataMask_NpcDataMask.Size(m)
}
func (m *Z_C_ENTITY_Entity_DataMask_NpcDataMask) XXX_DiscardUnknown() {
	xxx_messageInfo_Z_C_ENTITY_Entity_DataMask_NpcDataMask.DiscardUnknown(m)
}

var xxx_messageInfo_Z_C_ENTITY_Entity_DataMask_NpcDataMask proto.InternalMessageInfo

func (m *Z_C_ENTITY_Entity_DataMask_NpcDataMask) GetDataId() int32 {
	if m != nil {
		return m.DataId
	}
	return 0
}

type Z_C_ENTITY_Entity_DataMask_SpellDataMask struct {
	DataId               int32    `protobuf:"varint,1,opt,name=DataId,proto3" json:"DataId,omitempty"`
	LifeTime             int32    `protobuf:"varint,2,opt,name=LifeTime,proto3" json:"LifeTime,omitempty"`
	FlySpeed             int32    `protobuf:"varint,3,opt,name=FlySpeed,proto3" json:"FlySpeed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Z_C_ENTITY_Entity_DataMask_SpellDataMask) Reset() {
	*m = Z_C_ENTITY_Entity_DataMask_SpellDataMask{}
}
func (m *Z_C_ENTITY_Entity_DataMask_SpellDataMask) String() string { return proto.CompactTextString(m) }
func (*Z_C_ENTITY_Entity_DataMask_SpellDataMask) ProtoMessage()    {}
func (*Z_C_ENTITY_Entity_DataMask_SpellDataMask) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{4, 0, 0, 1}
}

func (m *Z_C_ENTITY_Entity_DataMask_SpellDataMask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Z_C_ENTITY_Entity_DataMask_SpellDataMask.Unmarshal(m, b)
}
func (m *Z_C_ENTITY_Entity_DataMask_SpellDataMask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Z_C_ENTITY_Entity_DataMask_SpellDataMask.Marshal(b, m, deterministic)
}
func (m *Z_C_ENTITY_Entity_DataMask_SpellDataMask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Z_C_ENTITY_Entity_DataMask_SpellDataMask.Merge(m, src)
}
func (m *Z_C_ENTITY_Entity_DataMask_SpellDataMask) XXX_Size() int {
	return xxx_messageInfo_Z_C_ENTITY_Entity_DataMask_SpellDataMask.Size(m)
}
func (m *Z_C_ENTITY_Entity_DataMask_SpellDataMask) XXX_DiscardUnknown() {
	xxx_messageInfo_Z_C_ENTITY_Entity_DataMask_SpellDataMask.DiscardUnknown(m)
}

var xxx_messageInfo_Z_C_ENTITY_Entity_DataMask_SpellDataMask proto.InternalMessageInfo

func (m *Z_C_ENTITY_Entity_DataMask_SpellDataMask) GetDataId() int32 {
	if m != nil {
		return m.DataId
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_DataMask_SpellDataMask) GetLifeTime() int32 {
	if m != nil {
		return m.LifeTime
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_DataMask_SpellDataMask) GetFlySpeed() int32 {
	if m != nil {
		return m.FlySpeed
	}
	return 0
}

type Z_C_ENTITY_Entity_MoveMask struct {
	Pos                  *Point3F `protobuf:"bytes,1,opt,name=Pos,proto3" json:"Pos,omitempty"`
	Rotation             float32  `protobuf:"fixed32,2,opt,name=Rotation,proto3" json:"Rotation,omitempty"`
	ForceUpdateFlag      bool     `protobuf:"varint,3,opt,name=ForceUpdateFlag,proto3" json:"ForceUpdateFlag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Z_C_ENTITY_Entity_MoveMask) Reset()         { *m = Z_C_ENTITY_Entity_MoveMask{} }
func (m *Z_C_ENTITY_Entity_MoveMask) String() string { return proto.CompactTextString(m) }
func (*Z_C_ENTITY_Entity_MoveMask) ProtoMessage()    {}
func (*Z_C_ENTITY_Entity_MoveMask) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{4, 0, 1}
}

func (m *Z_C_ENTITY_Entity_MoveMask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Z_C_ENTITY_Entity_MoveMask.Unmarshal(m, b)
}
func (m *Z_C_ENTITY_Entity_MoveMask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Z_C_ENTITY_Entity_MoveMask.Marshal(b, m, deterministic)
}
func (m *Z_C_ENTITY_Entity_MoveMask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Z_C_ENTITY_Entity_MoveMask.Merge(m, src)
}
func (m *Z_C_ENTITY_Entity_MoveMask) XXX_Size() int {
	return xxx_messageInfo_Z_C_ENTITY_Entity_MoveMask.Size(m)
}
func (m *Z_C_ENTITY_Entity_MoveMask) XXX_DiscardUnknown() {
	xxx_messageInfo_Z_C_ENTITY_Entity_MoveMask.DiscardUnknown(m)
}

var xxx_messageInfo_Z_C_ENTITY_Entity_MoveMask proto.InternalMessageInfo

func (m *Z_C_ENTITY_Entity_MoveMask) GetPos() *Point3F {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *Z_C_ENTITY_Entity_MoveMask) GetRotation() float32 {
	if m != nil {
		return m.Rotation
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_MoveMask) GetForceUpdateFlag() bool {
	if m != nil {
		return m.ForceUpdateFlag
	}
	return false
}

type Z_C_ENTITY_Entity_StatsMask struct {
	HP                   int32    `protobuf:"varint,1,opt,name=HP,proto3" json:"HP,omitempty"`
	MP                   int32    `protobuf:"varint,2,opt,name=MP,proto3" json:"MP,omitempty"`
	MaxHP                int32    `protobuf:"varint,3,opt,name=MaxHP,proto3" json:"MaxHP,omitempty"`
	MaxMP                int32    `protobuf:"varint,4,opt,name=MaxMP,proto3" json:"MaxMP,omitempty"`
	PhyDamage            int32    `protobuf:"varint,5,opt,name=PhyDamage,proto3" json:"PhyDamage,omitempty"`
	PhyDefence           int32    `protobuf:"varint,6,opt,name=PhyDefence,proto3" json:"PhyDefence,omitempty"`
	SplDamage            int32    `protobuf:"varint,7,opt,name=SplDamage,proto3" json:"SplDamage,omitempty"`
	SplDefence           int32    `protobuf:"varint,8,opt,name=SplDefence,proto3" json:"SplDefence,omitempty"`
	Heal                 int32    `protobuf:"varint,9,opt,name=Heal,proto3" json:"Heal,omitempty"`
	CriticalTimes        int32    `protobuf:"varint,10,opt,name=CriticalTimes,proto3" json:"CriticalTimes,omitempty"`
	Critical             int32    `protobuf:"varint,11,opt,name=Critical,proto3" json:"Critical,omitempty"`
	AntiCriticalTimes    int32    `protobuf:"varint,12,opt,name=AntiCriticalTimes,proto3" json:"AntiCriticalTimes,omitempty"`
	AntiCritical         int32    `protobuf:"varint,13,opt,name=AntiCritical,proto3" json:"AntiCritical,omitempty"`
	Dodge                int32    `protobuf:"varint,14,opt,name=Dodge,proto3" json:"Dodge,omitempty"`
	Hit                  int32    `protobuf:"varint,15,opt,name=Hit,proto3" json:"Hit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Z_C_ENTITY_Entity_StatsMask) Reset()         { *m = Z_C_ENTITY_Entity_StatsMask{} }
func (m *Z_C_ENTITY_Entity_StatsMask) String() string { return proto.CompactTextString(m) }
func (*Z_C_ENTITY_Entity_StatsMask) ProtoMessage()    {}
func (*Z_C_ENTITY_Entity_StatsMask) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{4, 0, 2}
}

func (m *Z_C_ENTITY_Entity_StatsMask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Z_C_ENTITY_Entity_StatsMask.Unmarshal(m, b)
}
func (m *Z_C_ENTITY_Entity_StatsMask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Z_C_ENTITY_Entity_StatsMask.Marshal(b, m, deterministic)
}
func (m *Z_C_ENTITY_Entity_StatsMask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Z_C_ENTITY_Entity_StatsMask.Merge(m, src)
}
func (m *Z_C_ENTITY_Entity_StatsMask) XXX_Size() int {
	return xxx_messageInfo_Z_C_ENTITY_Entity_StatsMask.Size(m)
}
func (m *Z_C_ENTITY_Entity_StatsMask) XXX_DiscardUnknown() {
	xxx_messageInfo_Z_C_ENTITY_Entity_StatsMask.DiscardUnknown(m)
}

var xxx_messageInfo_Z_C_ENTITY_Entity_StatsMask proto.InternalMessageInfo

func (m *Z_C_ENTITY_Entity_StatsMask) GetHP() int32 {
	if m != nil {
		return m.HP
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_StatsMask) GetMP() int32 {
	if m != nil {
		return m.MP
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_StatsMask) GetMaxHP() int32 {
	if m != nil {
		return m.MaxHP
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_StatsMask) GetMaxMP() int32 {
	if m != nil {
		return m.MaxMP
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_StatsMask) GetPhyDamage() int32 {
	if m != nil {
		return m.PhyDamage
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_StatsMask) GetPhyDefence() int32 {
	if m != nil {
		return m.PhyDefence
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_StatsMask) GetSplDamage() int32 {
	if m != nil {
		return m.SplDamage
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_StatsMask) GetSplDefence() int32 {
	if m != nil {
		return m.SplDefence
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_StatsMask) GetHeal() int32 {
	if m != nil {
		return m.Heal
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_StatsMask) GetCriticalTimes() int32 {
	if m != nil {
		return m.CriticalTimes
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_StatsMask) GetCritical() int32 {
	if m != nil {
		return m.Critical
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_StatsMask) GetAntiCriticalTimes() int32 {
	if m != nil {
		return m.AntiCriticalTimes
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_StatsMask) GetAntiCritical() int32 {
	if m != nil {
		return m.AntiCritical
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_StatsMask) GetDodge() int32 {
	if m != nil {
		return m.Dodge
	}
	return 0
}

func (m *Z_C_ENTITY_Entity_StatsMask) GetHit() int32 {
	if m != nil {
		return m.Hit
	}
	return 0
}

type C_Z_LoginCopyMap struct {
	PacketHead           *Ipacket `protobuf:"bytes,1,opt,name=PacketHead,proto3" json:"PacketHead,omitempty"`
	DataId               int32    `protobuf:"varint,2,opt,name=DataId,proto3" json:"DataId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C_Z_LoginCopyMap) Reset()         { *m = C_Z_LoginCopyMap{} }
func (m *C_Z_LoginCopyMap) String() string { return proto.CompactTextString(m) }
func (*C_Z_LoginCopyMap) ProtoMessage()    {}
func (*C_Z_LoginCopyMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{5}
}

func (m *C_Z_LoginCopyMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C_Z_LoginCopyMap.Unmarshal(m, b)
}
func (m *C_Z_LoginCopyMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C_Z_LoginCopyMap.Marshal(b, m, deterministic)
}
func (m *C_Z_LoginCopyMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C_Z_LoginCopyMap.Merge(m, src)
}
func (m *C_Z_LoginCopyMap) XXX_Size() int {
	return xxx_messageInfo_C_Z_LoginCopyMap.Size(m)
}
func (m *C_Z_LoginCopyMap) XXX_DiscardUnknown() {
	xxx_messageInfo_C_Z_LoginCopyMap.DiscardUnknown(m)
}

var xxx_messageInfo_C_Z_LoginCopyMap proto.InternalMessageInfo

func (m *C_Z_LoginCopyMap) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *C_Z_LoginCopyMap) GetDataId() int32 {
	if m != nil {
		return m.DataId
	}
	return 0
}

func init() {
	proto.RegisterType((*Point3F)(nil), "message.Point3F")
	proto.RegisterType((*C_Z_Move)(nil), "message.C_Z_Move")
	proto.RegisterType((*C_Z_Move_Move)(nil), "message.C_Z_Move.Move")
	proto.RegisterType((*C_Z_Move_Move_Normal)(nil), "message.C_Z_Move.Move.Normal")
	proto.RegisterType((*C_Z_Skill)(nil), "message.C_Z_Skill")
	proto.RegisterType((*Z_C_LoginMap)(nil), "message.Z_C_LoginMap")
	proto.RegisterType((*Z_C_ENTITY)(nil), "message.Z_C_ENTITY")
	proto.RegisterType((*Z_C_ENTITY_Entity)(nil), "message.Z_C_ENTITY.Entity")
	proto.RegisterType((*Z_C_ENTITY_Entity_DataMask)(nil), "message.Z_C_ENTITY.Entity.DataMask")
	proto.RegisterType((*Z_C_ENTITY_Entity_DataMask_NpcDataMask)(nil), "message.Z_C_ENTITY.Entity.DataMask.NpcDataMask")
	proto.RegisterType((*Z_C_ENTITY_Entity_DataMask_SpellDataMask)(nil), "message.Z_C_ENTITY.Entity.DataMask.SpellDataMask")
	proto.RegisterType((*Z_C_ENTITY_Entity_MoveMask)(nil), "message.Z_C_ENTITY.Entity.MoveMask")
	proto.RegisterType((*Z_C_ENTITY_Entity_StatsMask)(nil), "message.Z_C_ENTITY.Entity.StatsMask")
	proto.RegisterType((*C_Z_LoginCopyMap)(nil), "message.C_Z_LoginCopyMap")
}

func init() { proto.RegisterFile("game.proto", fileDescriptor_38fc58335341d769) }

var fileDescriptor_38fc58335341d769 = []byte{
	// 794 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x8a, 0x23, 0x45,
	0x14, 0x26, 0xdd, 0xf9, 0xe9, 0x9c, 0x64, 0x76, 0xc7, 0x42, 0x96, 0xa6, 0x51, 0x59, 0xda, 0x15,
	0x06, 0x91, 0x5e, 0xcd, 0x20, 0xc2, 0x7a, 0xa5, 0xc9, 0x86, 0x04, 0xb6, 0xc7, 0xa6, 0x12, 0x61,
	0x13, 0x84, 0x50, 0x9b, 0xd4, 0xc4, 0x66, 0x3a, 0xdd, 0x4d, 0x52, 0x3b, 0x9a, 0xb7, 0xf0, 0xd2,
	0x7b, 0xaf, 0x7d, 0x0b, 0xc1, 0xf7, 0xf0, 0x49, 0xe4, 0x9c, 0xaa, 0xee, 0x74, 0x1c, 0x67, 0x18,
	0x72, 0x93, 0xd4, 0xf7, 0xd5, 0xf9, 0xce, 0xa9, 0xf3, 0x53, 0xd5, 0x00, 0x6b, 0xb1, 0x91, 0x41,
	0xbe, 0xcd, 0x54, 0xc6, 0x5a, 0x1b, 0xb9, 0xdb, 0x89, 0xb5, 0xf4, 0xda, 0xf9, 0xfb, 0x77, 0x9a,
	0xf3, 0x2f, 0xa1, 0x15, 0x65, 0x71, 0xaa, 0x2e, 0x87, 0xac, 0x0b, 0xb5, 0xb7, 0x6e, 0xed, 0x79,
	0xed, 0xc2, 0xe2, 0xb5, 0xb7, 0x88, 0x66, 0xae, 0xa5, 0xd1, 0x0c, 0xd1, 0xdc, 0xb5, 0x35, 0x9a,
	0xfb, 0xbf, 0x5b, 0xe0, 0xf4, 0x17, 0xf3, 0x45, 0x98, 0xdd, 0x4a, 0xf6, 0x25, 0x40, 0x24, 0x96,
	0x37, 0x52, 0x8d, 0xa4, 0x58, 0x91, 0xbe, 0xd3, 0x3b, 0x0f, 0x4c, 0xa8, 0x60, 0x9c, 0xd3, 0x1e,
	0xaf, 0xd8, 0xb0, 0xcf, 0xa1, 0xbe, 0xc9, 0x6e, 0x25, 0x79, 0xef, 0xf4, 0x9e, 0x95, 0xb6, 0x85,
	0xcb, 0x00, 0x7f, 0x38, 0xd9, 0x78, 0x7f, 0xd6, 0xa0, 0x4e, 0x61, 0x18, 0xfe, 0xaf, 0x24, 0x05,
	0x68, 0x70, 0x5a, 0xb3, 0xaf, 0xa1, 0x99, 0x66, 0xdb, 0x8d, 0x48, 0x8c, 0xab, 0x8f, 0xff, 0xdf,
	0x55, 0x70, 0x45, 0x46, 0xdc, 0x18, 0x7b, 0x73, 0x68, 0x6a, 0x86, 0xf9, 0x60, 0x47, 0xd9, 0xee,
	0xce, 0xa1, 0x4d, 0x45, 0x38, 0x6e, 0xb2, 0x73, 0xb0, 0x67, 0xe2, 0x17, 0x53, 0x0a, 0x5c, 0x32,
	0x0f, 0x9c, 0xc1, 0xfb, 0xad, 0x50, 0x71, 0x96, 0x9a, 0x9a, 0x94, 0xd8, 0xdf, 0x41, 0x1b, 0x63,
	0x4f, 0x6e, 0xe2, 0x24, 0x39, 0xa1, 0x34, 0x2e, 0xb4, 0x48, 0x3a, 0x5e, 0x51, 0xc0, 0x06, 0x2f,
	0x20, 0x06, 0x9d, 0x8a, 0xed, 0x5a, 0xaa, 0xf1, 0x8a, 0x82, 0xda, 0xbc, 0xc4, 0xfe, 0x6f, 0x35,
	0xe8, 0xce, 0x17, 0xfd, 0xc5, 0x9b, 0x6c, 0x1d, 0xa7, 0xa1, 0xc8, 0x4f, 0x08, 0xfc, 0x04, 0x2c,
	0x13, 0xd3, 0xe6, 0xd6, 0x78, 0x55, 0x54, 0xc6, 0x7e, 0xa8, 0x32, 0x1e, 0x38, 0x3c, 0x53, 0xba,
	0x0e, 0x75, 0x5d, 0x87, 0x02, 0xfb, 0x7f, 0xb5, 0x01, 0xf0, 0x48, 0xaf, 0xaf, 0xa6, 0xe3, 0xe9,
	0xec, 0x84, 0x03, 0xbd, 0x02, 0x78, 0x9d, 0xaa, 0x58, 0xed, 0xc7, 0xe9, 0x75, 0xe6, 0x5a, 0xcf,
	0xed, 0x8b, 0x4e, 0xcf, 0x2b, 0x15, 0x07, 0xd7, 0x81, 0xb6, 0xe2, 0x15, 0x6b, 0xef, 0x6f, 0x07,
	0x9a, 0x1a, 0x9a, 0xbc, 0x6a, 0x65, 0x5e, 0xdf, 0x40, 0x7d, 0x20, 0x94, 0x30, 0x03, 0xf3, 0xe9,
	0xfd, 0x0e, 0x03, 0x34, 0x0b, 0xc5, 0xee, 0x86, 0x93, 0x00, 0x85, 0x38, 0x4b, 0xa6, 0x22, 0x0f,
	0x09, 0xd1, 0x4c, 0x0b, 0x69, 0x70, 0x5f, 0x41, 0x63, 0xa2, 0x84, 0xda, 0x51, 0x89, 0x3a, 0xbd,
	0x17, 0x0f, 0x28, 0xc9, 0x8e, 0xa4, 0x5a, 0xe2, 0xfd, 0x63, 0x81, 0x53, 0x9c, 0x03, 0x6f, 0xc0,
	0x74, 0x9f, 0x97, 0x37, 0x00, 0xd7, 0xec, 0x13, 0x00, 0x2e, 0xf1, 0xa2, 0x0c, 0x13, 0xb1, 0xa6,
	0xa4, 0x1c, 0x5e, 0x61, 0xd8, 0x18, 0x5a, 0x57, 0xf9, 0x92, 0x32, 0xd6, 0x07, 0x7f, 0xf9, 0x88,
	0x8c, 0x03, 0x23, 0xa1, 0x93, 0x14, 0x7a, 0xf6, 0x03, 0xb4, 0x27, 0xb9, 0x4c, 0x12, 0x72, 0xa6,
	0x73, 0xf9, 0xea, 0x31, 0xce, 0x4a, 0x11, 0xb9, 0x3b, 0xf8, 0xf0, 0x3e, 0x83, 0x4e, 0x25, 0x10,
	0x7b, 0x06, 0x4d, 0x5c, 0x9b, 0x6e, 0x35, 0xb8, 0x41, 0xde, 0x02, 0xce, 0x8e, 0x5c, 0xdc, 0x67,
	0x88, 0xe3, 0xf8, 0x26, 0xbe, 0x96, 0xd3, 0x78, 0x23, 0xcd, 0xe5, 0x29, 0x31, 0xee, 0x0d, 0x93,
	0xfd, 0x24, 0x97, 0x52, 0xdf, 0x9e, 0x06, 0x2f, 0xb1, 0xa7, 0xc0, 0x29, 0x5a, 0xf6, 0xa8, 0x07,
	0xa1, 0x3a, 0xf6, 0xd6, 0xf1, 0xd8, 0xb3, 0x0b, 0x78, 0x3a, 0xcc, 0xb6, 0x4b, 0xf9, 0x63, 0xbe,
	0x12, 0x4a, 0x37, 0xc5, 0xa6, 0xa6, 0xfc, 0x97, 0xf6, 0xfe, 0xb0, 0xa1, 0x5d, 0xf6, 0x1b, 0xc7,
	0x74, 0x14, 0x99, 0x7c, 0xac, 0x51, 0x84, 0x38, 0x8c, 0x4c, 0x16, 0x56, 0x18, 0xb1, 0x0f, 0xa1,
	0x11, 0x8a, 0x5f, 0x47, 0x91, 0x39, 0xbc, 0x06, 0x86, 0x0d, 0x23, 0x6a, 0x87, 0x66, 0xc3, 0x88,
	0x7d, 0x04, 0xed, 0xe8, 0xe7, 0xfd, 0x40, 0x6c, 0xc4, 0x5a, 0xba, 0x0d, 0xda, 0x39, 0x10, 0x38,
	0x31, 0x08, 0xe4, 0xb5, 0x4c, 0x97, 0xd2, 0x6d, 0xd2, 0x76, 0x85, 0x41, 0xf5, 0x24, 0x4f, 0x8c,
	0xba, 0xa5, 0xd5, 0x25, 0x81, 0x6a, 0x04, 0x46, 0xed, 0x68, 0xf5, 0x81, 0xc1, 0x19, 0x1d, 0x49,
	0x91, 0xb8, 0x6d, 0x3d, 0xa3, 0xb8, 0x66, 0x2f, 0xe0, 0xac, 0xbf, 0x8d, 0x55, 0xbc, 0x14, 0x09,
	0xf6, 0x62, 0xe7, 0x02, 0x6d, 0x1e, 0x93, 0x58, 0xd5, 0x82, 0x70, 0x3b, 0xba, 0x43, 0x05, 0x66,
	0x5f, 0xc0, 0x07, 0xdf, 0xa5, 0x2a, 0x3e, 0xf6, 0xd2, 0x25, 0xa3, 0xbb, 0x1b, 0xcc, 0x87, 0x6e,
	0x95, 0x74, 0xcf, 0xc8, 0xf0, 0x88, 0xc3, 0xca, 0x0d, 0xb2, 0xd5, 0x5a, 0xba, 0x4f, 0x74, 0xe5,
	0x08, 0xe0, 0x53, 0x3f, 0x8a, 0x95, 0xfb, 0x94, 0x38, 0x5c, 0xfa, 0x3f, 0xc1, 0x39, 0x3e, 0xe7,
	0xf4, 0xb0, 0xf6, 0xb3, 0x7c, 0x7f, 0xda, 0xe3, 0x7a, 0x98, 0x58, 0xab, 0x3a, 0xb1, 0xdf, 0x77,
	0xe7, 0x10, 0xbc, 0xfc, 0xd6, 0x28, 0xdf, 0x35, 0xe9, 0x8b, 0x7c, 0xf9, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf9, 0x82, 0x92, 0x27, 0xb3, 0x07, 0x00, 0x00,
}
