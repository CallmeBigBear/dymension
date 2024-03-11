// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymension/denommetadata/gov_denommetadata.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/bank/types"
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

type CreateDenomMetadataProposal struct {
	Title         string         `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string         `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	TokenMetadata types.Metadata `protobuf:"bytes,3,opt,name=tokenMetadata,proto3" json:"tokenMetadata"`
}

func (m *CreateDenomMetadataProposal) Reset()      { *m = CreateDenomMetadataProposal{} }
func (*CreateDenomMetadataProposal) ProtoMessage() {}
func (*CreateDenomMetadataProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba1b4bedb498208a, []int{0}
}
func (m *CreateDenomMetadataProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateDenomMetadataProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateDenomMetadataProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateDenomMetadataProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDenomMetadataProposal.Merge(m, src)
}
func (m *CreateDenomMetadataProposal) XXX_Size() int {
	return m.Size()
}
func (m *CreateDenomMetadataProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDenomMetadataProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDenomMetadataProposal proto.InternalMessageInfo

type UpdateDenomMetadataProposal struct {
	Title         string         `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string         `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	TokenMetadata types.Metadata `protobuf:"bytes,3,opt,name=tokenMetadata,proto3" json:"tokenMetadata"`
}

func (m *UpdateDenomMetadataProposal) Reset()      { *m = UpdateDenomMetadataProposal{} }
func (*UpdateDenomMetadataProposal) ProtoMessage() {}
func (*UpdateDenomMetadataProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba1b4bedb498208a, []int{1}
}
func (m *UpdateDenomMetadataProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateDenomMetadataProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateDenomMetadataProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateDenomMetadataProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDenomMetadataProposal.Merge(m, src)
}
func (m *UpdateDenomMetadataProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateDenomMetadataProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDenomMetadataProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDenomMetadataProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateDenomMetadataProposal)(nil), "dymensionxyz.dymension.denommetadata.CreateDenomMetadataProposal")
	proto.RegisterType((*UpdateDenomMetadataProposal)(nil), "dymensionxyz.dymension.denommetadata.UpdateDenomMetadataProposal")
}

func init() {
	proto.RegisterFile("dymension/denommetadata/gov_denommetadata.proto", fileDescriptor_ba1b4bedb498208a)
}

var fileDescriptor_ba1b4bedb498208a = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0xa9, 0xcc, 0x4d,
	0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0xd3, 0x4f, 0x49, 0xcd, 0xcb, 0xcf, 0xcd, 0x4d, 0x2d, 0x49, 0x4c,
	0x49, 0x2c, 0x49, 0xd4, 0x4f, 0xcf, 0x2f, 0x8b, 0x47, 0x11, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x52, 0x81, 0x6b, 0xa8, 0xa8, 0xac, 0xd2, 0x83, 0x73, 0xf4, 0x50, 0xd4, 0x4a, 0x89, 0xa4,
	0xe7, 0xa7, 0xe7, 0x83, 0x35, 0xe8, 0x83, 0x58, 0x10, 0xbd, 0x52, 0x72, 0xc9, 0xf9, 0xc5, 0xb9,
	0xf9, 0xc5, 0xfa, 0x49, 0x89, 0x79, 0xd9, 0xfa, 0x65, 0x86, 0x49, 0xa9, 0x25, 0x89, 0x86, 0x60,
	0x0e, 0x44, 0x5e, 0x69, 0x15, 0x23, 0x97, 0xb4, 0x73, 0x51, 0x6a, 0x62, 0x49, 0xaa, 0x0b, 0xc8,
	0x34, 0x5f, 0xa8, 0x69, 0x01, 0x45, 0xf9, 0x05, 0xf9, 0xc5, 0x89, 0x39, 0x42, 0x22, 0x5c, 0xac,
	0x25, 0x99, 0x25, 0x39, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x02,
	0x17, 0x77, 0x4a, 0x6a, 0x71, 0x72, 0x51, 0x66, 0x41, 0x49, 0x66, 0x7e, 0x9e, 0x04, 0x13, 0x58,
	0x0e, 0x59, 0x48, 0xc8, 0x93, 0x8b, 0xb7, 0x24, 0x3f, 0x3b, 0x35, 0x0f, 0x66, 0xa0, 0x04, 0xb3,
	0x02, 0xa3, 0x06, 0xb7, 0x91, 0xac, 0x1e, 0xc4, 0x3d, 0x7a, 0x60, 0x27, 0x40, 0xdd, 0xa3, 0x07,
	0x53, 0xe4, 0xc4, 0x72, 0xe2, 0x9e, 0x3c, 0x43, 0x10, 0xaa, 0x4e, 0x2b, 0x8e, 0x8e, 0x05, 0xf2,
	0x0c, 0x33, 0x16, 0xc8, 0x33, 0x80, 0x1d, 0x1b, 0x5a, 0x90, 0x32, 0x24, 0x1c, 0xeb, 0x14, 0x72,
	0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7,
	0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x56, 0xe9, 0x99, 0x25, 0x19, 0xa5,
	0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc8, 0x51, 0x8b, 0x94, 0x30, 0xca, 0x8c, 0xf5, 0x2b, 0xd0,
	0x52, 0x47, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0xda, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x36, 0xd0, 0x4f, 0xd6, 0x45, 0x02, 0x00, 0x00,
}

func (m *CreateDenomMetadataProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateDenomMetadataProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateDenomMetadataProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TokenMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGovDenommetadata(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGovDenommetadata(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGovDenommetadata(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateDenomMetadataProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateDenomMetadataProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateDenomMetadataProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TokenMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGovDenommetadata(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGovDenommetadata(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGovDenommetadata(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGovDenommetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovGovDenommetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreateDenomMetadataProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGovDenommetadata(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGovDenommetadata(uint64(l))
	}
	l = m.TokenMetadata.Size()
	n += 1 + l + sovGovDenommetadata(uint64(l))
	return n
}

func (m *UpdateDenomMetadataProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGovDenommetadata(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGovDenommetadata(uint64(l))
	}
	l = m.TokenMetadata.Size()
	n += 1 + l + sovGovDenommetadata(uint64(l))
	return n
}

func sovGovDenommetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGovDenommetadata(x uint64) (n int) {
	return sovGovDenommetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateDenomMetadataProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGovDenommetadata
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
			return fmt.Errorf("proto: CreateDenomMetadataProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateDenomMetadataProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovDenommetadata
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
				return ErrInvalidLengthGovDenommetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovDenommetadata
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
					return ErrIntOverflowGovDenommetadata
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
				return ErrInvalidLengthGovDenommetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovDenommetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovDenommetadata
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
				return ErrInvalidLengthGovDenommetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGovDenommetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGovDenommetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGovDenommetadata
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
func (m *UpdateDenomMetadataProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGovDenommetadata
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
			return fmt.Errorf("proto: UpdateDenomMetadataProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateDenomMetadataProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovDenommetadata
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
				return ErrInvalidLengthGovDenommetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovDenommetadata
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
					return ErrIntOverflowGovDenommetadata
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
				return ErrInvalidLengthGovDenommetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovDenommetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovDenommetadata
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
				return ErrInvalidLengthGovDenommetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGovDenommetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGovDenommetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGovDenommetadata
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
func skipGovDenommetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGovDenommetadata
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
					return 0, ErrIntOverflowGovDenommetadata
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
					return 0, ErrIntOverflowGovDenommetadata
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
				return 0, ErrInvalidLengthGovDenommetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGovDenommetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGovDenommetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGovDenommetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGovDenommetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGovDenommetadata = fmt.Errorf("proto: unexpected end of group")
)
