// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/proto/v0/accounts.proto

package proto

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

type Record struct {
	Key                  string    `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Payload              *Settings `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c84319968a576b, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Record) GetPayload() *Settings {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Query struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c84319968a576b, []int{1}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Settings struct {
	Theme                string   `protobuf:"bytes,2,opt,name=theme,proto3" json:"theme,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c84319968a576b, []int{2}
}

func (m *Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings.Unmarshal(m, b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
}
func (m *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(m, src)
}
func (m *Settings) XXX_Size() int {
	return xxx_messageInfo_Settings.Size(m)
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

func (m *Settings) GetTheme() string {
	if m != nil {
		return m.Theme
	}
	return ""
}

func init() {
	proto.RegisterType((*Record)(nil), "settings.Record")
	proto.RegisterType((*Query)(nil), "settings.Query")
	proto.RegisterType((*Settings)(nil), "settings.Settings")
}

func init() { proto.RegisterFile("pkg/proto/v0/accounts.proto", fileDescriptor_e3c84319968a576b) }

var fileDescriptor_e3c84319968a576b = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0xc8, 0x4e, 0xd7,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x33, 0xd0, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b,
	0x29, 0xd6, 0x03, 0x8b, 0x08, 0x71, 0x14, 0xa7, 0x96, 0x94, 0x64, 0xe6, 0xa5, 0x17, 0x2b, 0x79,
	0x70, 0xb1, 0x05, 0xa5, 0x26, 0xe7, 0x17, 0xa5, 0x08, 0x09, 0x70, 0x31, 0x67, 0xa7, 0x56, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42, 0x3a, 0x5c, 0xec, 0x05, 0x89, 0x95, 0x39,
	0xf9, 0x89, 0x29, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x42, 0x7a, 0x30, 0x7d, 0x7a, 0xc1,
	0x50, 0x46, 0x10, 0x4c, 0x89, 0x92, 0x38, 0x17, 0x6b, 0x60, 0x69, 0x6a, 0x51, 0xa5, 0x10, 0x1f,
	0x17, 0x53, 0x66, 0x0a, 0xd4, 0x1c, 0xa6, 0xcc, 0x14, 0x25, 0x05, 0x2e, 0x0e, 0x98, 0x6a, 0x21,
	0x11, 0x2e, 0xd6, 0x92, 0x8c, 0xd4, 0xdc, 0x54, 0xb0, 0x81, 0x9c, 0x41, 0x10, 0x8e, 0x51, 0x1a,
	0x17, 0x3f, 0x4c, 0x45, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x26, 0x17, 0x73, 0x70,
	0x6a, 0x89, 0x90, 0x00, 0xc2, 0x46, 0x88, 0x33, 0xa5, 0x30, 0x44, 0x84, 0x34, 0xb8, 0x98, 0xdd,
	0x53, 0x4b, 0x84, 0xf8, 0x11, 0x12, 0x60, 0x77, 0x60, 0xaa, 0x74, 0x62, 0x8f, 0x62, 0x05, 0xfb,
	0x3f, 0x89, 0x0d, 0x4c, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x28, 0x61, 0x84, 0x13, 0x25,
	0x01, 0x00, 0x00,
}
