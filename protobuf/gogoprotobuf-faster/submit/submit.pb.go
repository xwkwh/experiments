// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: submit.proto

package submit

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

type Request struct {
	Recvtime int64  `protobuf:"varint,1,opt,name=recvtime,proto3" json:"recvtime,omitempty"`
	Uniqueid string `protobuf:"bytes,2,opt,name=uniqueid,proto3" json:"uniqueid,omitempty"`
	Token    string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Content  string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Sign     string `protobuf:"bytes,6,opt,name=sign,proto3" json:"sign,omitempty"`
	Type     string `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Extend   string `protobuf:"bytes,8,opt,name=extend,proto3" json:"extend,omitempty"`
	Version  string `protobuf:"bytes,9,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f5b4b0271dafc7e, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetRecvtime() int64 {
	if m != nil {
		return m.Recvtime
	}
	return 0
}

func (m *Request) GetUniqueid() string {
	if m != nil {
		return m.Uniqueid
	}
	return ""
}

func (m *Request) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Request) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Request) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Request) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *Request) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Request) GetExtend() string {
	if m != nil {
		return m.Extend
	}
	return ""
}

func (m *Request) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "submit.request")
}

func init() { proto.RegisterFile("submit.proto", fileDescriptor_0f5b4b0271dafc7e) }

var fileDescriptor_0f5b4b0271dafc7e = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xd0, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x80, 0xe1, 0x98, 0xb6, 0x4e, 0x6a, 0x31, 0x59, 0x08, 0x9d, 0x18, 0xac, 0xaa, 0x53, 0xa7,
	0x2e, 0x8c, 0x6c, 0x3c, 0x42, 0x47, 0xc6, 0xb6, 0x27, 0xb0, 0x50, 0xcf, 0xa9, 0x7d, 0xae, 0xe0,
	0x2d, 0x78, 0x2c, 0xc6, 0x8c, 0x8c, 0x90, 0xbc, 0x08, 0xb2, 0x9d, 0x66, 0xbb, 0xef, 0x3f, 0xe9,
	0x86, 0x53, 0xb7, 0x21, 0xee, 0x4f, 0x96, 0xb7, 0xad, 0x77, 0xec, 0xb4, 0x2c, 0x5a, 0xff, 0x09,
	0x55, 0x7b, 0x3c, 0x47, 0x0c, 0xac, 0x1f, 0x54, 0xe3, 0xf1, 0x70, 0x61, 0x7b, 0x42, 0x10, 0x2b,
	0xb1, 0x99, 0xed, 0x26, 0xa7, 0x5d, 0x24, 0x7b, 0x8e, 0x68, 0x8f, 0x70, 0xb3, 0x12, 0x9b, 0xe5,
	0x6e, 0xb2, 0xbe, 0x53, 0x0b, 0x76, 0xef, 0x48, 0x30, 0xcb, 0x8b, 0x82, 0x54, 0xdb, 0x37, 0x47,
	0x08, 0xf3, 0x52, 0x33, 0x34, 0xa8, 0xfa, 0xe0, 0x88, 0x91, 0x18, 0x16, 0xb9, 0x5f, 0xa9, 0xb5,
	0x9a, 0x07, 0xfb, 0x4a, 0x20, 0x73, 0xce, 0x73, 0x6a, 0xfc, 0xd9, 0x22, 0xd4, 0xa5, 0xa5, 0x59,
	0xdf, 0x2b, 0x89, 0x1f, 0x8c, 0x74, 0x84, 0x26, 0xd7, 0x51, 0xe9, 0xf2, 0x05, 0x7d, 0xb0, 0x8e,
	0x60, 0x59, 0x2e, 0x8f, 0x7c, 0x5e, 0x7f, 0xf7, 0x46, 0x74, 0xbd, 0x11, 0xbf, 0xbd, 0x11, 0x5f,
	0x83, 0xa9, 0xba, 0xc1, 0x54, 0x3f, 0x83, 0xa9, 0x5e, 0x9a, 0xed, 0x53, 0xf9, 0xc3, 0x5e, 0xe6,
	0xb7, 0x3c, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x7c, 0x66, 0xac, 0x26, 0x01, 0x00, 0x00,
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintSubmit(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Extend) > 0 {
		i -= len(m.Extend)
		copy(dAtA[i:], m.Extend)
		i = encodeVarintSubmit(dAtA, i, uint64(len(m.Extend)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintSubmit(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Sign) > 0 {
		i -= len(m.Sign)
		copy(dAtA[i:], m.Sign)
		i = encodeVarintSubmit(dAtA, i, uint64(len(m.Sign)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintSubmit(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Phone) > 0 {
		i -= len(m.Phone)
		copy(dAtA[i:], m.Phone)
		i = encodeVarintSubmit(dAtA, i, uint64(len(m.Phone)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintSubmit(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Uniqueid) > 0 {
		i -= len(m.Uniqueid)
		copy(dAtA[i:], m.Uniqueid)
		i = encodeVarintSubmit(dAtA, i, uint64(len(m.Uniqueid)))
		i--
		dAtA[i] = 0x12
	}
	if m.Recvtime != 0 {
		i = encodeVarintSubmit(dAtA, i, uint64(m.Recvtime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSubmit(dAtA []byte, offset int, v uint64) int {
	offset -= sovSubmit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Recvtime != 0 {
		n += 1 + sovSubmit(uint64(m.Recvtime))
	}
	l = len(m.Uniqueid)
	if l > 0 {
		n += 1 + l + sovSubmit(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovSubmit(uint64(l))
	}
	l = len(m.Phone)
	if l > 0 {
		n += 1 + l + sovSubmit(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovSubmit(uint64(l))
	}
	l = len(m.Sign)
	if l > 0 {
		n += 1 + l + sovSubmit(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovSubmit(uint64(l))
	}
	l = len(m.Extend)
	if l > 0 {
		n += 1 + l + sovSubmit(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovSubmit(uint64(l))
	}
	return n
}

func sovSubmit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSubmit(x uint64) (n int) {
	return sovSubmit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubmit
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
			return fmt.Errorf("proto: request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recvtime", wireType)
			}
			m.Recvtime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubmit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Recvtime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uniqueid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubmit
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
				return ErrInvalidLengthSubmit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubmit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uniqueid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubmit
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
				return ErrInvalidLengthSubmit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubmit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubmit
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
				return ErrInvalidLengthSubmit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubmit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Phone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubmit
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
				return ErrInvalidLengthSubmit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubmit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sign", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubmit
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
				return ErrInvalidLengthSubmit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubmit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sign = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubmit
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
				return ErrInvalidLengthSubmit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubmit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extend", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubmit
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
				return ErrInvalidLengthSubmit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubmit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extend = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubmit
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
				return ErrInvalidLengthSubmit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubmit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubmit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSubmit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSubmit
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
func skipSubmit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSubmit
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
					return 0, ErrIntOverflowSubmit
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
					return 0, ErrIntOverflowSubmit
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
				return 0, ErrInvalidLengthSubmit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSubmit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSubmit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSubmit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSubmit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSubmit = fmt.Errorf("proto: unexpected end of group")
)