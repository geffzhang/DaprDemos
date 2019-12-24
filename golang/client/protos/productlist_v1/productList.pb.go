// Code generated by protoc-gen-go. DO NOT EDIT.
// source: productList.proto

package productlist_v1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type ProductListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductListRequest) Reset()         { *m = ProductListRequest{} }
func (m *ProductListRequest) String() string { return proto.CompactTextString(m) }
func (*ProductListRequest) ProtoMessage()    {}
func (*ProductListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1487a50ec114d0d3, []int{0}
}

func (m *ProductListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductListRequest.Unmarshal(m, b)
}
func (m *ProductListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductListRequest.Marshal(b, m, deterministic)
}
func (m *ProductListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductListRequest.Merge(m, src)
}
func (m *ProductListRequest) XXX_Size() int {
	return xxx_messageInfo_ProductListRequest.Size(m)
}
func (m *ProductListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProductListRequest proto.InternalMessageInfo

type ProductList struct {
	Results              []*Product `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProductList) Reset()         { *m = ProductList{} }
func (m *ProductList) String() string { return proto.CompactTextString(m) }
func (*ProductList) ProtoMessage()    {}
func (*ProductList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1487a50ec114d0d3, []int{1}
}

func (m *ProductList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductList.Unmarshal(m, b)
}
func (m *ProductList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductList.Marshal(b, m, deterministic)
}
func (m *ProductList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductList.Merge(m, src)
}
func (m *ProductList) XXX_Size() int {
	return xxx_messageInfo_ProductList.Size(m)
}
func (m *ProductList) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductList.DiscardUnknown(m)
}

var xxx_messageInfo_ProductList proto.InternalMessageInfo

func (m *ProductList) GetResults() []*Product {
	if m != nil {
		return m.Results
	}
	return nil
}

type Product struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_1487a50ec114d0d3, []int{2}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*ProductListRequest)(nil), "productlist.v1.ProductListRequest")
	proto.RegisterType((*ProductList)(nil), "productlist.v1.ProductList")
	proto.RegisterType((*Product)(nil), "productlist.v1.Product")
}

func init() { proto.RegisterFile("productList.proto", fileDescriptor_1487a50ec114d0d3) }

var fileDescriptor_1487a50ec114d0d3 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x4f,
	0x29, 0x4d, 0x2e, 0xf1, 0xc9, 0x2c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83,
	0x0a, 0xe5, 0x80, 0x84, 0xca, 0x0c, 0x95, 0x44, 0xb8, 0x84, 0x02, 0x10, 0x8a, 0x82, 0x52, 0x0b,
	0x4b, 0x53, 0x8b, 0x4b, 0x94, 0x1c, 0xb8, 0xb8, 0x91, 0x44, 0x85, 0x0c, 0xb9, 0xd8, 0x8b, 0x52,
	0x8b, 0x4b, 0x73, 0x4a, 0x8a, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0xc4, 0xf5, 0x50, 0x8d,
	0xd1, 0x83, 0xaa, 0x0e, 0x82, 0xa9, 0x53, 0x92, 0xe4, 0x62, 0x87, 0x8a, 0x09, 0xf1, 0x71, 0x31,
	0x79, 0xba, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x79, 0xba, 0x18, 0xa5, 0x71, 0x09,
	0xc2, 0x94, 0x07, 0x38, 0x07, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x05, 0x72, 0xf1, 0xb9,
	0xa7, 0x96, 0x38, 0xe6, 0xe4, 0x40, 0xa5, 0x8a, 0x85, 0x94, 0x70, 0xd8, 0x81, 0xe4, 0x4e, 0x29,
	0x69, 0x3c, 0x6a, 0x9c, 0x04, 0x57, 0x31, 0xf1, 0x21, 0xf1, 0xf5, 0xc2, 0x0c, 0x93, 0xd8, 0xc0,
	0x81, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x7e, 0x23, 0xa2, 0x19, 0x01, 0x00, 0x00,
}