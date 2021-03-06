// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/diademnetwork/go-diadem/examples/plugins/evmproxy/types/types.proto

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

type EthTransaction struct {
	Version              uint32   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EthTransaction) Reset()         { *m = EthTransaction{} }
func (m *EthTransaction) String() string { return proto.CompactTextString(m) }
func (*EthTransaction) ProtoMessage()    {}
func (*EthTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_d7c6a2e89cbc3ccf, []int{0}
}
func (m *EthTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EthTransaction.Unmarshal(m, b)
}
func (m *EthTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EthTransaction.Marshal(b, m, deterministic)
}
func (dst *EthTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthTransaction.Merge(dst, src)
}
func (m *EthTransaction) XXX_Size() int {
	return xxx_messageInfo_EthTransaction.Size(m)
}
func (m *EthTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_EthTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_EthTransaction proto.InternalMessageInfo

func (m *EthTransaction) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *EthTransaction) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type EthCall struct {
	Version              uint32   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EthCall) Reset()         { *m = EthCall{} }
func (m *EthCall) String() string { return proto.CompactTextString(m) }
func (*EthCall) ProtoMessage()    {}
func (*EthCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_d7c6a2e89cbc3ccf, []int{1}
}
func (m *EthCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EthCall.Unmarshal(m, b)
}
func (m *EthCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EthCall.Marshal(b, m, deterministic)
}
func (dst *EthCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthCall.Merge(dst, src)
}
func (m *EthCall) XXX_Size() int {
	return xxx_messageInfo_EthCall.Size(m)
}
func (m *EthCall) XXX_DiscardUnknown() {
	xxx_messageInfo_EthCall.DiscardUnknown(m)
}

var xxx_messageInfo_EthCall proto.InternalMessageInfo

func (m *EthCall) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *EthCall) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type EthCallResult struct {
	Version              uint32   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EthCallResult) Reset()         { *m = EthCallResult{} }
func (m *EthCallResult) String() string { return proto.CompactTextString(m) }
func (*EthCallResult) ProtoMessage()    {}
func (*EthCallResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_d7c6a2e89cbc3ccf, []int{2}
}
func (m *EthCallResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EthCallResult.Unmarshal(m, b)
}
func (m *EthCallResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EthCallResult.Marshal(b, m, deterministic)
}
func (dst *EthCallResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthCallResult.Merge(dst, src)
}
func (m *EthCallResult) XXX_Size() int {
	return xxx_messageInfo_EthCallResult.Size(m)
}
func (m *EthCallResult) XXX_DiscardUnknown() {
	xxx_messageInfo_EthCallResult.DiscardUnknown(m)
}

var xxx_messageInfo_EthCallResult proto.InternalMessageInfo

func (m *EthCallResult) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *EthCallResult) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*EthTransaction)(nil), "EthTransaction")
	proto.RegisterType((*EthCall)(nil), "EthCall")
	proto.RegisterType((*EthCallResult)(nil), "EthCallResult")
}

func init() {
	proto.RegisterFile("github.com/diademnetwork/go-diadem/examples/plugins/evmproxy/types/types.proto", fileDescriptor_types_d7c6a2e89cbc3ccf)
}

var fileDescriptor_types_d7c6a2e89cbc3ccf = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xc9, 0xcf, 0xcf, 0xcd, 0x4b, 0x2d, 0x29, 0xcf,
	0x2f, 0xca, 0xd6, 0x4f, 0xcf, 0xd7, 0x05, 0x71, 0xf5, 0x53, 0x2b, 0x12, 0x73, 0x0b, 0x72, 0x52,
	0x8b, 0xf5, 0x0b, 0x72, 0x4a, 0xd3, 0x33, 0xf3, 0x8a, 0xf5, 0x53, 0xcb, 0x72, 0x0b, 0x8a, 0xf2,
	0x2b, 0x2a, 0xf5, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x21, 0xa4, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe,
	0x92, 0x1d, 0x17, 0x9f, 0x6b, 0x49, 0x46, 0x48, 0x51, 0x62, 0x5e, 0x71, 0x62, 0x72, 0x49, 0x66,
	0x7e, 0x9e, 0x90, 0x04, 0x17, 0x7b, 0x59, 0x6a, 0x51, 0x71, 0x66, 0x7e, 0x9e, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0x6f, 0x10, 0x8c, 0x2b, 0x24, 0xc4, 0xc5, 0x92, 0x92, 0x58, 0x92, 0x28, 0xc1, 0xa4,
	0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x2b, 0x99, 0x73, 0xb1, 0xbb, 0x96, 0x64, 0x38, 0x27, 0xe6,
	0xe4, 0x90, 0xa8, 0xd1, 0x96, 0x8b, 0x17, 0xaa, 0x31, 0x28, 0xb5, 0xb8, 0x34, 0xa7, 0x84, 0x34,
	0xed, 0x49, 0x6c, 0x60, 0xe7, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xc5, 0xf8, 0xdf,
	0x0c, 0x01, 0x00, 0x00,
}
