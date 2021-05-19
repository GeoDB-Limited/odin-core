// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auction/params.proto

package types

import (
	fmt "fmt"
	types1 "github.com/GeoDB-Limited/odin-core/x/coinswap/types"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Params is the data structure that keeps the parameters of the auction module.
type Params struct {
	// Threshold is the threshold at which the auction starts
	Threshold github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=threshold,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"threshold"`
	// ExchangeRate is a rate for buying coins throw the auction
	ExchangeRate types1.Exchange `protobuf:"bytes,2,opt,name=exchange_rate,json=exchangeRate,proto3" json:"exchange_rate"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_e96a95233ccbd0c2, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetThreshold() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Threshold
	}
	return nil
}

func (m *Params) GetExchangeRate() types1.Exchange {
	if m != nil {
		return m.ExchangeRate
	}
	return types1.Exchange{}
}

func init() {
	proto.RegisterType((*Params)(nil), "auction.Params")
}

func init() { proto.RegisterFile("auction/params.proto", fileDescriptor_e96a95233ccbd0c2) }

var fileDescriptor_e96a95233ccbd0c2 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x13, 0xa8, 0x8a, 0x48, 0x61, 0x89, 0x8a, 0xd4, 0x76, 0x70, 0x2b, 0xa6, 0x2e, 0xf5,
	0xa3, 0x65, 0x43, 0x62, 0x09, 0x20, 0x96, 0x0e, 0xa8, 0x23, 0x0b, 0x72, 0x92, 0x47, 0x62, 0xd1,
	0xe4, 0x45, 0xb1, 0x0b, 0xed, 0x2d, 0x38, 0x02, 0x33, 0x27, 0xe0, 0x08, 0x1d, 0x3b, 0x32, 0x01,
	0x6a, 0x17, 0x8e, 0x81, 0x92, 0x3a, 0x02, 0x26, 0xdb, 0xff, 0xff, 0xdb, 0xdf, 0x2f, 0x3f, 0xa7,
	0x29, 0x66, 0x81, 0x96, 0x94, 0x42, 0x26, 0x72, 0x91, 0x28, 0x9e, 0xe5, 0xa4, 0xc9, 0xdd, 0x33,
	0x6a, 0xa7, 0x19, 0x51, 0x44, 0xa5, 0x06, 0xc5, 0x6e, 0x6b, 0x77, 0x58, 0x40, 0x2a, 0x21, 0x05,
	0xbe, 0x50, 0x08, 0x8f, 0x43, 0x1f, 0xb5, 0x18, 0x42, 0x40, 0x32, 0x35, 0x7e, 0x3b, 0x22, 0x8a,
	0xa6, 0x08, 0xe5, 0xc9, 0x9f, 0xdd, 0x83, 0x48, 0x17, 0xc6, 0x3a, 0x2a, 0x62, 0xea, 0x49, 0x64,
	0xff, 0x80, 0xc7, 0x6f, 0xb6, 0x53, 0xbf, 0x29, 0x05, 0x57, 0x3a, 0xfb, 0x3a, 0xce, 0x51, 0xc5,
	0x34, 0x0d, 0x5b, 0x76, 0x6f, 0xb7, 0xdf, 0x18, 0xb5, 0xf9, 0x16, 0xc8, 0x0b, 0x20, 0x37, 0x40,
	0x7e, 0x41, 0x32, 0xf5, 0x4e, 0x96, 0x1f, 0x5d, 0xeb, 0xf5, 0xb3, 0xdb, 0x8f, 0xa4, 0x8e, 0x67,
	0x3e, 0x0f, 0x28, 0x01, 0xd3, 0x6e, 0xbb, 0x0c, 0x54, 0xf8, 0x00, 0x7a, 0x91, 0xa1, 0x2a, 0x2f,
	0xa8, 0xc9, 0xef, 0xeb, 0xee, 0xb9, 0x73, 0x88, 0xf3, 0x20, 0x16, 0x69, 0x84, 0x77, 0xb9, 0xd0,
	0xd8, 0xda, 0xe9, 0xd9, 0xfd, 0xc6, 0xc8, 0xe5, 0x55, 0x49, 0x7e, 0x65, 0x6c, 0xaf, 0x56, 0x70,
	0x26, 0x07, 0x55, 0x7c, 0x22, 0x34, 0x9e, 0xd5, 0xbe, 0x5f, 0xba, 0x96, 0x37, 0x5e, 0xae, 0x99,
	0xbd, 0x5a, 0x33, 0xfb, 0x6b, 0xcd, 0xec, 0xe7, 0x0d, 0xb3, 0x56, 0x1b, 0x66, 0xbd, 0x6f, 0x98,
	0x75, 0x3b, 0xfa, 0xd3, 0xe9, 0x1a, 0xe9, 0xd2, 0x1b, 0x8c, 0x65, 0x22, 0x35, 0x86, 0x40, 0xa1,
	0x4c, 0x07, 0x01, 0xe5, 0x08, 0x73, 0xa8, 0x06, 0x50, 0x76, 0xf4, 0xeb, 0xe5, 0x7f, 0x9c, 0xfe,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x02, 0x62, 0x0e, 0x50, 0x98, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ExchangeRate.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Threshold) > 0 {
		for iNdEx := len(m.Threshold) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Threshold[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Threshold) > 0 {
		for _, e := range m.Threshold {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = m.ExchangeRate.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Threshold = append(m.Threshold, types.Coin{})
			if err := m.Threshold[len(m.Threshold)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExchangeRate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExchangeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
