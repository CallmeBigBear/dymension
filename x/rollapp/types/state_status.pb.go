// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymension/rollapp/state_status.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

type StateStatus int32

const (
	// zero-value for status ordering
	STATE_STATUS_UNSPECIFIED StateStatus = 0
	// STATE_STATUS_RECEIVED defines a rollapp state where the state update transaction was published on dYmension chain
	STATE_STATUS_RECEIVED StateStatus = 1
	// STATE_STATUS_FINALIZED defines a rollapp state where the the "Dispute Period" has ended and this state is considered final
	STATE_STATUS_FINALIZED StateStatus = 2
)

var StateStatus_name = map[int32]string{
	0: "STATE_STATUS_UNSPECIFIED",
	1: "STATE_STATUS_RECEIVED",
	2: "STATE_STATUS_FINALIZED",
}

var StateStatus_value = map[string]int32{
	"STATE_STATUS_UNSPECIFIED": 0,
	"STATE_STATUS_RECEIVED":    1,
	"STATE_STATUS_FINALIZED":   2,
}

func (x StateStatus) String() string {
	return proto.EnumName(StateStatus_name, int32(x))
}

func (StateStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b039e93a8767dffd, []int{0}
}

func init() {
	proto.RegisterEnum("dymensionxyz.dymension.rollapp.StateStatus", StateStatus_name, StateStatus_value)
}

func init() {
	proto.RegisterFile("dymension/rollapp/state_status.proto", fileDescriptor_b039e93a8767dffd)
}

var fileDescriptor_b039e93a8767dffd = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xa9, 0xcc, 0x4d,
	0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0xd3, 0x2f, 0xca, 0xcf, 0xc9, 0x49, 0x2c, 0x28, 0xd0, 0x2f, 0x2e,
	0x49, 0x2c, 0x49, 0x8d, 0x07, 0x91, 0xa5, 0xc5, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x72,
	0x70, 0x55, 0x15, 0x95, 0x55, 0x7a, 0x70, 0x8e, 0x1e, 0x54, 0x8b, 0x94, 0x48, 0x7a, 0x7e, 0x7a,
	0x3e, 0x58, 0xa9, 0x3e, 0x88, 0x05, 0xd1, 0xa5, 0x95, 0xc1, 0xc5, 0x1d, 0x0c, 0x32, 0x2b, 0x18,
	0x6c, 0x94, 0x90, 0x0c, 0x97, 0x44, 0x70, 0x88, 0x63, 0x88, 0x6b, 0x3c, 0x88, 0x0c, 0x0d, 0x8e,
	0x0f, 0xf5, 0x0b, 0x0e, 0x70, 0x75, 0xf6, 0x74, 0xf3, 0x74, 0x75, 0x11, 0x60, 0x10, 0x92, 0xe4,
	0x12, 0x45, 0x91, 0x0d, 0x72, 0x75, 0x76, 0xf5, 0x0c, 0x73, 0x75, 0x11, 0x60, 0x14, 0x92, 0xe2,
	0x12, 0x43, 0x91, 0x72, 0xf3, 0xf4, 0x73, 0xf4, 0xf1, 0x8c, 0x72, 0x75, 0x11, 0x60, 0x92, 0x62,
	0xe9, 0x58, 0x2c, 0xc7, 0xe0, 0xe4, 0x7d, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f,
	0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c,
	0x51, 0x86, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc8, 0x9e, 0x40,
	0x70, 0xf4, 0x2b, 0xe0, 0x3e, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xbb, 0xde, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xa1, 0x76, 0x35, 0x1b, 0x01, 0x00, 0x00,
}
