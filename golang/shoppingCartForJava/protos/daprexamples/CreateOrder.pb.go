// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CreateOrder.proto

package daprexamples

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateOrderRequest struct {
	ProductID            string   `protobuf:"bytes,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	Amount               int32    `protobuf:"varint,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
	CustomerID           string   `protobuf:"bytes,3,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderRequest) Reset()         { *m = CreateOrderRequest{} }
func (m *CreateOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrderRequest) ProtoMessage()    {}
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2476da9e0746a307, []int{0}
}

func (m *CreateOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderRequest.Unmarshal(m, b)
}
func (m *CreateOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderRequest.Marshal(b, m, deterministic)
}
func (m *CreateOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderRequest.Merge(m, src)
}
func (m *CreateOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOrderRequest.Size(m)
}
func (m *CreateOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderRequest proto.InternalMessageInfo

func (m *CreateOrderRequest) GetProductID() string {
	if m != nil {
		return m.ProductID
	}
	return ""
}

func (m *CreateOrderRequest) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *CreateOrderRequest) GetCustomerID() string {
	if m != nil {
		return m.CustomerID
	}
	return ""
}

type CreateOrderResponse struct {
	Succeed              bool     `protobuf:"varint,1,opt,name=Succeed,proto3" json:"Succeed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderResponse) Reset()         { *m = CreateOrderResponse{} }
func (m *CreateOrderResponse) String() string { return proto.CompactTextString(m) }
func (*CreateOrderResponse) ProtoMessage()    {}
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2476da9e0746a307, []int{1}
}

func (m *CreateOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderResponse.Unmarshal(m, b)
}
func (m *CreateOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderResponse.Marshal(b, m, deterministic)
}
func (m *CreateOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderResponse.Merge(m, src)
}
func (m *CreateOrderResponse) XXX_Size() int {
	return xxx_messageInfo_CreateOrderResponse.Size(m)
}
func (m *CreateOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderResponse proto.InternalMessageInfo

func (m *CreateOrderResponse) GetSucceed() bool {
	if m != nil {
		return m.Succeed
	}
	return false
}

type RetrieveOrderRequest struct {
	OrderID              string   `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveOrderRequest) Reset()         { *m = RetrieveOrderRequest{} }
func (m *RetrieveOrderRequest) String() string { return proto.CompactTextString(m) }
func (*RetrieveOrderRequest) ProtoMessage()    {}
func (*RetrieveOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2476da9e0746a307, []int{2}
}

func (m *RetrieveOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveOrderRequest.Unmarshal(m, b)
}
func (m *RetrieveOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveOrderRequest.Marshal(b, m, deterministic)
}
func (m *RetrieveOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveOrderRequest.Merge(m, src)
}
func (m *RetrieveOrderRequest) XXX_Size() int {
	return xxx_messageInfo_RetrieveOrderRequest.Size(m)
}
func (m *RetrieveOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveOrderRequest proto.InternalMessageInfo

func (m *RetrieveOrderRequest) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

type RetrieveOrderResponse struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=Order,proto3" json:"Order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveOrderResponse) Reset()         { *m = RetrieveOrderResponse{} }
func (m *RetrieveOrderResponse) String() string { return proto.CompactTextString(m) }
func (*RetrieveOrderResponse) ProtoMessage()    {}
func (*RetrieveOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2476da9e0746a307, []int{3}
}

func (m *RetrieveOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveOrderResponse.Unmarshal(m, b)
}
func (m *RetrieveOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveOrderResponse.Marshal(b, m, deterministic)
}
func (m *RetrieveOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveOrderResponse.Merge(m, src)
}
func (m *RetrieveOrderResponse) XXX_Size() int {
	return xxx_messageInfo_RetrieveOrderResponse.Size(m)
}
func (m *RetrieveOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveOrderResponse proto.InternalMessageInfo

func (m *RetrieveOrderResponse) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type GetOrderListRequest struct {
	CustomerID           string   `protobuf:"bytes,1,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderListRequest) Reset()         { *m = GetOrderListRequest{} }
func (m *GetOrderListRequest) String() string { return proto.CompactTextString(m) }
func (*GetOrderListRequest) ProtoMessage()    {}
func (*GetOrderListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2476da9e0746a307, []int{4}
}

func (m *GetOrderListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderListRequest.Unmarshal(m, b)
}
func (m *GetOrderListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderListRequest.Marshal(b, m, deterministic)
}
func (m *GetOrderListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderListRequest.Merge(m, src)
}
func (m *GetOrderListRequest) XXX_Size() int {
	return xxx_messageInfo_GetOrderListRequest.Size(m)
}
func (m *GetOrderListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderListRequest proto.InternalMessageInfo

func (m *GetOrderListRequest) GetCustomerID() string {
	if m != nil {
		return m.CustomerID
	}
	return ""
}

type GetOrderListResponse struct {
	Orders               []*Order `protobuf:"bytes,1,rep,name=Orders,proto3" json:"Orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderListResponse) Reset()         { *m = GetOrderListResponse{} }
func (m *GetOrderListResponse) String() string { return proto.CompactTextString(m) }
func (*GetOrderListResponse) ProtoMessage()    {}
func (*GetOrderListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2476da9e0746a307, []int{5}
}

func (m *GetOrderListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderListResponse.Unmarshal(m, b)
}
func (m *GetOrderListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderListResponse.Marshal(b, m, deterministic)
}
func (m *GetOrderListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderListResponse.Merge(m, src)
}
func (m *GetOrderListResponse) XXX_Size() int {
	return xxx_messageInfo_GetOrderListResponse.Size(m)
}
func (m *GetOrderListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderListResponse proto.InternalMessageInfo

func (m *GetOrderListResponse) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type Order struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ProductID            string   `protobuf:"bytes,2,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	Amount               int32    `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	CustomerID           string   `protobuf:"bytes,4,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_2476da9e0746a307, []int{6}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Order) GetProductID() string {
	if m != nil {
		return m.ProductID
	}
	return ""
}

func (m *Order) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Order) GetCustomerID() string {
	if m != nil {
		return m.CustomerID
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateOrderRequest)(nil), "daprexamples.CreateOrderRequest")
	proto.RegisterType((*CreateOrderResponse)(nil), "daprexamples.CreateOrderResponse")
	proto.RegisterType((*RetrieveOrderRequest)(nil), "daprexamples.RetrieveOrderRequest")
	proto.RegisterType((*RetrieveOrderResponse)(nil), "daprexamples.RetrieveOrderResponse")
	proto.RegisterType((*GetOrderListRequest)(nil), "daprexamples.GetOrderListRequest")
	proto.RegisterType((*GetOrderListResponse)(nil), "daprexamples.GetOrderListResponse")
	proto.RegisterType((*Order)(nil), "daprexamples.Order")
}

func init() { proto.RegisterFile("CreateOrder.proto", fileDescriptor_2476da9e0746a307) }

var fileDescriptor_2476da9e0746a307 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x61, 0x4b, 0x02, 0x41,
	0x14, 0xe4, 0xce, 0xd4, 0x7c, 0x5a, 0xd1, 0x6a, 0x71, 0x48, 0xc4, 0xb5, 0xf5, 0xc1, 0x08, 0x2c,
	0x8c, 0x7e, 0x40, 0x2a, 0x84, 0x10, 0x24, 0x27, 0x41, 0x5f, 0xaf, 0xbb, 0x47, 0x18, 0x9d, 0x7b,
	0xed, 0xee, 0x49, 0xff, 0xa3, 0x3f, 0x1c, 0xec, 0xad, 0xb6, 0x7b, 0xea, 0x7d, 0xf3, 0x8d, 0x33,
	0x3b, 0xb3, 0x6f, 0x6e, 0xe1, 0x78, 0xc4, 0x31, 0x94, 0xf8, 0xc2, 0x63, 0xe4, 0xfd, 0x94, 0x33,
	0xc9, 0x48, 0x2b, 0x0e, 0x53, 0x8e, 0x3f, 0x61, 0x92, 0x7e, 0xa1, 0xa0, 0x9f, 0x40, 0x0c, 0x4a,
	0x80, 0xdf, 0x19, 0x0a, 0x49, 0xce, 0xa0, 0x31, 0xe5, 0x2c, 0xce, 0x22, 0x39, 0x19, 0x7b, 0x8e,
	0xef, 0xf4, 0x1a, 0xc1, 0x3f, 0x40, 0x4e, 0xa1, 0xf6, 0x98, 0xb0, 0x6c, 0x21, 0x3d, 0xd7, 0x77,
	0x7a, 0xd5, 0x40, 0x4f, 0xe4, 0x1c, 0x60, 0x94, 0x09, 0xc9, 0x12, 0xe4, 0x93, 0xb1, 0x57, 0x51,
	0x32, 0x03, 0xa1, 0xb7, 0xd0, 0xb6, 0xbc, 0x44, 0xca, 0x16, 0x02, 0x89, 0x07, 0xf5, 0x59, 0x16,
	0x45, 0x88, 0xb1, 0xb2, 0xda, 0x0f, 0x56, 0x23, 0xbd, 0x83, 0x4e, 0x80, 0x92, 0xcf, 0x71, 0x69,
	0xc7, 0xf3, 0xa0, 0xae, 0xe6, 0x75, 0xb8, 0xd5, 0x48, 0x87, 0x70, 0x52, 0x50, 0x68, 0x93, 0x6b,
	0xa8, 0x2a, 0x40, 0x09, 0x9a, 0x83, 0x76, 0xdf, 0xdc, 0x42, 0x3f, 0xe7, 0xe6, 0x0c, 0xfa, 0x00,
	0xed, 0x27, 0x94, 0xea, 0xf7, 0xf3, 0x5c, 0xc8, 0x95, 0xa9, 0x7d, 0x3b, 0x67, 0xe3, 0x76, 0x23,
	0xe8, 0xd8, 0x32, 0xed, 0x7c, 0x03, 0x35, 0x05, 0x0a, 0xcf, 0xf1, 0x2b, 0xbb, 0xac, 0x35, 0x85,
	0x26, 0x3a, 0x26, 0x39, 0x04, 0x77, 0xed, 0xe2, 0x4e, 0xc6, 0x76, 0x23, 0xee, 0xee, 0x46, 0x2a,
	0x25, 0x8d, 0xec, 0x15, 0x33, 0x0f, 0x7e, 0x5d, 0x68, 0x29, 0xbf, 0x19, 0xf2, 0xe5, 0x3c, 0x42,
	0x12, 0x40, 0xd3, 0xa8, 0x88, 0xf8, 0x76, 0xd6, 0xcd, 0x2f, 0xa5, 0x7b, 0x51, 0xc2, 0xd0, 0x0b,
	0x78, 0x83, 0x03, 0xab, 0x13, 0x42, 0x6d, 0xcd, 0xb6, 0x8a, 0xbb, 0x97, 0xa5, 0x1c, 0x7d, 0xf2,
	0x2b, 0xb4, 0xcc, 0x95, 0x93, 0x42, 0x98, 0x2d, 0x2d, 0x76, 0x69, 0x19, 0x25, 0x3f, 0x76, 0x78,
	0x05, 0x47, 0x1f, 0xb8, 0x40, 0x1e, 0x4a, 0xcc, 0xdf, 0x8c, 0x18, 0x9a, 0xef, 0x68, 0xaa, 0xa0,
	0xf7, 0x9a, 0xfa, 0xeb, 0xfe, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xf2, 0xbe, 0xcc, 0x63, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	RetrieveOrder(ctx context.Context, in *RetrieveOrderRequest, opts ...grpc.CallOption) (*RetrieveOrderResponse, error)
	GetOrderList(ctx context.Context, in *GetOrderListRequest, opts ...grpc.CallOption) (*GetOrderListResponse, error)
}

type orderServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrderServiceClient(cc *grpc.ClientConn) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, "/daprexamples.OrderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) RetrieveOrder(ctx context.Context, in *RetrieveOrderRequest, opts ...grpc.CallOption) (*RetrieveOrderResponse, error) {
	out := new(RetrieveOrderResponse)
	err := c.cc.Invoke(ctx, "/daprexamples.OrderService/RetrieveOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrderList(ctx context.Context, in *GetOrderListRequest, opts ...grpc.CallOption) (*GetOrderListResponse, error) {
	out := new(GetOrderListResponse)
	err := c.cc.Invoke(ctx, "/daprexamples.OrderService/GetOrderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	RetrieveOrder(context.Context, *RetrieveOrderRequest) (*RetrieveOrderResponse, error)
	GetOrderList(context.Context, *GetOrderListRequest) (*GetOrderListResponse, error)
}

// UnimplementedOrderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (*UnimplementedOrderServiceServer) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (*UnimplementedOrderServiceServer) RetrieveOrder(ctx context.Context, req *RetrieveOrderRequest) (*RetrieveOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveOrder not implemented")
}
func (*UnimplementedOrderServiceServer) GetOrderList(ctx context.Context, req *GetOrderListRequest) (*GetOrderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderList not implemented")
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daprexamples.OrderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_RetrieveOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).RetrieveOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daprexamples.OrderService/RetrieveOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).RetrieveOrder(ctx, req.(*RetrieveOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daprexamples.OrderService/GetOrderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderList(ctx, req.(*GetOrderListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "daprexamples.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "RetrieveOrder",
			Handler:    _OrderService_RetrieveOrder_Handler,
		},
		{
			MethodName: "GetOrderList",
			Handler:    _OrderService_GetOrderList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "CreateOrder.proto",
}
