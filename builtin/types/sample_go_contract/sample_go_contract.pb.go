// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/diademnetwork/go-diadem/builtin/types/sample_go_contract/sample_go_contract.proto

package sample_go_contract

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/diademnetwork/go-diadem/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SampleGoContractInitRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SampleGoContractInitRequest) Reset()         { *m = SampleGoContractInitRequest{} }
func (m *SampleGoContractInitRequest) String() string { return proto.CompactTextString(m) }
func (*SampleGoContractInitRequest) ProtoMessage()    {}
func (*SampleGoContractInitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sample_go_contract_c611cfe49aefe1d2, []int{0}
}
func (m *SampleGoContractInitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SampleGoContractInitRequest.Unmarshal(m, b)
}
func (m *SampleGoContractInitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SampleGoContractInitRequest.Marshal(b, m, deterministic)
}
func (dst *SampleGoContractInitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampleGoContractInitRequest.Merge(dst, src)
}
func (m *SampleGoContractInitRequest) XXX_Size() int {
	return xxx_messageInfo_SampleGoContractInitRequest.Size(m)
}
func (m *SampleGoContractInitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SampleGoContractInitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SampleGoContractInitRequest proto.InternalMessageInfo

type SampleGoContractNestedEvmRequest struct {
	InnerEmitter         *types.Address `protobuf:"bytes,1,opt,name=inner_emitter,json=innerEmitter" json:"inner_emitter,omitempty"`
	OuterEmitter         *types.Address `protobuf:"bytes,2,opt,name=outer_emitter,json=outerEmitter" json:"outer_emitter,omitempty"`
	InnerEmitterValue    uint64         `protobuf:"varint,3,opt,name=inner_emitter_value,json=innerEmitterValue,proto3" json:"inner_emitter_value,omitempty"`
	OuterEmitterValue    uint64         `protobuf:"varint,4,opt,name=outer_emitter_value,json=outerEmitterValue,proto3" json:"outer_emitter_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SampleGoContractNestedEvmRequest) Reset()         { *m = SampleGoContractNestedEvmRequest{} }
func (m *SampleGoContractNestedEvmRequest) String() string { return proto.CompactTextString(m) }
func (*SampleGoContractNestedEvmRequest) ProtoMessage()    {}
func (*SampleGoContractNestedEvmRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sample_go_contract_c611cfe49aefe1d2, []int{1}
}
func (m *SampleGoContractNestedEvmRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SampleGoContractNestedEvmRequest.Unmarshal(m, b)
}
func (m *SampleGoContractNestedEvmRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SampleGoContractNestedEvmRequest.Marshal(b, m, deterministic)
}
func (dst *SampleGoContractNestedEvmRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampleGoContractNestedEvmRequest.Merge(dst, src)
}
func (m *SampleGoContractNestedEvmRequest) XXX_Size() int {
	return xxx_messageInfo_SampleGoContractNestedEvmRequest.Size(m)
}
func (m *SampleGoContractNestedEvmRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SampleGoContractNestedEvmRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SampleGoContractNestedEvmRequest proto.InternalMessageInfo

func (m *SampleGoContractNestedEvmRequest) GetInnerEmitter() *types.Address {
	if m != nil {
		return m.InnerEmitter
	}
	return nil
}

func (m *SampleGoContractNestedEvmRequest) GetOuterEmitter() *types.Address {
	if m != nil {
		return m.OuterEmitter
	}
	return nil
}

func (m *SampleGoContractNestedEvmRequest) GetInnerEmitterValue() uint64 {
	if m != nil {
		return m.InnerEmitterValue
	}
	return 0
}

func (m *SampleGoContractNestedEvmRequest) GetOuterEmitterValue() uint64 {
	if m != nil {
		return m.OuterEmitterValue
	}
	return 0
}

func init() {
	proto.RegisterType((*SampleGoContractInitRequest)(nil), "SampleGoContractInitRequest")
	proto.RegisterType((*SampleGoContractNestedEvmRequest)(nil), "SampleGoContractNestedEvmRequest")
}

func init() {
	proto.RegisterFile("github.com/diademnetwork/go-diadem/builtin/types/sample_go_contract/sample_go_contract.proto", fileDescriptor_sample_go_contract_c611cfe49aefe1d2)
}

var fileDescriptor_sample_go_contract_c611cfe49aefe1d2 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xbd, 0x4a, 0x04, 0x31,
	0x10, 0x80, 0x59, 0x3d, 0x44, 0xa2, 0x16, 0x6a, 0x73, 0x28, 0xc2, 0x72, 0xd5, 0x35, 0x97, 0x88,
	0x3e, 0x81, 0xc8, 0x21, 0x36, 0x16, 0x2b, 0x88, 0x5d, 0xd8, 0x9f, 0x61, 0x0d, 0x26, 0x99, 0x35,
	0x99, 0x9c, 0xf8, 0xb6, 0x3e, 0x8a, 0x24, 0x59, 0x61, 0xef, 0x10, 0x6c, 0x02, 0x99, 0xef, 0x9b,
	0xaf, 0x18, 0xf6, 0xda, 0x2b, 0x7a, 0x0b, 0x0d, 0x6f, 0xd1, 0x08, 0x8d, 0x68, 0x2c, 0xd0, 0x27,
	0xba, 0x77, 0xd1, 0xe3, 0x2a, 0x7e, 0x45, 0x13, 0x94, 0x26, 0x65, 0x05, 0x7d, 0x0d, 0xe0, 0x85,
	0xaf, 0xcd, 0xa0, 0x41, 0xf6, 0x28, 0x5b, 0xb4, 0xe4, 0xea, 0x96, 0xfe, 0x18, 0xf1, 0xc1, 0x21,
	0xe1, 0xc5, 0xf5, 0x3f, 0xe5, 0x5c, 0x4c, 0x6f, 0xde, 0x58, 0x5c, 0xb1, 0xcb, 0xe7, 0x54, 0x7b,
	0xc0, 0xfb, 0xb1, 0xf5, 0x68, 0x15, 0x55, 0xf0, 0x11, 0xc0, 0xd3, 0xe2, 0xbb, 0x60, 0xe5, 0x2e,
	0x7f, 0x02, 0x4f, 0xd0, 0xad, 0x37, 0x66, 0x94, 0xce, 0x56, 0xec, 0x44, 0x59, 0x0b, 0x4e, 0x82,
	0x51, 0x44, 0xe0, 0xe6, 0x45, 0x59, 0x2c, 0x8f, 0x6e, 0x0e, 0xf9, 0x5d, 0xd7, 0x39, 0xf0, 0xbe,
	0x3a, 0x4e, 0x78, 0x9d, 0x69, 0xd4, 0x31, 0xd0, 0x44, 0xdf, 0xdb, 0xd5, 0x13, 0xfe, 0xd5, 0x39,
	0x3b, 0xdf, 0xaa, 0xcb, 0x4d, 0xad, 0x03, 0xcc, 0xf7, 0xcb, 0x62, 0x39, 0xab, 0x4e, 0xa7, 0xe5,
	0x97, 0x08, 0xa2, 0xbf, 0x95, 0x1f, 0xfd, 0x59, 0xf6, 0xa7, 0xe9, 0xe4, 0x37, 0x07, 0xe9, 0x10,
	0xb7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xe4, 0xa0, 0xd1, 0x96, 0x01, 0x00, 0x00,
}
