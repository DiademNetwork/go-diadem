// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/diademnetwork/go-diadem/examples/types/types.proto

package types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MapEntry struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MapEntry) Reset()         { *m = MapEntry{} }
func (m *MapEntry) String() string { return proto.CompactTextString(m) }
func (*MapEntry) ProtoMessage()    {}
func (*MapEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_9f4cdff418172726, []int{0}
}
func (m *MapEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapEntry.Unmarshal(m, b)
}
func (m *MapEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapEntry.Marshal(b, m, deterministic)
}
func (dst *MapEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapEntry.Merge(dst, src)
}
func (m *MapEntry) XXX_Size() int {
	return xxx_messageInfo_MapEntry.Size(m)
}
func (m *MapEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_MapEntry.DiscardUnknown(m)
}

var xxx_messageInfo_MapEntry proto.InternalMessageInfo

func (m *MapEntry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MapEntry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type HelloRequest struct {
	In                   string   `protobuf:"bytes,1,opt,name=in,proto3" json:"in,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_9f4cdff418172726, []int{1}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (dst *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(dst, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetIn() string {
	if m != nil {
		return m.In
	}
	return ""
}

type HelloResponse struct {
	Out                  string   `protobuf:"bytes,1,opt,name=out,proto3" json:"out,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_9f4cdff418172726, []int{2}
}
func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (dst *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(dst, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetOut() string {
	if m != nil {
		return m.Out
	}
	return ""
}

func init() {
	proto.RegisterType((*MapEntry)(nil), "MapEntry")
	proto.RegisterType((*HelloRequest)(nil), "HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "HelloResponse")
}

func init() {
	proto.RegisterFile("github.com/diademnetwork/go-diadem/examples/types/types.proto", fileDescriptor_types_9f4cdff418172726)
}

var fileDescriptor_types_9f4cdff418172726 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xc9, 0xcf, 0xcf, 0xcd, 0x4b, 0x2d, 0x29, 0xcf,
	0x2f, 0xca, 0xd6, 0x4f, 0xcf, 0xd7, 0x05, 0x71, 0xf5, 0x53, 0x2b, 0x12, 0x73, 0x0b, 0x72, 0x52,
	0x8b, 0xf5, 0x4b, 0x2a, 0x0b, 0x60, 0xa4, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x92, 0x11, 0x17,
	0x87, 0x6f, 0x62, 0x81, 0x6b, 0x5e, 0x49, 0x51, 0xa5, 0x90, 0x00, 0x17, 0x73, 0x76, 0x6a, 0xa5,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53,
	0x9a, 0x2a, 0xc1, 0x04, 0x16, 0x83, 0x70, 0x94, 0xe4, 0xb8, 0x78, 0x3c, 0x52, 0x73, 0x72, 0xf2,
	0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0xf3, 0xa0, 0xda, 0x98,
	0x32, 0xf3, 0x94, 0x14, 0xb9, 0x78, 0xa1, 0xf2, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x20, 0x83,
	0xf3, 0x4b, 0x4b, 0x60, 0x06, 0xe7, 0x97, 0x96, 0x24, 0xb1, 0x81, 0x6d, 0x37, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xe8, 0xfb, 0x64, 0x17, 0xba, 0x00, 0x00, 0x00,
}