// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evmos/fees/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgRegisterContract defines a Msg to register a contract deployer to receive
// tx fees.
type MsgRegisterContract struct {
	// contract hex address
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// nonce (sequence) number that generates the contract address
	Nonce uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// contract deployer hex address
	DeployerAddress string `protobuf:"bytes,3,opt,name=deployer_address,json=deployerAddress,proto3" json:"deployer_address,omitempty"`
}

func (m *MsgRegisterContract) Reset()         { *m = MsgRegisterContract{} }
func (m *MsgRegisterContract) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterContract) ProtoMessage()    {}
func (*MsgRegisterContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_8839b3a6b237f9b6, []int{0}
}
func (m *MsgRegisterContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterContract.Merge(m, src)
}
func (m *MsgRegisterContract) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterContract) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterContract.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterContract proto.InternalMessageInfo

func (m *MsgRegisterContract) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *MsgRegisterContract) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *MsgRegisterContract) GetDeployerAddress() string {
	if m != nil {
		return m.DeployerAddress
	}
	return ""
}

// MsgRegisterContractResponse returns no fields
type MsgRegisterContractResponse struct {
}

func (m *MsgRegisterContractResponse) Reset()         { *m = MsgRegisterContractResponse{} }
func (m *MsgRegisterContractResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterContractResponse) ProtoMessage()    {}
func (*MsgRegisterContractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8839b3a6b237f9b6, []int{1}
}
func (m *MsgRegisterContractResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterContractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterContractResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterContractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterContractResponse.Merge(m, src)
}
func (m *MsgRegisterContractResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterContractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterContractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterContractResponse proto.InternalMessageInfo

// MsgUpdateWithdawAddress defines a Msg to update the withdraw address to which
// tx fees will be allocated.
type MsgUpdateWithdawAddress struct {
	// contract hex address registered on the fee distribution module
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// new hex address to receive tx fees.
	NewWithdrawAddress string `protobuf:"bytes,2,opt,name=new_withdraw_address,json=newWithdrawAddress,proto3" json:"new_withdraw_address,omitempty"`
	// hex address from the registered owner/withdrawal address
	WithdrawAddress string `protobuf:"bytes,3,opt,name=withdraw_address,json=withdrawAddress,proto3" json:"withdraw_address,omitempty"`
}

func (m *MsgUpdateWithdawAddress) Reset()         { *m = MsgUpdateWithdawAddress{} }
func (m *MsgUpdateWithdawAddress) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateWithdawAddress) ProtoMessage()    {}
func (*MsgUpdateWithdawAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_8839b3a6b237f9b6, []int{2}
}
func (m *MsgUpdateWithdawAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateWithdawAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateWithdawAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateWithdawAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateWithdawAddress.Merge(m, src)
}
func (m *MsgUpdateWithdawAddress) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateWithdawAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateWithdawAddress.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateWithdawAddress proto.InternalMessageInfo

func (m *MsgUpdateWithdawAddress) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *MsgUpdateWithdawAddress) GetNewWithdrawAddress() string {
	if m != nil {
		return m.NewWithdrawAddress
	}
	return ""
}

func (m *MsgUpdateWithdawAddress) GetWithdrawAddress() string {
	if m != nil {
		return m.WithdrawAddress
	}
	return ""
}

// MsgUpdateWithdawAddressResponse returns no fields
type MsgUpdateWithdawAddressResponse struct {
}

func (m *MsgUpdateWithdawAddressResponse) Reset()         { *m = MsgUpdateWithdawAddressResponse{} }
func (m *MsgUpdateWithdawAddressResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateWithdawAddressResponse) ProtoMessage()    {}
func (*MsgUpdateWithdawAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8839b3a6b237f9b6, []int{3}
}
func (m *MsgUpdateWithdawAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateWithdawAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateWithdawAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateWithdawAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateWithdawAddressResponse.Merge(m, src)
}
func (m *MsgUpdateWithdawAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateWithdawAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateWithdawAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateWithdawAddressResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegisterContract)(nil), "evmos.fees.v1.MsgRegisterContract")
	proto.RegisterType((*MsgRegisterContractResponse)(nil), "evmos.fees.v1.MsgRegisterContractResponse")
	proto.RegisterType((*MsgUpdateWithdawAddress)(nil), "evmos.fees.v1.MsgUpdateWithdawAddress")
	proto.RegisterType((*MsgUpdateWithdawAddressResponse)(nil), "evmos.fees.v1.MsgUpdateWithdawAddressResponse")
}

func init() { proto.RegisterFile("evmos/fees/v1/tx.proto", fileDescriptor_8839b3a6b237f9b6) }

var fileDescriptor_8839b3a6b237f9b6 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xee, 0x64, 0x55, 0x70, 0x40, 0x5c, 0xc6, 0xa0, 0x4b, 0xd4, 0xb8, 0x06, 0x94, 0x55, 0x31,
	0x63, 0xed, 0xc9, 0x93, 0xa8, 0xe7, 0x5e, 0x02, 0xb2, 0xe0, 0xa5, 0x4c, 0x93, 0xe7, 0x74, 0x60,
	0x77, 0x26, 0xcc, 0x4c, 0x9b, 0xf6, 0xa8, 0xbf, 0x40, 0xf0, 0xe0, 0xd9, 0xa3, 0xff, 0x44, 0xf0,
	0x52, 0xf0, 0xe2, 0x51, 0x5a, 0x7f, 0x88, 0x24, 0x93, 0x04, 0x4b, 0x52, 0xd4, 0xdb, 0xe4, 0x7b,
	0xdf, 0x7c, 0xf3, 0xbd, 0xef, 0xbd, 0xe0, 0xeb, 0xb0, 0x38, 0x57, 0x86, 0xbe, 0x05, 0x30, 0x74,
	0x31, 0xa4, 0x76, 0x19, 0xe7, 0x5a, 0x59, 0x45, 0xae, 0x54, 0x78, 0x5c, 0xe2, 0xf1, 0x62, 0x18,
	0xf8, 0x5c, 0x71, 0x55, 0x55, 0x68, 0x79, 0x72, 0xa4, 0xe0, 0x16, 0x57, 0x8a, 0x9f, 0x01, 0x65,
	0xb9, 0xa0, 0x4c, 0x4a, 0x65, 0x99, 0x15, 0x4a, 0x1a, 0x57, 0x8d, 0xde, 0x21, 0x7c, 0x6d, 0x6c,
	0x78, 0x02, 0x5c, 0x18, 0x0b, 0xfa, 0x95, 0x92, 0x56, 0xb3, 0xd4, 0x92, 0x07, 0xf8, 0x30, 0xad,
	0xcf, 0x13, 0x96, 0x65, 0x1a, 0x8c, 0x39, 0x42, 0xc7, 0xe8, 0xe4, 0x72, 0x72, 0xb5, 0xc1, 0x5f,
	0x38, 0x98, 0xf8, 0xf8, 0xa2, 0x54, 0x32, 0x85, 0x23, 0xef, 0x18, 0x9d, 0x5c, 0x48, 0xdc, 0x47,
	0x29, 0x90, 0x41, 0x7e, 0xa6, 0x56, 0xa0, 0x5b, 0x81, 0x03, 0x27, 0xd0, 0xe0, 0xb5, 0x40, 0x74,
	0x1b, 0xdf, 0xec, 0xb1, 0x90, 0x80, 0xc9, 0x95, 0x34, 0x10, 0x7d, 0x46, 0xf8, 0xc6, 0xd8, 0xf0,
	0xd7, 0x79, 0xc6, 0x2c, 0x9c, 0x0a, 0x3b, 0xcb, 0x58, 0xd1, 0xbc, 0xfd, 0x1f, 0x36, 0x9f, 0x60,
	0x5f, 0x42, 0x31, 0x29, 0x4a, 0x01, 0xcd, 0x8a, 0x96, 0xee, 0x55, 0x74, 0x22, 0xa1, 0x38, 0xad,
	0x4b, 0x7f, 0x88, 0x77, 0xd8, 0x75, 0x0b, 0xc5, 0x2e, 0x35, 0xba, 0x8b, 0xef, 0xec, 0xb1, 0xd8,
	0xb4, 0xf1, 0xf4, 0x9b, 0x87, 0x0f, 0xc6, 0x86, 0x93, 0x4f, 0x08, 0x1f, 0x76, 0xe2, 0x8e, 0xe2,
	0x9d, 0x51, 0xc6, 0x3d, 0x79, 0x04, 0x0f, 0xff, 0xce, 0x69, 0x33, 0x1b, 0xbd, 0xff, 0xfe, 0xeb,
	0xa3, 0xf7, 0x98, 0x3c, 0xa2, 0x6e, 0x75, 0x32, 0x61, 0xac, 0x16, 0xd3, 0x79, 0x39, 0x7a, 0xb7,
	0x42, 0x54, 0xd7, 0x77, 0x27, 0x4d, 0x50, 0xe4, 0x0b, 0xc2, 0x7e, 0x6f, 0xca, 0xf7, 0xbb, 0x2f,
	0xf7, 0xf1, 0x82, 0xf8, 0xdf, 0x78, 0xad, 0xcb, 0x67, 0x95, 0xcb, 0x11, 0x19, 0xee, 0x77, 0x39,
	0xaf, 0xee, 0x77, 0xa6, 0xf6, 0xf2, 0xf9, 0xd7, 0x4d, 0x88, 0xd6, 0x9b, 0x10, 0xfd, 0xdc, 0x84,
	0xe8, 0xc3, 0x36, 0x1c, 0xac, 0xb7, 0xe1, 0xe0, 0xc7, 0x36, 0x1c, 0xbc, 0xb9, 0xc7, 0x85, 0x9d,
	0xcd, 0xa7, 0x71, 0xaa, 0xce, 0xa9, 0x9d, 0x31, 0x6d, 0x84, 0xa9, 0xe5, 0x97, 0xee, 0x0f, 0xb2,
	0xab, 0x1c, 0xcc, 0xf4, 0x52, 0xb5, 0xff, 0xa3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x7a,
	0x6f, 0x50, 0x5c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// RegisterContract registers the contract on the fee distribution module in order for the owner
	// to recieve txs fees.
	RegisterContract(ctx context.Context, in *MsgRegisterContract, opts ...grpc.CallOption) (*MsgRegisterContractResponse, error)
	// UpdateWithdawAddress updates the withdrawal address for the contract that has been
	UpdateWithdawAddress(ctx context.Context, in *MsgUpdateWithdawAddress, opts ...grpc.CallOption) (*MsgUpdateWithdawAddressResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterContract(ctx context.Context, in *MsgRegisterContract, opts ...grpc.CallOption) (*MsgRegisterContractResponse, error) {
	out := new(MsgRegisterContractResponse)
	err := c.cc.Invoke(ctx, "/evmos.fees.v1.Msg/RegisterContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateWithdawAddress(ctx context.Context, in *MsgUpdateWithdawAddress, opts ...grpc.CallOption) (*MsgUpdateWithdawAddressResponse, error) {
	out := new(MsgUpdateWithdawAddressResponse)
	err := c.cc.Invoke(ctx, "/evmos.fees.v1.Msg/UpdateWithdawAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// RegisterContract registers the contract on the fee distribution module in order for the owner
	// to recieve txs fees.
	RegisterContract(context.Context, *MsgRegisterContract) (*MsgRegisterContractResponse, error)
	// UpdateWithdawAddress updates the withdrawal address for the contract that has been
	UpdateWithdawAddress(context.Context, *MsgUpdateWithdawAddress) (*MsgUpdateWithdawAddressResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RegisterContract(ctx context.Context, req *MsgRegisterContract) (*MsgRegisterContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterContract not implemented")
}
func (*UnimplementedMsgServer) UpdateWithdawAddress(ctx context.Context, req *MsgUpdateWithdawAddress) (*MsgUpdateWithdawAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWithdawAddress not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegisterContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evmos.fees.v1.Msg/RegisterContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterContract(ctx, req.(*MsgRegisterContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateWithdawAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateWithdawAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateWithdawAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evmos.fees.v1.Msg/UpdateWithdawAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateWithdawAddress(ctx, req.(*MsgUpdateWithdawAddress))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "evmos.fees.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterContract",
			Handler:    _Msg_RegisterContract_Handler,
		},
		{
			MethodName: "UpdateWithdawAddress",
			Handler:    _Msg_UpdateWithdawAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "evmos/fees/v1/tx.proto",
}

func (m *MsgRegisterContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DeployerAddress) > 0 {
		i -= len(m.DeployerAddress)
		copy(dAtA[i:], m.DeployerAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DeployerAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Nonce != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterContractResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterContractResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterContractResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateWithdawAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateWithdawAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateWithdawAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WithdrawAddress) > 0 {
		i -= len(m.WithdrawAddress)
		copy(dAtA[i:], m.WithdrawAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.WithdrawAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NewWithdrawAddress) > 0 {
		i -= len(m.NewWithdrawAddress)
		copy(dAtA[i:], m.NewWithdrawAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.NewWithdrawAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateWithdawAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateWithdawAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateWithdawAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRegisterContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovTx(uint64(m.Nonce))
	}
	l = len(m.DeployerAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRegisterContractResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateWithdawAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.NewWithdrawAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.WithdrawAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateWithdawAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRegisterContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRegisterContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeployerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeployerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRegisterContractResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRegisterContractResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterContractResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateWithdawAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateWithdawAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateWithdawAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewWithdrawAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewWithdrawAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateWithdawAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateWithdawAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateWithdawAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
