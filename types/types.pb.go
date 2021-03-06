// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/types.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_quantumsys_quantmint_types "github.com/quantumsys/quantmint/types"
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

// EthAccount implements the auth.Account interface and embeds an
// auth.BaseAccount type. It is compatible with the auth.AccountKeeper.
type EthAccount struct {
	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty" yaml:"base_account"`
	CodeHash           []byte                                                   `protobuf:"bytes,2,opt,name=code_hash,json=codeHash,proto3" json:"code_hash,omitempty" yaml:"code_hash"`
	Name               string                                                   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
	SubAddresses       map[string]github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,4,rep,name=sub_addresses,json=subAddresses,proto3,castvalue=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sub_addresses,omitempty" yaml:"sub_addresses" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *EthAccount) Reset()      { *m = EthAccount{} }
func (*EthAccount) ProtoMessage() {}
func (*EthAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{0}
}
func (m *EthAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EthAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthAccount.Merge(m, src)
}
func (m *EthAccount) XXX_Size() int {
	return m.Size()
}
func (m *EthAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_EthAccount.DiscardUnknown(m)
}

var xxx_messageInfo_EthAccount proto.InternalMessageInfo

type SubAccount struct {
	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty" yaml:"base_account"`
	MainAddress        github_com_cosmos_cosmos_sdk_types.AccAddress     `protobuf:"bytes,2,opt,name=main_address,json=mainAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"main_address,omitempty" yaml:"main_address"`
	Perms              []github_com_quantumsys_quantmint_types.PermValue `protobuf:"varint,3,rep,packed,name=perms,proto3,casttype=github.com/quantumsys/quantmint/types.PermValue" json:"perms,omitempty" yaml:"perms"`
}

func (m *SubAccount) Reset()      { *m = SubAccount{} }
func (*SubAccount) ProtoMessage() {}
func (*SubAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{1}
}
func (m *SubAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubAccount.Merge(m, src)
}
func (m *SubAccount) XXX_Size() int {
	return m.Size()
}
func (m *SubAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_SubAccount.DiscardUnknown(m)
}

var xxx_messageInfo_SubAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EthAccount)(nil), "ethermint.v1.EthAccount")
	proto.RegisterMapType((map[string]github_com_cosmos_cosmos_sdk_types.AccAddress)(nil), "ethermint.v1.EthAccount.SubAddressesEntry")
	proto.RegisterType((*SubAccount)(nil), "ethermint.v1.SubAccount")
}

func init() { proto.RegisterFile("types/types.proto", fileDescriptor_2c0f90c600ad7e2e) }

var fileDescriptor_2c0f90c600ad7e2e = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x41, 0x8b, 0xd3, 0x4e,
	0x14, 0xc0, 0x3b, 0xcd, 0xf6, 0xcf, 0x76, 0x92, 0x3f, 0xee, 0xc6, 0x1e, 0x42, 0x85, 0x24, 0x46,
	0x94, 0x20, 0xec, 0x84, 0xae, 0x08, 0xd2, 0xcb, 0xd2, 0xc8, 0xc2, 0x7a, 0x93, 0x08, 0x82, 0x5e,
	0xc2, 0x24, 0x19, 0x9a, 0xd2, 0x4d, 0x52, 0x33, 0x93, 0xb2, 0xf9, 0x06, 0xe2, 0xc9, 0xa3, 0x78,
	0xda, 0x6f, 0xe0, 0x67, 0xf0, 0xe6, 0xb1, 0x47, 0x4f, 0x51, 0xda, 0x6f, 0x90, 0xa3, 0x27, 0x49,
	0x26, 0x6b, 0x22, 0xbd, 0xe8, 0xc5, 0x4b, 0x79, 0x6f, 0xde, 0x7b, 0xbf, 0x47, 0x7f, 0x8f, 0xc0,
	0x63, 0x96, 0xaf, 0x08, 0xb5, 0xea, 0x5f, 0xb4, 0x4a, 0x13, 0x96, 0xc8, 0x12, 0x61, 0x21, 0x49,
	0xa3, 0x45, 0xcc, 0xd0, 0x7a, 0x32, 0x7e, 0xc0, 0xc2, 0x45, 0x1a, 0xb8, 0x2b, 0x9c, 0xb2, 0xdc,
	0xaa, 0x1b, 0xac, 0x79, 0x32, 0x4f, 0xda, 0x88, 0x4f, 0x8d, 0x1f, 0xef, 0xf7, 0xf9, 0x09, 0x8d,
	0x12, 0x7a, 0x42, 0x83, 0xa5, 0x75, 0x65, 0xe1, 0x8c, 0x85, 0xd6, 0xde, 0x32, 0xe3, 0x93, 0x00,
	0xe1, 0x39, 0x0b, 0x67, 0xbe, 0x9f, 0x64, 0x31, 0x93, 0x31, 0x94, 0x3c, 0x4c, 0x89, 0x8b, 0x79,
	0xae, 0x00, 0x1d, 0x98, 0xe2, 0xe9, 0x5d, 0xc4, 0x51, 0x2e, 0x0d, 0x96, 0xe8, 0x0a, 0x55, 0x28,
	0xb4, 0x9e, 0x20, 0x1b, 0x53, 0xd2, 0x0c, 0xda, 0x77, 0x36, 0x85, 0x06, 0xca, 0x42, 0xbb, 0x9d,
	0xe3, 0xe8, 0x72, 0x6a, 0x74, 0x21, 0x86, 0x23, 0x7a, 0x6d, 0xa7, 0x3c, 0x81, 0x43, 0x3f, 0x09,
	0x88, 0x1b, 0x62, 0x1a, 0x2a, 0x7d, 0x1d, 0x98, 0x92, 0x3d, 0x2a, 0x0b, 0xed, 0x88, 0x0f, 0xfe,
	0x2a, 0x19, 0xce, 0x61, 0x15, 0x5f, 0x60, 0x1a, 0xca, 0xf7, 0xe0, 0x41, 0x8c, 0x23, 0xa2, 0x08,
	0x3a, 0x30, 0x87, 0xf6, 0xad, 0xb2, 0xd0, 0x44, 0xde, 0x5d, 0xbd, 0x1a, 0x4e, 0x5d, 0x94, 0x3f,
	0x02, 0xf8, 0x3f, 0xcd, 0x3c, 0x17, 0x07, 0x41, 0x4a, 0x28, 0x25, 0x54, 0x39, 0xd0, 0x05, 0x53,
	0x3c, 0x7d, 0x88, 0xba, 0x3e, 0x51, 0xfb, 0x67, 0xd1, 0x8b, 0xcc, 0x9b, 0xdd, 0x34, 0x9f, 0xc7,
	0x2c, 0xcd, 0xed, 0x67, 0x65, 0xa1, 0x8d, 0x38, 0xfa, 0x37, 0x94, 0xf1, 0xee, 0x9b, 0x76, 0x32,
	0x5f, 0xb0, 0x30, 0xf3, 0x90, 0x9f, 0x44, 0x8d, 0xd9, 0xae, 0x60, 0xee, 0x74, 0xe6, 0xfb, 0x0d,
	0xd1, 0x91, 0x68, 0x87, 0x3e, 0x3e, 0x83, 0xc7, 0x7b, 0xdb, 0xe4, 0x23, 0x28, 0x2c, 0x49, 0x5e,
	0x3b, 0x1e, 0x3a, 0x55, 0x28, 0x8f, 0xe0, 0x60, 0x8d, 0x2f, 0x33, 0xc2, 0xbd, 0x38, 0x3c, 0x99,
	0xf6, 0x9f, 0x80, 0xe9, 0xe1, 0xdb, 0x6b, 0xad, 0xf7, 0xe1, 0x5a, 0xeb, 0x19, 0x9f, 0xfb, 0x10,
	0x56, 0xac, 0x7f, 0x77, 0xb1, 0x25, 0x94, 0x22, 0xbc, 0x88, 0x6f, 0x74, 0x34, 0x47, 0xbb, 0x68,
	0x67, 0xbb, 0x55, 0xe3, 0x47, 0xf1, 0xb7, 0xaa, 0xc4, 0x6a, 0xbe, 0x49, 0xe4, 0x57, 0x70, 0xb0,
	0x22, 0x69, 0x44, 0x15, 0x41, 0x17, 0xcc, 0x81, 0xfd, 0xb4, 0x2c, 0x34, 0x89, 0x6f, 0xa9, 0x9f,
	0x2b, 0xbc, 0xd5, 0xc1, 0xbf, 0xc9, 0x70, 0xcc, 0xb2, 0x88, 0xe6, 0x94, 0x87, 0xd5, 0x99, 0x9b,
	0x0d, 0xcf, 0x49, 0x1a, 0xbd, 0xac, 0x04, 0x3a, 0x9c, 0xd8, 0x3a, 0xb4, 0xcf, 0xbe, 0x6c, 0x55,
	0xb0, 0xd9, 0xaa, 0xe0, 0xfb, 0x56, 0x05, 0xef, 0x77, 0x6a, 0x6f, 0xb3, 0x53, 0x7b, 0x5f, 0x77,
	0x6a, 0xef, 0xf5, 0xfd, 0x3f, 0x62, 0x7b, 0xff, 0xd5, 0x5f, 0xcf, 0xa3, 0x9f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x34, 0x5c, 0x78, 0xc9, 0xbf, 0x03, 0x00, 0x00,
}

func (m *EthAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SubAddresses) > 0 {
		for k := range m.SubAddresses {
			v := m.SubAddresses[k]
			baseI := i
			if len(v) > 0 {
				i -= len(v)
				copy(dAtA[i:], v)
				i = encodeVarintTypes(dAtA, i, uint64(len(v)))
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintTypes(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintTypes(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CodeHash) > 0 {
		i -= len(m.CodeHash)
		copy(dAtA[i:], m.CodeHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.CodeHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SubAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Perms) > 0 {
		dAtA3 := make([]byte, len(m.Perms)*10)
		var j2 int
		for _, num1 := range m.Perms {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		i -= j2
		copy(dAtA[i:], dAtA3[:j2])
		i = encodeVarintTypes(dAtA, i, uint64(j2))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MainAddress) > 0 {
		i -= len(m.MainAddress)
		copy(dAtA[i:], m.MainAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.MainAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EthAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.CodeHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.SubAddresses) > 0 {
		for k, v := range m.SubAddresses {
			_ = k
			_ = v
			l = 0
			if len(v) > 0 {
				l = 1 + len(v) + sovTypes(uint64(len(v)))
			}
			mapEntrySize := 1 + len(k) + sovTypes(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovTypes(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *SubAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.MainAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Perms) > 0 {
		l = 0
		for _, e := range m.Perms {
			l += sovTypes(uint64(e))
		}
		n += 1 + sovTypes(uint64(l)) + l
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EthAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: EthAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseAccount == nil {
				m.BaseAccount = &types.BaseAccount{}
			}
			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeHash = append(m.CodeHash[:0], dAtA[iNdEx:postIndex]...)
			if m.CodeHash == nil {
				m.CodeHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubAddresses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SubAddresses == nil {
				m.SubAddresses = make(map[string]github_com_cosmos_cosmos_sdk_types.AccAddress)
			}
			var mapkey string
			mapvalue := []byte{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTypes
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthTypes
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapbyteLen uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapbyteLen |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intMapbyteLen := int(mapbyteLen)
					if intMapbyteLen < 0 {
						return ErrInvalidLengthTypes
					}
					postbytesIndex := iNdEx + intMapbyteLen
					if postbytesIndex < 0 {
						return ErrInvalidLengthTypes
					}
					if postbytesIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = make([]byte, mapbyteLen)
					copy(mapvalue, dAtA[iNdEx:postbytesIndex])
					iNdEx = postbytesIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTypes(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthTypes
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.SubAddresses[mapkey] = ((github_com_cosmos_cosmos_sdk_types.AccAddress)(mapvalue))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *SubAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: SubAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseAccount == nil {
				m.BaseAccount = &types.BaseAccount{}
			}
			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MainAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MainAddress = append(m.MainAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.MainAddress == nil {
				m.MainAddress = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v github_com_quantumsys_quantmint_types.PermValue
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= github_com_quantumsys_quantmint_types.PermValue(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Perms = append(m.Perms, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTypes
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTypes
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Perms) == 0 {
					m.Perms = make([]github_com_quantumsys_quantmint_types.PermValue, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v github_com_quantumsys_quantmint_types.PermValue
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= github_com_quantumsys_quantmint_types.PermValue(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Perms = append(m.Perms, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Perms", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
