// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/product_channel_exclusivity.proto

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

// Enum describing the availability of a product offer.
type ProductChannelExclusivityEnum_ProductChannelExclusivity int32

const (
	// Not specified.
	ProductChannelExclusivityEnum_UNSPECIFIED ProductChannelExclusivityEnum_ProductChannelExclusivity = 0
	// Used for return value only. Represents value unknown in this version.
	ProductChannelExclusivityEnum_UNKNOWN ProductChannelExclusivityEnum_ProductChannelExclusivity = 1
	// The item is sold through one channel only, either local stores or online
	// as indicated by its ProductChannel.
	ProductChannelExclusivityEnum_SINGLE_CHANNEL ProductChannelExclusivityEnum_ProductChannelExclusivity = 2
	// The item is matched to its online or local stores counterpart, indicating
	// it is available for purchase in both ShoppingProductChannels.
	ProductChannelExclusivityEnum_MULTI_CHANNEL ProductChannelExclusivityEnum_ProductChannelExclusivity = 3
)

var ProductChannelExclusivityEnum_ProductChannelExclusivity_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SINGLE_CHANNEL",
	3: "MULTI_CHANNEL",
}

var ProductChannelExclusivityEnum_ProductChannelExclusivity_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"SINGLE_CHANNEL": 2,
	"MULTI_CHANNEL":  3,
}

func (x ProductChannelExclusivityEnum_ProductChannelExclusivity) String() string {
	return proto.EnumName(ProductChannelExclusivityEnum_ProductChannelExclusivity_name, int32(x))
}

func (ProductChannelExclusivityEnum_ProductChannelExclusivity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4e6325aa75f52106, []int{0, 0}
}

// Availability of a product offer.
type ProductChannelExclusivityEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductChannelExclusivityEnum) Reset()         { *m = ProductChannelExclusivityEnum{} }
func (m *ProductChannelExclusivityEnum) String() string { return proto.CompactTextString(m) }
func (*ProductChannelExclusivityEnum) ProtoMessage()    {}
func (*ProductChannelExclusivityEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e6325aa75f52106, []int{0}
}

func (m *ProductChannelExclusivityEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductChannelExclusivityEnum.Unmarshal(m, b)
}
func (m *ProductChannelExclusivityEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductChannelExclusivityEnum.Marshal(b, m, deterministic)
}
func (m *ProductChannelExclusivityEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductChannelExclusivityEnum.Merge(m, src)
}
func (m *ProductChannelExclusivityEnum) XXX_Size() int {
	return xxx_messageInfo_ProductChannelExclusivityEnum.Size(m)
}
func (m *ProductChannelExclusivityEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductChannelExclusivityEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ProductChannelExclusivityEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.ProductChannelExclusivityEnum_ProductChannelExclusivity", ProductChannelExclusivityEnum_ProductChannelExclusivity_name, ProductChannelExclusivityEnum_ProductChannelExclusivity_value)
	proto.RegisterType((*ProductChannelExclusivityEnum)(nil), "google.ads.googleads.v2.enums.ProductChannelExclusivityEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/product_channel_exclusivity.proto", fileDescriptor_4e6325aa75f52106)
}

var fileDescriptor_4e6325aa75f52106 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0xb5, 0x1d, 0x28, 0x64, 0xa8, 0xb3, 0x37, 0xc5, 0x09, 0xdb, 0x0f, 0x48, 0xa1, 0xde, 0xe2,
	0x41, 0xb2, 0x19, 0x67, 0x71, 0xc6, 0xc2, 0xdc, 0x04, 0x29, 0xcc, 0xd8, 0x94, 0x58, 0xe8, 0x92,
	0xb2, 0xb4, 0x43, 0x8f, 0xfe, 0x15, 0x8f, 0xfe, 0x14, 0x7f, 0x8a, 0x57, 0xff, 0x80, 0x34, 0x71,
	0xf5, 0x34, 0x2f, 0xe1, 0x91, 0xef, 0x7d, 0xef, 0x7d, 0xef, 0x81, 0x73, 0xa1, 0x94, 0xc8, 0x53,
	0x9f, 0x71, 0xed, 0x5b, 0x58, 0xa3, 0x55, 0xe0, 0xa7, 0xb2, 0x5a, 0x68, 0xbf, 0x58, 0x2a, 0x5e,
	0x25, 0xe5, 0x3c, 0x79, 0x66, 0x52, 0xa6, 0xf9, 0x3c, 0x7d, 0x49, 0xf2, 0x4a, 0x67, 0xab, 0xac,
	0x7c, 0x85, 0xc5, 0x52, 0x95, 0xca, 0xeb, 0xda, 0x2d, 0xc8, 0xb8, 0x86, 0x8d, 0x00, 0x5c, 0x05,
	0xd0, 0x08, 0x1c, 0x1d, 0xaf, 0xf5, 0x8b, 0xcc, 0x67, 0x52, 0xaa, 0x92, 0x95, 0x99, 0x92, 0xda,
	0x2e, 0xf7, 0xdf, 0x1c, 0xd0, 0x8d, 0xac, 0xc5, 0xd0, 0x3a, 0x90, 0x3f, 0x03, 0x22, 0xab, 0x45,
	0xff, 0x11, 0x1c, 0x6e, 0x24, 0x78, 0xfb, 0xa0, 0x3d, 0xa5, 0x93, 0x88, 0x0c, 0xc3, 0xcb, 0x90,
	0x5c, 0x74, 0xb6, 0xbc, 0x36, 0xd8, 0x99, 0xd2, 0x6b, 0x7a, 0x7b, 0x4f, 0x3b, 0x8e, 0xe7, 0x81,
	0xbd, 0x49, 0x48, 0x47, 0x63, 0x32, 0x1f, 0x5e, 0x61, 0x4a, 0xc9, 0xb8, 0xe3, 0x7a, 0x07, 0x60,
	0xf7, 0x66, 0x3a, 0xbe, 0x0b, 0x9b, 0xaf, 0xd6, 0xe0, 0xdb, 0x01, 0xbd, 0x44, 0x2d, 0xe0, 0xbf,
	0x39, 0x06, 0x27, 0x1b, 0xaf, 0x88, 0xea, 0x24, 0x91, 0xf3, 0x30, 0xf8, 0x15, 0x10, 0x2a, 0x67,
	0x52, 0x40, 0xb5, 0x14, 0xbe, 0x48, 0xa5, 0xc9, 0xb9, 0x6e, 0xb6, 0xc8, 0xf4, 0x86, 0xa2, 0xcf,
	0xcc, 0xfb, 0xee, 0xb6, 0x46, 0x18, 0x7f, 0xb8, 0xdd, 0x91, 0x95, 0xc2, 0x5c, 0x43, 0x0b, 0x6b,
	0x34, 0x0b, 0x60, 0x5d, 0x89, 0xfe, 0x5c, 0xcf, 0x63, 0xcc, 0x75, 0xdc, 0xcc, 0xe3, 0x59, 0x10,
	0x9b, 0xf9, 0x97, 0xdb, 0xb3, 0x9f, 0x08, 0x61, 0xae, 0x11, 0x6a, 0x18, 0x08, 0xcd, 0x02, 0x84,
	0x0c, 0xe7, 0x69, 0xdb, 0x1c, 0x76, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x3c, 0x5f, 0x0b,
	0x00, 0x02, 0x00, 0x00,
}
