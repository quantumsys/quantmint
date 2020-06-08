// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: codec/codec.proto

package codec

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_x_auth_exported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
	types1 "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	types2 "github.com/cosmos/cosmos-sdk/x/supply/types"
	proto "github.com/gogo/protobuf/proto"
	types3 "github.com/quantumsys/quantmint/types"
	_ "github.com/regen-network/cosmos-proto"
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

// Account defines the application-level Account type.
type Account struct {
	// sum defines a list of all acceptable concrete Account implementations.
	//
	// Types that are valid to be assigned to Sum:
	//	*Account_BaseAccount
	//	*Account_ContinuousVestingAccount
	//	*Account_DelayedVestingAccount
	//	*Account_PeriodicVestingAccount
	//	*Account_ModuleAccount
	//	*Account_EthAccount
	Sum isAccount_Sum `protobuf_oneof:"sum"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_2557dd8a93a64b89, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Account.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

type isAccount_Sum interface {
	isAccount_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Account_BaseAccount struct {
	BaseAccount *types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,oneof" json:"base_account,omitempty"`
}
type Account_ContinuousVestingAccount struct {
	ContinuousVestingAccount *types1.ContinuousVestingAccount `protobuf:"bytes,2,opt,name=continuous_vesting_account,json=continuousVestingAccount,proto3,oneof" json:"continuous_vesting_account,omitempty"`
}
type Account_DelayedVestingAccount struct {
	DelayedVestingAccount *types1.DelayedVestingAccount `protobuf:"bytes,3,opt,name=delayed_vesting_account,json=delayedVestingAccount,proto3,oneof" json:"delayed_vesting_account,omitempty"`
}
type Account_PeriodicVestingAccount struct {
	PeriodicVestingAccount *types1.PeriodicVestingAccount `protobuf:"bytes,4,opt,name=periodic_vesting_account,json=periodicVestingAccount,proto3,oneof" json:"periodic_vesting_account,omitempty"`
}
type Account_ModuleAccount struct {
	ModuleAccount *types2.ModuleAccount `protobuf:"bytes,5,opt,name=module_account,json=moduleAccount,proto3,oneof" json:"module_account,omitempty"`
}
type Account_EthAccount struct {
	EthAccount *types3.EthAccount `protobuf:"bytes,6,opt,name=eth_account,json=ethAccount,proto3,oneof" json:"eth_account,omitempty"`
}

func (*Account_BaseAccount) isAccount_Sum()              {}
func (*Account_ContinuousVestingAccount) isAccount_Sum() {}
func (*Account_DelayedVestingAccount) isAccount_Sum()    {}
func (*Account_PeriodicVestingAccount) isAccount_Sum()   {}
func (*Account_ModuleAccount) isAccount_Sum()            {}
func (*Account_EthAccount) isAccount_Sum()               {}

func (m *Account) GetSum() isAccount_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Account) GetBaseAccount() *types.BaseAccount {
	if x, ok := m.GetSum().(*Account_BaseAccount); ok {
		return x.BaseAccount
	}
	return nil
}

func (m *Account) GetContinuousVestingAccount() *types1.ContinuousVestingAccount {
	if x, ok := m.GetSum().(*Account_ContinuousVestingAccount); ok {
		return x.ContinuousVestingAccount
	}
	return nil
}

func (m *Account) GetDelayedVestingAccount() *types1.DelayedVestingAccount {
	if x, ok := m.GetSum().(*Account_DelayedVestingAccount); ok {
		return x.DelayedVestingAccount
	}
	return nil
}

func (m *Account) GetPeriodicVestingAccount() *types1.PeriodicVestingAccount {
	if x, ok := m.GetSum().(*Account_PeriodicVestingAccount); ok {
		return x.PeriodicVestingAccount
	}
	return nil
}

func (m *Account) GetModuleAccount() *types2.ModuleAccount {
	if x, ok := m.GetSum().(*Account_ModuleAccount); ok {
		return x.ModuleAccount
	}
	return nil
}

func (m *Account) GetEthAccount() *types3.EthAccount {
	if x, ok := m.GetSum().(*Account_EthAccount); ok {
		return x.EthAccount
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Account) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Account_BaseAccount)(nil),
		(*Account_ContinuousVestingAccount)(nil),
		(*Account_DelayedVestingAccount)(nil),
		(*Account_PeriodicVestingAccount)(nil),
		(*Account_ModuleAccount)(nil),
		(*Account_EthAccount)(nil),
	}
}

func init() {
	proto.RegisterType((*Account)(nil), "ethermint.codec.v1.Account")
}

func init() { proto.RegisterFile("codec/codec.proto", fileDescriptor_2557dd8a93a64b89) }

var fileDescriptor_2557dd8a93a64b89 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0xeb, 0xd3, 0x30,
	0x18, 0xc7, 0x5b, 0xf7, 0x47, 0xc8, 0x54, 0x58, 0x40, 0x2d, 0x3b, 0x14, 0x1d, 0x08, 0xa2, 0x2c,
	0x65, 0xce, 0x29, 0xfe, 0x3b, 0x38, 0xff, 0xe0, 0x45, 0x11, 0x0f, 0x1e, 0xbc, 0x94, 0x36, 0x09,
	0x6b, 0xd9, 0xda, 0x84, 0x26, 0x29, 0xeb, 0xbb, 0xf0, 0xc5, 0xf8, 0x22, 0xc4, 0xd3, 0x8e, 0x1e,
	0x65, 0x7b, 0x15, 0xde, 0x64, 0x49, 0x69, 0x27, 0xed, 0xf6, 0xfb, 0x5d, 0x02, 0xcf, 0xf3, 0x7c,
	0xbf, 0xdf, 0x4f, 0x20, 0x4f, 0xc0, 0x10, 0x33, 0x42, 0xb1, 0xa7, 0x4f, 0xc4, 0x33, 0x26, 0x19,
	0x84, 0x54, 0x46, 0x34, 0x4b, 0xe2, 0x54, 0x22, 0xd3, 0xce, 0xa7, 0xa3, 0xa1, 0x2c, 0x38, 0x15,
	0x9e, 0x3e, 0x8d, 0x6c, 0xf4, 0x50, 0x46, 0x71, 0x46, 0x7c, 0x1e, 0x64, 0xb2, 0xf0, 0x74, 0xcb,
	0xc3, 0x4c, 0x24, 0x4c, 0x4c, 0x8e, 0x8b, 0x52, 0x3c, 0x3f, 0x29, 0x16, 0x64, 0xe5, 0x6d, 0xbc,
	0x40, 0xc9, 0xc8, 0x6b, 0x32, 0x5e, 0x5d, 0xc6, 0x96, 0x53, 0x21, 0xe3, 0x74, 0xd9, 0x62, 0x7f,
	0x7a, 0x81, 0x5d, 0x28, 0xce, 0xd7, 0x45, 0xd3, 0x38, 0xfe, 0xdb, 0x05, 0x57, 0x5f, 0x63, 0xcc,
	0x54, 0x2a, 0xe1, 0x7b, 0x70, 0x2d, 0x0c, 0x04, 0xf5, 0x03, 0x53, 0x3b, 0xf6, 0x1d, 0xfb, 0xfe,
	0xe0, 0xd1, 0x5d, 0x64, 0x92, 0x7c, 0x41, 0x56, 0x68, 0x83, 0x0e, 0x17, 0x41, 0xf9, 0x14, 0x2d,
	0x02, 0x41, 0x4b, 0xe3, 0x07, 0xeb, 0xcb, 0x20, 0xac, 0x4b, 0x98, 0x83, 0x11, 0x66, 0xa9, 0x8c,
	0x53, 0xc5, 0x94, 0xf0, 0xcb, 0x4b, 0x57, 0xa9, 0x57, 0x74, 0xea, 0x93, 0xb6, 0x54, 0xa3, 0x3c,
	0xa4, 0xbf, 0xa9, 0xfc, 0x5f, 0x4d, 0xb3, 0x46, 0x39, 0xf8, 0xc4, 0x0c, 0x26, 0xe0, 0x36, 0xa1,
	0xeb, 0xa0, 0xa0, 0xa4, 0x01, 0xed, 0x68, 0xe8, 0xec, 0x3c, 0xf4, 0xad, 0x31, 0x37, 0x88, 0x37,
	0x49, 0xdb, 0x00, 0x72, 0xe0, 0x70, 0x9a, 0xc5, 0x8c, 0xc4, 0xb8, 0xc1, 0xeb, 0x6a, 0xde, 0xe3,
	0xf3, 0xbc, 0xcf, 0xa5, 0xbb, 0x01, 0xbc, 0xc5, 0x5b, 0x27, 0xf0, 0x13, 0xb8, 0x91, 0x30, 0xa2,
	0xd6, 0xf5, 0x13, 0xf5, 0x34, 0xe7, 0xde, 0xff, 0x1c, 0xf3, 0xd8, 0x07, 0xc2, 0x47, 0xad, 0xae,
	0x83, 0xaf, 0x27, 0xc7, 0x0d, 0xf8, 0x02, 0x0c, 0xa8, 0x8c, 0xaa, 0xb0, 0xbe, 0x0e, 0x73, 0x50,
	0xfd, 0x2b, 0xf2, 0x29, 0x7a, 0x27, 0xa3, 0xda, 0x0f, 0x68, 0x55, 0x3d, 0x7f, 0xf6, 0xeb, 0xc7,
	0x64, 0xfe, 0x60, 0x19, 0xcb, 0x48, 0x85, 0x08, 0xb3, 0xa4, 0x5c, 0xb8, 0x96, 0xb5, 0xa5, 0x1b,
	0xce, 0x32, 0x49, 0x09, 0x2a, 0xad, 0x8b, 0x1e, 0xe8, 0x08, 0x95, 0x2c, 0x5e, 0xfe, 0xdc, 0xb9,
	0xf6, 0x76, 0xe7, 0xda, 0x7f, 0x76, 0xae, 0xfd, 0x7d, 0xef, 0x5a, 0xdb, 0xbd, 0x6b, 0xfd, 0xde,
	0xbb, 0xd6, 0xb7, 0xf1, 0xd9, 0x58, 0xfd, 0x57, 0xc3, 0xbe, 0x5e, 0xe0, 0xd9, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd5, 0xd2, 0x26, 0x71, 0xd8, 0x03, 0x00, 0x00,
}

func (this *Account) GetAccount() github_com_cosmos_cosmos_sdk_x_auth_exported.Account {
	if x := this.GetBaseAccount(); x != nil {
		return x
	}
	if x := this.GetContinuousVestingAccount(); x != nil {
		return x
	}
	if x := this.GetDelayedVestingAccount(); x != nil {
		return x
	}
	if x := this.GetPeriodicVestingAccount(); x != nil {
		return x
	}
	if x := this.GetModuleAccount(); x != nil {
		return x
	}
	if x := this.GetEthAccount(); x != nil {
		return x
	}
	return nil
}

func (this *Account) SetAccount(value github_com_cosmos_cosmos_sdk_x_auth_exported.Account) error {
	if value == nil {
		this.Sum = nil
		return nil
	}
	switch vt := value.(type) {
	case *types.BaseAccount:
		this.Sum = &Account_BaseAccount{vt}
		return nil
	case *types1.ContinuousVestingAccount:
		this.Sum = &Account_ContinuousVestingAccount{vt}
		return nil
	case *types1.DelayedVestingAccount:
		this.Sum = &Account_DelayedVestingAccount{vt}
		return nil
	case *types1.PeriodicVestingAccount:
		this.Sum = &Account_PeriodicVestingAccount{vt}
		return nil
	case *types2.ModuleAccount:
		this.Sum = &Account_ModuleAccount{vt}
		return nil
	case *types3.EthAccount:
		this.Sum = &Account_EthAccount{vt}
		return nil
	}
	return fmt.Errorf("can't encode value of type %T as message Account", value)
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Account_BaseAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account_BaseAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Account_ContinuousVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account_ContinuousVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ContinuousVestingAccount != nil {
		{
			size, err := m.ContinuousVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Account_DelayedVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account_DelayedVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DelayedVestingAccount != nil {
		{
			size, err := m.DelayedVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Account_PeriodicVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account_PeriodicVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PeriodicVestingAccount != nil {
		{
			size, err := m.PeriodicVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Account_ModuleAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account_ModuleAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ModuleAccount != nil {
		{
			size, err := m.ModuleAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *Account_EthAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account_EthAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EthAccount != nil {
		{
			size, err := m.EthAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	offset -= sovCodec(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *Account_BaseAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Account_ContinuousVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ContinuousVestingAccount != nil {
		l = m.ContinuousVestingAccount.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Account_DelayedVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DelayedVestingAccount != nil {
		l = m.DelayedVestingAccount.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Account_PeriodicVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PeriodicVestingAccount != nil {
		l = m.PeriodicVestingAccount.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Account_ModuleAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ModuleAccount != nil {
		l = m.ModuleAccount.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Account_EthAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EthAccount != nil {
		l = m.EthAccount.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Account) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types.BaseAccount{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Account_BaseAccount{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContinuousVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types1.ContinuousVestingAccount{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Account_ContinuousVestingAccount{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelayedVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types1.DelayedVestingAccount{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Account_DelayedVestingAccount{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodicVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types1.PeriodicVestingAccount{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Account_PeriodicVestingAccount{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types2.ModuleAccount{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Account_ModuleAccount{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types3.EthAccount{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Account_EthAccount{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
				return 0, ErrInvalidLengthCodec
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCodec
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCodec
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCodec        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCodec = fmt.Errorf("proto: unexpected end of group")
)
