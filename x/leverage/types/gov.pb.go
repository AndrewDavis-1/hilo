// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hilo/leverage/v1beta1/gov.proto

package types

import (
	fmt "fmt"
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

// UpdateRegistryProposal defines a governance proposal type where the token
// registry can be updated in the Hilo capital facility. Note, the registry
// defined in the proposal replaces the current registry in its entirety.
type UpdateRegistryProposal struct {
	Title       string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Registry    []Token `protobuf:"bytes,3,rep,name=registry,proto3" json:"registry"`
}

func (m *UpdateRegistryProposal) Reset()      { *m = UpdateRegistryProposal{} }
func (*UpdateRegistryProposal) ProtoMessage() {}
func (*UpdateRegistryProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_d46e4b240ab13a5c, []int{0}
}
func (m *UpdateRegistryProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateRegistryProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateRegistryProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateRegistryProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRegistryProposal.Merge(m, src)
}
func (m *UpdateRegistryProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateRegistryProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRegistryProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRegistryProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UpdateRegistryProposal)(nil), "hilonetwork.hilo.leverage.v1beta1.UpdateRegistryProposal")
}

func init() { proto.RegisterFile("hilo/leverage/v1beta1/gov.proto", fileDescriptor_d46e4b240ab13a5c) }

var fileDescriptor_d46e4b240ab13a5c = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0xcd, 0x4d, 0x4d,
	0xd5, 0xcf, 0x49, 0x2d, 0x4b, 0x2d, 0x4a, 0x4c, 0x4f, 0xd5, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49,
	0x34, 0xd4, 0x4f, 0xcf, 0x2f, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x04, 0x29, 0xc8,
	0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x03, 0xb1, 0xf5, 0x60, 0x8a, 0xf5, 0xa0, 0x8a, 0xa5,
	0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xaa, 0xf5, 0x41, 0x2c, 0x88, 0x46, 0x29, 0x15, 0xec, 0x26,
	0xc3, 0x75, 0x83, 0x55, 0x29, 0xad, 0x62, 0xe4, 0x12, 0x0b, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0x0d,
	0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x29, 0xaa, 0x0c, 0x28, 0xca, 0x2f, 0xc8, 0x2f, 0x4e, 0xcc, 0x11,
	0x12, 0xe1, 0x62, 0x2d, 0xc9, 0x2c, 0xc9, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82,
	0x70, 0x84, 0x14, 0xb8, 0xb8, 0x53, 0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0x32, 0xf3, 0xf3,
	0x24, 0x98, 0xc0, 0x72, 0xc8, 0x42, 0x42, 0x5e, 0x5c, 0x1c, 0x45, 0x50, 0xb3, 0x24, 0x98, 0x15,
	0x98, 0x35, 0xb8, 0x8d, 0x34, 0xf4, 0x08, 0x7a, 0x42, 0x2f, 0x24, 0x3f, 0x3b, 0x35, 0xcf, 0x89,
	0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20, 0xb8, 0x7e, 0x2b, 0x9e, 0x8e, 0x05, 0xf2, 0x0c, 0x33, 0x16,
	0xc8, 0x33, 0xbc, 0x58, 0x20, 0xcf, 0xe8, 0xe4, 0x7e, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72,
	0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7,
	0x72, 0x0c, 0x51, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x20,
	0xf3, 0x75, 0xa1, 0x96, 0x81, 0x39, 0xfa, 0x15, 0x88, 0x60, 0x28, 0xa9, 0x2c, 0x48, 0x2d, 0x4e,
	0x62, 0x03, 0x7b, 0xde, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xba, 0xb5, 0xbc, 0x2d, 0x7e, 0x01,
	0x00, 0x00,
}

func (this *UpdateRegistryProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateRegistryProposal)
	if !ok {
		that2, ok := that.(UpdateRegistryProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.Registry) != len(that1.Registry) {
		return false
	}
	for i := range this.Registry {
		if !this.Registry[i].Equal(&that1.Registry[i]) {
			return false
		}
	}
	return true
}
func (m *UpdateRegistryProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateRegistryProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateRegistryProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Registry) > 0 {
		for iNdEx := len(m.Registry) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Registry[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UpdateRegistryProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.Registry) > 0 {
		for _, e := range m.Registry {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UpdateRegistryProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: UpdateRegistryProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateRegistryProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Registry = append(m.Registry, Token{})
			if err := m.Registry[len(m.Registry)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
