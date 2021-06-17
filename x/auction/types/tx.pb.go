// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auction/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgBuyCoins struct {
	From      string     `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To        string     `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount    types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
	Requester string     `protobuf:"bytes,4,opt,name=requester,proto3" json:"requester,omitempty"`
}

func (m *MsgBuyCoins) Reset()         { *m = MsgBuyCoins{} }
func (m *MsgBuyCoins) String() string { return proto.CompactTextString(m) }
func (*MsgBuyCoins) ProtoMessage()    {}
func (*MsgBuyCoins) Descriptor() ([]byte, []int) {
	return fileDescriptor_085dc3627a5d85eb, []int{0}
}
func (m *MsgBuyCoins) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBuyCoins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBuyCoins.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBuyCoins) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBuyCoins.Merge(m, src)
}
func (m *MsgBuyCoins) XXX_Size() int {
	return m.Size()
}
func (m *MsgBuyCoins) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBuyCoins.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBuyCoins proto.InternalMessageInfo

func (m *MsgBuyCoins) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *MsgBuyCoins) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *MsgBuyCoins) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *MsgBuyCoins) GetRequester() string {
	if m != nil {
		return m.Requester
	}
	return ""
}

type MsgBuyCoinsResponse struct {
}

func (m *MsgBuyCoinsResponse) Reset()         { *m = MsgBuyCoinsResponse{} }
func (m *MsgBuyCoinsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgBuyCoinsResponse) ProtoMessage()    {}
func (*MsgBuyCoinsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_085dc3627a5d85eb, []int{1}
}
func (m *MsgBuyCoinsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBuyCoinsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBuyCoinsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBuyCoinsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBuyCoinsResponse.Merge(m, src)
}
func (m *MsgBuyCoinsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgBuyCoinsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBuyCoinsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBuyCoinsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgBuyCoins)(nil), "auction.MsgBuyCoins")
	proto.RegisterType((*MsgBuyCoinsResponse)(nil), "auction.MsgBuyCoinsResponse")
}

func init() { proto.RegisterFile("auction/tx.proto", fileDescriptor_085dc3627a5d85eb) }

var fileDescriptor_085dc3627a5d85eb = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xe3, 0x36, 0xea, 0xf7, 0xd5, 0x95, 0x10, 0x32, 0x45, 0x0a, 0x55, 0xe5, 0x56, 0x9d,
	0xba, 0xd4, 0x56, 0xcb, 0x80, 0xc4, 0xc0, 0x10, 0x40, 0x2c, 0xed, 0xd2, 0x91, 0x2d, 0x49, 0x4d,
	0xf0, 0x10, 0xdf, 0x10, 0x3b, 0xa8, 0x7d, 0x09, 0xc4, 0x23, 0xf0, 0x38, 0x1d, 0x3b, 0x32, 0x21,
	0x94, 0x2c, 0x3c, 0x06, 0xca, 0x1f, 0x4a, 0x07, 0xb6, 0xab, 0xdf, 0xb5, 0xcf, 0x39, 0xf7, 0xe0,
	0x63, 0x2f, 0x0d, 0x8c, 0x04, 0xc5, 0xcd, 0x9a, 0xc5, 0x09, 0x18, 0x20, 0xff, 0x6a, 0xd2, 0xeb,
	0x86, 0x10, 0x42, 0xc9, 0x78, 0x31, 0x55, 0xeb, 0x1e, 0x0d, 0x40, 0x47, 0xa0, 0xb9, 0xef, 0x69,
	0xc1, 0x9f, 0xa7, 0xbe, 0x30, 0xde, 0x94, 0x07, 0x20, 0x55, 0xb5, 0x1f, 0xbd, 0x20, 0xdc, 0x59,
	0xe8, 0xd0, 0x4d, 0x37, 0xd7, 0x20, 0x95, 0x26, 0x04, 0xdb, 0x0f, 0x09, 0x44, 0x0e, 0x1a, 0xa2,
	0x71, 0x7b, 0x59, 0xce, 0xe4, 0x08, 0x37, 0x0c, 0x38, 0x8d, 0x92, 0x34, 0x0c, 0x90, 0x0b, 0xdc,
	0xf2, 0x22, 0x48, 0x95, 0x71, 0x9a, 0x43, 0x34, 0xee, 0xcc, 0xce, 0x58, 0x65, 0xc2, 0x0a, 0x13,
	0x56, 0x9b, 0xb0, 0x42, 0xcf, 0xb5, 0xb7, 0x1f, 0x03, 0x6b, 0x59, 0x3f, 0x27, 0x7d, 0xdc, 0x4e,
	0xc4, 0x53, 0x2a, 0xb4, 0x11, 0x89, 0x63, 0x97, 0x7a, 0xbf, 0xe0, 0xd2, 0xfe, 0x7a, 0x1b, 0xa0,
	0xd1, 0x29, 0x3e, 0x39, 0xc8, 0xb3, 0x14, 0x3a, 0x06, 0xa5, 0xc5, 0xec, 0x16, 0x37, 0x17, 0x3a,
	0x24, 0x57, 0xf8, 0xff, 0x3e, 0x6a, 0x97, 0xd5, 0xa7, 0xb3, 0x83, 0x0f, 0xbd, 0xfe, 0x5f, 0xf4,
	0x47, 0xc6, 0x9d, 0x6f, 0x33, 0x8a, 0x76, 0x19, 0x45, 0x9f, 0x19, 0x45, 0xaf, 0x39, 0xb5, 0x76,
	0x39, 0xb5, 0xde, 0x73, 0x6a, 0xdd, 0xcf, 0x42, 0x69, 0x1e, 0x53, 0x9f, 0x05, 0x10, 0xf1, 0x3b,
	0x01, 0x37, 0xee, 0x64, 0x2e, 0x23, 0x69, 0xc4, 0x8a, 0xc3, 0x4a, 0xaa, 0x49, 0x00, 0x89, 0xe0,
	0x6b, 0xbe, 0xaf, 0x7f, 0x13, 0x0b, 0xed, 0xb7, 0xca, 0x0e, 0xcf, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x20, 0x06, 0x1e, 0xe5, 0x96, 0x01, 0x00, 0x00,
}

func (this *MsgBuyCoins) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgBuyCoins)
	if !ok {
		that2, ok := that.(MsgBuyCoins)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.From != that1.From {
		return false
	}
	if this.To != that1.To {
		return false
	}
	if !this.Amount.Equal(&that1.Amount) {
		return false
	}
	if this.Requester != that1.Requester {
		return false
	}
	return true
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
	BuyCoins(ctx context.Context, in *MsgBuyCoins, opts ...grpc.CallOption) (*MsgBuyCoinsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) BuyCoins(ctx context.Context, in *MsgBuyCoins, opts ...grpc.CallOption) (*MsgBuyCoinsResponse, error) {
	out := new(MsgBuyCoinsResponse)
	err := c.cc.Invoke(ctx, "/auction.Msg/BuyCoins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	BuyCoins(context.Context, *MsgBuyCoins) (*MsgBuyCoinsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) BuyCoins(ctx context.Context, req *MsgBuyCoins) (*MsgBuyCoinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyCoins not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_BuyCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBuyCoins)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BuyCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auction.Msg/BuyCoins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BuyCoins(ctx, req.(*MsgBuyCoins))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auction.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuyCoins",
			Handler:    _Msg_BuyCoins_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auction/tx.proto",
}

func (m *MsgBuyCoins) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBuyCoins) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBuyCoins) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Requester) > 0 {
		i -= len(m.Requester)
		copy(dAtA[i:], m.Requester)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Requester)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintTx(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBuyCoinsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBuyCoinsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBuyCoinsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgBuyCoins) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Requester)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgBuyCoinsResponse) Size() (n int) {
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
func (m *MsgBuyCoins) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgBuyCoins: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBuyCoins: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
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
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
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
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requester", wireType)
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
			m.Requester = string(dAtA[iNdEx:postIndex])
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
func (m *MsgBuyCoinsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgBuyCoinsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBuyCoinsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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