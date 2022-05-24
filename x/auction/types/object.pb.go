// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auction/object.proto

package types

import (
	fmt "fmt"
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

type Object struct {
	Id               uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Description      string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Owner            string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Value            string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	PrevOwnSign      string `protobuf:"bytes,5,opt,name=prevOwnSign,proto3" json:"prevOwnSign,omitempty"`
	StartBlockHeight uint64 `protobuf:"varint,6,opt,name=startBlockHeight,proto3" json:"startBlockHeight,omitempty"`
	EndBlockHeight   uint64 `protobuf:"varint,7,opt,name=endBlockHeight,proto3" json:"endBlockHeight,omitempty"`
	AuctionDuration  uint64 `protobuf:"varint,8,opt,name=auctionDuration,proto3" json:"auctionDuration,omitempty"`
	MarginBlocks     uint64 `protobuf:"varint,9,opt,name=marginBlocks,proto3" json:"marginBlocks,omitempty"`
	HighestBid       string `protobuf:"bytes,10,opt,name=highestBid,proto3" json:"highestBid,omitempty"`
	HighestBidder    string `protobuf:"bytes,11,opt,name=highestBidder,proto3" json:"highestBidder,omitempty"`
	Creator          string `protobuf:"bytes,12,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_48bb363b70b28b59, []int{0}
}
func (m *Object) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Object.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return m.Size()
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Object) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Object) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Object) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Object) GetPrevOwnSign() string {
	if m != nil {
		return m.PrevOwnSign
	}
	return ""
}

func (m *Object) GetStartBlockHeight() uint64 {
	if m != nil {
		return m.StartBlockHeight
	}
	return 0
}

func (m *Object) GetEndBlockHeight() uint64 {
	if m != nil {
		return m.EndBlockHeight
	}
	return 0
}

func (m *Object) GetAuctionDuration() uint64 {
	if m != nil {
		return m.AuctionDuration
	}
	return 0
}

func (m *Object) GetMarginBlocks() uint64 {
	if m != nil {
		return m.MarginBlocks
	}
	return 0
}

func (m *Object) GetHighestBid() string {
	if m != nil {
		return m.HighestBid
	}
	return ""
}

func (m *Object) GetHighestBidder() string {
	if m != nil {
		return m.HighestBidder
	}
	return ""
}

func (m *Object) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Object)(nil), "deepstack.auction.auction.Object")
}

func init() { proto.RegisterFile("auction/object.proto", fileDescriptor_48bb363b70b28b59) }

var fileDescriptor_48bb363b70b28b59 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xbb, 0x6e, 0x2a, 0x31,
	0x10, 0x86, 0x59, 0xae, 0x87, 0x81, 0xc3, 0x39, 0xb2, 0x28, 0x9c, 0xc6, 0x42, 0x28, 0x8a, 0x50,
	0x14, 0x41, 0x91, 0x37, 0x40, 0x51, 0x94, 0x0e, 0x89, 0x74, 0xe9, 0x8c, 0x3d, 0xda, 0x75, 0x80,
	0xf5, 0xca, 0xeb, 0x85, 0xe4, 0x15, 0x52, 0xe5, 0xb1, 0x52, 0x52, 0xa6, 0x8c, 0xe0, 0x45, 0x22,
	0x86, 0x4b, 0x16, 0x52, 0xd9, 0xff, 0x37, 0x9f, 0x47, 0x1e, 0x0d, 0xb4, 0x65, 0xa6, 0xbc, 0xb1,
	0xf1, 0xc0, 0x4e, 0x9e, 0x51, 0xf9, 0x7e, 0xe2, 0xac, 0xb7, 0xec, 0x42, 0x23, 0x26, 0xa9, 0x97,
	0x6a, 0xda, 0xdf, 0xd7, 0x0f, 0x67, 0xf7, 0xad, 0x04, 0xd5, 0x11, 0xb9, 0xac, 0x05, 0x45, 0xa3,
	0x79, 0xd0, 0x09, 0x7a, 0xe5, 0x71, 0xd1, 0x68, 0xd6, 0x81, 0x86, 0xc6, 0x54, 0x39, 0x93, 0x6c,
	0x4d, 0x5e, 0xec, 0x04, 0xbd, 0xfa, 0x38, 0x8f, 0x58, 0x1b, 0x2a, 0x76, 0x19, 0xa3, 0xe3, 0x25,
	0xaa, 0xed, 0xc2, 0x96, 0x2e, 0xe4, 0x2c, 0x43, 0x5e, 0xde, 0x51, 0x0a, 0xdb, 0x6e, 0x89, 0xc3,
	0xc5, 0x68, 0x19, 0x3f, 0x9a, 0x30, 0xe6, 0x95, 0x5d, 0xb7, 0x1c, 0x62, 0xd7, 0xf0, 0x3f, 0xf5,
	0xd2, 0xf9, 0xe1, 0xcc, 0xaa, 0xe9, 0x03, 0x9a, 0x30, 0xf2, 0xbc, 0x4a, 0xbf, 0xf9, 0xc5, 0xd9,
	0x15, 0xb4, 0x30, 0xd6, 0x79, 0xb3, 0x46, 0xe6, 0x19, 0x65, 0x3d, 0xf8, 0xb7, 0x9f, 0xf4, 0x2e,
	0x73, 0x92, 0xe6, 0xf8, 0x43, 0xe2, 0x39, 0x66, 0x5d, 0x68, 0xce, 0xa5, 0x0b, 0x4d, 0x4c, 0xcf,
	0x53, 0x5e, 0x27, 0xed, 0x84, 0x31, 0x01, 0x10, 0x99, 0x30, 0xc2, 0xd4, 0x0f, 0x8d, 0xe6, 0x40,
	0x23, 0xe4, 0x08, 0xbb, 0x84, 0xbf, 0x3f, 0x49, 0xa3, 0xe3, 0x0d, 0x52, 0x4e, 0x21, 0xe3, 0x50,
	0x53, 0x0e, 0xa5, 0xb7, 0x8e, 0x37, 0xa9, 0x7e, 0x88, 0xc3, 0xfb, 0x8f, 0xb5, 0x08, 0x56, 0x6b,
	0x11, 0x7c, 0xad, 0x45, 0xf0, 0xbe, 0x11, 0x85, 0xd5, 0x46, 0x14, 0x3e, 0x37, 0xa2, 0xf0, 0x74,
	0x13, 0x1a, 0x1f, 0x65, 0x93, 0xbe, 0xb2, 0xf3, 0xc1, 0x71, 0x99, 0x83, 0xc3, 0xb2, 0x5f, 0x8e,
	0x37, 0xff, 0x9a, 0x60, 0x3a, 0xa9, 0xd2, 0xda, 0x6f, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x22,
	0xa0, 0xbe, 0xf5, 0x0e, 0x02, 0x00, 0x00,
}

func (m *Object) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Object) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Object) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintObject(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.HighestBidder) > 0 {
		i -= len(m.HighestBidder)
		copy(dAtA[i:], m.HighestBidder)
		i = encodeVarintObject(dAtA, i, uint64(len(m.HighestBidder)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.HighestBid) > 0 {
		i -= len(m.HighestBid)
		copy(dAtA[i:], m.HighestBid)
		i = encodeVarintObject(dAtA, i, uint64(len(m.HighestBid)))
		i--
		dAtA[i] = 0x52
	}
	if m.MarginBlocks != 0 {
		i = encodeVarintObject(dAtA, i, uint64(m.MarginBlocks))
		i--
		dAtA[i] = 0x48
	}
	if m.AuctionDuration != 0 {
		i = encodeVarintObject(dAtA, i, uint64(m.AuctionDuration))
		i--
		dAtA[i] = 0x40
	}
	if m.EndBlockHeight != 0 {
		i = encodeVarintObject(dAtA, i, uint64(m.EndBlockHeight))
		i--
		dAtA[i] = 0x38
	}
	if m.StartBlockHeight != 0 {
		i = encodeVarintObject(dAtA, i, uint64(m.StartBlockHeight))
		i--
		dAtA[i] = 0x30
	}
	if len(m.PrevOwnSign) > 0 {
		i -= len(m.PrevOwnSign)
		copy(dAtA[i:], m.PrevOwnSign)
		i = encodeVarintObject(dAtA, i, uint64(len(m.PrevOwnSign)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintObject(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintObject(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintObject(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintObject(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintObject(dAtA []byte, offset int, v uint64) int {
	offset -= sovObject(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Object) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovObject(uint64(m.Id))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovObject(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovObject(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovObject(uint64(l))
	}
	l = len(m.PrevOwnSign)
	if l > 0 {
		n += 1 + l + sovObject(uint64(l))
	}
	if m.StartBlockHeight != 0 {
		n += 1 + sovObject(uint64(m.StartBlockHeight))
	}
	if m.EndBlockHeight != 0 {
		n += 1 + sovObject(uint64(m.EndBlockHeight))
	}
	if m.AuctionDuration != 0 {
		n += 1 + sovObject(uint64(m.AuctionDuration))
	}
	if m.MarginBlocks != 0 {
		n += 1 + sovObject(uint64(m.MarginBlocks))
	}
	l = len(m.HighestBid)
	if l > 0 {
		n += 1 + l + sovObject(uint64(l))
	}
	l = len(m.HighestBidder)
	if l > 0 {
		n += 1 + l + sovObject(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func sovObject(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozObject(x uint64) (n int) {
	return sovObject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Object) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: Object: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Object: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevOwnSign", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevOwnSign = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartBlockHeight", wireType)
			}
			m.StartBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlockHeight", wireType)
			}
			m.EndBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionDuration", wireType)
			}
			m.AuctionDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionDuration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarginBlocks", wireType)
			}
			m.MarginBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MarginBlocks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HighestBid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HighestBid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HighestBidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HighestBidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthObject
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
func skipObject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowObject
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
					return 0, ErrIntOverflowObject
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
					return 0, ErrIntOverflowObject
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
				return 0, ErrInvalidLengthObject
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupObject
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthObject
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthObject        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowObject          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupObject = fmt.Errorf("proto: unexpected end of group")
)