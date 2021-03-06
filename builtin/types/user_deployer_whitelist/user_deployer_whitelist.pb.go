// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/diademnetwork/go-diadem/builtin/types/user_deployer_whitelist/user_deployer_whitelist.proto

package user_deployer_whitelist

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
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

type TierID int32

const (
	TierID_DEFAULT TierID = 0
)

var TierID_name = map[int32]string{
	0: "DEFAULT",
}
var TierID_value = map[string]int32{
	"DEFAULT": 0,
}

func (x TierID) String() string {
	return proto.EnumName(TierID_name, int32(x))
}
func (TierID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93, []int{0}
}

type Tier struct {
	TierID               TierID         `protobuf:"varint,1,opt,name=id,proto3,enum=user_deployer_whitelist.TierID" json:"id,omitempty"`
	Fee                  *types.BigUInt `protobuf:"bytes,2,opt,name=fee" json:"fee,omitempty"`
	Name                 string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	BlockRange           uint64         `protobuf:"varint,4,opt,name=block_range,json=blockRange,proto3" json:"block_range,omitempty"`
	MaxTxs               uint64         `protobuf:"varint,5,opt,name=max_txs,json=maxTxs,proto3" json:"max_txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Tier) Reset()         { *m = Tier{} }
func (m *Tier) String() string { return proto.CompactTextString(m) }
func (*Tier) ProtoMessage()    {}
func (*Tier) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93, []int{0}
}
func (m *Tier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tier.Unmarshal(m, b)
}
func (m *Tier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tier.Marshal(b, m, deterministic)
}
func (dst *Tier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tier.Merge(dst, src)
}
func (m *Tier) XXX_Size() int {
	return xxx_messageInfo_Tier.Size(m)
}
func (m *Tier) XXX_DiscardUnknown() {
	xxx_messageInfo_Tier.DiscardUnknown(m)
}

var xxx_messageInfo_Tier proto.InternalMessageInfo

func (m *Tier) GetTierID() TierID {
	if m != nil {
		return m.TierID
	}
	return TierID_DEFAULT
}

func (m *Tier) GetFee() *types.BigUInt {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *Tier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tier) GetBlockRange() uint64 {
	if m != nil {
		return m.BlockRange
	}
	return 0
}

func (m *Tier) GetMaxTxs() uint64 {
	if m != nil {
		return m.MaxTxs
	}
	return 0
}

type TierInfo struct {
	TierID               TierID   `protobuf:"varint,1,opt,name=id,proto3,enum=user_deployer_whitelist.TierID" json:"id,omitempty"`
	Fee                  uint64   `protobuf:"varint,2,opt,name=fee,proto3" json:"fee,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	BlockRange           uint64   `protobuf:"varint,4,opt,name=block_range,json=blockRange,proto3" json:"block_range,omitempty"`
	MaxTxs               uint64   `protobuf:"varint,5,opt,name=max_txs,json=maxTxs,proto3" json:"max_txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TierInfo) Reset()         { *m = TierInfo{} }
func (m *TierInfo) String() string { return proto.CompactTextString(m) }
func (*TierInfo) ProtoMessage()    {}
func (*TierInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93, []int{1}
}
func (m *TierInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TierInfo.Unmarshal(m, b)
}
func (m *TierInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TierInfo.Marshal(b, m, deterministic)
}
func (dst *TierInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TierInfo.Merge(dst, src)
}
func (m *TierInfo) XXX_Size() int {
	return xxx_messageInfo_TierInfo.Size(m)
}
func (m *TierInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TierInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TierInfo proto.InternalMessageInfo

func (m *TierInfo) GetTierID() TierID {
	if m != nil {
		return m.TierID
	}
	return TierID_DEFAULT
}

func (m *TierInfo) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *TierInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TierInfo) GetBlockRange() uint64 {
	if m != nil {
		return m.BlockRange
	}
	return 0
}

func (m *TierInfo) GetMaxTxs() uint64 {
	if m != nil {
		return m.MaxTxs
	}
	return 0
}

type InitRequest struct {
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	TierInfo             []*TierInfo    `protobuf:"bytes,2,rep,name=tier_info,json=tierInfo" json:"tier_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *InitRequest) Reset()         { *m = InitRequest{} }
func (m *InitRequest) String() string { return proto.CompactTextString(m) }
func (*InitRequest) ProtoMessage()    {}
func (*InitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93, []int{2}
}
func (m *InitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitRequest.Unmarshal(m, b)
}
func (m *InitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitRequest.Marshal(b, m, deterministic)
}
func (dst *InitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitRequest.Merge(dst, src)
}
func (m *InitRequest) XXX_Size() int {
	return xxx_messageInfo_InitRequest.Size(m)
}
func (m *InitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitRequest proto.InternalMessageInfo

func (m *InitRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *InitRequest) GetTierInfo() []*TierInfo {
	if m != nil {
		return m.TierInfo
	}
	return nil
}

type DeployerContract struct {
	ContractAddress      *types.Address `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress" json:"contract_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeployerContract) Reset()         { *m = DeployerContract{} }
func (m *DeployerContract) String() string { return proto.CompactTextString(m) }
func (*DeployerContract) ProtoMessage()    {}
func (*DeployerContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93, []int{3}
}
func (m *DeployerContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployerContract.Unmarshal(m, b)
}
func (m *DeployerContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployerContract.Marshal(b, m, deterministic)
}
func (dst *DeployerContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployerContract.Merge(dst, src)
}
func (m *DeployerContract) XXX_Size() int {
	return xxx_messageInfo_DeployerContract.Size(m)
}
func (m *DeployerContract) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployerContract.DiscardUnknown(m)
}

var xxx_messageInfo_DeployerContract proto.InternalMessageInfo

func (m *DeployerContract) GetContractAddress() *types.Address {
	if m != nil {
		return m.ContractAddress
	}
	return nil
}

type WhitelistUserDeployerRequest struct {
	DeployerAddr         *types.Address `protobuf:"bytes,1,opt,name=deployer_addr,json=deployerAddr" json:"deployer_addr,omitempty"`
	TierID               TierID         `protobuf:"varint,2,opt,name=tier_id,json=tierId,proto3,enum=user_deployer_whitelist.TierID" json:"tier_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *WhitelistUserDeployerRequest) Reset()         { *m = WhitelistUserDeployerRequest{} }
func (m *WhitelistUserDeployerRequest) String() string { return proto.CompactTextString(m) }
func (*WhitelistUserDeployerRequest) ProtoMessage()    {}
func (*WhitelistUserDeployerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93, []int{4}
}
func (m *WhitelistUserDeployerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhitelistUserDeployerRequest.Unmarshal(m, b)
}
func (m *WhitelistUserDeployerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhitelistUserDeployerRequest.Marshal(b, m, deterministic)
}
func (dst *WhitelistUserDeployerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhitelistUserDeployerRequest.Merge(dst, src)
}
func (m *WhitelistUserDeployerRequest) XXX_Size() int {
	return xxx_messageInfo_WhitelistUserDeployerRequest.Size(m)
}
func (m *WhitelistUserDeployerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WhitelistUserDeployerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WhitelistUserDeployerRequest proto.InternalMessageInfo

func (m *WhitelistUserDeployerRequest) GetDeployerAddr() *types.Address {
	if m != nil {
		return m.DeployerAddr
	}
	return nil
}

func (m *WhitelistUserDeployerRequest) GetTierID() TierID {
	if m != nil {
		return m.TierID
	}
	return TierID_DEFAULT
}

type RemoveUserDeployerRequest struct {
	DeployerAddr         *types.Address `protobuf:"bytes,1,opt,name=deployer_addr,json=deployerAddr" json:"deployer_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RemoveUserDeployerRequest) Reset()         { *m = RemoveUserDeployerRequest{} }
func (m *RemoveUserDeployerRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveUserDeployerRequest) ProtoMessage()    {}
func (*RemoveUserDeployerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93, []int{5}
}
func (m *RemoveUserDeployerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveUserDeployerRequest.Unmarshal(m, b)
}
func (m *RemoveUserDeployerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveUserDeployerRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveUserDeployerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveUserDeployerRequest.Merge(dst, src)
}
func (m *RemoveUserDeployerRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveUserDeployerRequest.Size(m)
}
func (m *RemoveUserDeployerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveUserDeployerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveUserDeployerRequest proto.InternalMessageInfo

func (m *RemoveUserDeployerRequest) GetDeployerAddr() *types.Address {
	if m != nil {
		return m.DeployerAddr
	}
	return nil
}

type UserState struct {
	UserAddr             *types.Address   `protobuf:"bytes,1,opt,name=user_addr,json=userAddr" json:"user_addr,omitempty"`
	Deployers            []*types.Address `protobuf:"bytes,2,rep,name=deployers" json:"deployers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UserState) Reset()         { *m = UserState{} }
func (m *UserState) String() string { return proto.CompactTextString(m) }
func (*UserState) ProtoMessage()    {}
func (*UserState) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93, []int{6}
}
func (m *UserState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserState.Unmarshal(m, b)
}
func (m *UserState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserState.Marshal(b, m, deterministic)
}
func (dst *UserState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserState.Merge(dst, src)
}
func (m *UserState) XXX_Size() int {
	return xxx_messageInfo_UserState.Size(m)
}
func (m *UserState) XXX_DiscardUnknown() {
	xxx_messageInfo_UserState.DiscardUnknown(m)
}

var xxx_messageInfo_UserState proto.InternalMessageInfo

func (m *UserState) GetUserAddr() *types.Address {
	if m != nil {
		return m.UserAddr
	}
	return nil
}

func (m *UserState) GetDeployers() []*types.Address {
	if m != nil {
		return m.Deployers
	}
	return nil
}

type UserDeployerState struct {
	Address              *types.Address      `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Contracts            []*DeployerContract `protobuf:"bytes,2,rep,name=contracts" json:"contracts,omitempty"`
	TierID               TierID              `protobuf:"varint,3,opt,name=tier_id,json=tierId,proto3,enum=user_deployer_whitelist.TierID" json:"tier_id,omitempty"`
	Inactive             bool                `protobuf:"varint,4,opt,name=inactive,proto3" json:"inactive,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserDeployerState) Reset()         { *m = UserDeployerState{} }
func (m *UserDeployerState) String() string { return proto.CompactTextString(m) }
func (*UserDeployerState) ProtoMessage()    {}
func (*UserDeployerState) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93, []int{7}
}
func (m *UserDeployerState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDeployerState.Unmarshal(m, b)
}
func (m *UserDeployerState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDeployerState.Marshal(b, m, deterministic)
}
func (dst *UserDeployerState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDeployerState.Merge(dst, src)
}
func (m *UserDeployerState) XXX_Size() int {
	return xxx_messageInfo_UserDeployerState.Size(m)
}
func (m *UserDeployerState) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDeployerState.DiscardUnknown(m)
}

var xxx_messageInfo_UserDeployerState proto.InternalMessageInfo

func (m *UserDeployerState) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *UserDeployerState) GetContracts() []*DeployerContract {
	if m != nil {
		return m.Contracts
	}
	return nil
}

func (m *UserDeployerState) GetTierID() TierID {
	if m != nil {
		return m.TierID
	}
	return TierID_DEFAULT
}

func (m *UserDeployerState) GetInactive() bool {
	if m != nil {
		return m.Inactive
	}
	return false
}

type GetUserDeployersRequest struct {
	UserAddr             *types.Address `protobuf:"bytes,1,opt,name=user_addr,json=userAddr" json:"user_addr,omitempty"`
	IncludeInactive      bool           `protobuf:"varint,2,opt,name=include_inactive,json=includeInactive,proto3" json:"include_inactive,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetUserDeployersRequest) Reset()         { *m = GetUserDeployersRequest{} }
func (m *GetUserDeployersRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserDeployersRequest) ProtoMessage()    {}
func (*GetUserDeployersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93, []int{8}
}
func (m *GetUserDeployersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserDeployersRequest.Unmarshal(m, b)
}
func (m *GetUserDeployersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserDeployersRequest.Marshal(b, m, deterministic)
}
func (dst *GetUserDeployersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserDeployersRequest.Merge(dst, src)
}
func (m *GetUserDeployersRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserDeployersRequest.Size(m)
}
func (m *GetUserDeployersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserDeployersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserDeployersRequest proto.InternalMessageInfo

func (m *GetUserDeployersRequest) GetUserAddr() *types.Address {
	if m != nil {
		return m.UserAddr
	}
	return nil
}

func (m *GetUserDeployersRequest) GetIncludeInactive() bool {
	if m != nil {
		return m.IncludeInactive
	}
	return false
}

type GetUserDeployersResponse struct {
	Deployers            []*UserDeployerState `protobuf:"bytes,1,rep,name=deployers" json:"deployers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetUserDeployersResponse) Reset()         { *m = GetUserDeployersResponse{} }
func (m *GetUserDeployersResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserDeployersResponse) ProtoMessage()    {}
func (*GetUserDeployersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93, []int{9}
}
func (m *GetUserDeployersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserDeployersResponse.Unmarshal(m, b)
}
func (m *GetUserDeployersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserDeployersResponse.Marshal(b, m, deterministic)
}
func (dst *GetUserDeployersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserDeployersResponse.Merge(dst, src)
}
func (m *GetUserDeployersResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserDeployersResponse.Size(m)
}
func (m *GetUserDeployersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserDeployersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserDeployersResponse proto.InternalMessageInfo

func (m *GetUserDeployersResponse) GetDeployers() []*UserDeployerState {
	if m != nil {
		return m.Deployers
	}
	return nil
}

type GetDeployedContractsRequest struct {
	DeployerAddr         *types.Address `protobuf:"bytes,1,opt,name=deployer_addr,json=deployerAddr" json:"deployer_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetDeployedContractsRequest) Reset()         { *m = GetDeployedContractsRequest{} }
func (m *GetDeployedContractsRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeployedContractsRequest) ProtoMessage()    {}
func (*GetDeployedContractsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93, []int{10}
}
func (m *GetDeployedContractsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeployedContractsRequest.Unmarshal(m, b)
}
func (m *GetDeployedContractsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeployedContractsRequest.Marshal(b, m, deterministic)
}
func (dst *GetDeployedContractsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeployedContractsRequest.Merge(dst, src)
}
func (m *GetDeployedContractsRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeployedContractsRequest.Size(m)
}
func (m *GetDeployedContractsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeployedContractsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeployedContractsRequest proto.InternalMessageInfo

func (m *GetDeployedContractsRequest) GetDeployerAddr() *types.Address {
	if m != nil {
		return m.DeployerAddr
	}
	return nil
}

type GetDeployedContractsResponse struct {
	ContractAddresses    []*DeployerContract `protobuf:"bytes,1,rep,name=contract_addresses,json=contractAddresses" json:"contract_addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetDeployedContractsResponse) Reset()         { *m = GetDeployedContractsResponse{} }
func (m *GetDeployedContractsResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeployedContractsResponse) ProtoMessage()    {}
func (*GetDeployedContractsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93, []int{11}
}
func (m *GetDeployedContractsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeployedContractsResponse.Unmarshal(m, b)
}
func (m *GetDeployedContractsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeployedContractsResponse.Marshal(b, m, deterministic)
}
func (dst *GetDeployedContractsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeployedContractsResponse.Merge(dst, src)
}
func (m *GetDeployedContractsResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeployedContractsResponse.Size(m)
}
func (m *GetDeployedContractsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeployedContractsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeployedContractsResponse proto.InternalMessageInfo

func (m *GetDeployedContractsResponse) GetContractAddresses() []*DeployerContract {
	if m != nil {
		return m.ContractAddresses
	}
	return nil
}

type GetTierInfoRequest struct {
	TierID               TierID   `protobuf:"varint,1,opt,name=id,proto3,enum=user_deployer_whitelist.TierID" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTierInfoRequest) Reset()         { *m = GetTierInfoRequest{} }
func (m *GetTierInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetTierInfoRequest) ProtoMessage()    {}
func (*GetTierInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93, []int{12}
}
func (m *GetTierInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTierInfoRequest.Unmarshal(m, b)
}
func (m *GetTierInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTierInfoRequest.Marshal(b, m, deterministic)
}
func (dst *GetTierInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTierInfoRequest.Merge(dst, src)
}
func (m *GetTierInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetTierInfoRequest.Size(m)
}
func (m *GetTierInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTierInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTierInfoRequest proto.InternalMessageInfo

func (m *GetTierInfoRequest) GetTierID() TierID {
	if m != nil {
		return m.TierID
	}
	return TierID_DEFAULT
}

type GetTierInfoResponse struct {
	Tier                 *Tier    `protobuf:"bytes,1,opt,name=tier" json:"tier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTierInfoResponse) Reset()         { *m = GetTierInfoResponse{} }
func (m *GetTierInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetTierInfoResponse) ProtoMessage()    {}
func (*GetTierInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93, []int{13}
}
func (m *GetTierInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTierInfoResponse.Unmarshal(m, b)
}
func (m *GetTierInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTierInfoResponse.Marshal(b, m, deterministic)
}
func (dst *GetTierInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTierInfoResponse.Merge(dst, src)
}
func (m *GetTierInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetTierInfoResponse.Size(m)
}
func (m *GetTierInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTierInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTierInfoResponse proto.InternalMessageInfo

func (m *GetTierInfoResponse) GetTier() *Tier {
	if m != nil {
		return m.Tier
	}
	return nil
}

type SetTierInfoRequest struct {
	TierID               TierID         `protobuf:"varint,1,opt,name=id,proto3,enum=user_deployer_whitelist.TierID" json:"id,omitempty"`
	Fee                  *types.BigUInt `protobuf:"bytes,2,opt,name=fee" json:"fee,omitempty"`
	Name                 string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	BlockRange           uint64         `protobuf:"varint,4,opt,name=block_range,json=blockRange,proto3" json:"block_range,omitempty"`
	MaxTxs               uint64         `protobuf:"varint,5,opt,name=max_txs,json=maxTxs,proto3" json:"max_txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SetTierInfoRequest) Reset()         { *m = SetTierInfoRequest{} }
func (m *SetTierInfoRequest) String() string { return proto.CompactTextString(m) }
func (*SetTierInfoRequest) ProtoMessage()    {}
func (*SetTierInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93, []int{14}
}
func (m *SetTierInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTierInfoRequest.Unmarshal(m, b)
}
func (m *SetTierInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTierInfoRequest.Marshal(b, m, deterministic)
}
func (dst *SetTierInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTierInfoRequest.Merge(dst, src)
}
func (m *SetTierInfoRequest) XXX_Size() int {
	return xxx_messageInfo_SetTierInfoRequest.Size(m)
}
func (m *SetTierInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTierInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetTierInfoRequest proto.InternalMessageInfo

func (m *SetTierInfoRequest) GetTierID() TierID {
	if m != nil {
		return m.TierID
	}
	return TierID_DEFAULT
}

func (m *SetTierInfoRequest) GetFee() *types.BigUInt {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *SetTierInfoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetTierInfoRequest) GetBlockRange() uint64 {
	if m != nil {
		return m.BlockRange
	}
	return 0
}

func (m *SetTierInfoRequest) GetMaxTxs() uint64 {
	if m != nil {
		return m.MaxTxs
	}
	return 0
}

type SwapUserDeployerRequest struct {
	OldDeployerAddr      *types.Address `protobuf:"bytes,1,opt,name=old_deployer_addr,json=oldDeployerAddr" json:"old_deployer_addr,omitempty"`
	NewDeployerAddr      *types.Address `protobuf:"bytes,2,opt,name=new_deployer_addr,json=newDeployerAddr" json:"new_deployer_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SwapUserDeployerRequest) Reset()         { *m = SwapUserDeployerRequest{} }
func (m *SwapUserDeployerRequest) String() string { return proto.CompactTextString(m) }
func (*SwapUserDeployerRequest) ProtoMessage()    {}
func (*SwapUserDeployerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93, []int{15}
}
func (m *SwapUserDeployerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwapUserDeployerRequest.Unmarshal(m, b)
}
func (m *SwapUserDeployerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwapUserDeployerRequest.Marshal(b, m, deterministic)
}
func (dst *SwapUserDeployerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwapUserDeployerRequest.Merge(dst, src)
}
func (m *SwapUserDeployerRequest) XXX_Size() int {
	return xxx_messageInfo_SwapUserDeployerRequest.Size(m)
}
func (m *SwapUserDeployerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SwapUserDeployerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SwapUserDeployerRequest proto.InternalMessageInfo

func (m *SwapUserDeployerRequest) GetOldDeployerAddr() *types.Address {
	if m != nil {
		return m.OldDeployerAddr
	}
	return nil
}

func (m *SwapUserDeployerRequest) GetNewDeployerAddr() *types.Address {
	if m != nil {
		return m.NewDeployerAddr
	}
	return nil
}

func init() {
	proto.RegisterType((*Tier)(nil), "user_deployer_whitelist.Tier")
	proto.RegisterType((*TierInfo)(nil), "user_deployer_whitelist.TierInfo")
	proto.RegisterType((*InitRequest)(nil), "user_deployer_whitelist.InitRequest")
	proto.RegisterType((*DeployerContract)(nil), "user_deployer_whitelist.DeployerContract")
	proto.RegisterType((*WhitelistUserDeployerRequest)(nil), "user_deployer_whitelist.WhitelistUserDeployerRequest")
	proto.RegisterType((*RemoveUserDeployerRequest)(nil), "user_deployer_whitelist.RemoveUserDeployerRequest")
	proto.RegisterType((*UserState)(nil), "user_deployer_whitelist.UserState")
	proto.RegisterType((*UserDeployerState)(nil), "user_deployer_whitelist.UserDeployerState")
	proto.RegisterType((*GetUserDeployersRequest)(nil), "user_deployer_whitelist.GetUserDeployersRequest")
	proto.RegisterType((*GetUserDeployersResponse)(nil), "user_deployer_whitelist.GetUserDeployersResponse")
	proto.RegisterType((*GetDeployedContractsRequest)(nil), "user_deployer_whitelist.GetDeployedContractsRequest")
	proto.RegisterType((*GetDeployedContractsResponse)(nil), "user_deployer_whitelist.GetDeployedContractsResponse")
	proto.RegisterType((*GetTierInfoRequest)(nil), "user_deployer_whitelist.GetTierInfoRequest")
	proto.RegisterType((*GetTierInfoResponse)(nil), "user_deployer_whitelist.GetTierInfoResponse")
	proto.RegisterType((*SetTierInfoRequest)(nil), "user_deployer_whitelist.SetTierInfoRequest")
	proto.RegisterType((*SwapUserDeployerRequest)(nil), "user_deployer_whitelist.SwapUserDeployerRequest")
	proto.RegisterEnum("user_deployer_whitelist.TierID", TierID_name, TierID_value)
}

func init() {
	proto.RegisterFile("github.com/diademnetwork/go-diadem/builtin/types/user_deployer_whitelist/user_deployer_whitelist.proto", fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93)
}

var fileDescriptor_user_deployer_whitelist_ea863dc08d17cf93 = []byte{
	// 717 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xdf, 0x6e, 0x12, 0x4d,
	0x14, 0xff, 0x16, 0x28, 0x7f, 0x0e, 0xdf, 0xf7, 0x15, 0xc6, 0x18, 0x10, 0xab, 0xc5, 0x4d, 0x34,
	0xb4, 0x49, 0x41, 0xdb, 0xde, 0x99, 0x98, 0xb4, 0xa2, 0x14, 0xd3, 0x1b, 0xb7, 0x6d, 0x34, 0xde,
	0x6c, 0x16, 0xf6, 0x40, 0x27, 0xdd, 0x9d, 0xc1, 0x9d, 0xa1, 0xd0, 0x07, 0xf0, 0x09, 0x7c, 0x08,
	0xaf, 0x7c, 0x02, 0xdf, 0xc5, 0x8b, 0x3e, 0x89, 0xd9, 0xd9, 0x5d, 0xd8, 0x42, 0xa9, 0xb6, 0x55,
	0x6f, 0xc8, 0xcc, 0xf9, 0xf7, 0xfb, 0x9d, 0xdf, 0x39, 0xcc, 0x42, 0xa7, 0x4f, 0xe5, 0xf1, 0xb0,
	0x53, 0xef, 0x72, 0xb7, 0xe1, 0x70, 0xee, 0x32, 0x94, 0x23, 0xee, 0x9d, 0x34, 0xfa, 0x7c, 0xc3,
	0xbf, 0x36, 0x3a, 0x43, 0xea, 0x48, 0xca, 0x1a, 0xf2, 0x6c, 0x80, 0xa2, 0x31, 0x14, 0xe8, 0x99,
	0x36, 0x0e, 0x1c, 0x7e, 0x86, 0x9e, 0x39, 0x3a, 0xa6, 0x12, 0x1d, 0x2a, 0xe4, 0x22, 0x7b, 0x7d,
	0xe0, 0x71, 0xc9, 0x49, 0x69, 0x81, 0xbb, 0xb2, 0x11, 0x03, 0xef, 0xf3, 0x3e, 0x6f, 0xa8, 0xf8,
	0xce, 0xb0, 0xa7, 0x6e, 0xea, 0xa2, 0x4e, 0x41, 0x9d, 0xca, 0xd3, 0x9f, 0x70, 0x0d, 0x38, 0xaa,
	0xdf, 0x20, 0x43, 0xff, 0xaa, 0x41, 0xea, 0x90, 0xa2, 0x47, 0x9e, 0x43, 0x82, 0xda, 0x65, 0xad,
	0xaa, 0xd5, 0xfe, 0xdf, 0x5c, 0xad, 0x2f, 0xa2, 0xeb, 0x87, 0xb6, 0x9b, 0xbb, 0x70, 0xfe, 0x7d,
	0x35, 0x1d, 0x9c, 0x8d, 0x04, 0xb5, 0x49, 0x05, 0x92, 0x3d, 0xc4, 0x72, 0xa2, 0xaa, 0xd5, 0xf2,
	0x9b, 0xd9, 0xfa, 0x2e, 0xed, 0x1f, 0xb5, 0x99, 0x34, 0x7c, 0x23, 0x21, 0x90, 0x62, 0x96, 0x8b,
	0xe5, 0x64, 0x55, 0xab, 0xe5, 0x0c, 0x75, 0x26, 0xab, 0x90, 0xef, 0x38, 0xbc, 0x7b, 0x62, 0x7a,
	0x16, 0xeb, 0x63, 0x39, 0x55, 0xd5, 0x6a, 0x29, 0x03, 0x94, 0xc9, 0xf0, 0x2d, 0xa4, 0x04, 0x19,
	0xd7, 0x1a, 0x9b, 0x72, 0x2c, 0xca, 0x4b, 0xca, 0x99, 0x76, 0xad, 0xf1, 0xe1, 0x58, 0xe8, 0x5f,
	0x34, 0xc8, 0x2a, 0x60, 0xd6, 0xe3, 0xb7, 0xe3, 0x5c, 0x98, 0x72, 0x4e, 0xfd, 0x09, 0xa6, 0x2e,
	0xe4, 0xdb, 0x8c, 0x4a, 0x03, 0x3f, 0x0e, 0x51, 0x48, 0xf2, 0x10, 0x96, 0xf8, 0x88, 0xa1, 0xa7,
	0xe8, 0xfa, 0x22, 0xed, 0xd8, 0xb6, 0x87, 0x42, 0x18, 0x81, 0x99, 0xbc, 0x80, 0x9c, 0xa4, 0xe8,
	0x99, 0x94, 0xf5, 0x78, 0x39, 0x51, 0x4d, 0xd6, 0xf2, 0x9b, 0x8f, 0xae, 0x6e, 0x89, 0xf5, 0xb8,
	0x91, 0x95, 0xe1, 0x49, 0x6f, 0x41, 0xa1, 0x19, 0x06, 0xbe, 0xe4, 0x4c, 0x7a, 0x56, 0x57, 0x92,
	0x2d, 0x28, 0x74, 0xc3, 0xb3, 0x69, 0x05, 0x70, 0x73, 0xf0, 0xcb, 0x51, 0x44, 0x68, 0xd0, 0x3f,
	0x6b, 0xb0, 0xf2, 0x2e, 0x42, 0x3a, 0x12, 0xe8, 0x45, 0x65, 0xa3, 0x4e, 0x36, 0xe0, 0xbf, 0x09,
	0x25, 0xbf, 0xea, 0x5c, 0xc9, 0x7f, 0x23, 0xb7, 0x6f, 0x20, 0x4d, 0xc8, 0x04, 0x8d, 0xd9, 0x4a,
	0xeb, 0x6b, 0x4e, 0x2a, 0xad, 0x1a, 0xb4, 0xf5, 0x37, 0x70, 0xcf, 0x40, 0x97, 0x9f, 0xe2, 0xed,
	0x19, 0xe9, 0x1f, 0x20, 0xe7, 0x57, 0x39, 0x90, 0x96, 0x44, 0xf2, 0x18, 0x72, 0x8a, 0xce, 0xa5,
	0x79, 0x59, 0xdf, 0xa5, 0xba, 0x78, 0x02, 0xb9, 0xa8, 0x86, 0x08, 0xc7, 0x33, 0x0d, 0x9b, 0xba,
	0xf4, 0x73, 0x0d, 0x8a, 0x71, 0x8a, 0x01, 0x88, 0x0e, 0x99, 0x45, 0xfa, 0x47, 0x0e, 0xd2, 0x82,
	0x5c, 0x34, 0x8a, 0x08, 0x61, 0x6d, 0xa1, 0x52, 0xb3, 0xa3, 0x36, 0xa6, 0xb9, 0x71, 0xc1, 0x93,
	0x37, 0x16, 0x9c, 0x54, 0x20, 0x4b, 0x99, 0xd5, 0x95, 0xf4, 0x34, 0xd8, 0xfa, 0xac, 0x31, 0xb9,
	0xeb, 0x27, 0x50, 0x6a, 0xe1, 0x85, 0xdd, 0x10, 0xd1, 0x28, 0x7e, 0x51, 0xce, 0x35, 0x28, 0x50,
	0xd6, 0x75, 0x86, 0x36, 0x9a, 0x13, 0x94, 0x84, 0x42, 0x59, 0x0e, 0xed, 0xed, 0x08, 0xcc, 0x86,
	0xf2, 0x3c, 0x98, 0x18, 0x70, 0x26, 0x90, 0xec, 0xc5, 0xa7, 0xa2, 0x29, 0xcd, 0xd6, 0x17, 0x36,
	0x3b, 0x37, 0x96, 0xf8, 0xdc, 0xf6, 0xe1, 0x7e, 0x0b, 0x65, 0xe8, 0xb6, 0x23, 0x59, 0xc5, 0x0d,
	0x37, 0x6c, 0x0c, 0x2b, 0x97, 0x57, 0x0b, 0x79, 0xbf, 0x07, 0x32, 0xfb, 0xc7, 0xc4, 0xa8, 0x81,
	0x6b, 0x0c, 0xbd, 0x38, 0xf3, 0xdf, 0x45, 0xa1, 0xbf, 0x05, 0xd2, 0x42, 0x39, 0x79, 0x1f, 0x42,
	0xfa, 0xb7, 0x79, 0x28, 0xf5, 0x3d, 0xb8, 0x73, 0xa1, 0x64, 0xd8, 0xc3, 0x33, 0x48, 0xf9, 0xab,
	0x12, 0x2a, 0xf1, 0xe0, 0xca, 0xaa, 0x86, 0x0a, 0xd5, 0xbf, 0x69, 0x40, 0x0e, 0x7e, 0x2f, 0xbb,
	0xbf, 0xf8, 0xe9, 0xf9, 0xa4, 0x41, 0xe9, 0x60, 0x64, 0x0d, 0x2e, 0x7b, 0x81, 0xb6, 0xa1, 0xc8,
	0x1d, 0xdb, 0xbc, 0x7a, 0x47, 0x96, 0xb9, 0x63, 0x37, 0xe3, 0x4f, 0xe3, 0x36, 0x14, 0x19, 0x8e,
	0x66, 0xb2, 0x12, 0xb3, 0x59, 0x0c, 0x47, 0xf1, 0xac, 0xf5, 0xbb, 0x10, 0xf6, 0x4f, 0xf2, 0x90,
	0x69, 0xbe, 0x7a, 0xbd, 0x73, 0xb4, 0x7f, 0x58, 0xf8, 0xa7, 0x93, 0x56, 0x1f, 0xf4, 0xad, 0x1f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x94, 0x54, 0xf2, 0x2e, 0xb0, 0x08, 0x00, 0x00,
}
