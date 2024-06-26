// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymension/eibc/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_39216408ffc70aa8, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39216408ffc70aa8, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryGetDemandOrderRequest is the request type for the Query/GetDemandOrder RPC method.
type QueryGetDemandOrderRequest struct {
	// id of the demand order to get
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetDemandOrderRequest) Reset()         { *m = QueryGetDemandOrderRequest{} }
func (m *QueryGetDemandOrderRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetDemandOrderRequest) ProtoMessage()    {}
func (*QueryGetDemandOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_39216408ffc70aa8, []int{2}
}
func (m *QueryGetDemandOrderRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDemandOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDemandOrderRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDemandOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDemandOrderRequest.Merge(m, src)
}
func (m *QueryGetDemandOrderRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDemandOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDemandOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDemandOrderRequest proto.InternalMessageInfo

func (m *QueryGetDemandOrderRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// QueryDemandOrdersByStatusRequest is the request type for the Query/GetDemandOrdersByStatus RPC method.
type QueryDemandOrdersByStatusRequest struct {
	// id of the demand order to get
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *QueryDemandOrdersByStatusRequest) Reset()         { *m = QueryDemandOrdersByStatusRequest{} }
func (m *QueryDemandOrdersByStatusRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDemandOrdersByStatusRequest) ProtoMessage()    {}
func (*QueryDemandOrdersByStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_39216408ffc70aa8, []int{3}
}
func (m *QueryDemandOrdersByStatusRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDemandOrdersByStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDemandOrdersByStatusRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDemandOrdersByStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDemandOrdersByStatusRequest.Merge(m, src)
}
func (m *QueryDemandOrdersByStatusRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDemandOrdersByStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDemandOrdersByStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDemandOrdersByStatusRequest proto.InternalMessageInfo

func (m *QueryDemandOrdersByStatusRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

// QueryGetDemandOrderResponse is the response type for the Query/GetDemandOrder RPC method.
type QueryGetDemandOrderResponse struct {
	// demand order with the given id
	DemandOrder *DemandOrder `protobuf:"bytes,1,opt,name=demand_order,json=demandOrder,proto3" json:"demand_order,omitempty"`
}

func (m *QueryGetDemandOrderResponse) Reset()         { *m = QueryGetDemandOrderResponse{} }
func (m *QueryGetDemandOrderResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetDemandOrderResponse) ProtoMessage()    {}
func (*QueryGetDemandOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39216408ffc70aa8, []int{4}
}
func (m *QueryGetDemandOrderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDemandOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDemandOrderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDemandOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDemandOrderResponse.Merge(m, src)
}
func (m *QueryGetDemandOrderResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDemandOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDemandOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDemandOrderResponse proto.InternalMessageInfo

func (m *QueryGetDemandOrderResponse) GetDemandOrder() *DemandOrder {
	if m != nil {
		return m.DemandOrder
	}
	return nil
}

// QueryDemandOrdersByStatusResponse is the response type for the Query/GetDemandOrdersByStatus RPC method.
type QueryDemandOrdersByStatusResponse struct {
	// A list of demand orders with the given status
	DemandOrders []*DemandOrder `protobuf:"bytes,1,rep,name=demand_orders,json=demandOrders,proto3" json:"demand_orders,omitempty"`
}

func (m *QueryDemandOrdersByStatusResponse) Reset()         { *m = QueryDemandOrdersByStatusResponse{} }
func (m *QueryDemandOrdersByStatusResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDemandOrdersByStatusResponse) ProtoMessage()    {}
func (*QueryDemandOrdersByStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39216408ffc70aa8, []int{5}
}
func (m *QueryDemandOrdersByStatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDemandOrdersByStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDemandOrdersByStatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDemandOrdersByStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDemandOrdersByStatusResponse.Merge(m, src)
}
func (m *QueryDemandOrdersByStatusResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDemandOrdersByStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDemandOrdersByStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDemandOrdersByStatusResponse proto.InternalMessageInfo

func (m *QueryDemandOrdersByStatusResponse) GetDemandOrders() []*DemandOrder {
	if m != nil {
		return m.DemandOrders
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "dymensionxyz.dymension.eibc.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "dymensionxyz.dymension.eibc.QueryParamsResponse")
	proto.RegisterType((*QueryGetDemandOrderRequest)(nil), "dymensionxyz.dymension.eibc.QueryGetDemandOrderRequest")
	proto.RegisterType((*QueryDemandOrdersByStatusRequest)(nil), "dymensionxyz.dymension.eibc.QueryDemandOrdersByStatusRequest")
	proto.RegisterType((*QueryGetDemandOrderResponse)(nil), "dymensionxyz.dymension.eibc.QueryGetDemandOrderResponse")
	proto.RegisterType((*QueryDemandOrdersByStatusResponse)(nil), "dymensionxyz.dymension.eibc.QueryDemandOrdersByStatusResponse")
}

func init() { proto.RegisterFile("dymension/eibc/query.proto", fileDescriptor_39216408ffc70aa8) }

var fileDescriptor_39216408ffc70aa8 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6a, 0x14, 0x31,
	0x1c, 0xdf, 0x59, 0xeb, 0x82, 0xff, 0x56, 0x85, 0xb8, 0x88, 0xcc, 0xca, 0xd8, 0x4e, 0x11, 0x8a,
	0xca, 0xa4, 0xdb, 0xc5, 0x0f, 0x14, 0x05, 0x17, 0x41, 0xa4, 0x88, 0xba, 0x5e, 0xc4, 0x8b, 0x64,
	0x36, 0x61, 0x8c, 0x38, 0x93, 0xe9, 0x24, 0x5b, 0x3a, 0x96, 0x5e, 0x7c, 0x02, 0xc1, 0x8b, 0x4f,
	0xe2, 0xc9, 0x07, 0xe8, 0xb1, 0xe8, 0xc5, 0x93, 0xc8, 0xae, 0x0f, 0x22, 0x93, 0xa4, 0x35, 0xd6,
	0x75, 0xfa, 0x71, 0xdb, 0x24, 0xbf, 0xcf, 0xe4, 0xbf, 0x03, 0x3e, 0x2d, 0x53, 0x96, 0x49, 0x2e,
	0x32, 0xcc, 0x78, 0x3c, 0xc4, 0x6b, 0x23, 0x56, 0x94, 0x51, 0x5e, 0x08, 0x25, 0x50, 0x67, 0xef,
	0x6c, 0xa3, 0x7c, 0x17, 0xed, 0x2d, 0xa2, 0x0a, 0xe8, 0xb7, 0x13, 0x91, 0x08, 0x8d, 0xc3, 0xd5,
	0x2f, 0x43, 0xf1, 0x2f, 0x26, 0x42, 0x24, 0x6f, 0x19, 0x26, 0x39, 0xc7, 0x24, 0xcb, 0x84, 0x22,
	0x8a, 0x8b, 0x4c, 0xda, 0xd3, 0x2b, 0x43, 0x21, 0x53, 0x21, 0x71, 0x4c, 0x24, 0x33, 0x4e, 0x78,
	0xbd, 0x1b, 0x33, 0x45, 0xba, 0x38, 0x27, 0x09, 0xcf, 0x34, 0xd8, 0x62, 0x3b, 0xfb, 0x82, 0xe5,
	0xa4, 0x20, 0xe9, 0xae, 0xd0, 0xc2, 0xbe, 0x43, 0xca, 0x52, 0x92, 0xd1, 0x57, 0xa2, 0xa0, 0xac,
	0x30, 0x90, 0xb0, 0x0d, 0xe8, 0x59, 0xe5, 0xf0, 0x54, 0xf3, 0x06, 0x6c, 0x6d, 0xc4, 0xa4, 0x0a,
	0x5f, 0xc0, 0xb9, 0xbf, 0x76, 0x65, 0x2e, 0x32, 0xc9, 0xd0, 0x7d, 0x68, 0x19, 0xfd, 0x0b, 0xde,
	0xbc, 0xb7, 0x34, 0xbb, 0xb2, 0x18, 0xd5, 0x54, 0x8f, 0x0c, 0xb9, 0x3f, 0xb3, 0xfd, 0xe3, 0x52,
	0x63, 0x60, 0x89, 0xe1, 0x35, 0xf0, 0xb5, 0xf2, 0x43, 0xa6, 0x1e, 0xe8, 0x34, 0x4f, 0xaa, 0x30,
	0xd6, 0x17, 0x9d, 0x81, 0x26, 0xa7, 0x5a, 0xfc, 0xd4, 0xa0, 0xc9, 0x69, 0x78, 0x1b, 0xe6, 0x35,
	0xda, 0x81, 0xca, 0x7e, 0xf9, 0x5c, 0x11, 0x35, 0xda, 0xcd, 0x8a, 0xce, 0x43, 0x4b, 0xea, 0x0d,
	0xcb, 0xb3, 0xab, 0xf0, 0x0d, 0x74, 0xa6, 0x3a, 0xd9, 0x2e, 0xab, 0x30, 0xe7, 0x5e, 0x87, 0x6d,
	0xb4, 0x54, 0xdb, 0xc8, 0xd5, 0x99, 0xa5, 0x7f, 0x16, 0x61, 0x01, 0x0b, 0x35, 0x39, 0xad, 0xe3,
	0x63, 0x38, 0xed, 0x3a, 0x56, 0x79, 0x4f, 0x1c, 0xc9, 0x72, 0xce, 0xb1, 0x94, 0x2b, 0x9f, 0x67,
	0xe0, 0xa4, 0x36, 0x45, 0x9f, 0x3c, 0x68, 0x99, 0xcb, 0x46, 0xb8, 0x56, 0xec, 0xdf, 0x97, 0xf6,
	0x97, 0x0f, 0x4f, 0x30, 0x35, 0xc2, 0xab, 0xef, 0xbf, 0xfd, 0xfa, 0xd8, 0xbc, 0x8c, 0x16, 0xb1,
	0xcb, 0xc4, 0x53, 0xe7, 0x10, 0x7d, 0xf1, 0xe0, 0xac, 0x53, 0xa1, 0x5f, 0x3e, 0xa2, 0xe8, 0xe6,
	0xc1, 0x96, 0x53, 0xa7, 0xc3, 0xbf, 0x75, 0x74, 0xa2, 0xcd, 0x7c, 0x43, 0x67, 0x5e, 0x46, 0x51,
	0x6d, 0x66, 0xf7, 0x75, 0xf0, 0x26, 0xa7, 0x5b, 0xe8, 0xab, 0x07, 0xed, 0x69, 0x6f, 0x8a, 0xee,
	0x1e, 0x1c, 0xa5, 0x66, 0x66, 0xfd, 0x7b, 0xc7, 0xa5, 0xdb, 0x3e, 0x77, 0x74, 0x9f, 0xeb, 0xa8,
	0x77, 0xe8, 0x3e, 0x12, 0x6f, 0x9a, 0xff, 0xc5, 0x56, 0x7f, 0x75, 0x7b, 0x1c, 0x78, 0x3b, 0xe3,
	0xc0, 0xfb, 0x39, 0x0e, 0xbc, 0x0f, 0x93, 0xa0, 0xb1, 0x33, 0x09, 0x1a, 0xdf, 0x27, 0x41, 0xe3,
	0x65, 0x37, 0xe1, 0xea, 0xf5, 0x28, 0x8e, 0x86, 0x22, 0xfd, 0x9f, 0xf0, 0x7a, 0x0f, 0x6f, 0x18,
	0x75, 0x55, 0xe6, 0x4c, 0xc6, 0x2d, 0xfd, 0x19, 0xe9, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x37,
	0x0c, 0x43, 0xa0, 0x21, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a Demand Order by id.
	DemandOrderById(ctx context.Context, in *QueryGetDemandOrderRequest, opts ...grpc.CallOption) (*QueryGetDemandOrderResponse, error)
	// Queries a list of demand orders by status.
	DemandOrdersByStatus(ctx context.Context, in *QueryDemandOrdersByStatusRequest, opts ...grpc.CallOption) (*QueryDemandOrdersByStatusResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.eibc.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DemandOrderById(ctx context.Context, in *QueryGetDemandOrderRequest, opts ...grpc.CallOption) (*QueryGetDemandOrderResponse, error) {
	out := new(QueryGetDemandOrderResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.eibc.Query/DemandOrderById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DemandOrdersByStatus(ctx context.Context, in *QueryDemandOrdersByStatusRequest, opts ...grpc.CallOption) (*QueryDemandOrdersByStatusResponse, error) {
	out := new(QueryDemandOrdersByStatusResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.eibc.Query/DemandOrdersByStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a Demand Order by id.
	DemandOrderById(context.Context, *QueryGetDemandOrderRequest) (*QueryGetDemandOrderResponse, error)
	// Queries a list of demand orders by status.
	DemandOrdersByStatus(context.Context, *QueryDemandOrdersByStatusRequest) (*QueryDemandOrdersByStatusResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) DemandOrderById(ctx context.Context, req *QueryGetDemandOrderRequest) (*QueryGetDemandOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DemandOrderById not implemented")
}
func (*UnimplementedQueryServer) DemandOrdersByStatus(ctx context.Context, req *QueryDemandOrdersByStatusRequest) (*QueryDemandOrdersByStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DemandOrdersByStatus not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.eibc.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DemandOrderById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetDemandOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DemandOrderById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.eibc.Query/DemandOrderById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DemandOrderById(ctx, req.(*QueryGetDemandOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DemandOrdersByStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDemandOrdersByStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DemandOrdersByStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.eibc.Query/DemandOrdersByStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DemandOrdersByStatus(ctx, req.(*QueryDemandOrdersByStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dymensionxyz.dymension.eibc.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "DemandOrderById",
			Handler:    _Query_DemandOrderById_Handler,
		},
		{
			MethodName: "DemandOrdersByStatus",
			Handler:    _Query_DemandOrdersByStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dymension/eibc/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetDemandOrderRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDemandOrderRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDemandOrderRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDemandOrdersByStatusRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDemandOrdersByStatusRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDemandOrdersByStatusRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetDemandOrderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDemandOrderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDemandOrderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DemandOrder != nil {
		{
			size, err := m.DemandOrder.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDemandOrdersByStatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDemandOrdersByStatusResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDemandOrdersByStatusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DemandOrders) > 0 {
		for iNdEx := len(m.DemandOrders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DemandOrders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetDemandOrderRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDemandOrdersByStatusRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetDemandOrderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DemandOrder != nil {
		l = m.DemandOrder.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDemandOrdersByStatusResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DemandOrders) > 0 {
		for _, e := range m.DemandOrders {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetDemandOrderRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetDemandOrderRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDemandOrderRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDemandOrdersByStatusRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDemandOrdersByStatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDemandOrdersByStatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetDemandOrderResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetDemandOrderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDemandOrderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DemandOrder", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DemandOrder == nil {
				m.DemandOrder = &DemandOrder{}
			}
			if err := m.DemandOrder.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDemandOrdersByStatusResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDemandOrdersByStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDemandOrdersByStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DemandOrders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DemandOrders = append(m.DemandOrders, &DemandOrder{})
			if err := m.DemandOrders[len(m.DemandOrders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
