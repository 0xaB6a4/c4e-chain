// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cfevesting/vesting_types.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type VestingTypes struct {
	VestingTypes []*VestingType `protobuf:"bytes,1,rep,name=vesting_types,json=vestingTypes,proto3" json:"vesting_types,omitempty"`
}

func (m *VestingTypes) Reset()         { *m = VestingTypes{} }
func (m *VestingTypes) String() string { return proto.CompactTextString(m) }
func (*VestingTypes) ProtoMessage()    {}
func (*VestingTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_67b318d31cacf480, []int{0}
}
func (m *VestingTypes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VestingTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VestingTypes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VestingTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VestingTypes.Merge(m, src)
}
func (m *VestingTypes) XXX_Size() int {
	return m.Size()
}
func (m *VestingTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_VestingTypes.DiscardUnknown(m)
}

var xxx_messageInfo_VestingTypes proto.InternalMessageInfo

func (m *VestingTypes) GetVestingTypes() []*VestingType {
	if m != nil {
		return m.VestingTypes
	}
	return nil
}

type VestingType struct {
	// vesting type name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// period of locked coins (minutes) from vesting start
	LockupPeriod time.Duration `protobuf:"bytes,2,opt,name=lockup_period,json=lockupPeriod,proto3,stdduration" json:"lockup_period"`
	// period of vesting coins (minutes) from lockup period end
	VestingPeriod time.Duration `protobuf:"bytes,3,opt,name=vesting_period,json=vestingPeriod,proto3,stdduration" json:"vesting_period"`
}

func (m *VestingType) Reset()         { *m = VestingType{} }
func (m *VestingType) String() string { return proto.CompactTextString(m) }
func (*VestingType) ProtoMessage()    {}
func (*VestingType) Descriptor() ([]byte, []int) {
	return fileDescriptor_67b318d31cacf480, []int{1}
}
func (m *VestingType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VestingType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VestingType.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VestingType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VestingType.Merge(m, src)
}
func (m *VestingType) XXX_Size() int {
	return m.Size()
}
func (m *VestingType) XXX_DiscardUnknown() {
	xxx_messageInfo_VestingType.DiscardUnknown(m)
}

var xxx_messageInfo_VestingType proto.InternalMessageInfo

func (m *VestingType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VestingType) GetLockupPeriod() time.Duration {
	if m != nil {
		return m.LockupPeriod
	}
	return 0
}

func (m *VestingType) GetVestingPeriod() time.Duration {
	if m != nil {
		return m.VestingPeriod
	}
	return 0
}

func init() {
	proto.RegisterType((*VestingTypes)(nil), "chain4energy.c4echain.cfevesting.VestingTypes")
	proto.RegisterType((*VestingType)(nil), "chain4energy.c4echain.cfevesting.VestingType")
}

func init() { proto.RegisterFile("cfevesting/vesting_types.proto", fileDescriptor_67b318d31cacf480) }

var fileDescriptor_67b318d31cacf480 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x4e, 0x4b, 0x2d,
	0x4b, 0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0xd7, 0x87, 0xd2, 0xf1, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x7a,
	0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x0a, 0xc9, 0x19, 0x89, 0x99, 0x79, 0x26, 0xa9, 0x79, 0xa9,
	0x45, 0xe9, 0x95, 0x7a, 0xc9, 0x26, 0xa9, 0x60, 0xbe, 0x1e, 0x42, 0x97, 0x94, 0x5c, 0x7a, 0x7e,
	0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x58, 0x7d, 0x52, 0x69, 0x9a, 0x7e, 0x4a, 0x69, 0x51, 0x62, 0x49,
	0x66, 0x7e, 0x1e, 0xc4, 0x04, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0x30, 0x53, 0x1f, 0xc4, 0x82,
	0x88, 0x2a, 0x25, 0x71, 0xf1, 0x84, 0x41, 0x0c, 0x08, 0x01, 0xd9, 0x26, 0x14, 0xc4, 0xc5, 0x8b,
	0x62, 0xbd, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0xae, 0x1e, 0x21, 0xfb, 0xf5, 0x90, 0x8c,
	0x09, 0xe2, 0x29, 0x43, 0x32, 0x53, 0x69, 0x3b, 0x23, 0x17, 0x37, 0x92, 0xac, 0x90, 0x10, 0x17,
	0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d, 0xe4, 0xc1,
	0xc5, 0x9b, 0x93, 0x9f, 0x9c, 0x5d, 0x5a, 0x10, 0x5f, 0x90, 0x5a, 0x94, 0x99, 0x9f, 0x22, 0xc1,
	0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa9, 0x07, 0xf1, 0x95, 0x1e, 0xcc, 0x57, 0x7a, 0x2e, 0x50,
	0x5f, 0x39, 0x71, 0x9c, 0xb8, 0x27, 0xcf, 0x30, 0xe3, 0xbe, 0x3c, 0x63, 0x10, 0x0f, 0x44, 0x67,
	0x00, 0x58, 0xa3, 0x90, 0x17, 0x17, 0x1f, 0xcc, 0x07, 0x50, 0xa3, 0x98, 0x89, 0x37, 0x0a, 0xe6,
	0x79, 0x88, 0x59, 0x4e, 0x7e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65,
	0x92, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x8f, 0x1c, 0x34, 0xfa, 0xc9,
	0x26, 0xa9, 0xba, 0x60, 0x01, 0xfd, 0x0a, 0x7d, 0xa4, 0x38, 0x05, 0x07, 0x66, 0x12, 0x1b, 0xd8,
	0x6e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xac, 0xdf, 0x1b, 0xee, 0x01, 0x00, 0x00,
}

func (m *VestingTypes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VestingTypes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VestingTypes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingTypes) > 0 {
		for iNdEx := len(m.VestingTypes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingTypes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVestingTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *VestingType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VestingType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VestingType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.VestingPeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.VestingPeriod):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintVestingTypes(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.LockupPeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.LockupPeriod):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintVestingTypes(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintVestingTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVestingTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovVestingTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VestingTypes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.VestingTypes) > 0 {
		for _, e := range m.VestingTypes {
			l = e.Size()
			n += 1 + l + sovVestingTypes(uint64(l))
		}
	}
	return n
}

func (m *VestingType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovVestingTypes(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.LockupPeriod)
	n += 1 + l + sovVestingTypes(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.VestingPeriod)
	n += 1 + l + sovVestingTypes(uint64(l))
	return n
}

func sovVestingTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVestingTypes(x uint64) (n int) {
	return sovVestingTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VestingTypes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVestingTypes
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
			return fmt.Errorf("proto: VestingTypes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VestingTypes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingTypes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingTypes
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
				return ErrInvalidLengthVestingTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVestingTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingTypes = append(m.VestingTypes, &VestingType{})
			if err := m.VestingTypes[len(m.VestingTypes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVestingTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVestingTypes
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
func (m *VestingType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVestingTypes
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
			return fmt.Errorf("proto: VestingType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VestingType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingTypes
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
				return ErrInvalidLengthVestingTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVestingTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockupPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingTypes
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
				return ErrInvalidLengthVestingTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVestingTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.LockupPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingTypes
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
				return ErrInvalidLengthVestingTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVestingTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.VestingPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVestingTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVestingTypes
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
func skipVestingTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVestingTypes
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
					return 0, ErrIntOverflowVestingTypes
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
					return 0, ErrIntOverflowVestingTypes
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
				return 0, ErrInvalidLengthVestingTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVestingTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVestingTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVestingTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVestingTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVestingTypes = fmt.Errorf("proto: unexpected end of group")
)
