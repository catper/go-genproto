// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/product_channel.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing the locality of a product offer.
type ProductChannelEnum_ProductChannel int32

const (
	// Not specified.
	ProductChannelEnum_UNSPECIFIED ProductChannelEnum_ProductChannel = 0
	// Used for return value only. Represents value unknown in this version.
	ProductChannelEnum_UNKNOWN ProductChannelEnum_ProductChannel = 1
	// The item is sold online.
	ProductChannelEnum_ONLINE ProductChannelEnum_ProductChannel = 2
	// The item is sold in local stores.
	ProductChannelEnum_LOCAL ProductChannelEnum_ProductChannel = 3
)

var ProductChannelEnum_ProductChannel_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ONLINE",
	3: "LOCAL",
}

var ProductChannelEnum_ProductChannel_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ONLINE":      2,
	"LOCAL":       3,
}

func (x ProductChannelEnum_ProductChannel) String() string {
	return proto.EnumName(ProductChannelEnum_ProductChannel_name, int32(x))
}

func (ProductChannelEnum_ProductChannel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8c7e543098a4fcf9, []int{0, 0}
}

// Locality of a product offer.
type ProductChannelEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductChannelEnum) Reset()         { *m = ProductChannelEnum{} }
func (m *ProductChannelEnum) String() string { return proto.CompactTextString(m) }
func (*ProductChannelEnum) ProtoMessage()    {}
func (*ProductChannelEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c7e543098a4fcf9, []int{0}
}

func (m *ProductChannelEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductChannelEnum.Unmarshal(m, b)
}
func (m *ProductChannelEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductChannelEnum.Marshal(b, m, deterministic)
}
func (m *ProductChannelEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductChannelEnum.Merge(m, src)
}
func (m *ProductChannelEnum) XXX_Size() int {
	return xxx_messageInfo_ProductChannelEnum.Size(m)
}
func (m *ProductChannelEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductChannelEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ProductChannelEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.ProductChannelEnum_ProductChannel", ProductChannelEnum_ProductChannel_name, ProductChannelEnum_ProductChannel_value)
	proto.RegisterType((*ProductChannelEnum)(nil), "google.ads.googleads.v2.enums.ProductChannelEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/product_channel.proto", fileDescriptor_8c7e543098a4fcf9)
}

var fileDescriptor_8c7e543098a4fcf9 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdf, 0x4a, 0xfb, 0x30,
	0x14, 0xfe, 0xad, 0xe3, 0x37, 0x31, 0x03, 0x2d, 0xf1, 0x4e, 0xdc, 0xc5, 0xf6, 0x00, 0x09, 0x74,
	0x77, 0xf1, 0x2a, 0x9b, 0x75, 0x0c, 0x47, 0x56, 0x90, 0x4d, 0xd0, 0x82, 0xc4, 0xa6, 0xc4, 0x42,
	0x9b, 0x94, 0xa6, 0xdd, 0x03, 0x79, 0xe9, 0xa3, 0xf8, 0x24, 0xe2, 0x53, 0x48, 0x13, 0x5b, 0xd8,
	0x85, 0xde, 0x84, 0x8f, 0xf3, 0xfd, 0xc9, 0x77, 0x0e, 0x98, 0x4b, 0xad, 0x65, 0x9e, 0x62, 0x2e,
	0x0c, 0x76, 0xb0, 0x45, 0x87, 0x00, 0xa7, 0xaa, 0x29, 0x0c, 0x2e, 0x2b, 0x2d, 0x9a, 0xa4, 0x7e,
	0x4e, 0x5e, 0xb9, 0x52, 0x69, 0x8e, 0xca, 0x4a, 0xd7, 0x1a, 0x4e, 0x9c, 0x12, 0x71, 0x61, 0x50,
	0x6f, 0x42, 0x87, 0x00, 0x59, 0xd3, 0xe5, 0x55, 0x97, 0x59, 0x66, 0x98, 0x2b, 0xa5, 0x6b, 0x5e,
	0x67, 0x5a, 0x19, 0x67, 0x9e, 0x3d, 0x01, 0x18, 0xb9, 0xd4, 0xa5, 0x0b, 0x0d, 0x55, 0x53, 0xcc,
	0x42, 0x70, 0x76, 0x3c, 0x85, 0xe7, 0x60, 0xbc, 0x63, 0xf7, 0x51, 0xb8, 0x5c, 0xdf, 0xae, 0xc3,
	0x1b, 0xff, 0x1f, 0x1c, 0x83, 0x93, 0x1d, 0xbb, 0x63, 0xdb, 0x07, 0xe6, 0x0f, 0x20, 0x00, 0xa3,
	0x2d, 0xdb, 0xac, 0x59, 0xe8, 0x7b, 0xf0, 0x14, 0xfc, 0xdf, 0x6c, 0x97, 0x74, 0xe3, 0x0f, 0x17,
	0x9f, 0x03, 0x30, 0x4d, 0x74, 0x81, 0xfe, 0x2c, 0xb8, 0xb8, 0x38, 0xfe, 0x2a, 0x6a, 0x7b, 0x45,
	0x83, 0xc7, 0xc5, 0x8f, 0x4b, 0xea, 0x9c, 0x2b, 0x89, 0x74, 0x25, 0xb1, 0x4c, 0x95, 0x6d, 0xdd,
	0xdd, 0xa6, 0xcc, 0xcc, 0x2f, 0xa7, 0xba, 0xb6, 0xef, 0x9b, 0x37, 0x5c, 0x51, 0xfa, 0xee, 0x4d,
	0x56, 0x2e, 0x8a, 0x0a, 0x83, 0x1c, 0x6c, 0xd1, 0x3e, 0x40, 0xed, 0xb2, 0xe6, 0xa3, 0xe3, 0x63,
	0x2a, 0x4c, 0xdc, 0xf3, 0xf1, 0x3e, 0x88, 0x2d, 0xff, 0xe5, 0x4d, 0xdd, 0x90, 0x10, 0x2a, 0x0c,
	0x21, 0xbd, 0x82, 0x90, 0x7d, 0x40, 0x88, 0xd5, 0xbc, 0x8c, 0x6c, 0xb1, 0xf9, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x52, 0x69, 0xdc, 0xfc, 0xc2, 0x01, 0x00, 0x00,
}
